package authentication

import (
	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/gin-gonic/gin"
)

func New(params *config.Params, api *gin.RouterGroup) {
	// init handler
	handler := NewHandler(params)

	auth := api.Group("auth")
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)
}
