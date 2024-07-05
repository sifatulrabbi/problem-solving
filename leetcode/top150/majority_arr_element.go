package leetcode

func majorityElement(nums []int) int {
	appearances := map[int]int{}
	for _, num := range nums {
		appearances[num]++
	}
	result := 0
	n := 0
	for num, appeared := range appearances {
		if n < appeared {
			n = appeared
			result = num
		}
	}
	return result
}
