package util_test

import (
	"os"
	"testing"

	"github.com/CotaPreco/Horus/util"
	"github.com/stretchr/testify/assert"
)

func EnvOrDefaultReturnsDefault(t *testing.T) {
	var val = util.EnvOrDefault("UNDEFINED", "default-value")

	assert.Equal(t, "default-value", val)
}

func EnvOrDefaultReturnsEnv(t *testing.T) {
	os.Setenv("DEFINED", "was I")

	var val = util.EnvOrDefault("DEFINED", "?")

	assert.Equal(t, "was I", val)
}
