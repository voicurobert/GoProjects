package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano() / 1e6

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Salut"
	}()

	time.Sleep(2 * time.Second)
	// use select to retrieve simulatinously data from channels

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		default:
			fmt.Println("No activity")
		}
	}

	end := time.Now().UnixNano() / 1e6
	fmt.Println("Time", end-start)
}
