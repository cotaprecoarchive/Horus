package tag_test

import (
	"strings"
	"testing"

	"github.com/CotaPreco/Horus/tag"
	"github.com/stretchr/testify/assert"
)

var tags = map[string]bool{
	"#invalid": false,
	"#":        false,
	"":         false,
	strings.Repeat("invalid", 38): false,
	"valid":     true,
	":valid":    true,
	"**valid**": true,
}

func TestTagCastsToString(t *testing.T) {
	tag, err := tag.NewTag(":user")

	assert.Nil(t, err)
	assert.Equal(t, ":user", tag.String())
}

func TestTag(t *testing.T) {
	for tagAsString, valid := range tags {
		_, err := tag.NewTag(tagAsString)

		if valid {
			assert.Nil(t, err)
		}

		if !valid {
			assert.NotNil(t, err)
		}
	}
}
