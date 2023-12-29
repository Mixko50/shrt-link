package types

import "github.com/gofiber/fiber/v2"

type Response[T any] struct {
	Success bool    `json:"success"`
	Payload *T      `json:"payload,omitempty"`
	Message *string `json:"message,omitempty"`
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message *string) error {
	return c.Status(statusCode).JSON(Response[*string]{
		Success: false,
		Payload: nil,
		Message: message,
	})
}

func SuccessResponse[T any](c *fiber.Ctx, data *T) error {
	return c.Status(fiber.StatusOK).JSON(Response[T]{
		Success: true,
		Payload: data,
	})
}
