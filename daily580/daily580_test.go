package daily580

import "testing"

type node struct{
	value int
	left, right *node
}

func TestTree(t *testing.T) {
	v := &node{
		10,
		&node{
			5,
			nil,
			&node{2, nil, nil}},
		&node{
			5,
			nil,
			&node{
				1,
				&node{-1, nil, nil},
				nil}}}
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