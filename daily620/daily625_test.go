package daily620

import "testing"

func TestStripe(t *testing.T) {
	AssertEqual(t,  0, LongestConsecutiveRunOfOnes(0))
	AssertEqual(t,  3, LongestConsecutiveRunOfOnes(156))
	AssertEqual(t,  4, LongestConsecutiveRunOfOnes(15))
	AssertEqual(t,  1, LongestConsecutiveRunOfOnes(0x55))
}

func LongestConsecutiveRunOfOnes(n int) int {
	r := 0
	l := 0
	for n > 0 {
		if n % 2 == 0 {
			l = 0
		} else {
			l++
			if l > r {
				r = l
			}
		}
		n /= 2
	}
	return r
}

func AssertEqual(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Fatal("expected", expected, "actual", actual)
	}
}
