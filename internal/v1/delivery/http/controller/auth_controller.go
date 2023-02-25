package controller

import (
	_const "elotus/const"
	"elotus/internal/v1/entity"
	"elotus/internal/v1/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	userUseCase usecase.AuthUseCase
}

func NewAuthController(userUseCase usecase.AuthUseCase) *AuthController {
	return &AuthController{userUseCase}
}

// RegisNewUser godoc
// @Summary register new user
// @Tags Authentication
// @ID register-new-user
// @Accept json
// @Produce json
// @Param json body entity.CreateUserRequest true "json body"
// @Success 200 {object} entity.CreateUserResponse
// @Failure 400 {object} entity.ResponseError
// @Router /api/v1/auth/register [post]
func (c *AuthController) RegisNewUser(ctx *fiber.Ctx) error {
	var req entity.CreateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	res, err := c.userUseCase.CreateUser(ctx.Context(), &req)
	if err != nil {
		errorCode := fiber.StatusInternalServerError
		if code, ok := _const.ErrorCode[err.Error()]; ok {
			errorCode = code
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Code:    errorCode,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

// Login godoc
// @Summary log in
// @Tags Authentication
// @ID login
// @Accept json
// @Produce json
// @Param json body entity.LoginRequest true "json body"
// @Success 200 {object} entity.LoginResponse
// @Failure 400 {object} entity.ResponseError
// @Router /api/v1/auth [post]
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req entity.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	res, err := c.userUseCase.Login(ctx.Context(), &req)
	if err != nil {
		errorCode := fiber.StatusInternalServerError
		if code, ok := _const.ErrorCode[err.Error()]; ok {
			errorCode = code
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Code:    errorCode,
			Message: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}
