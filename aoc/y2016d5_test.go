package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2016d5Input = `abc
`

func Test_2016_Day_5_Part_1_Example(t *testing.T) {
	result := y2016d5part1(y2016d5Input)
	assert.Equal(t, "18f47a30", result)
}

func Test_2016_Day_5_Part_2_Example(t *testing.T) {
	result := y2016d5part2(y2016d5Input)
	assert.Equal(t, "05ace8e3", result)
}
