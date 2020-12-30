package euler010

import (
	"math/big"
	"testing"
)

func Test15(t *testing.T){
	actual := SumDigitsPowerOfTwo(15)
	AssertEqual(t, actual, 26)
}

func Test1000(t *testing.T){
	actual := SumDigitsPowerOfTwo(1000)
	AssertEqual(t, actual, 1366)
}

func AssertEqual(t *testing.T, actual int, expected int) {
	if expected != actual {
		t.Fatal("exoected", expected, "actual", actual)
	}
}

func SumDigitsPowerOfTwo(power int) int {
	// two:= big.Int(2)
	x:= &big.Int{}
	x.SetBit(x, power, 1)
	ten := (&big.Int{}).SetInt64(10)
	rem:= &big.Int{}
	var sum int64
	for x.Sign() > 0 {
		x.QuoRem(x, ten, rem)
		sum += rem.Int64()
	}
	return int(sum)
}