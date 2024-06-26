package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
mainloop:
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			break mainloop
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
