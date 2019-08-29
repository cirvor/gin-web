package main

import (
	"web/middleware"
	"web/route"

	_ "web/task"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化引擎
	engine := gin.Default()
	// 捕获全局异常
	engine.Use(middleware.HandleErrors())
	// 路由
	route.Web(engine)

	// 绑定端口并启动应用
	engine.Run(":9205")
}
