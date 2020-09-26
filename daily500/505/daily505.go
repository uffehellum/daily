package main

import "fmt"

// Rotate an array in place n steps to the right

func main() {
	a := []int{}
	for i := 1; i <= 50; i++ {
		a = append(a, i)
	}
	a = rotate2(a, 49)
	fmt.Println(a)
}

func rotate2(a []int, r int) []int {
	i := len(a) - r
	return append(a[i:], a[:i]...)
}

func rotate(a []int, r int) {
	n := len(a)
	p0 := -1
	p := p0
	t := 0
	for i := n; i > 0; i-- {
		if p == p0 {
			p0++
			p = p0
			t = a[p]
		}
		p = (p + r) % n
		t, a[p] = a[p], t
	}
}
