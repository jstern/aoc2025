package junk

import "strconv"

// StringToInt converts a string to an int.
type StringToInt func(s string) (int, error)

// Hex converts a hex value to an int.
func Hex(s string) (int, error) {
	i, err := strconv.ParseInt(s, 16, 64)
	return int(i), err
}

// Abs returns abs value of n.
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Min returns the smaller of a and b.
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Max returns the larger of a and b.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
