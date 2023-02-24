package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(logLevel int64) *zap.Logger {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.Level(logLevel)),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level",
		},
	}
	return zap.Must(cfg.Build())
}

func Close(logger *zap.Logger) {
	logger.Sync()
}

func TraceID(ctx context.Context) zap.Field {
	return zap.Any("trace_id", ctx.Value("requestid"))
}
