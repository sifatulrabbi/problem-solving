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

// TODO:
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	p := ""
	for i := 0; i < len(s); i++ {
		initial := []byte{s[i]}
		for j := i + 1; j < len(s); j++ {
		}
	}
	return p
}
