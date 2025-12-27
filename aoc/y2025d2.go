package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jstern/aoc2025/aoc/junk"
)

func init() {
	registerSolution("2025:2:1", y2025d2part1)
	registerSolution("2025:2:2", y2025d2part2)
}

func sillyIDValue(id int) int {
	str := fmt.Sprint(id)
	if len(str)%2 != 0 {
		return 0
	}
	first, second := str[0:len(str)/2], str[len(str)/2:]
	if first == second {
		return id
	}
	return 0
}

func reallySillyIDValue(id int, memo func(int) []int) int {
	str := fmt.Sprint(id)
	if len(str) == 1 {
		return 0
	}
	sizes := memo(len(str))
	sizes = append(sizes, len(str))
	//fmt.Printf("need to check %s %v\n", str, sizes)
	for _, size := range sizes {
		chunks := junk.Chunks(str, len(str)/size)
		//fmt.Printf("checking %d %v\n", size, chunks)
		unique := junk.NewSet(chunks...)
		if len(unique) == 1 {
			return id
		}
	}
	return 0
}

func y2025d2part1(input string) string {
	total := 0
	input = strings.TrimSpace(input)
	for idRange := range strings.SplitSeq(input, ",") {
		startEnd := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])
		for i := start; i <= end; i++ {
			total += sillyIDValue(i)
		}
	}
	return fmt.Sprint(total)
}

func y2025d2part2(input string) string {
	total := 0
	input = strings.TrimSpace(input)
	for idRange := range strings.SplitSeq(input, ",") {
		startEnd := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])
		for i := start; i <= end; i++ {
			total += reallySillyIDValue(i, junk.Memo(junk.Factors))
		}
	}
	return fmt.Sprint(total)
}
