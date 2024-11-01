package initialize

import (
	"go-ecommerce-be/internal/controllers"
	"go-ecommerce-be/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthMiddleware())

	v1 := r.Group("/v1")
	{
		v1.GET("/register", controllers.InitUserController().Register)
	}
	return r
}
