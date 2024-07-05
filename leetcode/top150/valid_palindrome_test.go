package leetcode

import (
	"regexp"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func isPalindrome(s string) bool {
	isvalid := true
	re := regexp.MustCompile("([a-z]|[0-9])*")
	trimmed := strings.Join(re.FindAllString(strings.ToLower(s), -1), "")
	lastidx := len(trimmed) - 1
	for i := 0; i <= lastidx; i++ {
		if trimmed[i] != trimmed[lastidx-i] {
			isvalid = false
			break
		}
	}
	return isvalid
}

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome("A man, a plan, a canal: Panama"))
}
