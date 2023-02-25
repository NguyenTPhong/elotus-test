package http

import (
	"elotus/config"
	"elotus/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func (h *Handler) initRoute() {

	docs.SwaggerInfo.Title = "API Documentations"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Domain
	docs.SwaggerInfo.Schemes = []string{"http"}

	// documentation public
	h.fiberApp.Get("/swagger/*", swagger.HandlerDefault)

	api := h.fiberApp.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]interface{}{"message": "service is active"})
	})

	// authentication
	authGroup := v1.Group("/auth")
	authGroup.Post("/register", h.authController.RegisNewUser)
	authGroup.Post("/", h.authController.Login)

	// upload
	v1.Post("/upload", h.authMiddleware.Authentication, h.uploadController.UploadFile)
}
