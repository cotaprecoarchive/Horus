package command

import (
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestNewSimpleTextCommand(t *testing.T) {
	var connection = &websocket.Conn{}
	var c = NewSimpleTextCommand("CMD", connection)

	assert.Equal(t, "CMD", c.String())
	assert.Equal(t, connection, c.GetFrom())
}
