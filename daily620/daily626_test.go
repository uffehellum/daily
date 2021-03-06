package daily620

import (
	"fmt"
	"sort"
	"testing"
)

func TestLargestProduct(t *testing.T) {
	AssertEqual(t, 500, LargestProduct([]int{-10, -10, 5, 2}))
	AssertEqual(t, -6, LargestProduct([]int{-1, -2, -3}))
	AssertEqual(t, -6, LargestProduct([]int{-1, -2, -3}))
}

func LargestProduct(a []int) int {
	sort.Ints(a)
	n := len(a)
	r := -int(^uint(0)/2)
	for i := 0; i < 4; i++ {
		x := product(a[0:i], a[n-3+i:n])
		if x > r {
			r = x
		}
	}
	return r
}

func product(a1, a2 []int) int {
	r := 1
	fmt.Println(a1, a2)
	for _, x := range a1 {
		r *= x
	}
	for _, x := range a2 {
		r *= x
	}
	return r
}
