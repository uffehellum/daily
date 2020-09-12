package main

import (
	"fmt"
	"testing"
)

func TestFindTreeDepth(t *testing.T) {
	assertFindTreeDepth(t, 1, "(00)")
	assertFindTreeDepth(t, 2, "((00)(00))")
	assertFindTreeDepth(t, 4, "((((00)0)0)0)")
}

func TestIsValidTree(t *testing.T) {
	assertValidTree(t, true, "0")
	assertValidTree(t, true, "(00)")
	assertValidTree(t, true, "((00)(00))")
	assertValidTree(t, true, "((((00)0)0)0)")
}

func assertValidTree(t *testing.T, expected bool, s string) {
	actual :=IsValidTree(s)
	if expected == actual {
		return
	}
	message := fmt.Sprintf("%s: got %v expected %v", s, actual, expected)
	t.Fatal(message)
}

func assertFindTreeDepth(t *testing.T, expected int, s string) {
	actual:=FindTreeDepth(s)
	if expected == actual {
		return
	}
	message := fmt.Sprintf("%s: got %v expected %v", s, actual, expected)
	t.Fatal(message)
}

func IsValidTree(s string) bool {
	s, ok := ParseValidTree(s)
	return ok && s == ""
}

func ParseValidTree(s string) (rest string, ok bool) {
	if s == "" {
		return s, false
	}
	if s[0] == '0' {
		return s[1:], true
	}
	if s[0] == '(' {
		s = s[1:]
		if s, ok = ParseValidTree(s); !ok {
			return s, ok
		}
		if s, ok = ParseValidTree(s); !ok {
			return s, ok
		}
		if s == "" {
			return s, false
		}
		if s[0] == ')' {
			return s[1:], true
		}
	}
	return s, false
}

func FindTreeDepth(s string) int {
	d := 0
	r := 0
	for _, c := range s {
		if c == '(' {
			d++
			if d > r {
				r = d
			}
		}
		if c == ')' {
			d--
		}
	}
	return r
}
