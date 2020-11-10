package daily570

import (
	"fmt"
	"strings"
	"testing"
)

func TestHanoi(t *testing.T) {
	AssertEqual(t, sHanoi(1),"1 to 3\n")
	AssertEqual(t, sHanoi(2),"1 to 2\n1 to 3\n2 to 3\n")
}

func AssertEqual(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Fatal("expected", expected, "found", actual)
	}
}

func sHanoi(n int) string {
	sb := &strings.Builder{}
	rHanoi(sb, n, 1, 2, 3)
	return sb.String()
}

func rHanoi(sb *strings.Builder, n, a, b, c int) {
	if n > 1 {
		rHanoi(sb, n - 1, a, c, b)
	}
	sb.WriteString(fmt.Sprintf("%d to %d\n", a, c))
	if n > 1 {
		rHanoi(sb, n - 1, b, a, c)
	}
}