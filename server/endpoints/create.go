package endpoints

import "github.com/gofiber/fiber/v2"

func CreateShortLink(c *fiber.Ctx) error {
	return c.SendString("OK")
}
