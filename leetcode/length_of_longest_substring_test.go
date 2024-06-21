package leetcode

import "testing"

type SubstringTestcase struct {
	Value  string
	Result int
}

func TestLengthOfLongestSubString(t *testing.T) {
	testcases := []SubstringTestcase{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"aab", 2},
		{"aabaab!bb", 3},
		{"dvdf", 3},
	}
	for _, v := range testcases {
		if result := lengthOfLongestSubstring(v.Value); result != v.Result {
			t.Errorf("result for '%s' should be %d, got: %d", v.Value, v.Result, result)
			t.FailNow()
		}
	}
}
