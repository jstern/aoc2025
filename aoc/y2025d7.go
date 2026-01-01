package aoc

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jstern/aoc2025/aoc/junk"
	"github.com/samber/lo"
)

func init() {
	registerSolution("2025:7:1", y2025d7part1)
	registerSolution("2025:7:2", y2025d7part2)
}

type tachyonManifold struct {
	origin    junk.Cell2D
	grid      [][]rune
	splits    []junk.Cell2D
	timelines map[int]int
}

func (m tachyonManifold) String() string {
	return strings.Join(lo.Map(m.grid, func(row []rune, _ int) string { return string(row) }), "\n")
}

func newTachyonManifold(input string) tachyonManifold {
	grid := junk.RuneGrid(strings.Split(input, "\n"))
	splits := make([]junk.Cell2D, 0)

	res := tachyonManifold{
		grid:      grid,
		splits:    splits,
		timelines: make(map[int]int),
	}

	prev := []junk.Cell2D{}
	beamAbove := func(cell junk.Cell2D) bool {
		above := cell.N()
		return slices.Contains(prev, above)
	}
	onGrid := func(cell junk.Cell2D) bool {
		return junk.OnGrid(grid, cell)
	}
	splitTimelines := func(cell junk.Cell2D) {
		active := res.timelines[cell.Col]
		fmt.Printf("Splitting %d active timelines at %s...\n", active, cell)
		res.timelines[cell.Col] = 0
		res.timelines[cell.E().Col] += active
		res.timelines[cell.W().Col] += active
	}

	for r, row := range grid {
		beams := make([]junk.Cell2D, 0)
		for c, col := range row {
			cell := junk.Cell2D{Row: r, Col: c}
			switch col {
			case 'S':
				beams = append(beams, cell)
				res.origin = cell
				res.timelines[cell.Col] = 1
			case '.':
				if beamAbove(cell) {
					grid[r][c] = '|'
					beams = append(beams, cell)
				}
			case '^':
				if beamAbove(cell) {
					res.splits = append(res.splits, cell)
					west := cell.W()
					if onGrid(west) {
						grid[west.Row][west.Col] = '|'
						beams = append(beams, west)
					}
					east := cell.E()
					if onGrid(east) {
						grid[east.Row][east.Col] = '|'
						beams = append(beams, east)
					}
					splitTimelines(cell)
				}
			}
		}
		prev = beams
		fmt.Println("parsed row", r)
		//fmt.Println(res)
		//fmt.Println(splits)
	}

	return res
}

func (m *tachyonManifold) countTimelines() int {
	res := 0
	for _, v := range m.timelines {
		res += v
	}
	return res
}

func y2025d7part1(input string) string {
	input = junk.TrimInput(input)
	manifold := newTachyonManifold(input)
	fmt.Println(manifold)
	return fmt.Sprint(len(manifold.splits))
}

func y2025d7part2(input string) string {
	input = junk.TrimInput(input)

	fmt.Println("parsing...")
	manifold := newTachyonManifold(input)

	return fmt.Sprint(manifold.countTimelines())
}
