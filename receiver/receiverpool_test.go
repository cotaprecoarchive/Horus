package receiver_test

import (
	"testing"

	"github.com/CotaPreco/Horus/receiver"
	"github.com/stretchr/testify/assert"
)

type R struct {
	Ch chan int
}

func (r *R) Receive() {
	r.Ch <- 1
}

func NewReceiver() *R {
	return &R{}
}

func TestNewReceiverPool(t *testing.T) {
	var pool = receiver.NewReceiverPool(make([]receiver.Receiver, 0))

	assert.Equal(t, 0, pool.Size())

	pool.AddReceiver(NewReceiver())
	pool.AddReceiver(NewReceiver())
	pool.AddReceiver(NewReceiver())
	pool.AddReceiver(NewReceiver())

	assert.Equal(t, 4, pool.Size())
}

func TestPoolReceive(t *testing.T) {
	var calls, i int
	ch := make(chan int)

	var a = &R{ch}
	var b = &R{ch}
	var c = &R{ch}

	var pool = receiver.NewReceiverPool([]receiver.Receiver{
		a,
		b,
		c,
	})

	assert.Equal(t, 3, pool.Size())

	pool.Receive()

	for i := i; i < pool.Size(); i++ {
		calls = calls + <-ch
	}

	assert.Equal(t, 3, calls)
}
