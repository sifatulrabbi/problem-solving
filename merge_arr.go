package main

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Slice(nums1, func(i, j int) bool {
		a := nums1[i]
		b := nums1[j]
		return a < b
	})
}
