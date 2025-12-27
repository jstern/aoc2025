package junk_test

import (
	"fmt"
	"testing"

	"github.com/jstern/aoc2025/aoc/junk"
	"github.com/stretchr/testify/assert"
)

func Test_AdjacentCells(t *testing.T) {
	grid := [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}

	tcs := []struct {
		row      int
		col      int
		expected []junk.Cell2D
	}{
		{0, 0, []junk.Cell2D{{0, 1}, {1, 0}, {1, 1}}},
		{0, 1, []junk.Cell2D{{0, 0}, {0, 2}, {1, 0}, {1, 1}, {1, 2}}},
		{0, 2, []junk.Cell2D{{0, 1}, {1, 1}, {1, 2}}},

		{1, 0, []junk.Cell2D{{0, 0}, {0, 1}, {1, 1}, {2, 1}, {2, 0}}},
		{1, 1, []junk.Cell2D{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
		{1, 2, []junk.Cell2D{{0, 2}, {0, 1}, {1, 1}, {2, 1}, {2, 2}}},

		{2, 0, []junk.Cell2D{{1, 0}, {1, 1}, {2, 1}}},
		{2, 1, []junk.Cell2D{{2, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 2}}},
		{2, 2, []junk.Cell2D{{1, 2}, {1, 1}, {2, 1}}},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("(%d,%d)", tc.row, tc.col), func(t *testing.T) {
			actual := junk.AdjacentCells(grid, tc.row, tc.col)
			t.Log(tc.expected)
			t.Log(actual)
			assert.ElementsMatch(t, tc.expected, actual)
		})

	}
}
