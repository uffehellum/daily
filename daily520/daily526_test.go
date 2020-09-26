package daily520

import (
	"fmt"
	"sort"
	"testing"
)

func Test526(t *testing.T) {
	assertSmallest(t, "daily", 1, "ailyd")
	assertSmallest(t, "daily", 0, "daily")
	assertSmallest(t, "daily", 2, "adily")
}

func assertSmallest(t *testing.T, s string, k int, expected string) {
	actual := Smallest(s, k)
	if actual != expected {
		t.Fatal(fmt.Sprintf("Smallest(s=%s, k=%d) = %s, expected %s", s, k, actual, expected))
	}
}

func Smallest(s string, k int) string {
	switch k {
		case 0:
			return s
		case 1:
			best := s
			for _ = range s {
				s = s[1:] + s[:1]
				if s < best { best = s}
			}
			return best
		default:
			a := []rune(s)
			sort.Sort(byRune(a))
			return string(a)
	}
}

type byRune []rune
func (s byRune) Len() int {
	return len(s)
}
func (s byRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byRune) Less(i, j int) bool {
	return s[i] < s[j]
}
