package message

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/internal/message/domain/model"
	"github.com/adityarifqyfauzan/go-chat/internal/message/domain/repository"
	"github.com/adityarifqyfauzan/go-chat/internal/message/dto"
	"github.com/adityarifqyfauzan/go-chat/pkg/exception"
	"github.com/adityarifqyfauzan/go-chat/pkg/utils"
	"gorm.io/gorm"
)

type Usecase interface {
	FindRooms(req dto.RoomRequest) utils.WebResponse
	CreateRoom(req dto.RoomCreateRequest) utils.WebResponse
	FindConversation(req dto.ConversationRequest) utils.WebResponse
}

type usecase struct {
	params       *config.Params
	roomRepo     repository.RoomRepository
	userRoomRepo repository.UserRoomRepository
	messageRepo  repository.MessageRepository
}

func NewUsecase(params *config.Params) Usecase {
	return &usecase{
		params:       params,
		roomRepo:     repository.NewRoomRepository(params.DB),
		userRoomRepo: repository.NewUserRoomRepository(params.DB),
		messageRepo:  repository.NewMessageRepository(params.DB),
	}
}

func (uc *usecase) FindRooms(req dto.RoomRequest) utils.WebResponse {
	rooms, err := uc.roomRepo.FindByUserID(req.AuthenticatedUser, req.Page, req.Size)
	if err != nil {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to find rooms: %v", err)))
	}

	userRooms := make([]*dto.RoomResponse, 0)
	for _, r := range rooms {
		roomResp := dto.RoomResponse{}
		roomResp.ModelToDto(r)
		userRooms = append(userRooms, &roomResp)
	}

	return utils.NewResponseWithPagination("here is your data", userRooms, &utils.MetaData{
		Page:  req.Page,
		Size:  req.Size,
		Total: int(uc.roomRepo.Count(req.AuthenticatedUser)),
	})
}

func (uc *usecase) CreateRoom(req dto.RoomCreateRequest) utils.WebResponse {
	existingRoom, err := uc.userRoomRepo.CheckExistingRoom(req.AuthenticatedUser, req.FriendID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to check existing room: %v", err)))
	}

	if existingRoom.ID != 0 {
		return utils.NewResponse(http.StatusOK, "room registered", existingRoom)
	}

	tx := uc.params.DB.Begin()
	defer tx.Rollback()

	// create room
	roomID, err := uc.roomRepo.Create(tx)
	if err != nil {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to create room: %v", err)))
	}

	// create user room	for sender
	senderUserRoom := &model.UserRoom{
		UserID:    req.AuthenticatedUser,
		RoomID:    roomID,
		CreatedAt: time.Now(),
	}
	err = uc.userRoomRepo.Create(tx, senderUserRoom)
	if err != nil {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to create sender user room: %v", err)))
	}

	receiverUserRoom := &model.UserRoom{
		UserID:    req.FriendID,
		RoomID:    roomID,
		CreatedAt: time.Now(),
	}
	err = uc.userRoomRepo.Create(tx, receiverUserRoom)
	if err != nil {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to create receiver user room: %v", err)))
	}

	tx.Commit()
	return utils.NewResponse(http.StatusCreated, "room registered", nil)
}

func (uc *usecase) FindConversation(req dto.ConversationRequest) utils.WebResponse {
	// find available room
	rooms, err := uc.roomRepo.FindRoomBy(map[string]interface{}{
		"id": req.RoomID,
	})
	if err != nil {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to find room: %v", err)))
	}

	if len(rooms) == 0 {
		panic(exception.NewUnprocessableEntityException("no room found!"))
	}

	// find conversation
	conversation, err := uc.messageRepo.FindConversation(req.RoomID, req.Page, req.Size)
	if err != nil {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to find conversation: %v", err)))
	}

	conversationResp := make([]*dto.ConversationResponse, 0)
	for _, c := range conversation {
		newConversation := dto.ConversationResponse{AuthenticatedUser: req.AuthenticatedUser}
		newConversation.ModelToDto(c)
		conversationResp = append(conversationResp, &newConversation)
	}

	return utils.NewResponseWithPagination("here is your data", conversationResp, &utils.MetaData{
		Page:  req.Page,
		Size:  req.Size,
		Total: int(uc.messageRepo.CountConversation(req.RoomID)),
	})
}
