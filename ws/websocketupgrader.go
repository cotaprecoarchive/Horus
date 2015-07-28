package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var WebSocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// @link https://godoc.org/github.com/gorilla/websocket#hdr-Origin_Considerations
		return true
	},
}
