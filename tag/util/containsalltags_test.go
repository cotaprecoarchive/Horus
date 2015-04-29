package util_test

import (
	"testing"

	"github.com/CotaPreco/Horus/tag"
	"github.com/CotaPreco/Horus/tag/util"
	"github.com/stretchr/testify/assert"
)

func NewTag(tstr string) tag.Tag {
	var t, _ = tag.NewTag(tstr)
	return t
}

func TestContainsAllReturnsTrue(t *testing.T) {
	var a = []tag.Tag{
		NewTag("a"),
		NewTag("b"),
	}

	var b = []tag.Tag{
		NewTag("a"),
		NewTag("b"),
		NewTag("c"),
	}

	assert.True(t, util.ContainsAllTags(a, b))
}

func TestContainsAllReturnsFalse(t *testing.T) {
	var a = []tag.Tag{
		NewTag("a"),
		NewTag("b"),
	}

	var b = []tag.Tag{
		NewTag("z"),
		NewTag("a"),
	}

	assert.False(t, util.ContainsAllTags(a, b))
}
