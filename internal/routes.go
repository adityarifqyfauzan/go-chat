package internal

import (
	"net/http"

	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/internal/authentication"
	"github.com/adityarifqyfauzan/go-chat/internal/message"
	"github.com/adityarifqyfauzan/go-chat/internal/user"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	params *config.Params
	engine *gin.Engine
}

func New(params *config.Params, engine *gin.Engine) *Routes {
	return &Routes{
		params: params,
		engine: engine,
	}
}

func (r *Routes) RegisterRoutes() {
	// register all routes here
	api := r.engine.Group("api")
	authentication.New(r.params, api)
	user.New(r.params, api)
	message.New(r.params, api)

	// health check
	api.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello, World!",
		})
	})

}
