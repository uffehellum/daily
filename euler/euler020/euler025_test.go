package euler020

import (
	"math/big"
	"testing"
)

func TestIndexOfFibonacci(t *testing.T){
	AssertEqual(t, IndexOfFibonacciWithNDigits(2), 7)
	AssertEqual(t, IndexOfFibonacciWithNDigits(3), 12)
	AssertEqual(t, IndexOfFibonacciWithNDigits(100), 476)
	AssertEqual(t, IndexOfFibonacciWithNDigits(1000), 4782)
}
func IndexOfFibonacciWithNDigits(digits int) int {
	a := big.NewInt(1)
	b := big.NewInt(1)
	i := 2
	limit :=big.NewInt(1)

	ten := big.NewInt(10)
	for i := 1; i < digits; i++ {
		limit.Mul(limit, ten)
	}

	for limit.Cmp(b) > 0 {
		c := a.Add(a, b)
		a = b
		b = c
		i ++
	}
	return i
}
