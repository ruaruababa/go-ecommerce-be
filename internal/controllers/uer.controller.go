package controllers

import (
	"go-ecommerce-be/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func InitUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Register(c *gin.Context) {
	response.SuccessResponse(c, response.ErrorCodeSuccess, "Register")
}

func (uc *UserController) Login(c *gin.Context) {
	response.SuccessResponse(c, response.ErrorCodeSuccess, "Login")
}
