package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test525(t *testing.T) {
	assertSpiral(t, 0, 0, []int{})
	assertSpiral(t, 1, 1, []int{1})
	assertSpiral(t, 2, 2, []int{1, 2, 4, 3})
	assertSpiral(t, 2, 3, []int{1, 2, 3, 6, 5, 4})
	assertSpiral(t, 4, 5, []int{
		1, 2, 3, 4, 5,
		10, 15,
		20, 19, 18, 17, 16,
		11,
		6, 7, 8, 9,
		14, 13, 12})
}

func assertSpiral(t *testing.T, n, m int, expected []int) {
	actual := Spiral(n, m)
	a, _ := json.Marshal(actual)
	e, _ := json.Marshal(expected)
	if string(a) != string(e) {
		t.Fatal(fmt.Sprintf("Spiral(n=%d, m=%d) = %v, expected=%v",
			n, m, actual, expected))
	}
}

func Spiral(n int, m int) []int {
	matrix := Matrix(n, m)
	return ClockwiseSpiral(matrix)
}

func ClockwiseSpiral(a [][]int) []int {
	r := []int{}
	if len(a) == 0 {
		return []int{}
	}
	top, bottom := 0, len(a)-1
	left, right := 0, len(a[0])-1
	dy, dx := 0, 1
	y, x := 0, 0
	for left <= right && top <= bottom {
		r = append(r, a[y][x])
		switch {
		case x+dx > right:
			top++
			dy, dx = dx, -dy
		case y+dy > bottom:
			right--
			dy, dx = dx, -dy
		case x+dx < left:
			bottom--
			dy, dx = dx, -dy
		case y+dy < top:
			left++
			dy, dx = dx, -dy
		}
		x += dx
		y += dy
	}
	return r
}

func Matrix(n int, m int) [][]int {
	a := make([][]int, n)
	x := 1
	for i := 0; i < n; i++ {
		r := make([]int, m)
		for j := 0; j < m; j++ {
			r[j] = x
			x++
		}
		a[i] = r
	}
	return a
}
