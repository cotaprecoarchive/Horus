package udp_test

import (
	"testing"

	"github.com/CotaPreco/Horus/message"
	"github.com/CotaPreco/Horus/receiver/udp"
	"github.com/stretchr/testify/assert"
)

func TestCanReceive(t *testing.T) {
	var strategy = new(udp.NullByteReceiveStrategy)

	assert.False(t, strategy.CanReceive([]byte("not\x00ok\x00sorry")))
	assert.False(t, strategy.CanReceive([]byte("")))

	assert.True(t, strategy.CanReceive([]byte("\x00ok")))
	assert.True(t, strategy.CanReceive([]byte("ok\x00")))
	assert.True(t, strategy.CanReceive([]byte("payload")))
	assert.True(t, strategy.CanReceive([]byte("tag\x00payload")))

	// ...can't receive if invalid tag was present (refs: gh:issue #8)
	assert.False(t, strategy.CanReceive([]byte("#invalid\x00payload")))
	assert.False(t, strategy.CanReceive([]byte("$invalid\x00payload")))
}

func TestReceive(t *testing.T) {
	var strategy = new(udp.NullByteReceiveStrategy)

	assert.IsType(t, new(message.Message), strategy.Receive([]byte("message")))
	assert.IsType(t, new(message.TaggedMessage), strategy.Receive([]byte("tag\x00message")))

	var message = strategy.Receive([]byte("tag\x00message")).(*message.TaggedMessage)

	assert.Equal(t, "tag", message.Tag.String())
	assert.Equal(t, "message", string(message.Payload))
}
