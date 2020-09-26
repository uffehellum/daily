package daily520

import (
	"fmt"
	"testing"
)

func Test529(t *testing.T) {
	assertClip(t, "bob", []string{"bob"})
	assertClip(t, "racecarannakayak",
		[]string{"racecar", "anna", "kayak"})
}

func assertClip(t *testing.T, source string, expected []string) {
	actual := ClipPalindromes(source)
	if fmt.Sprint(actual) != fmt.Sprint(expected) {
		t.Fatal(fmt.Sprintf(
			"SplitPalindromes(source=%s) = %v, expected %v",
			source, actual, expected))
	}
}

func ClipPalindromes(s string) []string {
	r := []string{}
	for s > "" {
		l:
		for n := len(s); n >= 1; n-- {
			for i := (n - 1) / 2; i >= 0; i-- {
				if s[i] != s[n-i-1] {
					continue l
				}
			}
			r = append(r, s[:n])
			s = s[n:]
			break
		}
	}
	return r
}