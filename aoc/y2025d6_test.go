package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2025d6Input = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`

func Test_2025_Day_6_Part_1_Example(t *testing.T) {
	result := y2025d6part1(y2025d6Input)
	assert.Equal(t, "4277556", result)
}

func Test_2025_Day_6_Part_2_Example(t *testing.T) {
	result := y2025d6part2(y2025d6Input)
	assert.Equal(t, "3263827", result)
}

func Test_2025_Day_6_col_toProblem(t *testing.T) {
	col := &cephCol{
		op:   cephOpAdd,
		args: [][]rune{{'6', '4', ' '}, {'2', '3', ' '}, {'3', '1', '4'}},
	}
	problem := col.toProblem()
	assert.Equal(t, problem.op, col.op)
	assert.Equal(t, []int{623, 431, 4}, problem.args)
}
