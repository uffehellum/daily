package daily580

import "testing"

type point struct {
	x, y int
}
type rect struct {
	topleft point
	dimensions point
}

func overlappingArea(r1, r2 rect) int {
	dx := overlappingDistance(
		r1.topleft.x,
		r1.topleft.x + r1.dimensions.x,
		r2.topleft.x,
		r2.topleft.x + r2.dimensions.x)
	dy := overlappingDistance(
		r1.topleft.y,
		r1.topleft.y + r1.dimensions.y,
		r2.topleft.y,
		r2.topleft.y + r2.dimensions.y)
	if dx < 0 || dy < 0 {
		return 0
	}
	return dx * dy
}

func overlappingDistance(min1, max1, min2, max2 int) int {
	if min1 < min2 {
		min1 = min2
	}
	if max1 > max2 {
		max1 = max2
	}
	return max1 - min1
}

func TestDaily581(t *testing.T){
	r1:= rect {
		topleft: point{1, 4},
		dimensions: point{3, 3},
	}
	r2 := rect{
		topleft: point{0, 5},
		dimensions: point{4, 3},
	}
	actual := overlappingArea(r1, r2)
	assertEqual(t, 6, actual)
}
