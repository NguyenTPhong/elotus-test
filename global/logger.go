package global

import (
	"elotus/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func initLogger() {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.Level(config.LogLevel)),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level",
		},
	}
	Logger = zap.Must(cfg.Build())
}

func deInitLogger() {
	Logger.Sync()
}
