package authentication

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/internal/authentication/domain/model"
	"github.com/adityarifqyfauzan/go-chat/internal/authentication/domain/repository"
	"github.com/adityarifqyfauzan/go-chat/internal/authentication/dto"
	"github.com/adityarifqyfauzan/go-chat/pkg/exception"
	"github.com/adityarifqyfauzan/go-chat/pkg/jwt"
	"github.com/adityarifqyfauzan/go-chat/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usecase interface {
	Register(req dto.RegisterRequest) utils.WebResponse
	Login(req dto.LoginRequest) utils.WebResponse
	Me(req dto.Me) utils.WebResponse
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

func (u *usecase) Register(req dto.RegisterRequest) utils.WebResponse {
	// check user by username
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"username": req.Username,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to find user: %v", err)))
	}

	// check if user exists
	if user.ID != 0 {
		panic(exception.NewUnprocessableEntityException("username has been taken!"))
	}

	// create new user
	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to hash password: %v", err)))
	}
	newUser := model.User{
		Username: req.Username,
		Name:     req.Name,
		Password: string(hashPass),
	}

	if err := u.userRepo.Create(&newUser); err != nil {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to create new user: %v", err)))
	}

	return utils.NewResponse(http.StatusCreated, "register successfully", nil)
}

func (u *usecase) Login(req dto.LoginRequest) utils.WebResponse {
	if req.Username == "" {
		panic(exception.NewBadRequestException("Username is required!"))
	}

	if req.Password == "" {
		panic(exception.NewBadRequestException("Password is required!"))
	}

	// check user by username
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"username": req.Username,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to find user: %v", err)))
	}

	// check if user exists
	if user.ID == 0 {
		panic(exception.NewUnprocessableEntityException("User is not registered!"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		panic(exception.NewUnprocessableEntityException("Password is wrong!"))
	}

	token, err := jwt.SignToken(u.params, int(user.ID))
	if err != nil {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to generate jwt: %v", err)))
	}

	return utils.NewResponse(http.StatusOK, "login successfully", map[string]interface{}{
		"token": token,
	})
}

func (u *usecase) Me(req dto.Me) utils.WebResponse {
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"id": req.AuthenticatedUser,
	})
	if err != nil {
		panic(exception.NewUnprocessableEntityException(fmt.Sprintf("unable to find user: %v", err)))
	}
	return utils.NewResponse(http.StatusOK, "here is your data", user)
}
