package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
	panic("done")
}

func Error(context *gin.Context, status int, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code": status,
		"msg":  msg,
	})
	panic("done")
}
