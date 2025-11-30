package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var y2016d4Input = `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]
`

func Test_2016_Day_4_Part_1_Example(t *testing.T) {
	result := y2016d4part1(y2016d4Input)
	assert.Equal(t, "1514", result)
}

func Test_2016_Day_4_Part_2_decodeName(t *testing.T) {
	desc := "qzmt-zixmtkozy-ivhz-343"
	res := decodeName(desc, 343)
	assert.Contains(t, res, "very encrypted name")
}
