package aoc

import (
	"fmt"
	"strings"

	"github.com/jstern/aoc2025/aoc/junk"
)

func init() {
	registerSolution("2025:4:1", y2025d4part1)
	registerSolution("2025:4:2", y2025d4part2)
}

const (
	noPaperRoll   = '.'
	paperRoll     = '@'
	reachableRoll = 'x'
)

func grid(input string) [][]rune {
	res := make([][]rune, 0)
	for line := range strings.SplitSeq(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		res = append(res, []rune(line))
	}
	return res
}

func printGrid(grid [][]rune) {
	for i := range grid {
		fmt.Println(string(grid[i]))
	}
}

func markReachable(grid [][]rune) int {
	found := 0
	for r, row := range grid {
		for c, val := range row {
			if val == paperRoll && isReachable(r, c, grid) {
				grid[r][c] = reachableRoll
				found++
			}
		}
	}
	return found
}

func isReachable(row, col int, grid [][]rune) bool {
	adj := junk.AdjacentCells(grid, row, col)
	ct := 0
	for _, cell := range adj {
		obj := grid[cell.Row][cell.Col]
		if obj == paperRoll || obj == reachableRoll {
			ct++
		}
	}
	return ct < 4
}

func removeReachable(grid [][]rune) {
	for r, row := range grid {
		for c := range row {
			if grid[r][c] == reachableRoll {
				grid[r][c] = noPaperRoll
			}
		}
	}
}

func y2025d4part1(input string) string {
	grid := grid(input)
	printGrid(grid)
	fmt.Println("----")
	reachable := markReachable(grid)
	printGrid(grid)
	return fmt.Sprint(reachable)
}

func y2025d4part2(input string) string {
	grid := grid(input)

	total := 0
	reachable := markReachable(grid)

	for reachable > 0 {
		total += reachable
		removeReachable(grid)
		reachable = markReachable(grid)
	}

	return fmt.Sprint(total)
}
