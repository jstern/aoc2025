package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jstern/aoc2025/aoc/junk"
)

func init() {
	registerSolution("2025:6:1", y2025d6part1)
	registerSolution("2025:6:2", y2025d6part2)
}

type cephOp string

const (
	cephOpAdd cephOp = "+"
	cephOpMul cephOp = "*"
)

func cephMul(xs ...int) int {
	res := 1
	for _, x := range xs {
		res *= x
	}
	return res
}

func cephAdd(xs ...int) int {
	res := 0
	for _, x := range xs {
		res += x
	}
	return res
}

type cephProblem struct {
	args []int
	op   cephOp
}

func (p *cephProblem) solve() int {
	switch p.op {
	case cephOpAdd:
		return cephAdd(p.args...)
	case cephOpMul:
		return cephMul(p.args...)
	default:
		panic("unknown op")
	}
}

type cephWorksheet struct {
	problems []cephProblem
}

func (w cephWorksheet) solveAll() int {
	total := 0
	for _, problem := range w.problems {
		total += problem.solve()
	}
	return total
}

func newCephWorksheet(input string) cephWorksheet {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	problems := make([]cephProblem, len(strings.Fields(rows[0])))
	for i := range problems {
		problems[i] = cephProblem{args: make([]int, 0)}
	}

	for _, row := range rows {
		for c, col := range strings.Fields(row) {
			problem := problems[c]
			switch col {
			case "*":
				problem.op = cephOpMul
			case "+":
				problem.op = cephOpAdd
			default:
				arg, err := strconv.Atoi(col)
				if err != nil {
					panic(err)
				}
				problem.args = append(problem.args, arg)
			}
			problems[c] = problem
		}
	}

	return cephWorksheet{problems: problems}
}

func y2025d6part1(input string) string {
	worksheet := newCephWorksheet(input)
	return fmt.Sprint(worksheet.solveAll())
}

type cephCol struct {
	op    cephOp
	start int // inclusive
	end   int // exclusive
	args  [][]rune
}

func parseCephOp(r rune) cephOp {
	switch r {
	case '*':
		return cephOpMul
	case '+':
		return cephOpAdd
	default:
		panic("bad op")
	}
}

func cephCols(input string) []*cephCol {
	input = strings.TrimPrefix(input, "\n")
	input = strings.TrimSuffix(input, "\n")

	rows := strings.Split(input, "\n")
	cols := make([]*cephCol, 0)

	// grab ops, starts, widths from last row
	var col *cephCol
	for i, char := range rows[len(rows)-1] {
		switch char {
		case '*', '+':
			col = &cephCol{op: parseCephOp(char), start: i, args: make([][]rune, len(rows)-1)}
			cols = append(cols, col)
		}
		col.end = i + 1
	}

	figures := junk.RuneGrid(rows[0 : len(rows)-1])
	for _, col := range cols {
		for row := range len(rows) - 1 {
			i := -1
			for s := col.start; s < col.end; s++ {
				i++
				col.args[row] = append(col.args[row], figures[row][s])
			}
		}
	}

	return cols
}

func (c *cephCol) toProblem() *cephProblem {
	transposed := junk.Transpose(c.args)
	args := make([]int, 0)
	for _, row := range transposed {
		raw := strings.TrimSpace(string(row))
		if raw == "" {
			continue // i'm doing something dumb somewhere else
		}
		val, err := strconv.Atoi(raw)
		if err != nil {
			panic(val)
		}
		args = append(args, val)
	}
	return &cephProblem{op: c.op, args: args}
}

func y2025d6part2(input string) string {
	cols := cephCols(input)
	total := 0
	for _, col := range cols {
		problem := col.toProblem()
		add := problem.solve()
		total += add
	}
	return fmt.Sprint(total)
}
