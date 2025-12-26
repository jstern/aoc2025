package aoc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2025d1Input = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func Test_2025_Day_1_Part_1_Example(t *testing.T) {
	result := y2025d1part1(y2025d1Input)
	assert.Equal(t, "3", result)
}

func Test_2025_Day_1_Part_2_Example(t *testing.T) {
	result := y2025d1part2(y2025d1Input)
	assert.Equal(t, "6", result)
}

func Test_2024_Day_1_dial_toZero(t *testing.T) {
	tcs := []struct {
		dial      *dial
		dist      int
		end       int
		clicks    int
		remaining int
	}{
		{dial: &dial{size: 12, position: 3}, dist: 3, end: 6, clicks: 0, remaining: 0},
		{dial: &dial{size: 12, position: 3}, dist: 9, end: 0, clicks: 1, remaining: 0},

		{dial: &dial{size: 12, position: 3}, dist: 10, end: 0, clicks: 1, remaining: 1},
		{dial: &dial{size: 12, position: 0, clicks: 1}, dist: 1, end: 1, clicks: 1, remaining: 0},

		{dial: &dial{size: 12, position: 6}, dist: -3, end: 3, clicks: 0, remaining: 0},
		{dial: &dial{size: 12, position: 6}, dist: -6, end: 0, clicks: 1, remaining: 0},

		{dial: &dial{size: 12, position: 6}, dist: -7, end: 0, clicks: 1, remaining: -1},
		{dial: &dial{size: 12, position: 0, clicks: 1}, dist: -1, end: 11, clicks: 1, remaining: 0},

		{dial: &dial{size: 100, position: 0, clicks: 59}, dist: -350, end: 0, clicks: 60, remaining: -250},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%d + %d", tc.dial, tc.dist), func(t *testing.T) {
			remaining := tc.dial.toZero(tc.dist)
			assert.Equal(t, tc.end, tc.dial.position)
			assert.Equal(t, tc.clicks, tc.dial.clicks)
			assert.Equal(t, tc.remaining, remaining)
		})
	}
}
