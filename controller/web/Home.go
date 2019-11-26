package web

/***
首页控制器
*/
import (
	"net/http"
	"web/database"
	"web/utils"

	"github.com/gin-gonic/gin"
)

/**
 */
func HomeIndex(context *gin.Context) {

	redis := database.Redis.Get()
	defer redis.Close()
	_, err := redis.Do("SET", "name", "HHHHHH")
	if err != nil {
		utils.Error(context, http.StatusBadGateway, "服务器异常")
	}

	utils.Success(context, "hello, world")
}

func HomeTest(context *gin.Context) {
	utils.Success(context, "Get Test")
}
