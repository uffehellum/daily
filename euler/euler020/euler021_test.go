package euler020

import "testing"

func TestD(t *testing.T){
	AssertEqual(t, d(220), 284)
	AssertEqual(t, d(284), 220)
}

func TestSumAmicable(t *testing.T){
	AssertEqual(t, SumAmicableUnderN(10), 0)
	AssertEqual(t, SumAmicableUnderN(100), 0)
	AssertEqual(t, SumAmicableUnderN(1000), 504)
	AssertEqual(t, SumAmicableUnderN(10000), 31626)
}

func AssertEqual(t *testing.T,  actual, expected int) {
	if actual != expected {
		t.Fatal("expected", expected, "actual", actual)
	}
}

func SumAmicableUnderN(n int) int {
	s := 0
	for a := 1; a <= n; a++{
		b := d(a)
		if b != a && d(b) == a {
			s += a
		}
	}
	return s
}
func d(n int) int {
	s := 0
	for i := 1; i < n; i++ {
		if n % i == 0 {
			s += i
		}
	}
	return s
}