package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2016d7Input = `abba[mnop]qrst
abcd[bddb]xyyx
aaaa[qwer]tyui
ioxxoj[asdfgh]zxcvbn
`

func Test_2016_Day_7_Part_1_Example(t *testing.T) {
	result := y2016d7part1(y2016d7Input)
	assert.Equal(t, "2", result)
}

func Test_2016_Day_7_Part_1_TLS(t *testing.T) {
	tcs := []struct {
		in  string
		out bool
	}{
		{"abba[mnop]qrst", true},
		{"abcd[bddb]xyyx", false},
		{"aaaa[qwer]tyui", false},
		{"ioxxoj[asdfgh]zxcvbn", true},
	}
	for _, tc := range tcs {
		assert.Equal(t, tc.out, supportsTLS(tc.in))
	}
}

func Test_2016_Day_7_Part_2_Example(t *testing.T) {
	//result := y2016d7part2(y2016d7Input)
	//assert.Equal(t, "still right!", result)
	t.Skip()
}
