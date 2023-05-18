package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	successJson struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	errorJson struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
	}
)

func ResponseSuccessJSON(c *gin.Context, message string, data interface{}) {

	if message == "" {
		message = "success"
	}

	res := successJson{
		Message: message,
		Success: true,
		Data:    data,
	}
	c.JSON(http.StatusOK, res)

}

func ResponseErrorJSON(c *gin.Context, code int, err error) {
	res := errorJson{
		Error: err.Error(),
	}
	c.JSON(code, res)

}

func ResponseValidationErrorJson(c *gin.Context, message string, detail interface{}) {
	res := errorJson{
		Message: message,
		Success: false,
		Error:   detail,
	}

	c.JSON(http.StatusBadRequest, res)
}
