package junk_test

import (
	"testing"

	"github.com/jstern/aoc2025/aoc/junk"
	"github.com/stretchr/testify/assert"
)

func Test_ReverseString(t *testing.T) {
	assert.Equal(t, "evil", junk.ReverseString("live"))
	assert.Equal(t, "界世 ,olleH", junk.ReverseString("Hello, 世界"))
}
