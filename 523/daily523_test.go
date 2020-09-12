package main

import (
	"fmt"
	"testing"
)

func TestCountPairs(t *testing.T) {
	assertCountPairs(t, 2, 0, 1)
	assertCountPairs(t, 2, 1, 0)
	assertCountPairs(t, 3, 3, 2)
	assertCountPairs(t, 3, 4, 0)
	assertCountPairs(t, 4, 0, 1)
	assertCountPairs(t, 4, 2, 2)
}

func assertCountPairs(t *testing.T, sum, xor, expected int) {
	actual := Count(sum, xor)
	if expected == actual {
		return
	}
	message := fmt.Sprintf(
		"Count(a + b = %d, a XOR b = %d) got %d expected %d",
		sum, xor, actual, expected)
	t.Fatal(message)
}

func Count(sum, xor int) int {
	r := 0
	for i := 1; i < sum; i++ {
		if i ^ (sum- i) == xor {
			r++
		}
	}
	return r
}
