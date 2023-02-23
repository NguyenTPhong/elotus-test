package controller

import (
	"elotus/internal/v1/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	userUseCase usecase.UserUseCase
}

func NewAuthController(userUseCase usecase.UserUseCase) *AuthController {
	return &AuthController{userUseCase}
}

func (c *AuthController) RegisNewUser(ctx *fiber.Ctx) error {
	panic("implement me")
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	panic("implement me")
}
