package main

import (
	"fmt"

	"problem-solving/go/data"
)

func run_binary_search() {
	fmt.Printf("found at index: %d\n", binary_search(data.Get_sorted_large_array(50), 34))
}

func binary_search(arr []int, val int) int {
	i := -1
	lo := 0
	hi := len(arr) - 1

	for lo < hi {
		mid := lo + (hi-lo)/2
		if arr[mid] == val {
			i = mid
			break
		} else if arr[mid] > val {
			hi = mid
		} else if arr[mid] < val {
			lo = mid + 1
		}
	}

	return i
}
