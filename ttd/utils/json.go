package utils

import (
	"net/http"

	"github.com/gofiber/fiber"
)

func JSON(ctx *fiber.Ctx, statusCode int, data interface{}) {
	ctx.Status(statusCode)
	err := ctx.JSON(data)
	if err != nil {
		Error(ctx, http.StatusUnprocessableEntity, err)
	}
}

func Error(ctx *fiber.Ctx, statusCode int, err error) {
	JSON(ctx, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
