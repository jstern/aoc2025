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

// RuneFreqs counts occurrences of runes. Pass in a map to build on previous results.
func RuneFreqs(s string, fs map[rune]int) map[rune]int {
	if fs == nil {
		fs = make(map[rune]int)
	}
	for _, r := range s {
		f := fs[r]
		fs[r] = f + 1
	}

	return fs
}
