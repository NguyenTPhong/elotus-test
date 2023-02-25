package middleware

import (
	"elotus/internal/v1/entity"
	"elotus/internal/v1/usecase"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthMiddleware(authUseCase usecase.AuthUseCase) *AuthMiddleware {
	return &AuthMiddleware{authUseCase}
}

func (c *AuthMiddleware) Authentication(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	token := headers["Authorization"]
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(entity.ResponseError{
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	token = strings.Replace(token, "Bearer ", "", 1)
	session, err := c.authUseCase.ValidateToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(entity.ResponseError{
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	ctx.Locals("session", session)
	return ctx.Next()
}
