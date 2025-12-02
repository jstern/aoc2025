package aoc

import (
	"fmt"
	"strings"
)

func init() {
	registerSolution("2016:7:1", y2016d7part1)
	registerSolution("2016:7:2", y2016d7part2)
}

func y2016d7part1(input string) string {
	total := 0
	for ip := range strings.SplitSeq(input, "\n") {
		if supportsTLS(ip) {
			total += 1
		}
	}
	return fmt.Sprint(total)
}

func supportsTLS(ip string) bool {
	res := false
	allow := true
	pair1 := [2]rune{' ', ' '}
	pair2 := [2]rune{' ', ' '}
	for _, r := range ip {
		if r == '[' || r == ']' {
			allow = !allow
			pair1 = [2]rune{' ', ' '}
			pair2 = [2]rune{' ', ' '}
			continue
		}

		pair1[0] = pair1[1]
		pair1[1] = pair2[0]
		pair2[0] = pair2[1]
		pair2[1] = r

		if (pair1[0] != pair1[1]) && (pair1[1] == pair2[0] && pair1[0] == pair2[1]) {
			if !allow {
				return false
			}
			res = true
		}
	}
	return res
}

func y2016d7part2(input string) string {
	return "wrong again"
}
