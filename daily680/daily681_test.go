package daily680

import (
	"fmt"
	"testing"
)

var snakes = map[int]int{16: 6, 48: 26, 49: 11, 56: 53, 62: 19, 64: 60, 87: 24, 93: 73, 95: 75, 98: 78}
var ladders = map[int]int{1: 38, 4: 14, 9: 31, 21: 42, 28: 84, 36: 44, 51: 67, 71: 91, 80: 100}

const size = 100

func TestDaily681(t *testing.T) {
	actual := findFewestSteps()
	expected := 7
	AssertEqual(t, actual, expected, "steps needed")
}

func AssertEqual(t *testing.T, actual int, expected int, message string) {
	if actual != expected {
		t.Fatal("expected", expected, "actual", actual, message)
	}
}

func findFewestSteps() int {
	visited := map[int]int{1: 0}
	queue := []int{1}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		steps := visited[p] + 1
		for i := p + 1; i <= p+6; i++ {
			x := i
			if y, ok := snakes[x]; ok {
				x = y
			}
			if y, ok := ladders[x]; ok {
				x = y
			}
			if x == size {
				fmt.Println(queue)
				fmt.Println(visited)
				return steps
			}
			if _, ok := visited[x]; !ok {
				visited[x] = steps
				queue = append(queue, x)
			}
		}
	}
	return -1
}
