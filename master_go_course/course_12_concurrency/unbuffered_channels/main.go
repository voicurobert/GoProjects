package main

import (
	"fmt"
	"time"
)

func main() {

	// unbuffered channel
	c1 := make(chan int)

	// buffered channel
	c2 := make(chan int, 3)
	_ = c2

	go func(c chan int) {
		fmt.Println("Starting function")
		c <- 10
		fmt.Println("finishing function")
	}(c1)

	fmt.Println("main sleeping for 2 seconds")
	time.Sleep(time.Second * 2)
	fmt.Println("data from channel", <-c1)

	time.Sleep(time.Second * 1)
}
