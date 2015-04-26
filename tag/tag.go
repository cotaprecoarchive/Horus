package tag

import (
	"errors"
	"regexp"
)

type Tag interface {
	String() string
}

func NewTag(t string) (Tag, error) {
	if matched, _ := regexp.MatchString("(?i)^[a-zA-Z0-9_\\-\\*:]{1,255}$", t); !matched {
		return nil, errors.New("A tag should contains only `a-zA-Z0-9_-*:` 1..255")
	}

	return &tag{
		tag: t,
	}, nil
}

type tag struct {
	tag string
}

func (t *tag) String() string {
	return t.tag
}
