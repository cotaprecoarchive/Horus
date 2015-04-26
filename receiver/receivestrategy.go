package receiver

import "github.com/CotaPreco/Horus/message"

type ReceiveStrategy interface {
	CanReceive(message []byte) bool
	Receive(message []byte) message.MessageInterface
}
