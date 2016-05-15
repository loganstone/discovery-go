package methods

import "strings"

// MultiSet is data structure
type MultiSet map[string]int

// Insert is up counting to val
func (m MultiSet) Insert(val string) {
	m[val]++
}

// Erase is down counting to val
func (m MultiSet) Erase(val string) {
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

// Count is return val count
func (m MultiSet) Count(val string) int {
	return m[val]
}

func (m MultiSet) String() string {
	s := "{ "
	for val, count := range m {
		s += strings.Repeat(val+" ", count)
	}
	return s + "}"
}
