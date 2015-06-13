package udp

import (
	"net"

	"github.com/CotaPreco/Horus/receiver"
	"github.com/CotaPreco/Horus/util"
)

type UdpReceiver struct {
	util.Observable
	host            string
	port            int
	maxPacketSize   int
	receiveStrategy receiver.ReceiveStrategy
}

func NewUdpReceiver(
	host string,
	port int,
	maxPacketSize int,
	receiveStrategy receiver.ReceiveStrategy,
) *UdpReceiver {
	return &UdpReceiver{
		host:            host,
		port:            port,
		maxPacketSize:   maxPacketSize,
		receiveStrategy: receiveStrategy,
	}
}

func (r *UdpReceiver) Receive() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(r.host),
		Port: r.port,
	})

	if err != nil {
		return
	}

	defer conn.Close()

	for {
		message := make([]byte, r.maxPacketSize)

		_, _, err := conn.ReadFromUDP(message)

		if err == nil && r.receiveStrategy.CanReceive(message) {
			r.NotifyAll(r.receiveStrategy.Receive(message))
		}
	}
}
