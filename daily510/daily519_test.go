package daily510

import (
	"fmt"
	"testing"
)

func TestBitOps(t *testing.T) {
	assertBitOps(t, 0, 0, 0, 0)
	assertBitOps(t, 7, 8, 1, 7)
	assertBitOps(t, 7, 8, 0, 8)
	assertBitOps(t, 1234566, 8765432, 1, 1234566)
	assertBitOps(t, 1234566, 8765432, 0, 8765432)
}

func assertBitOps(t *testing.T,x, y, b, expected int) {
	actual := BitOps(x, y, b)
	if actual != expected {
		t.Fatal(fmt.Sprintf(
			"BitOps(x=%d, y=%d, b=%b) = %d expected %d",
			x, y, b, actual, expected))
	}
}

func BitOps(x, y, b int) int {
	return x & -b | y & (b - 1)
	// return x * b + y - y * b
}