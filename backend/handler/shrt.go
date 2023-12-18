package handler

import "github.com/gofiber/fiber/v2"

type ShrtHandler interface {
	CreateShrtLink(c *fiber.Ctx) error
	GetOriginalURL(c *fiber.Ctx) error
	GetOriginalURLToRedirect(c *fiber.Ctx) error
}
