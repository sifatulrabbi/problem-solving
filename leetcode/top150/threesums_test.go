package leetcode

import (
	"fmt"
	"testing"
)

func threeSum(nums []int) [][]int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	triplets := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1
		for j < k {
			total := nums[i] + nums[j] + nums[k]
			if total > 0 {
				k--
			} else if total < 0 {
				j++
			} else {
				triplets = append(triplets, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return triplets
}

func TestThreeSums(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
