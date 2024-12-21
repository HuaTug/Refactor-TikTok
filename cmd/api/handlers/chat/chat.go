package chat

import (
	"context"
	"encoding/json"
	"fmt"

	jwt "HuaTug.com/pkg"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/websocket"
)

type Message struct {
	Text string `json:"text"`
}

var upgrader = websocket.HertzUpgrader{
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true // 允许所有来源连接
	},
}

var (
	BadConnection = []byte(`bad connection`)
)

func Handler(ctx context.Context, c *app.RequestContext) {
	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		uid, err := jwt.ConvertJWTPayloadToString(ctx, c)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, BadConnection)
			return
		}
		conn.WriteMessage(websocket.TextMessage, []byte(`Welcome, `+fmt.Sprint(uid)))

		s := NewChatService(ctx, c, conn)

		if err := s.Login(); err != nil {
			conn.WriteMessage(websocket.TextMessage, BadConnection)
			return
		}
		defer s.Logout()

		if err := s.ReadOfflineMessage(); err != nil {
			hlog.Info(err)
			conn.WriteMessage(websocket.TextMessage, BadConnection)
			return
		}
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Error reading message"))
				return
			}
			hlog.Info("Received message:", string(message))

			var msg Message
			// Unmarshal the received message into the Message struct
			if err := json.Unmarshal(message, &msg); err != nil {
				// Send an error message back if unmarshalling fails
				conn.WriteMessage(websocket.TextMessage, []byte("Error unmarshalling message"))
				return
			}

			// Call SendMessage with the unmarshalled message (msg), not the raw message
			if err := s.SendMessage(msg); err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Error sending message"))
				return
			}
		}
	})
	if err != nil {
		c.JSON(consts.StatusOK, `error`)
		return
	}
}
