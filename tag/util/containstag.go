package util

import "github.com/CotaPreco/Horus/tag"

func ContainsTag(tag tag.Tag, tags []tag.Tag) bool {
	for _, candidate := range tags {
		if candidate.String() == tag.String() {
			return true
		}
	}

	return false
}
