package globals

import (
	"go-ecommerce-be/pkg/logger"
	"go-ecommerce-be/pkg/settings"
)

/*
 * Config is a global variable that holds the configuration settings for the application.
 * It is initialized in the main function using the settings package.
 * This allows the configuration settings to be accessed from any part of the application.
 * This is useful for accessing configuration settings in different parts of the application without having to pass them around as parameters.
 */

var (
	Config settings.Config
	Logger *logger.LoggerZap
)
