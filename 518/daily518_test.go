package main

import (
	"fmt"
	"testing"
)

func TestSumOfThree(t *testing.T) {
	assertSumOfThree(t, []int{}, 49, false)
	assertSumOfThree(t, []int{20, 303, 3, 4, 25}, 49, true)
	assertSumOfThree(t, []int{20, 303, 3, 5, 25}, 49, false)
}

func assertSumOfThree(t *testing.T, a []int, k int, expected bool){
	actual := hasSumOfThree(a, k)
	if actual != expected {
		t.Fatal(fmt.Sprintf(
			"hasSumOfThree(a=%v, k=%d) = %v expected %v",
			a, k, actual, expected))
	}
}
func hasSumOfThree(a []int, k int) bool {
	for i, x := range a {
		for j, y:= range a[i+1:] {
			for _, z := range a[j+1:] {
				if x+y+z == k{
					return true
				}
			}
		}
	}
	return false
}