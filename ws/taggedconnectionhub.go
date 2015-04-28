package ws

import (
	"github.com/CotaPreco/Horus/command"
	"github.com/CotaPreco/Horus/message"
	"github.com/CotaPreco/Horus/tag"
	tagUtil "github.com/CotaPreco/Horus/tag/util"
	"github.com/CotaPreco/Horus/util"
	wst "github.com/CotaPreco/Horus/ws/tag"
	"github.com/gorilla/websocket"
)

type TaggedConnectionHub struct {
	util.Observer
	command.CommandHandler
	connections map[*websocket.Conn][]tag.Tag
}

func NewTaggedConnectionHub() *TaggedConnectionHub {
	return &TaggedConnectionHub{
		connections: make(map[*websocket.Conn][]tag.Tag),
	}
}

func (h *TaggedConnectionHub) Unsubscribe(connection *websocket.Conn) {
	if h.hasConnection(connection) {
		delete(h.connections, connection)
	}
}

func (h *TaggedConnectionHub) Subscribe(connection *websocket.Conn) {
	h.connections[connection] = make([]tag.Tag, 0)
}

func (h *TaggedConnectionHub) Send(msg message.MessageInterface) {
	switch msg.(type) {
	case *message.Message:
		var m = msg.(*message.Message)

		for connection, _ := range h.connections {
			connection.WriteMessage(websocket.TextMessage, m.Payload)
		}
		break
	case *message.TaggedMessage:
		var m = msg.(*message.TaggedMessage)

		for connection, tags := range h.connections {
			for _, tag := range tags {
				if tag.String() == m.Tag.String() {
					connection.WriteMessage(websocket.TextMessage, m.Payload)
				}
			}
		}
		break
	}
}

// `util.Observer`
func (h *TaggedConnectionHub) Update(args ...interface{}) {
	h.Send(args[0].(message.MessageInterface))
}

// `command.CommandHandler`
func (h *TaggedConnectionHub) CanHandle(cmd command.Command) bool {
	switch cmd.(type) {
	case *wst.ATAGCommand:
		return true
	case *wst.RTAGCommand:
		return true
	}

	return false
}

// `command.CommandHandler`
func (h *TaggedConnectionHub) Handle(cmd command.Command) {
	switch cmd.(type) {
	case *wst.ATAGCommand:
		var c = cmd.(*wst.ATAGCommand)

		for _, tag := range h.collectTagsToAdd(c.Connection, c.Tags) {
			h.connections[c.Connection] = append(
				h.connections[c.Connection],
				tag,
			)
		}
		break
	case *wst.RTAGCommand:
		var c = cmd.(*wst.RTAGCommand)

		if !h.hasConnection(c.Connection) {
			return
		}

		// > x y z a b c
		// RTAG x b c
		// > y z a

		var retain []tag.Tag

		for _, remove := range c.Tags {
			for _, tag := range h.connections[c.Connection] {
				if tag.String() != remove.String() {
					retain = append(retain, tag)
				}
			}
		}

		h.connections[c.Connection] = retain
		break
	}
}

func (h *TaggedConnectionHub) collectTagsToAdd(
	connection *websocket.Conn,
	tags []tag.Tag,
) []tag.Tag {
	var add []tag.Tag

	if !h.hasConnection(connection) {
		return add
	}

	for _, candidate := range tags {
		if !tagUtil.ContainsTag(candidate, h.connections[connection]) {
			add = append(add, candidate)
		}
	}

	return add
}

func (h *TaggedConnectionHub) hasConnection(connection *websocket.Conn) bool {
	for conn, _ := range h.connections {
		if conn == connection {
			return true
		}
	}

	return false
}