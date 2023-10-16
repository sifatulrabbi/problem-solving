package main

import "fmt"

func RunJumpGame() {
	arr := []int{2, 3, 1, 1, 4} // true
	fmt.Println(canJump(arr), arr)
	arr = []int{3, 2, 1, 0, 4} // false
	fmt.Println(canJump(arr), arr)
	arr = []int{2, 3, 1, 1, 4} // true
	fmt.Println(canJump(arr), arr)
	arr = []int{2, 0} // true
	fmt.Println(canJump(arr), arr)
	arr = []int{2, 0, 0} // true
	fmt.Println(canJump(arr), arr)
	arr = []int{3, 0, 8, 2, 0, 0, 1} // true
	fmt.Println(canJump(arr), arr)
}

func canJump(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return true
	}
	jump := 0
	for i := 0; i < n; i++ {
		if j := nums[i] + i; j > jump {
			jump = j
		}
		if jump >= n-1 {
			break
		} else if nums[i] == 0 && i == jump {
			return false
		}
	}
	return true
}
