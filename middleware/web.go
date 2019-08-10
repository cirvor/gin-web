package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Middleware1(context *gin.Context) {
	log.Println("exec middleware1111111111")

	//你可以写一些逻辑代码

	// 执行该中间件之后的逻辑
	context.Next()

	//等 handler执行完成后，跳回执行
	log.Println("test")
}
