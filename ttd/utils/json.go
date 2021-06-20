package utils

import (
	"github.com/gofiber/fiber/v2"
)

func JSON(ctx *fiber.Ctx, statusCode int, data interface{}) error {
	ctx.Status(statusCode)
	return ctx.JSON(data)
}

func Error(ctx *fiber.Ctx, statusCode int, err error) error {
	return JSON(ctx, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
