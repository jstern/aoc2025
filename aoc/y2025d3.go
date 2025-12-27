package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/jstern/aoc2025/aoc/junk"
)

func init() {
	registerSolution("2025:3:1", y2025d3part1)
	registerSolution("2025:3:2", y2025d3part2)
}

type batteryBank struct {
	batteries []int
	on        []bool
}

func (b *batteryBank) String() string {
	res := make([]string, len(b.batteries))
	yellow := color.New(color.FgYellow, color.Bold).SprintFunc()

	for i, v := range b.batteries {
		if b.on[i] {
			res[i] = yellow(v)
		} else {
			res[i] = fmt.Sprint(v)
		}
	}
	return strings.Join(res, "")
}

func (b *batteryBank) joltage() int {
	j := 0
	c := 1

	for i := len(b.batteries) - 1; i >= 0; i-- {
		if b.on[i] {
			j += c * b.batteries[i]
			c *= 10
		}
	}

	return j
}

func newBatteryBank(desc string) *batteryBank {
	desc = strings.TrimSpace(desc)
	if desc == "" {
		return nil
	}
	ints := junk.Chunks(desc, 1)
	batteries := make([]int, len(ints))
	for i, b := range ints {
		v, err := strconv.Atoi(b)
		if err != nil || v < 0 || v > 9 {
			panic("bad battery " + b + "")
		}
		batteries[i] = v
	}
	return &batteryBank{batteries: batteries, on: make([]bool, len(batteries))}
}

func setupBank(desc string, ct int) *batteryBank {
	bank := newBatteryBank(desc)
	if bank == nil {
		return nil
	}

	after := -1
	remaining := ct

	for ; remaining > 0; remaining-- {
		after = turnOnNext(bank, remaining, after)
	}

	return bank
}

func turnOnNext(b *batteryBank, remaining, prev int) int {
	// find the index of the next highest digit in the window between prev and len(bank.batteries) - remaining
	// activate the battery at that index and return the index
	hi := 0
	var flip int
	for i := prev + 1; i <= len(b.batteries)-remaining; i++ {
		if b.batteries[i] > hi {
			hi = b.batteries[i]
			flip = i
		}
	}
	b.on[flip] = true
	return flip

}

func y2025d3part1(input string) string {
	total := 0
	for desc := range strings.SplitSeq(input, "\n") {
		bank := setupBank(desc, 2)
		if bank != nil {
			joltage := bank.joltage()
			fmt.Printf("%s : %d\n", bank, joltage)
			total += joltage
		}
	}
	return fmt.Sprint(total)
}

func y2025d3part2(input string) string {
	total := 0
	for desc := range strings.SplitSeq(input, "\n") {
		bank := setupBank(desc, 12)
		if bank != nil {
			joltage := bank.joltage()
			fmt.Printf("%s : %d\n", bank, joltage)
			total += joltage
		}
	}
	return fmt.Sprint(total)
}
