package main

import "fmt"

func main() {
	var ch chan int // geting values of int through channels

	fmt.Println(ch)

	ch = make(chan int) // initializing a channel

	fmt.Println(ch)

	c := make(chan int)

	// <- channel operator

	// send operation
	c <- 10

	// receive operation
	num := <-c

	fmt.Println(num)

	close(c) // closing a channel

}
