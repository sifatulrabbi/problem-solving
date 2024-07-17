package hackerrank

import (
	"fmt"
	"math"
	"strconv"
)

func plusMinus(arr []int32) {
	total := float32(len(arr))
	pos := 0
	neg := 0
	zero := 0
	for _, v := range arr {
		if v == 0 {
			zero++
		} else if v < 0 {
			neg++
		} else {
			pos++
		}
	}
	fmt.Printf("%.6f\n", float32(pos)/total)
	fmt.Printf("%.6f\n", float32(neg)/total)
	fmt.Printf("%.6f\n", float32(zero)/total)
}

func miniMaxSum(arr []int32) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	var (
		sum1    int32
		sum2    int32
		lastidx = len(arr) - 1
	)
	for i := 0; i < 4; i++ {
		sum1 += arr[i]
		sum2 += arr[lastidx-i]
	}
	fmt.Println(sum1, sum2)
}

func timeConversion(s string) string {
	hr24 := "00"
	_hr12, _ := strconv.ParseInt(s[:2], 10, 32)
	hr12 := int(_hr12)
	if s[8:10] == "AM" {
		if hr12 == 12 {
			hr24 = "00"
		} else if hr12 < 10 {
			hr24 = fmt.Sprintf("0%d", hr12)
		} else {
			hr24 = fmt.Sprintf("%d", hr12)
		}
	} else if hr12 == 12 {
		hr24 = "12"
	} else {
		hr24 = fmt.Sprintf("%d", 12+hr12)
	}
	return fmt.Sprintf("%s:%s:%s", hr24, s[3:5], s[6:8])
}

func findMedian(arr []int32) int32 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr[len(arr)-1]
}

func lonelyinteger(a []int32) int32 {
	// Write your code here
	table := map[int32]int{}
	for _, v := range a {
		if _, ok := table[v]; !ok {
			table[v] = 1
		} else {
			table[v]++
		}
	}
	var val int32
	for v, c := range table {
		if c == 1 {
			val = v
			break
		}
	}
	return val
}

func diagonalDifference(arr [][]int32) int32 {
	var (
		sideWidth         = len(arr)
		leftToRight int32 = 0
		rightToLeft int32 = 0
	)
	for i := 0; i < sideWidth; i++ {
		l, r := i, (sideWidth-1)-i
		leftToRight += arr[i][l]
		rightToLeft += arr[i][r]
	}
	return int32(math.Abs(float64(leftToRight) - float64(rightToLeft)))
}

func countingSort(arr []int32) []int32 {
	return arr
}

func flippingMatrix(matrix [][]int32) int32 {
	length := len(matrix)

	for col := 0; col < length; col++ {
		if matrix[0][col] < matrix[length-1][col] {
			colReverse(matrix, col)
		}
	}

	var sum int32 = 0
	for _, v := range matrix[0] {
		sum += v
	}
	return sum
}

func rowReverse(matrix [][]int32, row int) {
	l := len(matrix[0]) - 1
	for j := 0; j < l/2; j++ {
		matrix[row][j], matrix[row][l-j] = matrix[row][l-j], matrix[row][j]
	}
}

func colReverse(matrix [][]int32, col int) {
	l := len(matrix) - 1
	for i := 0; i < l/2; i++ {
		matrix[i][col], matrix[l-i][col] = matrix[l-i][col], matrix[i][col]
	}
}

func findZigZagSquence(arr []int) {
	k := (len(arr) + 1) / 2
}
