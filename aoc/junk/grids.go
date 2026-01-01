package junk

import "fmt"

type Cell2D struct {
	Row int
	Col int
}

func (c Cell2D) String() string {
	return fmt.Sprintf("(%d,%d)", c.Row, c.Col)
}

func (c Cell2D) N() Cell2D {
	return Cell2D{c.Row - 1, c.Col}
}

func (c Cell2D) NE() Cell2D {
	return Cell2D{c.Row - 1, c.Col + 1}
}

func (c Cell2D) NW() Cell2D {
	return Cell2D{c.Row - 1, c.Col - 1}
}

func (c Cell2D) W() Cell2D {
	return Cell2D{c.Row, c.Col - 1}
}

func (c Cell2D) E() Cell2D {
	return Cell2D{c.Row, c.Col + 1}
}

func (c Cell2D) SW() Cell2D {
	return Cell2D{c.Row + 1, c.Col - 1}
}

func (c Cell2D) SE() Cell2D {
	return Cell2D{c.Row + 1, c.Col + 1}
}

func (c Cell2D) S() Cell2D {
	return Cell2D{c.Row + 1, c.Col}
}

// AdjacentCells for (row, col) x on a finite 2D grid.
func AdjacentCells[T any](grid [][]T, row, col int) []Cell2D {
	res := make([]Cell2D, 0)

	ref := Cell2D{Row: row, Col: col}

	candidates := []Cell2D{
		ref.NW(),
		ref.N(),
		ref.NE(),

		ref.W(),
		ref.E(),

		ref.SW(),
		ref.S(),
		ref.SE(),
	}

	for _, c := range candidates {
		if c.Row >= 0 && c.Row < len(grid) && c.Col >= 0 && c.Col < len(grid[0]) {
			res = append(res, c)
		}
	}
	return res
}

// RuneGrid converts a slice of strings into a slice of rune slices.
func RuneGrid(rows []string) [][]rune {
	res := make([][]rune, len(rows))
	for r, row := range rows {
		res[r] = []rune(row)
	}
	return res
}

// OnGrid returns true if the cell exists on the grid.
func OnGrid[T any](grid [][]T, cell Cell2D) bool {
	return (cell.Row >= 0 &&
		cell.Row < len(grid) &&
		cell.Col >= 0 &&
		cell.Col < len(grid[0]))
}
