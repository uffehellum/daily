package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
// https://leetcode.com/problems/sort-list/

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var half *ListNode;
	for {
		half = findHalf(head, half)
		if half == head {
			break
		}
		p1 := head
		p2 := half
		for p2 != nil {
			if p1.Val > p2.Val {
				p1.Val, p2.Val = p2.Val, p1.Val
			}
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return head
}

func findHalf(head, target *ListNode) *ListNode {
	i := 0
	for p := head; p != target; p = p.Next {
		i++
	}
	if i == 1 {
		return head
	}
	for i = (i + 1) / 2; i != 0; i-- {
		head = head.Next
	}
	return head
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
