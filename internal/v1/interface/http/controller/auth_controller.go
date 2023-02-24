package controller

import (
	"elotus/internal/v1/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	userUseCase usecase.AuthUseCase
}

func NewAuthController(userUseCase usecase.AuthUseCase) *AuthController {
	return &AuthController{userUseCase}
}

func (c *AuthController) RegisNewUser(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	panic("implement me")
}
