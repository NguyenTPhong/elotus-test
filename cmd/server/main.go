package main

import (
	"elotus/config"
	"elotus/global"
	"go.uber.org/zap"
)

func main() {
	global.Init()
	defer global.DeInit()

	global.Logger.Info("start application", zap.Any("environment", config.Environment))

}
