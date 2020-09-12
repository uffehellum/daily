package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(indexes( "ægøægægæ", "æg"))
}

func indexes(s, pattern string) []int {
	a := []int{}
	p := 0
	for {
		i := strings.Index(s[p:], pattern)
		if i < 0 {
			break;
		}
		p += i
		a = append(a, len([]rune(s[:p])))
		p += len(pattern)
	}
	return a
}
