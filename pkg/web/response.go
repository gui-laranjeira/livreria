package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	Response(c, status, data)
}

func Error(c *gin.Context, status int, formatString string, args ...interface{}) {
	Response(c, status, ErrorResponse{
		Status:  status,
		Code:    http.StatusText(status),
		Message: fmt.Sprintf(formatString, args...),
	})
}
