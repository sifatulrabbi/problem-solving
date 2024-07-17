package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text := sc.Text()
		if !strings.Contains(text, ",") {
			continue
		}
		trimmed := strings.TrimFunc(text, func(r rune) bool {
			return !strings.Contains("1234567890,", string(r))
		})
		initialArr := strings.Split(trimmed, ",")
		arr := initialArr
		for i := len(initialArr) - 1; i > 0; i-- {
			if initialArr[i] == "" {
				arr = initialArr[:i]
			}
		}
		fmt.Println("array:", arr)
	}
}
