package leetcode

import (
	"fmt"
)

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	if len(nums) < k {
		k = k % len(nums)
	}
	values := map[int]int{}
	idx := 0
	for k > 0 {
		values[idx] = nums[len(nums)-k]
		k--
		idx++
	}
	for i := 0; i < len(nums)-idx; i++ {
		values[idx+i] = nums[i]
	}
	fmt.Println(idx)
	for key, val := range values {
		nums[key] = val
	}
}
