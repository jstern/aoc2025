package aoc

import (
	"cmp"
	"maps"
	"slices"
	"strings"
)

func init() {
	registerSolution("2016:6:1", y2016d6part1)
	registerSolution("2016:6:2", y2016d6part2)
}

func y2016d6part1(input string) string {
	return repDecode(input, true)
}

func repDecode(input string, most bool) string {
	var freqs []map[rune]int
	for line := range strings.SplitSeq(input, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if len(freqs) == 0 {
			freqs = make([]map[rune]int, len(line))
			for i := range len(line) {
				freqs[i] = make(map[rune]int)
			}
		}
		for i, r := range line {
			fs := freqs[i]
			fr := fs[r]
			fs[r] = fr + 1
		}
	}
	res := make([]rune, len(freqs))
	for i, fs := range freqs {
		chars := slices.SortedFunc(maps.Keys(fs), func(a, b rune) int {
			res := cmp.Compare(fs[b], fs[a]) // desc
			if res == 0 {
				return cmp.Compare(a, b)
			}
			return res
		})
		if most {
			res[i] = chars[0]
		} else {
			res[i] = chars[len(chars)-1]
		}
	}
	return string(res)
}

func y2016d6part2(input string) string {
	return repDecode(input, false)
}
