package leetcode

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	var arr []byte
	for i := 0; i < len(s); i++ {
		char := s[i]
		if idx := includes(arr, char); idx > -1 {
			if idx+1 < len(arr) {
				arr = arr[idx+1:]
			} else {
				arr = []byte{}
			}
		}
		arr = append(arr, char)
		if len(arr) > maxLen {
			maxLen = len(arr)
		}
	}
	return maxLen
}

func includes(arr []byte, v byte) int {
	idx := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			idx = i
			break
		}
	}
	return idx
}
