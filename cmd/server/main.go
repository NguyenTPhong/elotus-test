package main

import (
	"elotus/internal/server/interface/http"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"elotus/config"
	"elotus/global"
)

func main() {
	global.Init()
	defer global.DeInit()

	global.Logger.Info("start application", zap.Any("environment", config.Environment))

	app := fiber.New()
	app.Use(cors.New())
	app.Use(requestid.New())

	http.InitRestfulApi(app)

	app.Listen(fmt.Sprintf(":%v", config.Port))
}
