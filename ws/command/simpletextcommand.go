package command

import (
	c "github.com/CotaPreco/Horus/command"
	"github.com/gorilla/websocket"
)

type SimpleTextCommand struct {
	c.Command
	CommandStr string
	Connection *websocket.Conn
}

func NewSimpleTextCommand(command string, connection *websocket.Conn) *SimpleTextCommand {
	return &SimpleTextCommand{
		CommandStr: command,
		Connection: connection,
	}
}
