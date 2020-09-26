package daily520

import (
	"fmt"
	"testing"
)

func Test524(t *testing.T) {
	assertMaxSum(t, []int {1, 2, 3}, 6)
	assertMaxSum(t, []int {34, -50, 42, 14, -5, 86}, 137)
}

func assertMaxSum(t *testing.T, a []int, expected int){
	actual := MaxSum(a)
	if actual != expected {
		t.Fatal(fmt.Printf(
			"MaxSum(%v) = %d, expected %d",
			a, actual, expected))
	}
}

func MaxSum(a []int) int {
	if len(a) == 0 { return 0 }
	sum := 0
	min := 0
	max := a[0]
	for _, x := range a {
		sum += x
		if sum < min { min = sum }
		if sum > max { max = sum }
	}
	return max - min
}
