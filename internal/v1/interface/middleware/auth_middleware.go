package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Authentication(ctx *fiber.Ctx) error {

	return ctx.Next()
}
