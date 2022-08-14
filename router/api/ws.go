package api

import (
	"net/http"
	"tconn/service/ws"
	"tconn/util/log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Error(err.Error())
		return
	}

	wsConn := ws.Conn{
		Conn: conn,
	}

	go wsConn.Read()
	go wsConn.Write()
}
