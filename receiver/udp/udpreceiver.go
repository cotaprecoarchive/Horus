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
	receiveStrategy receiver.ReceiveStrategy
}

const (
	PACKET_SIZE = 512
)

func NewUdpReceiver(
	host string,
	port int,
	receiveStrategy receiver.ReceiveStrategy,
) *UdpReceiver {
	return &UdpReceiver{
		host:            host,
		port:            port,
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

	// util.Invariant(
	// 	err == nil,
	// 	"...unexpected error: `%s` (ListenUDP)",
	// 	err,
	// )

	defer conn.Close()

	for {
		message := make([]byte, PACKET_SIZE)

		_, _, err := conn.ReadFromUDP(message)

		if err == nil && r.receiveStrategy.CanReceive(message) {
			r.NotifyAll(r.receiveStrategy.Receive(message))
		}
	}
}
