package main

import (
	"fmt"
	"sort"
)

func RunHIndex() {
	arr := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(arr))
	arr = []int{0, 1}
	fmt.Println(hIndex(arr))
	arr = []int{100}
	fmt.Println(hIndex(arr))
	arr = []int{0}
	fmt.Println(hIndex(arr))
	arr = []int{1, 3, 1}
	fmt.Println(hIndex(arr))
}

func hIndex(citations []int) int {
	sort.Ints(citations)
	h := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if h >= citations[i] {
			return h
		}
		if citations[i] > 0 {
			h++
		}
	}
	return h
}
