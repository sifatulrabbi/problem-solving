package leetcode

import "fmt"

func RunBubbleSort() {
	arr := []int{64, 25, 12, 22, 11, 4, 54, 67, 2, 34, 12, 76, 23, 12, 50}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j < n-1-i; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
