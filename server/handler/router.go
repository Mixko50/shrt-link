package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"shrt-server/pkg/db"
	"shrt-server/repository"
	"shrt-server/service"
)

func Router(app fiber.Router) {
	database := db.Init()
	payloadValidator := validator.New()

	shrtRepository := repository.NewShrtRepository(database)
	shrtService := service.NewShrtService(shrtRepository)
	shrtRoute := NewShrtHandler(shrtService, payloadValidator)

	app.Get("/:slug", shrtRoute.GetOriginalURLToRedirect)

	apiGroup := app.Group("/api/v1/")
	apiGroup.Post("/create", shrtRoute.CreateShrtLink)
	apiGroup.Get("/retrieve", shrtRoute.GetOriginalURL)
}
