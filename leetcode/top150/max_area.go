package leetcode

func maxArea(height []int) int {
	maxArea := 0
	curr := 0
	left, right := 0, len(height)-1
	for left < right {
		curr = min(height[left], height[right]) * (right - left)
		maxArea = max(curr, maxArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
