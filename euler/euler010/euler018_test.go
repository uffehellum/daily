package euler010

import (
	"strconv"
	"strings"
	"testing"
)


/*
Maximum path sum I
Submit

 Show HTML problem content
Problem 18
By starting at the top of the triangle below and moving to adjacent numbers on the row below,
the maximum total from top to bottom is 23.

   3
  7 4
 2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

   0
  1 2
 3 4 5
6 7 8 9

   1
  2 3
 4 5 6
7 8 9 10
Find the maximum total from top to bottom of the triangle below:
*/

const largeTriangle = `
              75
             95 64
            17 47 82
           18 35 87 10
          20 04 82 47 65
         19 01 23 75 03 34
        88 02 77 73 07 63 67
       99 65 04 28 06 16 70 92
      41 41 26 56 83 40 80 70 33
     41 48 72 33 47 32 37 16 94 29
    53 71 44 65 25 43 91 52 97 51 14
   70 11 33 28 77 73 17 78 39 68 17 57
  91 71 52 38 17 14 91 43 58 50 27 29 48
 63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

/*
NOTE: As there are only 16384 routes, it is possible to solve this problem by
trying every route. However, Problem 67, is the same challenge with a triangle
containing one-hundred rows; it cannot be solved by brute force, and requires
a clever method! ;o)
 */

const smallTriangle  = `
   3
  7 4
 2 4 6
8 5 9 3
`

func TestLoadTriangle(t *testing.T){
	small := LoadTriangle(smallTriangle)
	AssertEqual(t, len(small), 10)
	AssertEqual(t, small[2], 4)
	large:= LoadTriangle(largeTriangle)
	AssertEqual(t, len(large), 120)
	AssertEqual(t, large[119], 23)
	AssertEqual(t, large[110], 9)
}

func LoadTriangle(src string) []int {
	a := []int{}
	for _, s := range strings.Fields(src) {
		i, _ := strconv.ParseInt(s, 10, 32)
		a = append(a, int(i))
	}
	return a
}

func TestMaximumPathSum(t *testing.T) {
	AssertEqual(t, MaximumPathSum(LoadTriangle(smallTriangle)), 23)
	AssertEqual(t, MaximumPathSum(LoadTriangle(largeTriangle)), 1074)
}

func MaximumPathSum(triangle []int) int {
	rowcount := TriangleRowCount(triangle)
	for row := rowcount - 2; row >= 0; row-- {
		top := TriangleRowPosition(row)
		bottom := TriangleRowPosition(row + 1)
		for col := 0; col <= row; col++ {
			a := triangle[bottom + col]
			b := triangle[bottom + col + 1]
			if a < b {
				a = b
			}
			triangle[top + col] += a
		}
	}
	return triangle[0]
}

func TestTriangleRowCount(t *testing.T) {
	AssertEqual(t, TriangleRowCount(make([]int, 0)), 0)
	AssertEqual(t, TriangleRowCount(make([]int, 1)), 1)
	AssertEqual(t, TriangleRowCount(LoadTriangle(smallTriangle)), 4)
	AssertEqual(t, TriangleRowCount(LoadTriangle(largeTriangle)), 15)
}

func TriangleRowCount(triangle []int) int {
	row := 0
	for TriangleRowPosition(row) < len(triangle) {
		row++
	}
	return row
}

func TriangleRowPosition(row int) int {
	return row * (row + 1) / 2
}
//
//func TriangleGet(triangle []int, row, col int) int {
//	return triangle[row * (row + 1) / 2 + col]
//}
//
//func TrianglePut(triangle []int, row, col, val int) {
//	triangle[row * (row + 1) / 2 + col] = val
//}
//
//func LoadTriangleScanner(s string) []int {
//	sc := scanner.Scanner{ Mode: scanner.ScanStrings }
//	sc.Init(strings.NewReader(s))
//	a := []int{}
//	for tok := sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
//		i, _ := strconv.ParseInt(sc.TokenText(), 10, 32)
//		a = append(a, int(i))
//	}
//	return a
//}