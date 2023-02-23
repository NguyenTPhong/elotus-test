package http

import (
	"github.com/gofiber/fiber/v2"
)

func InitRestfulApi(app *fiber.App) {
	// init controller
	handler := NewHandler(
		WithEngine(app),
	)

	// create route
	handler.InitV1Route()
}

type Handler struct {
	fiberApp *fiber.App
}

type HandlerOption func(*Handler)

func NewHandler(options ...HandlerOption) *Handler {
	handler := &Handler{}
	for _, option := range options {
		option(handler)
	}
	return handler
}

func WithEngine(r *fiber.App) HandlerOption {
	return func(handler *Handler) {
		handler.fiberApp = r
	}
}
