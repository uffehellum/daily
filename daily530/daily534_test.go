package daily530

import "testing"

func TestDaily534(t *testing.T) {
	assertPalindrome(t, "carrace", true)
	assertPalindrome(t, "daily", false)
	assertPalindrome(t, "ab", false)
	assertPalindrome(t, "aab", true)
}

func assertPalindrome(t *testing.T, s string, expected bool) {
	actual := CanPermutateToPalindrome(s)
	if actual != expected {
		t.Fatalf("Expected %v found %v for %s", expected, actual, s)
	}
}
func CanPermutateToPalindrome(s string) bool {
	h := make(map[rune]bool)
	for _, c:= range s {
		h[c] = !h[c]
	}
	odds := 0
	for _, b := range h {
		if b { odds++ }
	}
	return odds < 2
}