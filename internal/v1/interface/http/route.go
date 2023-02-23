package http

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initRoute() {
	api := h.fiberApp.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]interface{}{"message": "service is active"})
	})

	authGroup := v1.Group("/auth")
	authGroup.Post("/register", h.authController.RegisNewUser)
	authGroup.Post("/", h.authController.Login)
}
