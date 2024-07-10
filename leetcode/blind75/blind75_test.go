package blind75

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLongestConsecutive(t *testing.T) {
	assert.Equal(t, 4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	assert.Equal(t, 9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	assert.Equal(t, 4, longestConsecutive([]int{100, 4, 200, 56, 89, 1, 22, 3, 2, 89, 35}))
}

func TestLongestValidPalindrome(t *testing.T) {
	testcases := []string{"abac", "bbbadcasd"}
	for _, v := range testcases {
		fmt.Printf("%s: %s\n", v, longestPalindrome(v))
	}
}
