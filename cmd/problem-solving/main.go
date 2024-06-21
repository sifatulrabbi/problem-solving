package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr[4], arr[4:])
}
