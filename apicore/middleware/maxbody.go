package middleware

import "github.com/gofiber/fiber/v2"

func MaxBody(size int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if len(c.Body()) >= size {
			return fiber.ErrRequestEntityTooLarge
		}
		return c.Next()
	}
}
