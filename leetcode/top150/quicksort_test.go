package leetcode

import (
	"fmt"
	"testing"
)

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)
	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	idx := lo - 1
	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			arr[idx], arr[i] = arr[i], arr[idx]
		}
	}

	idx++
	arr[hi], arr[idx] = arr[idx], pivot
	return idx
}

func TestQuicksort(t *testing.T) {
	arr := []int{9, 3, 4, 2, 6, 7, 5, 8}
	quickSort(arr)
	fmt.Println(arr)

	arr = []int{71, 15, 94, 32, 81, 29, 41, 36, 79}
	quickSort(arr)
	fmt.Println(arr)
}
