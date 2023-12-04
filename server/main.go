package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"shrt-server/db"
	"shrt-server/endpoints"
	"shrt-server/utilities/config"
)

func main() {
	config.InitConfig()
	db.Init()

	app := fiber.New(fiber.Config{
		AppName: "Shrt service",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: config.C.AllowedOrigins,
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/create", endpoints.CreateShortLink)
	app.Get("/:id", endpoints.Redirect)
	app.Post("/retrieve", endpoints.RetrieveOriginalLink)

	app.Use(recover2.New())

	err := app.Listen(config.C.ServerPort)
	if err != nil {
		panic(err)
	}

}
