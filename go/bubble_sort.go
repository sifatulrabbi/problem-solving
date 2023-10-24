package main

import "fmt"

func RunBubbleSort() {
	arr := []int{64, 25, 12, 22, 11}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		fmt.Printf("Iteration(%d) arr: %v\n", i+1, arr)
	}
}
