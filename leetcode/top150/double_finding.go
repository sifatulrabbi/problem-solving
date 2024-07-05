package leetcode

import (
	"fmt"
	"math/rand"
	"time"

	"problem-solving/data"
)

func RunDoubleFinding() {
	nums := []int{}
	for i := 0; i < 10; i++ {
		num := data.Array2k[rand.Intn(len(data.Array2k))]
		nums = append(nums, num)
	}
	for _, v := range nums {
		doubleFinding(data.Array2k, v)
		regularFinding(data.Array2k, v)
		fmt.Println()
	}
}

func doublePartFinding(arr []int, item int, found chan bool) {
	mid := len(arr) / 2
	go func() {
		for i := 0; i < mid; i++ {
			if arr[i] == item {
				break
			}
		}
		found <- true
	}()
	go func() {
		for j := len(arr) - 1; j >= mid; j-- {
			if arr[j] == item {
				break
			}
		}
		found <- true
	}()
}

func doubleFinding(arr []int, item int) {
	found := make(chan bool, 1)
	st := time.Now()
	mid := len(arr) / 2

	go doublePartFinding(data.Array2k[:mid], item, found)
	go doublePartFinding(data.Array2k[mid:], item, found)

	<-found
	et := time.Now()
	diff := et.Sub(st)
	fmt.Printf("finding: %d; double: %dns\n", item, diff.Nanoseconds())
}

func regularFinding(arr []int, item int) {
	st := time.Now()

	for _, v := range arr {
		if v == item {
			break
		}
	}

	et := time.Now()
	diff := et.Sub(st)
	fmt.Printf("finding: %d; regular: %dns\n", item, diff.Nanoseconds())
}
