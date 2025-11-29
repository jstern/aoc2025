package junk

import "strconv"

// StringToInt converts a string to an int.
type StringToInt func(s string) (int, error)

// Hex converts a hex value to an int.
func Hex(s string) (int, error) {
	i, err := strconv.ParseInt(s, 16, 64)
	return int(i), err
}
