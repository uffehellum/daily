package daily680

import (
	"fmt"
	"sort"
	"testing"
)

/*
Given a list of elements, find the majority element, which appears more than half the time (> floor(len(lst) / 2.0)).

You can assume that such element exists.

For example, given [1, 2, 1, 1, 3, 4, 0], return 1.
*/

var sample = []int{1, 2, 1, 1, 3, 4, 0}

func TestDaily683Sort(t *testing.T) {
	actual := FindMajoritySort(sample)
	AssertEqual(t, actual, 1, fmt.Sprint("majority", sample))
}
func TestDaily683Hash(t *testing.T) {
	actual := FindMajorityHash(sample)
	AssertEqual(t, actual, 1, fmt.Sprint("majority", sample))
}

func FindMajoritySort(a []int) int {
	sort.Ints(a)
	return a[len(a)/2]
}

func FindMajorityHash(a []int) int {
	h := map[int]int{}
	for _, x := range a {
		i, _ := h[x]
		h[x] = i+1
	}
	x, n := 0, 0
	for y, m := range h{
		if m > n {
			x, n = y, m
		}
	}
	return x
}
