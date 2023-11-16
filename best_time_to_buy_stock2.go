package main

import (
	"fmt"
)

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.
Find and return the maximum profit you can achieve.
*/
func RunMaxProfit2() {
	arr := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arr), arr)
	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(arr), arr)
	arr = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(arr), arr)
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	bought := 0
	totalProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] > prices[bought] {
			totalProfit += prices[i] - prices[bought]
			bought = i
		} else if prices[i] < prices[bought] {
			bought = i
		}
	}
	return totalProfit
}
