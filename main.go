package main

import (
	"web/route"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化引擎
	engine := gin.Default()
	route.Web(engine)

	// 绑定端口，然后启动应用
	engine.Run(":9205")
}
