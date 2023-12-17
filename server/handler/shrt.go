package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"shrt-server/service"
	"shrt-server/types"
	"shrt-server/types/request"
	"shrt-server/utilities/config"
	"shrt-server/utilities/text"
)

type shrtHandler struct {
	shrtService service.ShrtService
	validator   *validator.Validate
}

func NewShrtHandler(shrtService service.ShrtService, validator *validator.Validate) shrtHandler {
	return shrtHandler{
		shrtService: shrtService,
		validator:   validator,
	}
}

func (h shrtHandler) CreateShrtLink(c *fiber.Ctx) error {
	body := new(request.CreateShortenLinkRequest)
	if err := c.BodyParser(body); err != nil {
		return text.FormatBodyParserErrorResponse(err, c)
	}

	if validateErr := h.validator.Struct(*body); validateErr != nil {
		return text.FormatValidatorErrorResponse(c, validateErr)
	}

	data, err := h.shrtService.CreateShrtLink(body)
	if err != nil {
		return types.ErrorResponse(c, fiber.StatusInternalServerError, text.Ptr(err.Error()))
	}

	return types.SuccessResponse(c, data)
}

func (h shrtHandler) GetOriginalURL(c *fiber.Ctx) error {
	slug := c.Query("slug")

	if slug == "" {
		return types.ErrorResponse(c, fiber.StatusBadRequest, text.Ptr(types.ErrSomethingWentWrong.Error()))
	}

	data, err := h.shrtService.GetOriginalURL(slug)
	if err != nil {
		return types.ErrorResponse(c, fiber.StatusInternalServerError, text.Ptr(types.ErrSomethingWentWrong.Error()))
	}

	return types.SuccessResponse(c, data)
}

func (h shrtHandler) GetOriginalURLToRedirect(c *fiber.Ctx) error {
	slug := c.Params("slug")

	if slug == "" {
		return c.Redirect(config.C.BaseUrl, fiber.StatusMovedPermanently)
	}

	data, err := h.shrtService.GetOriginalURLToRedirect(slug)
	if err != nil {
		return c.Redirect(config.C.BaseUrl, fiber.StatusMovedPermanently)
	}

	return c.Redirect(data, fiber.StatusMovedPermanently)
}
