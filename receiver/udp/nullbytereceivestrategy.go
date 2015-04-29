package udp

import (
	"bytes"

	"github.com/CotaPreco/Horus/message"
	"github.com/CotaPreco/Horus/tag"
)

type NullByteReceiveStrategy struct {
}

func (s *NullByteReceiveStrategy) CanReceive(message []byte) bool {
	message = s.trim(message)

	if len(message) == 0 {
		return false
	}

	var countOfNullByte = bytes.Count(message, []byte{0})

	// ...tagged
	if countOfNullByte == 1 {
		var tag = s.receiveTag(message, bytes.Index(message, []byte{0}))

		return tag != nil
	}

	return countOfNullByte <= 1
}

func (s *NullByteReceiveStrategy) Receive(msg []byte) message.MessageInterface {
	msg = s.trim(msg)

	var posOfDelim = bytes.Index(msg, []byte{0})

	if posOfDelim != -1 {
		var tag = s.receiveTag(msg, posOfDelim)

		return message.NewTaggedMessage(tag, s.trim(msg[posOfDelim:]))
	}

	return message.NewMessage(msg)
}

func (s *NullByteReceiveStrategy) receiveTag(msg []byte, posOfDelim int) tag.Tag {
	var tag, err = tag.NewTag(string(msg[0:posOfDelim]))

	if err != nil {
		return nil
	}

	return tag
}

func (s *NullByteReceiveStrategy) trim(message []byte) []byte {
	return bytes.Trim(message, string([]byte{0}))
}
