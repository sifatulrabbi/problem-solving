package leetcode

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		table[nums[i]] = i + 1
	}
	for i := 0; i < len(nums); i++ {
		dif := target - nums[i]
		if j, ok := table[dif]; ok {
			result[0] = i + 1
			result[1] = j
			break
		}
	}
	return result
}
