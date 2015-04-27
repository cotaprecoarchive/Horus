package tag

import (
	"github.com/CotaPreco/Horus/command"
	"github.com/CotaPreco/Horus/tag"
	"github.com/gorilla/websocket"
)

type ATAGCommand struct {
	command.Command
	Connection *websocket.Conn
	Tags       []tag.Tag
}

func NewATAGCommand(connection *websocket.Conn, tags []tag.Tag) *ATAGCommand {
	return &ATAGCommand{
		Connection: connection,
		Tags:       tags,
	}
}
