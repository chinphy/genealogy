package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// WsConnection websocket连接
func WsConnection(c *gin.Context) {
	var (
		wsConn *websocket.Conn
		err    error
	)
	if wsConn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	defer wsConn.Close()

	for {
		//读取ws中的数据
		if _, _, err := wsConn.ReadMessage(); err != nil {
			break
		}
		//写入ws数据
		if err = wsConn.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
			break
		}
	}
}
