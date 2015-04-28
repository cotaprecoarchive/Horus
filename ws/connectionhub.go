package ws

import (
	"github.com/CotaPreco/Horus/message"
	"github.com/gorilla/websocket"
)

type ConnectionHub interface {
	Unsubscribe(connection *websocket.Conn)
	Subscribe(connection *websocket.Conn)
	Send(message message.MessageInterface)
}
