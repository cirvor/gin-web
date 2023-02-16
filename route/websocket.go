package route

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func Websocket(route *gin.Engine) {
	ws := melody.New()
	route.GET("/ws", func(c *gin.Context) {
		err := ws.HandleRequest(c.Writer, c.Request)
		if err != nil {
			return
		}
	})
	//处理连接
	ws.HandleConnect(func(session *melody.Session) {
		_, err := fmt.Fprint(os.Stdout, "connected")
		if err != nil {
			return
		}
	})

	//处理发送事件
	ws.HandleMessage(func(s *melody.Session, msg []byte) {
		err := ws.BroadcastOthers(msg, s)
		if err != nil {
			return
		}
		err = ws.Close()
		if err != nil {
			return
		}
	})

	//处理断开
	ws.HandleDisconnect(func(session *melody.Session) {
		_, err := fmt.Fprint(os.Stdout, "disconnected")
		if err != nil {
			return
		}
	})
}
