package util_test

import (
	"testing"

	"github.com/CotaPreco/Horus/tag"
	"github.com/CotaPreco/Horus/tag/util"
	"github.com/stretchr/testify/assert"
)

func TestContainsTagReturnsFalse(t *testing.T) {
	var tags []tag.Tag
	var tag, _ = tag.NewTag("user")

	assert.False(t, util.ContainsTag(tag, tags))
}

func TestContainsTagReturnsTrue(t *testing.T) {
	var user, _ = tag.NewTag("user")
	var guest, _ = tag.NewTag("guest")
	var admin, _ = tag.NewTag("admin")
	var mod, _ = tag.NewTag("mod")

	var tags = []tag.Tag{
		user,
		guest,
		admin,
	}

	assert.True(t, util.ContainsTag(user, tags))
	assert.True(t, util.ContainsTag(guest, tags))
	assert.True(t, util.ContainsTag(admin, tags))
	assert.False(t, util.ContainsTag(mod, tags))
}
