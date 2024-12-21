package websocket

import (
	handler_ws_chat "HuaTug.com/cmd/api/handlers/chat"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func register(h *server.Hertz) {
	h.GET(`/`, append(_homeMW(), handler_ws_chat.Handler)...)
}
