package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2025d3Input = `987654321111111
811111111111119
234234234234278
818181911112111
`

func Test_2025_Day_3_Part_1_Example(t *testing.T) {
	result := y2025d3part1(y2025d3Input)
	assert.Equal(t, "357", result)
}

func Test_2025_Day_3_Part_1_joltage(t *testing.T) {
	tcs := []struct {
		bank     *batteryBank
		expected int
	}{
		{newBatteryBank("9999"), 0},
		{&batteryBank{[]int{9, 1, 8, 2}, []bool{true, false, true, false}}, 98},
	}
	for _, tc := range tcs {
		assert.Equal(t, tc.expected, tc.bank.joltage())
	}
}

func Test_2025_Day_3_Part_1_setupBank(t *testing.T) {
	desc := "9182"
	bank := setupBank(desc, 2)
	assert.Equal(t, []bool{true, false, true, false}, bank.on)

	desc = "234234234234278"
	bank = setupBank(desc, 2)
	assert.Equal(t, 78, bank.joltage())
}

func Test_2025_Day_3_Part_2_Example(t *testing.T) {
	result := y2025d3part2(y2025d3Input)
	assert.Equal(t, "3121910778619", result)
}
