package maps

// Count is counting word in string
func Count(s string, refCnt map[rune]int) {
	for _, r := range s {
		refCnt[r]++
	}
}

// HasDupeRune is Confirmed the presence of the key
func HasDupeRune(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		if _, exists := runeSet[r]; exists {
			return true
		}
		runeSet[r] = struct{}{}
	}
	return false
}
