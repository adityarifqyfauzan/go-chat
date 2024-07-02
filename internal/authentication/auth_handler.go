package authentication

import (
	"net/http"

	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/internal/authentication/dto"
	"github.com/adityarifqyfauzan/go-chat/pkg/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	params *config.Params
	u      Usecase
}

func NewHandler(params *config.Params) *handler {
	return &handler{
		params: params,
		u:      NewUsecase(params),
	}
}

func (h *handler) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	resp := h.u.Register(req)
	utils.WriteToResponseBody(ctx, resp)
}

func (h *handler) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	resp := h.u.Login(req)
	utils.WriteToResponseBody(ctx, resp)
}
