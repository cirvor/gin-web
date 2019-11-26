package route

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func Websocket(route *gin.Engine) {
	ws := melody.New()
	route.GET("/ws", func(c *gin.Context) {
		ws.HandleRequest(c.Writer, c.Request)
	})
	//处理连接
	ws.HandleConnect(func(session *melody.Session) {
		fmt.Fprint(os.Stdout, "connected")
	})

	//处理发送事件
	ws.HandleMessage(func(s *melody.Session, msg []byte) {
		ws.BroadcastOthers(msg, s)
		ws.Close()
	})

	//处理断开
	ws.HandleDisconnect(func(session *melody.Session) {
		fmt.Fprint(os.Stdout, "disconnected")
	})
}
