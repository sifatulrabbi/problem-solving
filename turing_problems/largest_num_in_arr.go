package turingproblems

import (
	"fmt"
)

/*
Given an integer n, follow the next rules to fill in a new array result:
  - result array is of size num + 1
  - result[0] = 0
  - result[1] = 1
  - result[2 * i] = result[i] when 2 <= 2 * i <= num
  - result[2 * i + 1] = result[i] + result[i + 1] when 2 <= 2 * i + 1 <= num
*/
func largestNumInArr(n int) int {
	if n+1 < 2 {
		return 0
	} else if n+1 < 2 {
		return 1
	}

	length := n + 1
	result := make([]int, length)
	result[0] = 0
	result[1] = 1

	for i := 2; i < length; i++ {
		if i%2 == 0 {
			result[i] = result[i/2]
		} else {
			result[i] = result[i/2] + result[i/2+1]
		}
	}
	fmt.Println("the array", result)

	largest := -1
	for _, v := range result {
		if v > largest {
			largest = v
		}
	}
	return largest
}
