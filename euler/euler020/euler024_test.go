package euler020

import (
	"fmt"
	"testing"
)

func TestMillionthPermutation(t *testing.T) {
	AssertEqualString(t, NthPermutation(1000000, "0123456789"), "2783915460")
}

func AssertEqualString(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Fatal("expected", expected, "actual", actual)
	}
}

func NthPermutation(n int, s string) string {
	n--
	x := 1
	for i, _ := range s {
		x *= i + 1
	}
	if n > x {
		panic(fmt.Sprintf("Permutations total %d is less than selected %d", x, n))
	}
	r := ""
	for s > "" {
		x /=  len(s)
		d := n / x
		n %= x
		r += s[d : d+1]
		s = s[:d] + s[d+1:]
	}
	return r
}
