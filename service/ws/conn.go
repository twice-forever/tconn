package ws

import (
	"encoding/json"
	"tconn/service/model"
	"tconn/service/protocol"
	"tconn/util/log"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	jsonMsg  []byte
	protoMsg []byte
)

const (
	loopTime = 10000000
)

const (
	position = iota
)

func Setup() {
	msg1 := &model.SendMsgJson{
		ID:   1,
		Type: position,
		Position: model.Position{
			X: 120.213213123,
			Y: 120.213213123,
			Z: 213123.213123123,
		},
	}
	jsonMsg, _ = json.Marshal(msg1)

	msg2 := &protocol.Msg{
		Id:   1,
		Type: position,
		Position: &protocol.Position{
			X: 120.213213123,
			Y: 120.213213123,
			Z: 213123.213123123,
		},
	}
	protoMsg, _ = proto.Marshal(msg2)

	d1 := &model.SendMsgJson{}
	d2 := &protocol.Msg{}

	start := time.Now()
	for i := 0; i <= loopTime; i++ {
		_ = json.Unmarshal(jsonMsg, d1)
	}
	log.Info(loopTime, " json: ", time.Since(start))

	start = time.Now()
	for i := 0; i <= loopTime; i++ {
		_ = proto.Unmarshal(protoMsg, d2)
	}
	log.Info(loopTime, " proto: ", time.Since(start))
}

// Conn 连接结构体
type Conn struct {
	Conn *websocket.Conn
}

// Read 读取Websocket发送数据
func (c *Conn) Read() {
	defer c.Conn.Close()

	for {
		msgType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Error(err.Error())
			return
		}

		switch msgType {
		case websocket.TextMessage:
			log.Info(string(msg))
			start := time.Now()
			d := &model.SendMsgJson{}
			_ = json.Unmarshal(msg, d)
			log.Info(*d)
			log.Info(time.Since(start))
			log.Info("json: ", time.Since(start))
		case websocket.BinaryMessage:
			log.Info(msg)
			start := time.Now()
			d := &protocol.Msg{}
			_ = proto.Unmarshal(msg, d)
			log.Info(*d)
			log.Info("proto: ", time.Since(start))
		}
	}
}

func (c *Conn) Write() {
	defer c.Conn.Close()

	ticker := time.NewTicker(1 * time.Second)
	for {
		<-ticker.C

		c.Conn.WriteMessage(websocket.TextMessage, jsonMsg)

		<-ticker.C
		c.Conn.WriteMessage(websocket.BinaryMessage, protoMsg)
	}
}
