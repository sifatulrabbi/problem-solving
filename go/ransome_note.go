package main

import "fmt"

func RunCanConstruct() {
	ransomNote, magazine := "a", "b"
	fmt.Println(canConstruct(ransomNote, magazine))
	ransomNote, magazine = "aa", "ab"
	fmt.Println(canConstruct(ransomNote, magazine))
	ransomNote, magazine = "aa", "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
	ransomNote, magazine = "cdcd", "abhytacgdvadcccdd"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	rn, m := map[string]int{}, map[string]int{}
	for _, v := range []rune(ransomNote) {
		rn[string(v)]++
	}
	for _, v := range []rune(magazine) {
		m[string(v)]++
	}
	for i, v := range rn {
		if m[i] < v {
			return false
		}
	}
	return true
}
