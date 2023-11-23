package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	converted := ""
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		converted += trimStr(arr[i])
	}
	return false
}

func trimStr(s rune) string {
	allowed := []rune("abcdefghijklmnopqrstuvwxyz")
	sl := strings.ToLower(string(s))
	for i := 0; i < len(allowed); i++ {
		if sl == string(allowed[i]) {
			return sl
		}
	}
	return ""
}

