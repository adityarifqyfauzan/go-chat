package user

import (
	"errors"
	"fmt"

	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/internal/user/domain/repository"
	"github.com/adityarifqyfauzan/go-chat/internal/user/dto"
	"github.com/adityarifqyfauzan/go-chat/pkg/exception"
	"github.com/adityarifqyfauzan/go-chat/pkg/utils"
	"gorm.io/gorm"
)

type Usecase interface {
	FindUser(req dto.UserRequest) utils.WebResponse
}

type usecase struct {
	params   *config.Params
	userRepo repository.UserRepository
}

func NewUsecase(params *config.Params) Usecase {
	return &usecase{
		params:   params,
		userRepo: repository.NewUserRepository(params.DB),
	}
}

func (u *usecase) FindUser(req dto.UserRequest) utils.WebResponse {
	users, err := u.userRepo.FindByUsername(req.Username, req.Page, req.Size)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to find users: %v", err)))
	}

	return utils.NewResponseWithPagination("here is your data", users, &utils.MetaData{
		Page:  req.Page,
		Size:  req.Size,
		Total: int(u.userRepo.CountByUsername(req.Username)),
	})
}
