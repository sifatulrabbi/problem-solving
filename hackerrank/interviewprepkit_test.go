package hackerrank

import (
	"fmt"
	"strconv"
	"testing"
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

func TestPlusMinusZeroRatio(t *testing.T) {
	plusMinus([]int32{1, 1, 0, -1, -1})
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

func TestMiniMaxSum(t *testing.T) {
	miniMaxSum([]int32{1, 3, 5, 7, 9})
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

func TestTimeConversion(t *testing.T) {
	fmt.Println(timeConversion("12:01:00AM"))
	fmt.Println(timeConversion("12:01:00PM"))
	fmt.Println(timeConversion("07:05:45PM"))
	fmt.Println(timeConversion("07:01:00AM"))
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
