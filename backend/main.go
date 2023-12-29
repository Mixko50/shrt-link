package main

import (
	"github.com/go-playground/validator/v10"
	"shrt-server/handler"
	"shrt-server/pkg/db"
	"shrt-server/repository"
	"shrt-server/server"
	"shrt-server/service"
	"shrt-server/utility"
	"shrt-server/utility/configuration"
)

func main() {
	config := configuration.InitConfig()
	database := db.Init(config)
	utilityText := utility.NewUtility()
	payloadValidator := validator.New()

	shrtRepository := repository.NewShrtRepository(database)
	shrtService := service.NewShrtService(shrtRepository, utilityText)
	shrtRoute := handler.NewShrtHandler(shrtService, payloadValidator, &config)

	shrtServer := server.NewServer(shrtRoute, config)
	shrtServer.Initialize()
}
