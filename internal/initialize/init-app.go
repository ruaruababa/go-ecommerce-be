package initialize

import (
	"fmt"

	"go-ecommerce-be/globals"

	"go.uber.org/zap"
)

func InitApp() {
	// router
	r := InitRouter()
	// config
	InitLoadConfig()
	// logger
	InitLogger()
	globals.Logger.Info("Starting application...", zap.Int("port", globals.Config.Server.Port))
	// mysql
	InitMysql()
	// redis
	InitRedis()

	r.Run(fmt.Sprintf(":%d", globals.Config.Server.Port))
}
