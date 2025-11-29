package junk

import (
	"strconv"
	"strings"

	"github.com/samber/lo"
)

// IntSlices carves up an input into a slice of slices of ints,
// assuming values on each line are separated by sep.
func IntSlices(input string, conv func(string) (int, error)) [][]int {
	res := make([][]int, 0)

	if conv == nil {
		conv = strconv.Atoi
	}

	loconv := func(s string, _ int) int {
		i, err := conv(s)
		if err != nil {
			panic(err)
		}
		return i
	}

	for line := range strings.SplitSeq(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		res = append(res, lo.Map(strings.Fields(line), loconv))
	}

	return res
}

// Transpose transposes a matrix.
func Transpose[T any](input [][]T) [][]T {
	rows := len(input)
	cols := len(input[0])

	res := make([][]T, cols)
	for i := range res {
		res[i] = make([]T, rows)
	}

	for i := range input {
		for j := range cols {
			res[j][i] = input[i][j]
		}
	}

	return res
}
