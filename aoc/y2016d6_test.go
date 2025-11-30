package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2016d6Input = `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar
`

func Test_2016_Day_6_Part_1_Example(t *testing.T) {
	result := y2016d6part1(y2016d6Input)
	assert.Equal(t, "easter", result)
}

func Test_2016_Day_6_Part_2_Example(t *testing.T) {
	result := y2016d6part2(y2016d6Input)
	assert.Equal(t, "advent", result)
}
