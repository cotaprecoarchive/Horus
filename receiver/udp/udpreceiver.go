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

	util.Invariant(
		err == nil,
		"...unexpected error: `%s` (ListenUDP)",
		err,
	)

	for {
		message := make([]byte, 1024)

		_, _, err := conn.ReadFromUDP(message)

		if err == nil && r.receiveStrategy.CanReceive(message) {
			r.NotifyAll(r.receiveStrategy.Receive(message))
		}
	}
}
