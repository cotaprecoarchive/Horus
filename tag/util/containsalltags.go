package util

import "github.com/CotaPreco/Horus/tag"

func ContainsAllTags(a []tag.Tag, b []tag.Tag) bool {
	for _, at := range a {
		if !ContainsTag(at, b) {
			return false
		}
	}

	return true
}
