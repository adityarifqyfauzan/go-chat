package user

import (
	"net/http"

	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/internal/user/dto"
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

func (h *handler) FindUser(ctx *gin.Context) {
	var req dto.UserRequest
	if err := ctx.BindQuery(&req); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	if err := ctx.BindUri(&req); err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	resp := h.u.FindUser(req)
	utils.WriteToResponseBody(ctx, resp)
}
