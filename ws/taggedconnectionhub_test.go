package ws_test

import (
	"reflect"
	"testing"

	"github.com/CotaPreco/Horus/ws"
	wst "github.com/CotaPreco/Horus/ws/tag"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestSubscribe(t *testing.T) {
	hub := ws.NewTaggedConnectionHub()
	hub.Subscribe(&websocket.Conn{})
	hub.Subscribe(&websocket.Conn{})
	hub.Subscribe(&websocket.Conn{})

	assert.Equal(
		t,
		3,
		reflect.Indirect(reflect.ValueOf(hub)).FieldByName("connections").Len(),
	)
}

func TestUnsubscribe(t *testing.T) {
	var a = &websocket.Conn{}
	var b = &websocket.Conn{}
	var c = &websocket.Conn{}

	hub := ws.NewTaggedConnectionHub()
	hub.Subscribe(a)
	hub.Subscribe(b)
	hub.Subscribe(c)

	hub.Unsubscribe(a)
	hub.Unsubscribe(c)

	assert.Equal(
		t,
		1,
		reflect.Indirect(reflect.ValueOf(hub)).FieldByName("connections").Len(),
	)
}

func TestCanHandle(t *testing.T) {
	hub := ws.NewTaggedConnectionHub()

	assert.True(t, hub.CanHandle(&wst.ATAGCommand{}))
	assert.True(t, hub.CanHandle(&wst.RTAGCommand{}))

	var command struct{}
	assert.False(t, hub.CanHandle(&command))
}

// func TestHandle(t *testing.T) {
// 	var tA, _ = tag.NewTag("A")
// 	var tB, _ = tag.NewTag("B")
// 	var tC, _ = tag.NewTag("C")
//
// 	hub := ws.NewTaggedConnectionHub()
// 	conn := &websocket.Conn{}
//
// 	hub.Subscribe(conn)
//
// 	hub.Handle(wst.NewATAGCommand(conn, []tag.Tag{
// 		tA,
// 		tB,
// 		tC,
// 	}))
// }
