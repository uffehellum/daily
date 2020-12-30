package euler010

import "testing"

func TestCountSundaysOnFirstOfMonthYearInterval(t *testing.T) {
	AssertEqual(t, CountSundaysOnFirstOfMonthYearInterval(1901, 2000), 171)
}

func CountSundaysOnFirstOfMonthYearInterval(startYear, endYear int) int {
	sum := 0
	for year := startYear; year <= endYear; year++ {
		sum += CountSundaysOnFirstOfMonthOneYear(year)
	}
	return sum
}

func CountSundaysOnFirstOfMonthOneYear(year int) int {
	n := 0
	for month := 1; month <= 12; month++ {
		if IsSunday(year, month, 1) {
			n++
		}
	}
	return n
}

func IsSunday(year, month, day int) bool {
	return Weekday(year, month, day) == 0
}

var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func Weekday(year int, month int, day int) int {
	w := WeekdayJanFirst(year)
	leap := IsLeapYear(year)
	for m := 1; m < month; m++ {
		w += monthDays[m - 1]
	}
	if leap && month > 2 {
		w++
	}
	w += day - 1
	return w % 7
}

func WeekdayJanFirst(year int) int {
	w := 1
	for y := 1900; y < year; y++ {
		w += 365
		if IsLeapYear(y) {
			w += 1
		}
	}
	return w % 7
}

func IsLeapYear(year int) bool {
	return year % 400 == 0 || year % 100 > 0 && year % 4 == 0
}
