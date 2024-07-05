package leetcode

import "fmt"

func RunMoveZeroes() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)

	arr = []int{0, 0, 1}
	moveZeroes(arr)
	fmt.Println(arr)
}

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != 0 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == 0 {
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
