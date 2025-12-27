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

// Chunks returns s cut into chunks of size n
func Chunks(s string, n int) []string {
	res := make([]string, 0)
	start := 0
	end := n
	max := len(s) + 1
	for end < max {
		res = append(res, s[start:end])
		start += n
		end += n
	}

	return res
}
