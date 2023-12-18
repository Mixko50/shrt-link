package middleware

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"shrt-server/types"
	"shrt-server/utilities/text"
)

func PathNotFoundError(c *fiber.Ctx) error {
	log.Error(c.Method() + " " + c.Path() + " Not found")
	return types.ErrorResponse(c, fiber.StatusNotFound, text.Ptr(c.Path()+" Not found"))
}
