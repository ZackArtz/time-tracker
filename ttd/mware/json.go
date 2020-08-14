package mware

import "github.com/gofiber/fiber"

func JsonMiddleware(c *fiber.Ctx) {
	c.Set("Content-Type", "application/json")
	c.Next()
}
