package tag

import (
	"errors"
	"regexp"
)

type Tag interface {
	String() string
}

func NewTag(t string) (Tag, error) {
	// RFC 3986 - Section 3.4
	// @link http://tools.ietf.org/html/rfc3986#page-23
	if matched, _ := regexp.MatchString("(?i)^[a-zA-Z0-9-\\._~!\\$&\\'\\(\\)\\*\\+\\,;=\\:@\\/\\?]{1,255}$", t); !matched {
		return nil, errors.New("A tag should contains only `A-Z a-z 0-9 - . _ ~ ! $ & ' ( ) * + , ; = : @ / ?` 1..255")
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
