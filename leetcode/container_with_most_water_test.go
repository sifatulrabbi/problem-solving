package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func maxArea(height []int) int {
	maxArea := 0
	ch := make(chan int, len(height))
	for i := 0; i < len(height); i++ {
		go func() {
			ch <- calcMaxArea(height, i)
		}()
	}
	for {
		area := <-ch
		if maxArea < area {
			maxArea = area
		}
	}
	return maxArea
}

func calcMaxArea(height []int, i int) int {
	maxArea := 0
	ch := make(chan int, len(height))
	for j := i + 1; j < len(height); j++ {
		s := height[j]
		area := 0
		if height[i] < s {
			area = (j - i) * height[i]
		} else {
			area = (j - i) * s
		}
		ch <- area
	}
	for {
		area := <-ch
		if maxArea < area {
			maxArea = area
		}
	}
	return maxArea
}

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 1, maxArea([]int{1, 1}))
}
