package junk_test

import (
	"testing"

	"github.com/jstern/aoc2025/aoc/junk"
	"github.com/stretchr/testify/assert"
)

func Test_Factors(t *testing.T) {
	assert.Equal(t, []int{2, 3}, junk.Factors(6))
	assert.Equal(t, []int{3}, junk.Factors(9))
	assert.Equal(t, []int{}, junk.Factors(0))
	assert.Equal(t, []int{}, junk.Factors(1))
	assert.Equal(t, []int{}, junk.Factors(2))

}
