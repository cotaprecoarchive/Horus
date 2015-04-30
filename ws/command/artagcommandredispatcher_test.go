package command

import (
	"testing"

	cmmd "github.com/CotaPreco/Horus/command"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CmdBus struct {
	mock.Mock
	cmmd.CommandBus
}

func (b *CmdBus) Dispatch(cmd cmmd.Command) {
	b.Called(cmd)
}

func (b *CmdBus) PushHandler(handler cmmd.CommandHandler) {
}

type Command struct {
	cmmd.Command
}

func TestCanHandle(t *testing.T) {
	var redispatcher = NewARTagCommandRedispatcher(&CmdBus{})

	assert.False(t, redispatcher.CanHandle(&Command{}))

	assert.False(t, redispatcher.CanHandle(
		NewSimpleTextCommand("CMD", &websocket.Conn{}),
	))

	assert.True(t, redispatcher.CanHandle(
		NewSimpleTextCommand("ATAG <tag>", &websocket.Conn{}),
	))

	assert.True(t, redispatcher.CanHandle(
		NewSimpleTextCommand("RTAG <tag>", &websocket.Conn{}),
	))
}

func TestHandle(t *testing.T) {
	var bus = &CmdBus{}
	bus.On("Dispatch", mock.Anything).Return(nil)

	var redispatcher = NewARTagCommandRedispatcher(bus)

	redispatcher.Handle(NewSimpleTextCommand("ATAG <tag>", &websocket.Conn{}))
	redispatcher.Handle(NewSimpleTextCommand("RTAG <tag>", &websocket.Conn{}))

	bus.AssertNumberOfCalls(t, "Dispatch", 2)
}
