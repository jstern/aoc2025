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

func Test_Chunks(t *testing.T) {
	assert.Equal(t, junk.Chunks("abcdef", 2), []string{"ab", "cd", "ef"})
	assert.Equal(t, junk.Chunks("abcdef", 3), []string{"abc", "def"})
	assert.Equal(t, junk.Chunks("abcdef", 1), []string{"a", "b", "c", "d", "e", "f"})
}
