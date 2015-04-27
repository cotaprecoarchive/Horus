package tag

import (
	"github.com/CotaPreco/Horus/command"
	"github.com/CotaPreco/Horus/tag"
	"github.com/gorilla/websocket"
)

type RTAGCommand struct {
	command.Command
	Connection *websocket.Conn
	Tags       []tag.Tag
}

func NewRTAGCommand(connection *websocket.Conn, tags []tag.Tag) *RTAGCommand {
	return &RTAGCommand{
		Connection: connection,
		Tags:       tags,
	}
}
