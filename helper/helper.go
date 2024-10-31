package helper

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func RespondJSON(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}

func RespondError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
	})
}
