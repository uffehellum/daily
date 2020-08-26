package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
// https://leetcode.com/problems/sort-list/

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	half := findHalf(head)
	for half != head {
		p1 := head
		p2 := half
		if p1.Next == half {
			half = p1
		}
		for p2 != nil {
			if p1.Val > p2.Val {
				p1.Val, p2.Val = p2.Val, p1.Val
			}
			p1 = p1.Next
			if p1.Next == half {
				half = p1
			}
			p2 = p2.Next
		}
	}
	return head
}

func findHalf(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			break
		}
		p2 = p2.Next
	}
	return p1
}

func println(head *ListNode) {
	fmt.Print("[")
	for head != nil {
		fmt.Print(" ", head.Val)
		head = head.Next
	}
	fmt.Println(" ]")
}

func main() {
	head := &ListNode{ 7, nil}
	head = &ListNode{ 9, head}
	head = &ListNode{ 13, head}
	println(head)
	head = sortList(head)
	println(head)
}
