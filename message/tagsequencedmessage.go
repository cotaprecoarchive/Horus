package message

import "github.com/CotaPreco/Horus/tag"

type TagSequencedMessage struct {
	Message
	Tags []tag.Tag
}

func NewTagSequencedMessage(tags []tag.Tag, payload []byte) *TagSequencedMessage {
	return &TagSequencedMessage{
		Message{
			payload,
		},
		tags,
	}
}
