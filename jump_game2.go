package main

import (
	"fmt"
)

/*
You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].
*/
func RunJumpGame2() {
	arr := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(arr), arr)
	arr = []int{2, 3, 0, 1, 4}
	fmt.Println(jump(arr), arr)
	arr = []int{1, 2, 3}
	fmt.Println(jump(arr), arr)
	arr = []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	fmt.Println(jump(arr), arr)
	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0} // 2
	fmt.Println(jump(arr), arr)
}

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	position, furthestIdx, jumps := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		// if the combination fo the current index and current jumps is greater or equal to the last index then the calc is over.
		if i+nums[i] >= len(nums)-1 {
			jumps++
			return jumps
		}
		// if the combination of current index and current jump count is more than current furthest point, and the jump count of that position is not 0 then update the furthest point.
		if i+nums[i] > furthestIdx {
			furthestIdx = i + nums[i]
		}
		// if the current index is similar to our current position then we need to make a jump.
		if position == i {
			position = furthestIdx
			jumps++
		}
	}
	return 0
}
