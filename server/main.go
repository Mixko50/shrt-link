package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"shrt-server/handler"
	"shrt-server/pkg/middleware"
	"shrt-server/utilities/config"
)

func main() {
	config.InitConfig()

	app := fiber.New(fiber.Config{
		AppName:       "Shrt service",
		CaseSensitive: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: config.C.AllowedOrigins,
	}))

	app.Use(middleware.PathLog)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Use(recover2.New())

	app.Route("/", handler.Router)

	app.Use(middleware.PathNotFoundError)

	err := app.Listen(config.C.ServerPort)
	if err != nil {
		panic(err)
	}

}
