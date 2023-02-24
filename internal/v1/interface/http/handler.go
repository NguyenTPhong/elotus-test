package http

import (
	"elotus/internal/v1/interface/http/controller"
	"elotus/internal/v1/repository"
	"elotus/internal/v1/usecase"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	fiberApp    *fiber.App
	db          *gorm.DB
	redisClient *redis.Client

	// controller
	authController *controller.AuthController
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

func WithDatabase(db *gorm.DB) HandlerOption {
	return func(handler *Handler) {
		handler.db = db
	}
}

func WithRedisClient(redisClient *redis.Client) HandlerOption {
	return func(handler *Handler) {
		handler.redisClient = redisClient
	}
}

func (h *Handler) CreateController() *Handler {
	// init repositories
	authRepo := repository.NewAuthRepository(h.db)
	authCacheRepo := repository.NewAuthCacheRepository(h.redisClient)

	// init use case
	userUseCase := usecase.NewUserUseCase(authRepo, authCacheRepo)

	// init child controller here
	h.authController = controller.NewAuthController(userUseCase)

	return h
}

func (h *Handler) StartHandling() {

	// init all route
	h.initRoute()
}
