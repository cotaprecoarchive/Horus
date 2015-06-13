package udp

import (
	"bytes"
	"net"
	"testing"
	"time"

	"github.com/CotaPreco/Horus/message"
	"github.com/stretchr/testify/assert"
)

const (
	RHOST = "0.0.0.0"
	RPORT = 10000
)

type ReceiveStrategy struct {
	Ch chan string
}

func (s *ReceiveStrategy) CanReceive(message []byte) bool {
	s.Ch <- "CanReceive:"
	return true
}

func (s *ReceiveStrategy) Receive(message []byte) message.MessageInterface {
	s.Ch <- " " + string(bytes.Trim(message, "\x00"))
	return nil
}

func TestReceive(t *testing.T) {
	ch := make(chan string)

	strategy := &ReceiveStrategy{
		Ch: ch,
	}

	receiver := NewUdpReceiver(RHOST, RPORT, 1, strategy)
	go receiver.Receive()

	client, _ := net.DialUDP(
		"udp",
		&net.UDPAddr{
			IP:   net.IPv4zero,
			Port: 0,
		},
		&net.UDPAddr{
			IP:   net.ParseIP(RHOST),
			Port: RPORT,
		},
	)

	time.Sleep(time.Millisecond + 100)

	client.Write([]byte("A"))
	client.Write([]byte("B"))
	client.Write([]byte("C"))
	client.Write([]byte("D"))

	client.Close()

	time.Sleep(time.Millisecond + 100)

	assert.Equal(t, <-ch+<-ch, "CanReceive: A")
	assert.Equal(t, <-ch+<-ch, "CanReceive: B")
	assert.Equal(t, <-ch+<-ch, "CanReceive: C")
	assert.Equal(t, <-ch+<-ch, "CanReceive: D")
}
