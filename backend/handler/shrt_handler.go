package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"shrt-server/service"
	"shrt-server/types"
	"shrt-server/types/request"
	"shrt-server/utility/configuration"
	text2 "shrt-server/utility/text"
)

type shrtHandler struct {
	shrtService service.ShrtService
	validator   *validator.Validate
	config      *configuration.Config
}

func NewShrtHandler(shrtService service.ShrtService, validator *validator.Validate, config *configuration.Config) shrtHandler {
	return shrtHandler{
		shrtService: shrtService,
		validator:   validator,
		config:      config,
	}
}

func (h shrtHandler) CreateShrtLink(c *fiber.Ctx) error {
	body := new(request.CreateShortenLinkRequest)
	if err := c.BodyParser(body); err != nil {
		return text2.FormatBodyParserErrorResponse(err, c)
	}

	if validateErr := h.validator.Struct(*body); validateErr != nil {
		return text2.FormatValidatorErrorResponse(c, validateErr)
	}

	data, err := h.shrtService.CreateShrtLink(body)
	if err != nil {
		return types.ErrorResponse(c, fiber.StatusInternalServerError, text2.Ptr(err.Error()))
	}

	return types.SuccessResponse(c, data)
}

func (h shrtHandler) GetOriginalURL(c *fiber.Ctx) error {
	slug := c.Query("slug")

	if slug == "" {
		return types.ErrorResponse(c, fiber.StatusBadRequest, text2.Ptr(types.ErrSlugIsRequired.Error()))
	}

	data, err := h.shrtService.GetOriginalURL(slug)
	if err != nil {
		return types.ErrorResponse(c, fiber.StatusInternalServerError, text2.Ptr(err.Error()))
	}

	return types.SuccessResponse(c, data)
}

func (h shrtHandler) GetOriginalURLToRedirect(c *fiber.Ctx) error {
	slug := c.Params("slug")

	data, err := h.shrtService.GetOriginalURLToRedirect(slug)
	if err != nil {
		return c.Redirect(h.config.BaseUrl, fiber.StatusMovedPermanently)
	}

	return c.Redirect(data, fiber.StatusMovedPermanently)
}
