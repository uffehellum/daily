package euler020

import (
	"fmt"
	"math/big"
	"testing"
)

func TestRecurringLength(t *testing.T) {
	AssertEqual(t, RecurringLength(big.NewRat(1, 3)), 1)
	AssertEqual(t, RecurringLength(big.NewRat(1, 6)), 1)
	AssertEqual(t, RecurringLength(big.NewRat(1, 7)), 6)
	AssertEqual(t, RecurringLength(big.NewRat(1, 8)), 0)
	AssertEqual(t, RecurringLength(big.NewRat(1, 30)), 1)
	AssertEqual(t, RecurringLength(big.NewRat(1, 99)), 2)
	AssertEqual(t, RecurringLength(big.NewRat(1, 17)), 16)
}

func TestRecurringLengthInt(t *testing.T) {
	AssertEqual(t, RecurringLengthInt(3), 1)
	AssertEqual(t, RecurringLengthInt( 6), 1)
	AssertEqual(t, RecurringLengthInt( 7), 6)
	AssertEqual(t, RecurringLengthInt( 8), 0)
	AssertEqual(t, RecurringLengthInt( 30), 1)
	AssertEqual(t, RecurringLengthInt( 99), 2)
	AssertEqual(t, RecurringLengthInt( 17), 16)
}

func TestMaxRecurringLength(t *testing.T) {
	AssertEqual(t, NumberWithLongestRecurringLength(), 983)
}

func TestMaxRecurringLengthInt(t *testing.T) {
	AssertEqual(t, NumberWithLongestRecurringLengthInt(), 983)
}

func NumberWithLongestRecurringLength() int {
	r := 0
	n := 0
	for i := int64(1); i < 1000; i++ {
		x := RecurringLength(big.NewRat(1, i))
		if x > r {
			fmt.Println(i, x)
			r = x
			n = int(i)
		}
	}
	return n
}

func NumberWithLongestRecurringLengthInt() int {
	r := 0
	n := 0
	for i := 1; i < 1000; i++ {
		x := RecurringLengthInt(i)
		if x > r {
			fmt.Println(i, x)
			n = i
			r = x
		}
	}
	return n
}

func RecurringLength(rat *big.Rat) int {
	tenInt := big.NewInt(10)
	tenInt.Exp(tenInt, big.NewInt(1000), nil)
	thousandRat := &big.Rat{}
	thousandRat.SetInt(tenInt)
	rat = rat.Mul(rat, thousandRat)
	l := 0
	one := big.NewInt(1)
	nine := big.NewRat(9, 1)
	ten := big.NewRat(10, 1)
	nines := big.NewRat(0, 1)
	for rat.Denom().Cmp(one) > 0 {
		l++
		nines = nines.Mul(nines, ten)
		nines = nines.Add(nines, nine)
		rat = rat.Mul(rat, nines)
	}
	return l
}

func RecurringLengthInt(n int) int {
	a := make([]int, n * 10)
	r := 1
	for p := 1; r > 0; p++ {
		if a[r] > 0 {
			// fmt.Println(n, p-a[r])
			return p - a[r]
		}
		a[r] = p
		r = r * 10 % n
	}
	return 0
}