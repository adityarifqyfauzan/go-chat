package internal

import (
	"net/http"

	"github.com/adityarifqyfauzan/go-chat/config"
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

	// health check
	r.engine.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello, World!",
		})
	})

}
