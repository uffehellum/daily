package euler020

import "testing"

func TestAbundant(t *testing.T) {
	AssertTrue(t, isAbundant(12), "12")
}

func TestSumNotSum(t *testing.T) {
	all := AllAbundantBelow(28123)
	AssertEqual(t, SumNotSumOfTwoAbundantBelow(all), 4179871)
}

func TestAllAbundantBelow(t *testing.T) {
	all := AllAbundantBelow(28123)
	AssertTrue(t, all[12], "12")
}

func AssertTrue(t *testing.T, test bool, message string) {
	if !test {
		t.Fatal("expected true", message)
	}
}

func SumNotSumOfTwoAbundantBelow(m []bool) int {
	r := make([]bool, len(m))
	for i, b := range m {
		if !b { continue }
		for j := i+1; j < len(r); j++ {
			if !r[j] && m[j - i] {
				r[j] = true
			}
		}
	}
	s := 0
	for i, b := range r {
		if !b {
			s += i
		}
	}
	return s
}
func AllAbundantBelow(n int) []bool {
	r := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if isAbundant(i) {
			r[i] = true
		}
	}
	return r
}
func isAbundant(n int) bool {
	s := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			s += i
		}
	}
	return s > n
}
