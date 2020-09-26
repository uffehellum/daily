package daily530

import "testing"

func TestEditDistance(t *testing.T){
	assertDistance(t, "kitten", "sitting", 3)
	assertDistance(t, "smørrebrød", "pølsebrød", 4)
	assertDistance(t, "xxaaa", "aaa", 2)
	assertDistance(t, "xxa", "aaa", 2)
}

func assertDistance(t *testing.T, a string, b string, expected int) {
	actual := EditDistance(a, b)
	if actual != expected {
		t.Fatalf("Expected %d found %d for %s and %s",
			expected, actual, a, b)
	}
}

func EditDistance(a string, b string) int {
	ra, rb := []rune(a), []rune(b)
	j := 0
	n:= 0
	for i := range ra {
		switch {
		case j < len(rb) && ra[i] == rb[j]:
			j++
			continue
		case j + 1 < len(rb) && ra[i] == rb[j+1]:
			j+=2
			n++
		case i + 1 < len(ra) && ra[i+1] == rb[j]:
			n++
		default:
			n++
			j++
		}
	}
	return n + len(rb) - j
}
