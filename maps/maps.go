package maps

// Count is counting word in string
func Count(s string, refCnt map[rune]int) {
	for _, r := range s {
		refCnt[r]++
	}
}
