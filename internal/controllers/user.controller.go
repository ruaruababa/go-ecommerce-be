package controllers

import (
	"go-ecommerce-be/pkg/response"

	"github.com/gin-gonic/gin"
)

const loginMessage = "Login"

const RegisterMessage = "Register"

type UserController struct{}

func InitUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Register(c *gin.Context) {
	response.SuccessResponse(c, response.ErrorCodeSuccess, "Register")
	response.SuccessResponse(c, response.ErrorCodeSuccess, RegisterMessage)
	response.SuccessResponse(c, response.ErrorCodeSuccess, loginMessage)
}

func (uc *UserController) Login(c *gin.Context) {
	response.SuccessResponse(c, response.ErrorCodeSuccess, "Login")
}

func (uc *UserController) Logout(c *gin.Context) {
	response.SuccessResponse(c, response.ErrorCodeSuccess, "Logout")
}
