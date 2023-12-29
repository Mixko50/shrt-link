package text

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"shrt-server/types"
)

func FormatBodyParserErrorResponse(err error, c *fiber.Ctx) error {
	var typeError *fiber.UnmarshalTypeError
	if errors.As(err, &typeError) {
		log.Error(typeError.Field + " should be " + typeError.Type.String())
		return types.ErrorResponse(c, fiber.StatusBadRequest, Ptr(typeError.Field+" should be "+typeError.Type.String()))
	}
	log.Error("no data provided")
	return types.ErrorResponse(c, fiber.StatusBadRequest, Ptr("no data provided"))
}
