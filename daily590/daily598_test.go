package daily590

import (
	"daily/daily590/daily598"
	"testing"
)

/*
You have n fair coins and you flip them all at the same time.
Any that come up tails you set aside.
The ones that come up heads you flip again.
How many rounds do you expect to play before only one coin remains?

Write a function that, given n, returns the number of rounds you'd
expect to play until one coin remains.
*/


func TestExpectedRounds(t *testing.T) {
	AssertEqual(t, 0, daily598.ExpectedRounds(1), "1 coins")
	AssertEqual(t, 1, daily598.ExpectedRounds(2), "2 coins")
	AssertEqual(t, 3, daily598.ExpectedRounds(7), "7 coins")
	AssertEqual(t, 8, daily598.ExpectedRounds(200), "200 coins")
}

func AssertEqual(t *testing.T, expected, actual int, message string) {
	if expected != actual {
		t.Fatal(message, "expected", expected, "actual", actual)
	}
}