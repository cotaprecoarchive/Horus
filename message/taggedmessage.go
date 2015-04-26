package message

import "github.com/CotaPreco/Horus/tag"

type TaggedMessage struct {
	Message
	Tag tag.Tag
}

func NewTaggedMessage(tag tag.Tag, payload []byte) *TaggedMessage {
	return &TaggedMessage{
		Message{payload},
		tag,
	}
}
