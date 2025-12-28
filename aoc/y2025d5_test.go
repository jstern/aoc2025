package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2025d5Input = `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func Test_2025_Day_5_Part_1_Example(t *testing.T) {
	result := y2025d5part1(y2025d5Input)
	assert.Equal(t, "3", result)
}

func Test_2025_Day_5_addRange(t *testing.T) {
	db := &ingredientDB{[]freshRange{}}

	assert.Equal(t, []freshRange{}, db.ranges)

	db.addRange(10, 14)
	assert.Equal(t, []freshRange{{10, 14}}, db.ranges)

	db.addRange(3, 5)
	assert.Equal(t, []freshRange{{3, 5}, {10, 14}}, db.ranges)

	db.addRange(12, 18)
	assert.Equal(t, []freshRange{{3, 5}, {10, 14}, {12, 18}}, db.ranges)
}

func Test_2025_Day_5_include(t *testing.T) {
	assert.True(t, freshRange{0, 0}.includes(0))
	assert.True(t, freshRange{0, 1}.includes(1))
	assert.True(t, freshRange{1, 10}.includes(5))
}

func Test_2025_Day_5_size(t *testing.T) {
	assert.Equal(t, 1, freshRange{0, 0}.size())
	assert.Equal(t, 2, freshRange{0, 1}.size())
	assert.Equal(t, 3, freshRange{0, 2}.size())
}

func Test_2025_Day_5_Part_2_Example(t *testing.T) {
	result := y2025d5part2(y2025d5Input)
	assert.Equal(t, "14", result)
}
