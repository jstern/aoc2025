package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2025d4Input = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func Test_2025_Day_4_Part_1_Example(t *testing.T) {
	result := y2025d4part1(y2025d4Input)
	assert.Equal(t, "13", result)
}

func Test_2025_Day_4_Part_2_Example(t *testing.T) {
	result := y2025d4part2(y2025d4Input)
	assert.Equal(t, "43", result)
}
