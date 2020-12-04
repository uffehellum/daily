package daily600

import (
	"math/cmplx"
	"testing"
)

func TestClosestPoints(t *testing.T){
	a, b := ClosestTwoPoints([]complex128{1+1i, -1-1i, 3+4i, 6+1i, -1-6i, -4-3i})
	AssertEqual(t, a, -1-1i, "a")
	AssertEqual(t, b, 1+1i, "b")
}

func AssertEqual(t *testing.T, expected complex128, actual complex128, message string) {
	if !AboutEqual(expected, actual)  {
		t.Fatal("expected", expected, "actual", actual, message)
	}
}

func AboutEqual(expected complex128, actual complex128) bool {
	return cmplx.Abs(expected - actual) / (cmplx.Abs(expected) + 1) < 1e-17
}

func ClosestTwoPoints(points []complex128) (a, b complex128) {
	a, b = points[0], points[1]
	best := cmplx.Abs(a - b)
	for i, x := range points[:len(points)-1] {
		for _, y := range points[i+1:]{
			d := cmplx.Abs(x - y)
			if d < best {
				best = d
				a = x
				b = y
			}
		}
	}
	if real(a) > real(b) || real(a) == real(b) && imag(a) > imag(b) {
		a, b = b, a
	}
	return
}
