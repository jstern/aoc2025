package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/jstern/aoc2025/aoc/junk"
)

func init() {
	registerSolution("2025:1:1", y2025d1part1)
	registerSolution("2025:1:2", y2025d1part2)
}

type dial struct {
	size     int
	position int
	clicks   int
}

func (d *dial) String() string {
	return fmt.Sprintf("%02d,%d", d.position, d.clicks)
}

func parseDay1(t string) int {
	t = strings.TrimSpace(t)
	if len(t) == 0 {
		return 0
	}
	dist, err := strconv.Atoi(t[1:])
	if err != nil {
		panic(err)
	}
	dir := string(t[0])
	if dir == "L" {
		dist = -dist
	} else if dir != "R" {
		panic("bad direction " + dir)
	}
	return dist
}

func (d *dial) wind(dist int) {
	// wind the dial to zero
	remainingDist := d.toZero(dist)

	for remainingDist != 0 {
		remainingDist = d.toZero(remainingDist)
	}
}

func (d *dial) toZero(dist int) int {
	var remainingDist int

	toZero := 0
	if dist > 0 {
		// e.g. 3 -> 12 clockwise = 9
		toZero = d.size - d.position
	}
	if dist < 0 {
		// e.g. 3 -> 0 counterclockwise = -3
		toZero = -d.position

		// e.g. 0 -> 0 counterclockwise = 12
		if d.position == 0 {
			toZero = -d.size
		}
	}
	if junk.Abs(toZero) > junk.Abs(dist) {
		// we can't get to zero
		d.position = d.position + dist
		remainingDist = 0
	} else {
		d.position = 0
		d.clicks++
		remainingDist = dist - toZero
	}
	if d.position < 0 {
		d.position = d.size + d.position
	}

	return remainingDist
}

func y2025d1part1(input string) string {
	d := &dial{size: 100, position: 50}
	zeroes := 0
	for t := range strings.SplitSeq(input, "\n") {
		if t == "" {
			break
		}
		fmt.Printf("\n%s(%s)\n", t, d)

		// (a % n + n) % n
		d.position = (d.position + parseDay1(t)) % d.size

		print := color.New(color.FgWhite).PrintfFunc()
		if d.position == 0 {
			zeroes++
			print = color.New(color.FgYellow).PrintfFunc()
		}

		print("  %s\n", d)
	}
	return fmt.Sprint(zeroes)
}

func y2025d1part2(input string) string {
	d := &dial{size: 100, position: 50}
	for t := range strings.SplitSeq(input, "\n") {
		if t == "" {
			break
		}
		fmt.Printf("\n%s(%s)", t, d)
		clicks := d.clicks

		d.wind(parseDay1(t))

		print := color.New(color.FgWhite).PrintfFunc()
		if d.clicks > clicks {
			print = color.New(color.FgYellow).PrintfFunc()
		}

		print("  %s\n", d)
	}
	return fmt.Sprint(d.clicks)
}
