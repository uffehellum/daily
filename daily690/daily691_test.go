package daily690

import "testing"

func TestDaily691(t *testing.T) {
	AssertEqual(t, 1, AndBetween(1, 1), "1-1")
	AssertEqual(t, 0, AndBetween(1, 100), "1-100")
	AssertEqual(t, 4, AndBetween(4, 7), "4-7")
}

func AssertEqual(t *testing.T, expected int, actual int, message string) {
	if expected != actual {
		t.Fatal("expected", expected, "actual", actual, message)
	}
}

func AndBetween(m int, n int) int {
	r := n
	for x := m; x < n; x++ {
		r &= x
	}
	return r
}
