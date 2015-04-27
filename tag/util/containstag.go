package util

import "github.com/CotaPreco/Horus/tag"

func ContainsTag(tag tag.Tag, tags []tag.Tag) bool {
	for _, t := range tags {
		if t.String() == tag.String() {
			return true
		}
	}

	return false
}
