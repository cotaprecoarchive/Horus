package util

import (
	"os"
)

func EnvOrDefault(key, def string) string {
	val := os.Getenv(key)

	if val == "" {
		return def
	}

	return val
}
