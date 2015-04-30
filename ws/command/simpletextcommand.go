package command

import (
	c "github.com/CotaPreco/Horus/command"
	"github.com/gorilla/websocket"
)

type SimpleTextCommand struct {
	c.Command
	cmd        string
	connection *websocket.Conn
}

func NewSimpleTextCommand(command string, connection *websocket.Conn) *SimpleTextCommand {
	return &SimpleTextCommand{
		cmd:        command,
		connection: connection,
	}
}

func (c *SimpleTextCommand) String() string {
	return c.cmd
}

func (c *SimpleTextCommand) GetFrom() *websocket.Conn {
	return c.connection
}
