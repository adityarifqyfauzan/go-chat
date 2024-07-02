package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/internal"
	"github.com/adityarifqyfauzan/go-chat/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	params *config.Params
}

func New(params *config.Params) *Server {
	return &Server{params: params}
}

func (r *Server) Start() {
	engine := gin.New()

	server := &http.Server{
		Addr:    r.params.Env.GetString("app.port.rest"),
		Handler: engine,
	}

	engine.Use(middleware.ExceptionMiddleware())

	// register routes
	routes := internal.New(r.params, engine)
	routes.RegisterRoutes()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
