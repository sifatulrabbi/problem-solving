package blind75

func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var (
		table  = make(map[int]bool)
		maxLen = 1
		curr   = 1
	)

	for i := 0; i < len(nums); i++ {
		table[nums[i]] = true
	}

	for k := range table {
		c := k + 1
		_, after := table[c]
		for after {
			curr++
			c++
			_, after = table[c]
			if maxLen < curr {
				maxLen = curr
			}
		}
		curr = 1
	}

	return maxLen
}
