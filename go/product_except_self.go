package main

import (
	"fmt"
)

func RunProductExceptSelf() {
	// fmt.Println(productExceptSelf(data.ArrWith1s))
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

// O(n^2)
func productExceptSelf(nums []int) []int {
	l := len(nums)
	arr := make([]int, len(nums))
	arr[0], arr[l-1] = 1, 1
	for i := 1; i < l; i++ {
		arr[i] = arr[i-1] * nums[i-1]
	}

	rp := 1
	for i := l - 2; i >= 0; i-- {
		rp *= nums[i+1]
		arr[i] *= rp
	}
	return arr
}
