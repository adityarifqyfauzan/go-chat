package user

import (
	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/middleware"
	"github.com/gin-gonic/gin"
)

func New(params *config.Params, api *gin.RouterGroup) {
	handler := NewHandler(params)

	users := api.Group("users")
	users.Use(middleware.AuthMiddleware(params))
	users.GET("", handler.FindUser)
}
