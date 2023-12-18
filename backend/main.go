package main

import (
	"github.com/go-playground/validator/v10"
	"shrt-server/handler"
	"shrt-server/pkg/db"
	"shrt-server/repository"
	"shrt-server/server"
	"shrt-server/service"
	"shrt-server/utilities/configuration"
)

func main() {
	config := configuration.InitConfig()
	database := db.Init(config)
	payloadValidator := validator.New()

	shrtRepository := repository.NewShrtRepository(database)
	shrtService := service.NewShrtService(shrtRepository)
	shrtRoute := handler.NewShrtHandler(shrtService, payloadValidator, &config)

	shrtServer := server.NewServer(shrtRoute, config)
	shrtServer.Initialize()
}
