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
		return types.ErrorResponse(c, fiber.StatusBadRequest, Ptr(types.ErrSomethingWentWrong.Error()))
	}
	log.Error("no body provided")
	return types.ErrorResponse(c, fiber.StatusBadRequest, Ptr(types.ErrSomethingWentWrong.Error()))
}
