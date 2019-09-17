package route

import (
	"net/http"
	"web/middleware"

	"web/controller/web"

	"github.com/gin-gonic/gin"
)

func Web(route *gin.Engine) {

	// 注册一个路由和处理函数
	route.GET("/", middleware.Middleware1, web.HomeIndex)
	route.GET("/get", web.HomeTest)
	user := route.Group("/user", middleware.Middleware1)
	{
		user.GET("/", web.UserIndex)
		user.GET("/name/:name", web.UserName)
		user.POST("/update", web.UserUpdate)
		user.GET("/welcome", func(context *gin.Context) {
			firstName := context.DefaultQuery("first_name", "Guest")
			lastName := context.Query("last_name")
			context.String(http.StatusOK, "Hello %s %s", firstName, lastName)
		})
	}
}
