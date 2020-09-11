package main

import "fmt"

func main() {
	for _, x := range (Zigzag("thisisazigzag", 4)) {
		fmt.Println(x)
	}
}

func Zigzag(s string, k int) []string {
	a := make([]string, k)
	d := 1
	r := 0
	for _, c:= range s {
		for i := 0; i < k; i++ {
			if i == r {
				a[i] += string(c)
			} else {
				a[i] += string(' ')
			}
		}
		r += d
		if r < 0 || r >= k {
			d = -d
			r += 2 * d
		}
	}
	return a
}