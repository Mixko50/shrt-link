package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"shrt-server/handler"
	"shrt-server/pkg/middleware"
	"shrt-server/utilities/configuration"
)

type shrtServer struct {
	shrtHandler handler.ShrtHandler
	config      configuration.Config
}

func NewServer(shrtHandler handler.ShrtHandler, config configuration.Config) *shrtServer {
	return &shrtServer{
		shrtHandler: shrtHandler,
		config:      config,
	}
}

func (s *shrtServer) Initialize() {
	app := fiber.New(fiber.Config{
		AppName:       "Shrt service",
		CaseSensitive: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: s.config.AllowedOrigins,
	}))

	app.Use(middleware.PathLog)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Use(recover2.New())

	app.Get("/:slug", s.shrtHandler.GetOriginalURLToRedirect)

	apiGroup := app.Group("/api/v1/")
	apiGroup.Post("/create", s.shrtHandler.CreateShrtLink)
	apiGroup.Get("/retrieve", s.shrtHandler.GetOriginalURL)

	app.Use(middleware.PathNotFoundError)

	err := app.Listen(s.config.ServerPort)
	if err != nil {
		panic(err)
	}
}
