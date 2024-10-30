package controllers

import (
	"go-ecommerce-be/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {}


func InitUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Register(c *gin.Context) {
	c.JSON(200, response.BaseResponse{
		Message: "Hello World",
		Status: 200,
		Data: []string{"Hello", "World"},
	})
}