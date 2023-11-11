package main

import (
	"fmt"
	"sync"
)

func RunProductExceptSelf() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(productExceptSelf(arr))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	arr := make([]int, n)
	c := make(chan [2]int, n)

	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(pos int) {
			defer wg.Done()
			r := 0
			for j := 0; j < n; j++ {
				if j != pos {
					if r == 0 {
						r = nums[j]
					} else {
						r *= nums[j]
					}
				}
			}
			c <- [2]int{pos, r}
		}(i)
	}

	wg.Wait()
	close(c)

	for pair := range c {
		arr[pair[0]] = pair[1]
	}

	return arr
}
