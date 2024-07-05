package leetcode

import (
	"fmt"
)

func RunQuicksort() {
	arr := []int{70, 50, 40, 30, 80, 20, 10, 60, 90}
	quickSort(arr)
	fmt.Println(arr)
}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	// Choose a pivot (e.g., the middle element)
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	// Partition the slice into two sub-slices
	for left <= right {
		for arr[left] < pivot {
			left++
		}
		fmt.Printf("left: %d\n", arr[left])
		for arr[right] > pivot {
			right--
		}
		fmt.Printf("right: %d\n", arr[right])
		fmt.Printf("before: %v\n", arr)
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
		fmt.Printf("after : %v\n", arr)
	}

	// Recursively sort the sub-slices
	if right > 0 {
		quickSort(arr[:right+1])
	}
	if left < len(arr)-1 {
		quickSort(arr[left:])
	}
}
