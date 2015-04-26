package tag_test

import (
	"strings"
	"testing"

	"github.com/CotaPreco/Horus/tag"
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

func TestTag(t *testing.T) {
	for tstr, valid := range tags {
		tag, err := tag.NewTag(tstr)

		if (err != nil && valid) || (err == nil && !valid) {
			t.Error()
		}

		if err == nil {
			if tag.String() != tstr {
				t.Error()
			}
		}
	}
}
