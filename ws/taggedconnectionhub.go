package ws

import (
	"sync"

	"github.com/CotaPreco/Horus/command"
	"github.com/CotaPreco/Horus/message"
	"github.com/CotaPreco/Horus/tag"
	tutil "github.com/CotaPreco/Horus/tag/util"
	wst "github.com/CotaPreco/Horus/ws/tag"
	"github.com/gorilla/websocket"
)

type TaggedConnectionHub struct {
	*sync.Mutex
	connections map[*websocket.Conn][]tag.Tag
}

func NewTaggedConnectionHub() *TaggedConnectionHub {
	return &TaggedConnectionHub{
		&sync.Mutex{},
		map[*websocket.Conn][]tag.Tag{},
	}
}

func (h *TaggedConnectionHub) Unsubscribe(connection *websocket.Conn) {
	h.Lock()
	defer h.Unlock()

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
			if tutil.ContainsTag(m.Tag, tags) {
				connection.WriteMessage(websocket.TextMessage, m.Payload)
			}
		}
		break

	case *message.TagSequencedMessage:
		var m = msg.(*message.TagSequencedMessage)

		for connection, tags := range h.connections {
			if tutil.ContainsAllTags(m.Tags, tags) {
				connection.WriteMessage(websocket.TextMessage, m.Payload)
			}
		}
		break
	}
}

func (h *TaggedConnectionHub) Update(args ...interface{}) {
	if args[0] == nil {
		return
	}

	h.Send(args[0].(message.MessageInterface))
}

func (h *TaggedConnectionHub) CanHandle(cmd command.Command) bool {
	switch cmd.(type) {
	case *wst.ATAGCommand:
		return true
	case *wst.RTAGCommand:
		return true
	}

	return false
}

func (h *TaggedConnectionHub) Handle(cmd command.Command) {
	h.Lock()
	defer h.Unlock()

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
		if !tutil.ContainsTag(candidate, h.connections[connection]) {
			add = append(add, candidate)
		}
	}

	return add
}

func (h *TaggedConnectionHub) hasConnection(candidate *websocket.Conn) bool {
	for connection, _ := range h.connections {
		if connection == candidate {
			return true
		}
	}

	return false
}
