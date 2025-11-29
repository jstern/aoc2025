package junk

// ReverseString reverses a string.
func ReverseString(s string) string {
	runes := []rune(s)
	ct := len(runes) - 1
	reversed := make([]rune, len(runes))
	for i, r := range runes {
		reversed[ct-i] = r
	}
	return string(reversed)
}
