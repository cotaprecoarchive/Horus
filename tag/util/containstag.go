package util

import "github.com/CotaPreco/Horus/tag"

func ContainsTag(tag tag.Tag, tags []tag.Tag) bool {
	for _, candidate := range tags {
		if candidate == tag {
			return true
		}
	}

	return false
}
