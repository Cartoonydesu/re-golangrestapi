package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func Success(c *gin.Context, st string, data any) {
	c.JSON(http.StatusOK, Response{
		Status: st,
		Data: data,
	})
}

func BadRequest(c *gin.Context, st string, msg string) {
	c.JSON(http.StatusBadRequest, Response{
		Status: st,
		Message: msg,
	})
}

func InternalServerErr(c *gin.Context, st string, msg string) {
	c.JSON(http.StatusInternalServerError, Response{
		Status: st,
		Message: msg,
	})
}