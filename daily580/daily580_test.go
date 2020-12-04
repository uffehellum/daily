package daily580

import "testing"

type node struct{
	value int
	left, right *node
}

func TestTree(t *testing.T) {
	v := &node{
		value: 10,
		left: &node{
			value: 5,
			right: &node{value: 2}},
		right: &node{
			value: 5,
			right:&node{
				value: 1,
				left: &node{value: -1}}}}
	actual := minPath(v)
	assertEqual(t, 15, actual)
}

func minPath(n *node) int {
	if n==nil {
		return 0
	}
	x := minPath(n.left)
	y := minPath(n.right)
	if x < y {
		return n.value + x
	}
	return n.value + y
}

func assertEqual(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Fatal("expected", expected, "actual", actual)
	}
}