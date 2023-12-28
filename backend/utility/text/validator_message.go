package text

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"shrt-server/types"
	"strings"
)

func FormatValidatorErrorResponse(c *fiber.Ctx, err error) error {
	var validateError validator.ValidationErrors
	errors.As(err, &validateError)
	if validateError[0].Tag() != "required" {
		log.Error(validateError[0].Error())
		return types.ErrorResponse(c, fiber.StatusBadRequest, Ptr(validateError[0].Field()+" should be "+validateError[0].Tag()))
	}
	log.Error(strings.ToLower(validateError[0].Error()))
	return types.ErrorResponse(c, fiber.StatusBadRequest, Ptr(validateError[0].Field()+" is required"))
}
