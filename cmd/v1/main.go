package main

import (
	redis2 "elotus/package/redis"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"elotus/config"
	"elotus/global"
	"elotus/internal/v1/interface/http"
	"elotus/package/db"
)

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

	app := fiber.New()
	app.Use(cors.New())
	app.Use(requestid.New())

	http.NewHandler(
		http.WithEngine(app),
		http.WithDatabase(database),
		http.WithRedisClient(redisClient),
	).CreateController().StartHandling()

	app.Listen(fmt.Sprintf(":%v", config.Port))
}
