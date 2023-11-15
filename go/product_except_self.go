package main

import (
	"fmt"
)

type Shirt struct {
	price int
}

func RunProductExceptSelf() {
	fmt.Println(productExceptSelf([]int{4, 3, 2, 1}))

	shirts := []Shirt{
		{price: 5},
		{price: 4},
		{price: 8},
		{price: 3},
	}
	fmt.Println(createProductArrayFromShirts(shirts))
}

// O(n)
func productExceptSelf(nums []int) []int {
	l := len(nums)
	arr := make([]int, l)
	arr[0], arr[l-1] = 1, 1
	for i := 1; i < l; i++ {
		arr[i] = arr[i-1] * nums[i-1]
	}

	rp := 1
	for i := l - 2; i >= 0; i-- {
		rp *= nums[i+1]
		arr[i] *= rp
	}
	return arr
}

// advance product
func createProductArrayFromShirts(shirts []Shirt) []int {
	l := len(shirts)
	arr := make([]int, l)
	arr[0], arr[l-1] = 1, 1 // to make sure we aren't getting 0.
	// firstly product all the items at the left side of the current index.
	// since we are placing the product at the current index we are replacing the current index item with the calculated value.
	for i := 1; i < l; i++ {
		arr[i] = arr[i-1] * shirts[i-1].price
	}
	// now product all the items at the right of the current index.
	// because we have to know the last calculated value when determining total product we will create a pointer to store that
	rightTotal := 1
	for i := l - 2; i >= 0; i-- {
		rightTotal *= shirts[i+1].price
		arr[i] *= rightTotal
	}
	return arr
}
