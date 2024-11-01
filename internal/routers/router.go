package routers

import (
	"go-ecommerce-be/internal/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/register", controllers.InitUserController().Register)
	}
	return r
}
