package hackerrank

import (
	"fmt"
	"testing"
)

func TestFlippingMatrix(t *testing.T) {
	matrix := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)
	sum := flippingMatrix(matrix)
	fmt.Println(sum, matrix)
}

func TestPlusMinusZeroRatio(t *testing.T) {
	t.Skip()
	plusMinus([]int32{1, 1, 0, -1, -1})
}

func TestMiniMaxSum(t *testing.T) {
	t.Skip()
	miniMaxSum([]int32{1, 3, 5, 7, 9})
}

func TestTimeConversion(t *testing.T) {
	t.Skip()
	fmt.Println(timeConversion("12:01:00AM"))
	fmt.Println(timeConversion("12:01:00PM"))
	fmt.Println(timeConversion("07:05:45PM"))
	fmt.Println(timeConversion("07:01:00AM"))
}
