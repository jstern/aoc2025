package aoc

import (
	"fmt"
	"testing"

	"github.com/jstern/aoc2025/aoc/junk"
	"github.com/stretchr/testify/assert"
)

var y2025d2Input = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124
`

func Test_2025_Day_2_Part_1_Example(t *testing.T) {
	result := y2025d2part1(y2025d2Input)
	assert.Equal(t, "1227775554", result)
}

func Test_2025_Day_2_sillyIDValue(t *testing.T) {
	tcs := []struct {
		id       int
		expected int
	}{
		{11, 11},
		{22, 22},
		{1, 0},
		{101, 0},
		{1188511885, 1188511885},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprint(tc.id), func(t *testing.T) {
			assert.Equal(t, tc.expected, sillyIDValue(tc.id))
		})
	}
}

func Test_2025_Day_2_Part_2_Example(t *testing.T) {
	result := y2025d2part2(y2025d2Input)
	assert.Equal(t, "4174379265", result)
}

func Test_2025_Day_2_reallySillyIDValue(t *testing.T) {
	tcs := []struct {
		id       int
		expected int
	}{
		{11, 11},
		{22, 22},
		{1, 0},
		{101, 0},
		{1188511885, 1188511885},
		{999, 999},
		{121212, 121212},
		{123412341234, 123412341234},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprint(tc.id), func(t *testing.T) {
			assert.Equal(t, tc.expected, reallySillyIDValue(tc.id, junk.Factors))
		})
	}
}
