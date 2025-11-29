package aoc

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/jstern/aoc2025/aoc/junk"
	"github.com/samber/lo"
)

func init() {
	registerSolution("2016:3:1", y2016d3part1)
	registerSolution("2016:3:2", y2016d3part2)
}

func y2016d3part1(input string) string {
	total := 0
	for _, sides := range junk.IntSlices(input, strconv.Atoi) {
		total += validTriangle(sides)
	}
	return fmt.Sprint(total)
}

func validTriangle(sides []int) int {
	sort.Ints(sides)
	if sides[0]+sides[1] <= sides[2] {
		return 0
	}
	return 1
}

func y2016d3part2(input string) string {
	total := 0
	for _, rows := range lo.Chunk(junk.IntSlices(input, strconv.Atoi), 3) {
		for _, sides := range junk.Transpose(rows) {
			total += validTriangle(sides)
		}
	}
	return fmt.Sprint(total)
}
