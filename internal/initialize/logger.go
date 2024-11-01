package initialize

import (
	"go-ecommerce-be/globals"
	"go-ecommerce-be/pkg/logger"
)

func InitLogger() {
	// logger
	globals.Logger = logger.NewLogger(globals.Config.Logger)
}
