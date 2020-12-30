package euler010

import "testing"

/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there
are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words,
how many letters would be used?


NOTE: Do not count spaces or hyphens.
For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters.
The use of "and" when writing out numbers is in compliance with British usage.
 */

func TestCountOnes(t *testing.T) {
	AssertEqual(t, CountLettersInOnes(0), 0)
	AssertEqual(t, CountLettersInOnes(1), 3)
	AssertEqual(t, CountLettersInOnes(2), 3)
	AssertEqual(t, CountLettersInOnes(3), 5)
	AssertEqual(t, CountLettersInOnes(4), 4)
	AssertEqual(t, CountLettersInOnes(5), 4)
	AssertEqual(t, CountLettersInOnes(6), 3)
	AssertEqual(t, CountLettersInOnes(7), 5)
	AssertEqual(t, CountLettersInOnes(8), 5)
	AssertEqual(t, CountLettersInOnes(9), 4)
}

func TestCountTens(t *testing.T) {
	AssertEqual(t, CountLettersInTens(0), 0)
	AssertEqual(t, CountLettersInTens(1), 3)
	AssertEqual(t, CountLettersInTens(2), 6)
	AssertEqual(t, CountLettersInTens(3), 6)
	AssertEqual(t, CountLettersInTens(4), 5)
	AssertEqual(t, CountLettersInTens(5), 5)
	AssertEqual(t, CountLettersInTens(6), 5)
	AssertEqual(t, CountLettersInTens(7), 7)
	AssertEqual(t, CountLettersInTens(8), 6)
	AssertEqual(t, CountLettersInTens(9), 6)
}

func TestCountLess50(t *testing.T) {
	AssertEqual(t, CountLettersInOneNumber(0), 0)
	AssertEqual(t, CountLettersInOneNumber(1), 3)
	AssertEqual(t, CountLettersInOneNumber(2), 3)
	AssertEqual(t, CountLettersInOneNumber(3), 5)
	AssertEqual(t, CountLettersInOneNumber(4), 4)
	AssertEqual(t, CountLettersInOneNumber(5), 4)
	AssertEqual(t, CountLettersInOneNumber(6), 3)
	AssertEqual(t, CountLettersInOneNumber(7), 5)
	AssertEqual(t, CountLettersInOneNumber(8), 5)
	AssertEqual(t, CountLettersInOneNumber(9), 4)
	AssertEqual(t, CountLettersInOneNumber(10), 3)
	AssertEqual(t, CountLettersInOneNumber(11), 6)
	AssertEqual(t, CountLettersInOneNumber(12), 6)
	AssertEqual(t, CountLettersInOneNumber(13), 8)
	AssertEqual(t, CountLettersInOneNumber(14), 8)
	AssertEqual(t, CountLettersInOneNumber(15), 7)
	AssertEqual(t, CountLettersInOneNumber(16), 7)
	AssertEqual(t, CountLettersInOneNumber(17), 9)
	AssertEqual(t, CountLettersInOneNumber(18), 8)
	AssertEqual(t, CountLettersInOneNumber(19), 8)
}


func TestCountHundreds(t *testing.T) {
	AssertEqual(t, CountLettersInHundreds(0), 0)
	AssertEqual(t, CountLettersInHundreds(1), 10)
	AssertEqual(t, CountLettersInHundreds(2), 10)
	AssertEqual(t, CountLettersInHundreds(3), 12)
	AssertEqual(t, CountLettersInHundreds(4), 11)
	AssertEqual(t, CountLettersInHundreds(5), 11)
	AssertEqual(t, CountLettersInHundreds(6), 10)
	AssertEqual(t, CountLettersInHundreds(7), 12)
	AssertEqual(t, CountLettersInHundreds(8), 12)
	AssertEqual(t, CountLettersInHundreds(9), 11)
}

func TestCountThousands(t *testing.T) {
	AssertEqual(t, CountLettersInThousands(0), 0)
	AssertEqual(t, CountLettersInThousands(1), 11)
}

func TestCountLettersInInterval(t *testing.T){
	AssertEqual(t, CountLettersInInterval(1, 1), 3)
	AssertEqual(t, CountLettersInInterval(1, 2), 6)
	AssertEqual(t, CountLettersInInterval(1, 3), 11)
	AssertEqual(t, CountLettersInInterval(1, 5), 19)
	AssertEqual(t, CountLettersInInterval(1, 1000), 21124)

}

func TestCountLettersIn342(t *testing.T) {
	AssertEqual(t, CountLettersInOneNumber(300), 12 )
	AssertEqual(t, CountLettersInOneNumber(42), 8 )
	AssertEqual(t, CountLettersInOneNumber(342), 23 )
}

func TestCountLettersIn115(t *testing.T) {
	AssertEqual(t, CountLettersInOneNumber(100), 10 )
	AssertEqual(t, CountLettersInOneNumber(15), 7 )
	AssertEqual(t, CountLettersInOneNumber(115), 20 )
}

func TestCountLettersInOneNumber(t *testing.T) {
	AssertEqual(t, CountLettersInOneNumber(1000), 11 )
}

func CountLettersInInterval(a, b int) int {
	s := 0
	for i := a; i <= b; i++ {
		s += CountLettersInOneNumber(i)
	}
	return s
}

func CountLettersInOneNumber(n int) int {
	s := 0
	s += CountLettersInThousands( n / 1000)
	s += CountLettersInHundreds( n / 100 % 10)
	s += CountLettersInOnesAndTens(n % 100)
	s += CountLettersInAnd(n)
	return s
}

func CountLettersInOnesAndTens(n int) int {
	switch {
	case n < 10:
		return CountLettersInOnes(n)
	case n < 20:
		return CountLettersInTeens(n)
	default:
		return CountLettersInTens(n/10) + CountLettersInOnes(n%10)
	}
}


func CountLettersInAnd(n int) int {
	if n % 100 > 0 && n / 100 > 0 {
		return 3
	}
	return 0
}

func CountLettersInOnes(n int) int {
	return []int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4}[n]
}

func CountLettersInTens(n int) int {
	return []int{0, 3, 6, 6, 5, 5, 5, 7, 6, 6}[n]
}

func CountLettersInTeens(n int) int {
	return []int{3, 6, 6, 8, 8, 7, 7, 9, 8, 8}[n - 10]
}

func CountLettersInHundreds(n int) int {
	if n == 0 {
		return 0
	}
	return 7 + CountLettersInOnes(n)
}

func CountLettersInThousands(n int) int {
	if n == 0 {
		return 0
	}
	return 8 + CountLettersInOnes(n)
}
