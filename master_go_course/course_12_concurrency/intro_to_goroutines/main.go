package main

import (
	"fmt"
	"time"
)

func doSomething() {
	fmt.Println("Hello")
}

func main() {
	doSomething()

	// spawning a goroutine
	go doSomething()

	time.Sleep(time.Second * 2)
}
