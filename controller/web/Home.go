package web

/***
首页控制器
*/

import (
	"net/http"
	"web/database"

	"github.com/gin-gonic/gin"
)

/**
 */
func HomeIndex(context *gin.Context) {

	redis := database.Redis.Get()
	redis.Do("SET", "name", "HHHHHH")

	context.String(http.StatusOK, "hello, worldfff")
}

func HomeTest(context *gin.Context) {
	context.String(http.StatusOK, "Get Test")
}
