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

func findPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	p := string(s[0])
	for i := 0; i < len(s); i++ {
		initial := []byte{s[i]}
		for j := i + 1; j < len(s); j++ {
			initial = append(initial, s[j])

			valid := true
			for k := 0; k < len(initial); k++ {
				if initial[k] != initial[len(initial)-1-k] {
					valid = false
					break
				}
			}

			if valid && len(initial) > len(p) {
				p = string(initial)
			}
		}
	}
	return p
}

func longestPlindrome(s string) string {
	return findPalindrome(s)
}
