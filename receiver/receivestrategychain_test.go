package receiver_test

import (
	"testing"

	"github.com/CotaPreco/Horus/message"
	"github.com/CotaPreco/Horus/receiver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ReceiveStrategy struct {
	mock.Mock
}

func (m *ReceiveStrategy) CanReceive(msg []byte) bool {
	var args = m.Called(msg)
	return args.Bool(0)
}

func (m *ReceiveStrategy) Receive(msg []byte) message.MessageInterface {
	m.Called(msg)
	return nil
}

func TestCanReceiveReturnsFalse(t *testing.T) {
	var chain = receiver.NewReceiveStrategyChain([]receiver.ReceiveStrategy{})

	assert.False(t, chain.CanReceive([]byte{}))
}

func TestCanReceiveReturnsTrue(t *testing.T) {
	var one = new(ReceiveStrategy)
	one.On("CanReceive", mock.Anything).Return(false)

	var two = new(ReceiveStrategy)
	two.On("CanReceive", mock.Anything).Return(true)

	var chain = receiver.NewReceiveStrategyChain(
		[]receiver.ReceiveStrategy{
			one,
			two,
		},
	)

	assert.True(t, chain.CanReceive([]byte{}))
}

func TestReceive(t *testing.T) {
	var one = new(ReceiveStrategy)
	one.On("CanReceive", mock.Anything).Return(true)
	one.On("Receive", mock.Anything).Return(nil)

	var two = new(ReceiveStrategy)
	two.On("CanReceive", mock.Anything).Return(false)
	two.On("Receive", mock.Anything).Return(nil)

	var chain = receiver.NewReceiveStrategyChain(
		[]receiver.ReceiveStrategy{
			one,
			two,
		},
	)

	assert.True(t, chain.CanReceive([]byte{}))
	assert.Nil(t, chain.Receive([]byte{}))

	one.AssertNumberOfCalls(t, "CanReceive", 2)
	one.AssertNumberOfCalls(t, "Receive", 1)

	two.AssertNotCalled(t, "CanReceive")
	two.AssertNotCalled(t, "Receive")
}
