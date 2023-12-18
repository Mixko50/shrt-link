package middleware

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func PathLog(c *fiber.Ctx) error {
	log.Info(c.Method() + " " + c.Path())
	return c.Next()
}
