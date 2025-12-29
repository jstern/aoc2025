package junk

type Cell2D struct {
	Row int
	Col int
}

// AdjacentCells for (row, col) x on a finite 2D grid.
func AdjacentCells[T any](grid [][]T, row, col int) []Cell2D {
	res := make([]Cell2D, 0)

	candidates := []Cell2D{
		{Row: row - 1, Col: col - 1},
		{Row: row - 1, Col: col + 0},
		{Row: row - 1, Col: col + 1},

		{Row: row + 0, Col: col - 1},
		{Row: row + 0, Col: col + 1},

		{Row: row + 1, Col: col - 1},
		{Row: row + 1, Col: col + 0},
		{Row: row + 1, Col: col + 1},
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
