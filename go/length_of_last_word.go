package main

import (
	"fmt"
	"strings"
)

func RunLengthOfLastWord() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}

func lengthOfLastWord(s string) int {
	words := strings.Split(strings.Trim(s, " "), " ")
	return len(words[len(words)-1])
}
