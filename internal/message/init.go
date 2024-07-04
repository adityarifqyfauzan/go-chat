package message

import (
	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/middleware"
	"github.com/gin-gonic/gin"
)

func New(params *config.Params, api *gin.RouterGroup) {
	handler := NewHandler(params)

	message := api.Group("message")
	message.Use(middleware.AuthMiddleware(params))
	message.GET("/room", handler.FindRooms)
	message.POST("/room", handler.CreateRoom)
	message.GET("/conversation/:room_id", handler.FindConversation)
}
