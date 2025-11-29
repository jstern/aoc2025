package junk_test

import (
	"strconv"
	"testing"

	"github.com/jstern/aoc2025/aoc/junk"
	"github.com/stretchr/testify/assert"
)

func TestIntSlices_Atoi(t *testing.T) {
	in := `5 6  6
7 8 9
`
	actual := junk.IntSlices(in, strconv.Atoi)
	expected := [][]int{
		{5, 6, 6},
		{7, 8, 9},
	}
	assert.Equal(t, expected, actual)
}

func TestIntSlices_Custom(t *testing.T) {
	in := `a1 b2 c3`

	actual := junk.IntSlices(in, junk.Hex)
	expected := [][]int{
		{161, 178, 195},
	}
	assert.Equal(t, expected, actual)
}

func TestTranspose(t *testing.T) {
	ints := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}

	assert.Equal(t, expected, junk.Transpose(ints))
}
