package daily570

import (
	"fmt"
	"sort"
	"testing"
)

func NextPerm(a []int) []int {
	if len(a) < 2 {
		return a
	}
	i := len(a) - 1
	for i > 0 && a[i] <= a[i - 1] {
		i--
	}
	if i > 0 {
		a[i], a[i - 1] = a[i - 1], a[i]
	}
	sort.Ints(a[i:])
	return a
}

func Test1(t *testing.T) {
	a := []int{1}
	b := []int{1}
	AssertEqualA(t, NextPerm(a), b)
}

func Test123(t *testing.T) {
	a := []int{1,2,3}
	b := []int{1,3,2}
	AssertEqualA(t, NextPerm(a), b)
}

func Test132(t *testing.T) {
	a := []int{1,3,2}
	b := []int{3,1,2}
	AssertEqualA(t, NextPerm(a), b)
}

func Test321(t *testing.T) {
	a := []int{3,2,1}
	b := []int{1,2,3}
	AssertEqualA(t, NextPerm(a), b)
}

func AssertEqualA(t *testing.T, a, b []int) {
	AssertEqual(t, fmt.Sprintf("%v", a),fmt.Sprintf("%v", b))
}