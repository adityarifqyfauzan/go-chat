package main

import (
	"github.com/adityarifqyfauzan/go-chat/cmd/server"
	"github.com/adityarifqyfauzan/go-chat/config"
	"github.com/adityarifqyfauzan/go-chat/pkg/database"
)

func main() {
	params := new(config.Params)
	params.Env = config.New("config", nil)
	params.DB = database.InitPostgreDB(params.Env)

	server := server.New(params)
	server.Start()
}
