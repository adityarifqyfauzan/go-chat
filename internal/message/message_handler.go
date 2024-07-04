package message

import (
	"net/http"

	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/internal/message/dto"
	"github.com/adityarifqyfauzan/go-chat/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	params *config.Params
	uc     Usecase
}

func NewHandler(params *config.Params) *handler {
	return &handler{
		params: params,
		uc:     NewUsecase(params),
	}
}

// find all rooms by user_id
func (h *handler) FindRooms(ctx *gin.Context) {
	var req dto.RoomRequest
	userID, _ := ctx.Get("user_id")
	req.AuthenticatedUser = uint(userID.(int))
	if err := ctx.BindQuery(&req); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	resp := h.uc.FindRooms(req)
	utils.WriteToResponseBody(ctx, resp)
}

// create new room
func (h *handler) CreateRoom(ctx *gin.Context) {
	var req dto.RoomCreateRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	resp := h.uc.CreateRoom(req)
	utils.WriteToResponseBody(ctx, resp)
}

// find conversation by room id
func (h *handler) FindConversation(ctx *gin.Context) {
	var req dto.ConversationRequest
	userID, _ := ctx.Get("user_id")
	req.AuthenticatedUser = uint(userID.(int))
	if err := ctx.BindUri(&req); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	if err := ctx.BindQuery(&req); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	resp := h.uc.FindConversation(req)
	utils.WriteToResponseBody(ctx, resp)
}
