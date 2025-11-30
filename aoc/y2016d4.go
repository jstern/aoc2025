package aoc

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/jstern/aoc2025/aoc/junk"
)

func init() {
	registerSolution("2016:4:1", y2016d4part1)
	registerSolution("2016:4:2", y2016d4part2)
}

func y2016d4part1(input string) string {
	total := 0
	for desc := range strings.SplitSeq(input, "\n") {
		desc = strings.TrimSpace(desc)
		if desc != "" {
			total += readTheRoom(desc)
		}
	}
	return fmt.Sprint(total)
}

func y2016d4part2(input string) string {
	var res int

	for desc := range strings.SplitSeq(input, "\n") {
		desc = strings.TrimSpace(desc)
		if desc == "" {
			continue
		}
		sectorID := readTheRoom(desc)
		if sectorID == 0 {
			continue
		}
		decoded := decodeName(desc, sectorID)
		name := color.MagentaString(decoded)
		if strings.Contains(decoded, "north") {
			name = color.CyanString(decoded)
			res = sectorID
		}
		fmt.Printf("%d %s %s\n", sectorID, desc, name)
	}
	return fmt.Sprint(res)
}

// Takes a day 4 room descriptor and returns a sector id if the descriptor is valid.
func readTheRoom(desc string) int {
	var sectorID int
	var checksum string
	var err error
	freqs := make(map[rune]int)

	for part := range strings.SplitSeq(desc, "-") {
		if strings.HasSuffix(part, "]") {
			parts := strings.Split(part, "[")
			sectorID, err = strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			checksum = strings.TrimSuffix(parts[1], "]")
			if checksum == roomChecksum(freqs) {
				return sectorID
			}
		}
		freqs = junk.RuneFreqs(part, freqs)
	}

	return 0
}

func roomChecksum(freqs map[rune]int) string {
	// top 5 most frequent letters, ties broken alphabetically
	keys := slices.SortedFunc(maps.Keys(freqs), func(a, b rune) int {
		res := cmp.Compare(freqs[b], freqs[a]) // desc
		if res == 0 {
			return cmp.Compare(a, b)
		}
		return res
	})
	return string(keys[:5])
}

func decodeName(desc string, id int) string {
	orig := []byte(desc)
	//fmt.Println(orig)
	res := make([]byte, len(desc))
	delta := id % 26
	// 97-122
	for i, b := range orig {
		if b == byte(45) {
			res[i] = byte(32)
			continue
		}
		start := int(b) - 97
		end := start + delta
		res[i] = byte((end % 26) + 97)
	}
	//fmt.Println(res)
	return string(res)
}
