package main

import "fmt"

func mostActive(customers []string) []string {
	activeCustomers := []string{}
	shareMap := map[string]int{}
	totalShares := len(customers)
	for i := 0; i < totalShares; i++ {
		c := customers[i]
		if _, exists := shareMap[c]; !exists {
			shareMap[c] = 1
		} else {
			shareMap[c] += 1
		}
	}
	for cus, nshare := range shareMap {
		if (nshare*100)/totalShares < 5 {
			continue
		}
		activeCustomers = append(activeCustomers, cus)
	}
	for i := 0; i < len(activeCustomers); i++ {
		for j := i + 1; j < len(activeCustomers); j++ {
			fmt.Println()
			fmt.Println(activeCustomers)
			fmt.Println(activeCustomers[j][0], string(activeCustomers[j][0]), activeCustomers[j], activeCustomers[i][0], string(activeCustomers[i][0]), activeCustomers[i])
			if activeCustomers[j][0] < activeCustomers[i][0] {
				activeCustomers[j], activeCustomers[i] = activeCustomers[i], activeCustomers[j]
			}
			fmt.Println(activeCustomers)
		}
	}
	return activeCustomers
}

// sample input []int32{1,1,1,3,3,2,2}
func longestSubarray(arr []int32) int32 {
	var l int32 = 1
	for i := 0; i < len(arr); i++ {
		var matches int32 = 1
		similar := 1
		b := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if b != arr[j] && b-arr[j] != 1 && b-arr[j] != -1 {
				continue
			}
			if b == arr[j] {
				similar++
			}
			matches++
		}
		if l < matches {
			l = matches
		}
	}
	return l
}
