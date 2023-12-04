package endpoints

import "github.com/gofiber/fiber/v2"

func RetrieveOriginalLink(c *fiber.Ctx) error {
	return c.SendString("OK")
}
