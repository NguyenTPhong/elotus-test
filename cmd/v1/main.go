package main

import (
	"elotus/database/migration"
	redis2 "elotus/package/redis"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"elotus/config"
	"elotus/global"
	"elotus/internal/v1/delivery/http"
	"elotus/package/db"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	global.Init()
	defer global.DeInit()

	database, err := db.NewDatabase(config.DbConnStr, int(config.DbMaxConn), int(config.DbMaxIdleConn), int(config.DBLogLevel))
	if err != nil {
		panic(err)
	}

	redisClient, err := redis2.NewClient(config.RedisHost, config.RedisPassword)
	if err != nil {
		panic(err)
	}

	// migrate database
	migration.CreateTable(database)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Accept-Language, Content-Length,Authorization",
	}))

	http.NewHandler(
		http.WithEngine(app),
		http.WithDatabase(database),
		http.WithRedisClient(redisClient),
	).CreateController().StartHandling()

	app.Listen(fmt.Sprintf(":%v", config.Port))
}
