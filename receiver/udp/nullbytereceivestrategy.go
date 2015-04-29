package udp

import (
	"bytes"

	"github.com/CotaPreco/Horus/message"
	"github.com/CotaPreco/Horus/tag"
)

type NullByteReceiveStrategy struct {
}

func (s *NullByteReceiveStrategy) CanReceive(message []byte) bool {
	message = trim(message)

	if bytes.Count(message, []byte{0}) == 0 && len(message) > 0 {
		return true
	}

	return isValid(message) && extractTags(message) != nil
}

func (s *NullByteReceiveStrategy) Receive(msg []byte) message.MessageInterface {
	msg = trim(msg)

	var last = bytes.LastIndex(msg, []byte{0})

	if last == -1 {
		return message.NewMessage(trim(msg))
	}

	var tags = extractTags(msg)
	var msgb = trim(msg[last:])

	// > 0 == 1 tagged
	if len(tags) == 1 {
		return message.NewTaggedMessage(tags[0], msgb)
	}

	// ...sequence (full-match)
	return message.NewTagSequencedMessage(tags, msgb)
}

func extractTags(msg []byte) []tag.Tag {
	var i, l int
	var ttag []byte
	var tags []tag.Tag

	var last = bytes.LastIndex(msg, []byte{0})

	for i, l = 0, len(msg); i < l-1; i++ {
		ttag = append(ttag, msg[i])

		var hn = msg[i] == 0 && msg[i+1] == 0
		var la = msg[i] != 0 && msg[i+1] == 0 && (i+1) == last

		if hn || la {
			t, err := tag.NewTag(string(trim(ttag)))

			if err != nil {
				return nil
			}

			tags = append(tags, t)
			ttag = []byte{}
		}
	}

	return tags
}

func isValid(message []byte) bool {
	message = trim(message)

	if len(message) == 0 {
		return false
	}

	var count = bytes.Count(message, []byte{0})

	if count == 0 || count == 1 {
		return true
	}

	if count > 2 && (count&1) == 0 || count == 2 {
		return false
	}

	var i, l, lo, le int

	for i, l = 0, len(message); i < l-1; i++ {
		if message[i] == message[i+1] {
			le++
		} else {
			if le > lo {
				lo = le
			}

			le = 1
		}
	}

	if le > lo {
		lo = le
	}

	return lo < 3
}

func trim(message []byte) []byte {
	return bytes.Trim(message, string([]byte{0}))
}
