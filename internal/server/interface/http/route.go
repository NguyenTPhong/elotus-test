package http

import (
	"elotus/internal/server/entity"
	"elotus/internal/server/interface/middleware"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) InitV1Route() {
	api := h.fiberApp.Group("/api")
	v1 := api.Group("/v1", middleware.Authentication)

	v1.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(entity.Response{Data: map[string]interface{}{"message": "service is active"}})
	})
}
