package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Println("Starting function:", i)
			c <- i
			fmt.Println("finishing function:", i)
		}
		close(c)
	}(c)

	fmt.Println("main goroutine sleeping for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c {
		fmt.Println("data from channel", v)
	}

	time.Sleep(time.Second * 1)

	c <- 10 // buffered channel will panic because it is closed
}
