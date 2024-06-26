package leetcode

import (
	"fmt"
	"testing"
)

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 3 {
		return []int{1, 2}
	}

	var (
		idx1 int
		idx2 int
	)
	for i := 0; i < len(numbers); i++ {
		tar := target - numbers[i]
		if idx := binarySearch(numbers[i+1:], tar); idx > -1 {
			idx1, idx2 = i+1, i+2+idx
			break
		}
	}
	return []int{idx1, idx2}
}

func binarySearch(arr []int, val int) int {
	var (
		i  = -1
		lo = 0
		hi = len(arr) - 1
	)
	for lo < hi {
		if arr[lo] == val {
			i = lo
			break
		}
		if arr[hi] == val {
			i = hi
			break
		}

		mid := (lo + hi) / 2
		if arr[mid] == val {
			i = mid
			break
		}

		if arr[mid] > val {
			hi = mid
		} else if arr[mid] < val {
			lo = mid + 1
		}
	}
	if lo == hi {
		if arr[lo] == val {
			i = lo
		}
	}
	return i
}

func TestTwoSum(t *testing.T) {
	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	// fmt.Println(twoSum([]int{2, 3, 4}, 6))
	// fmt.Println(twoSum([]int{-1, 0}, -1))
	// fmt.Println(twoSum([]int{5, 25, 75}, 100))
	fmt.Println(twoSum([]int{3, 24, 50, 79, 88, 150, 345}, 200))

	// assert.Equal(t, 4, binarySearch([]int{1, 2, 3, 4, 5, 6, 7}, 5))
	// assert.Equal(t, 7, binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 8))
	// assert.Equal(t, 2, binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3))
	// assert.Equal(t, 0, binarySearch([]int{1, 2, 3, 4}, 1))
}
