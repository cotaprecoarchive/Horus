package udp

import (
	"bytes"

	"github.com/CotaPreco/Horus/message"
	"github.com/CotaPreco/Horus/tag"
)

type NullByteReceiveStrategy struct {
}

func (s *NullByteReceiveStrategy) CanReceive(message []byte) bool {
	// shouldn't start with \0 or even ends with \0 exclusively
	if bytes.HasSuffix(message, []byte{0}) || bytes.HasPrefix(message, []byte{0}) {
		return false
	}

	message = s.trim(message)

	// empty? can't
	if len(message) == 0 {
		return false
	}

	return bytes.Count(message, []byte{0}) <= 1
}

func (s *NullByteReceiveStrategy) Receive(msg []byte) message.MessageInterface {
	msg = s.trim(msg)

	var positionOfDelimiter = bytes.Index(msg, []byte{0})

	if positionOfDelimiter != -1 {
		var tag, err = tag.NewTag(string(msg[0:positionOfDelimiter]))

		if err != nil {
			panic(err)
		}

		return message.NewTaggedMessage(tag, s.trim(msg[positionOfDelimiter:]))
	}

	return message.NewMessage(msg)
}

func (s *NullByteReceiveStrategy) trim(message []byte) []byte {
	return bytes.Trim(message, string([]byte{0}))
}
