package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr2IntReturnsZero(t *testing.T) {
	assert.Equal(t, Str2int("ABC"), 0)
}

func TestStr2IntReturnsInt(t *testing.T) {
	assert.Equal(t, Str2int("1234"), 1234)
	assert.Equal(t, Str2int("999999"), 999999)
}
