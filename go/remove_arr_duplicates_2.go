package main

import "fmt"

func RunRemoveDuplicates2() {
	arr := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(arr), arr)

	arr = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(arr), arr)
}

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return l
	}

	k := 2
	for i := 2; i < l; i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
