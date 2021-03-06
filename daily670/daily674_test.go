package daily670

import (
	"fmt"
	"testing"
)
/*
This problem was asked by Google.

A girl is walking along an apple orchard with a bag in each hand. She likes to pick apples from each tree as she goes along, but is meticulous about not putting different kinds of apples in the same bag.

Given an input describing the types of apples she will pass on her path, in order, determine the length of the longest portion of her path that consists of just two types of apple trees.

For example, given the input [2, 1, 2, 3, 3, 1, 3, 5], the longest portion will involve types 1 and 3, with a length of four.


*/

func TestDaily674(t *testing.T){
	AssertEqual(t, 4, LongestStreakOfTwoKindsOfApples, []int{1,1,2,2})
	AssertEqual(t, 4, LongestStreakOfTwoKindsOfApples, []int{2, 1, 2, 3, 3, 1, 3, 5})
}

func AssertEqual(t *testing.T, expected int, apples func(trees []int) int, input []int) {
	actual := apples(input)
	if actual != expected {
		t.Fatal("Expected=", expected, "actual=", actual, "input=", input )
	}
}

func LongestStreakOfTwoKindsOfApples(trees []int) int {
	if len(trees) < 3 {
		return len(trees)
	}
	p1 := 0
	p2 := 0
	v1 := trees[p1]
	v2 := trees[p2]
	longest := 2
	for p, v := range trees {
		fmt.Println(p1, v1, p2, v2, p, v)
		if v != v1 {
			if v != v2 {
				v1, v2 = v, v1
				p1, p2 = p, p1
			} else {
				v1, v2 = v2, v1
				p1 = p
			}
		}
		l := p - p2 + 1
		if l > longest {
			longest = l
		}
	}
	return longest
}