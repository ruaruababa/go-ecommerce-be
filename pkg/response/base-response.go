package response

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, BaseResponse{
		Message: msg[code],
		Status:  code,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message error) {
	c.JSON(code, BaseResponse{
		Message: msg[code],
		Status:  code,
		Data:    nil,
	})
}
