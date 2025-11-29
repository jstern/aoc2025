package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2016d3Input = `  5  10  25
  5 5  5
  3   4  5
  25  5   10
`

func Test_2016_Day_3_Part_1_Example(t *testing.T) {
	result := y2016d3part1(y2016d3Input)
	assert.Equal(t, "2", result)
}

var y2016d3part2Input = `  5  3  5
  5  4  10
  5  5  25
`

func Test_2016_Day_3_Part_2_Example(t *testing.T) {
	result := y2016d3part2(y2016d3part2Input)
	assert.Equal(t, "2", result)
}
