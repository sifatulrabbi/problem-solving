package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func isSubsequence(s string, t string) bool {
	if len(s) < 1 {
		return true
	}
	j := 0
	for i := 0; i < len(t); i++ {
		if s[j] == t[i] {
			j++
		}
		if j == len(s) {
			break
		}
	}
	return j == len(s)
}

func TestIsSubsequence(t *testing.T) {
	assert.Equal(t, true, isSubsequence("abc", "ahbgdc"))
	assert.Equal(t, false, isSubsequence("acb", "ahbgdc"))
}
