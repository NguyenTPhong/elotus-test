package global

import (
	"elotus/config"
	"elotus/package/logger"
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

func Init() {
	Logger = logger.NewLogger(config.LogLevel)
}

func DeInit() {
	logger.Close(Logger)
}
