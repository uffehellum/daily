package euler010

import (
	"math/big"
	"testing"
)

/*
	https://projecteuler.net/problem=15
	Starting in the top left corner of a 2×2 grid, and only being
	able to move to the right and down, there are exactly 6 routes
	to the bottom right corner.
	How many such routes are there through a 20×20 grid?
 */
func TestGrid2(t *testing.T) {
	actual := CountGridPaths(2)
	expected := &big.Int{}
	expected.SetInt64(6)
	AssertEqualBig(t, actual, expected)
}

func AssertEqualBig(t *testing.T, actual *big.Int, expected *big.Int) {
	if actual.Cmp(expected) != 0 {
		t.Fatal("exoected", expected, "actual", actual)
	}
}

func TestGrid20(t *testing.T) {
	actual := CountGridPaths(20)
	expected := &big.Int{}
	expected.SetString("1378465288200", 10)
	AssertEqualBig(t, actual, expected)
}

func CountGridPaths(gridSize int64) *big.Int {
	// spread grid steps into grid+1 buckets
	// = p grid + 1 - 1(grid)
	// = 2*grid over grid
	// = (2 * grid)! / grid! / grid!

	a := bigFactorial(gridSize * 2)
	b := bigFactorial(gridSize)
	a = a.Div(a, b)
	a = a.Div(a, b)
	return a
}

func bigFactorial(n int64) *big.Int {
	a := &big.Int{}
	b := &big.Int{}
	a.SetInt64(1)
	for i := int64(2); i < n; i++ {
		b.SetInt64(i)
		a = a.Mul(a, b)
	}
	return a
}


