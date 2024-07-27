package leetcode

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

func getConcatenation(arr []int) []int {
	return append(arr, arr...)
}

func validParenthesis(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	table := map[string]int{
		"[": 3,
		"]": 3,
		"{": 2,
		"}": 2,
		"(": 1,
		")": 1,
	}

	valid := true
	last := len(s) - 1
	for i := 0; i <= last; i++ {
		if table[string(s[i])] != table[string(s[last-i])] {
			valid = false
			break
		}
	}
	return valid
}

func TestValidParenthesis(t *testing.T) {
	fmt.Printf("valid parenthesis? %q -> %v\n", "[{()}]", validParenthesis("[{()}]"))
	assert.Equal(t, validParenthesis("[{()}]"), true)
	fmt.Printf("valid parenthesis? %q -> %v\n", "[(])", validParenthesis("[{()}]"))
	assert.Equal(t, validParenthesis("[(])"), false)
	fmt.Printf("valid parenthesis? %q -> %v\n", "[])", validParenthesis("[{()}]"))
	assert.Equal(t, validParenthesis("[])"), false)
}
