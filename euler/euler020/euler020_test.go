package euler020

import (
	"math/big"
	"testing"
)

func TestDigitSum(t *testing.T) {
	AssertSame(t, FactorialDigitSum(1), 1, "1")
	AssertSame(t, FactorialDigitSum(2), 2, "2")
	AssertSame(t, FactorialDigitSum(4), 6, "4")
	AssertSame(t, FactorialDigitSum(100), 648, "100")
}

func AssertSame(t *testing.T, actual, expected int, message string) {
	if expected != actual {
		t.Fatal(message, "expected", expected, "actual", actual)
	}
}

func FactorialDigitSum(n int) int {
	d := big.NewInt(1)
	d = d.MulRange(2, int64(n))
	//zero := big.NewInt(0)
	ten := big.NewInt(10)
	rem := big.NewInt(0)
	sum := 0
	for d.Sign() > 0 {
		d.DivMod(d, ten, rem)
		sum += int(rem.Int64())
	}
	return sum
}