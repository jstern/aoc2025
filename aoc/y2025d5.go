package aoc

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jstern/aoc2025/aoc/junk"
)

func init() {
	registerSolution("2025:5:1", y2025d5part1)
	registerSolution("2025:5:2", y2025d5part2)
}

type freshRange struct {
	min int
	max int
}

func (r freshRange) String() string {
	return fmt.Sprintf("%d-%d", r.min, r.max)
}

func (r freshRange) includes(id int) bool {
	return id >= r.min && id <= r.max
}

func (r freshRange) size() int {
	res := r.max - r.min + 1
	if res < 1 {
		panic("size less than one")
	}
	return res
}

type ingredientDB struct {
	ranges []freshRange
}

func (db *ingredientDB) addRange(min, max int) {
	if max < min {
		panic("aaah")
	}
	db.ranges = append(db.ranges, freshRange{min, max})
	slices.SortFunc(db.ranges, func(a, b freshRange) int {
		return a.min - b.min
	})
}

func (db *ingredientDB) fresh(id int) bool {
	for _, fr := range db.ranges {
		if fr.includes(id) {
			return true
		}
	}
	return false
}

func y2025d5part1(input string) string {
	parts := strings.Split(input, "\n\n")
	ranges := parts[0]
	ids := strings.TrimSpace(parts[1])

	db := &ingredientDB{[]freshRange{}}
	for fr := range strings.SplitSeq(ranges, "\n") {
		parts := strings.Split(fr, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])
		db.addRange(min, max)
	}

	fresh := 0
	for idnum := range strings.SplitSeq(ids, "\n") {
		id, _ := strconv.Atoi(idnum)
		if db.fresh(id) {
			fresh++
		}
	}

	return fmt.Sprint(fresh)
}

func y2025d5part2(input string) string {
	parts := strings.Split(input, "\n\n")
	ranges := parts[0]

	db := &ingredientDB{[]freshRange{}}
	for fr := range strings.SplitSeq(ranges, "\n") {
		parts := strings.Split(fr, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])
		db.addRange(min, max)
	}

	fresh := 0
	maxID := 0
	for _, cur := range db.ranges {
		if cur.max <= maxID {
			continue
		}
		// we know current range ends after max counted id
		toCount := freshRange{
			min: junk.Max(cur.min, maxID+1),
			max: cur.max, // and we know it ends after the previous maxID from cur.max <= maxID above
		}
		fresh += toCount.size()
		maxID = toCount.max
	}
	return fmt.Sprint(fresh)
}
