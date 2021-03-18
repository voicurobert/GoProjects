package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 go routine started")
	for i := 1; i < 3; i++ {
		fmt.Println("f1: i-", i)
	}
	fmt.Println("f1 go routine finished")
}

func f2() {
	fmt.Println("f2 go routine started")
	for i := 0; i < 3; i++ {
		fmt.Println("f2: i-", i)
	}
	fmt.Println("f2 go routine finished")
}

func main() {
	fmt.Println("main execution started")
	fmt.Println("No of cpus", runtime.NumCPU())
	fmt.Println("No of go routines", runtime.NumGoroutine())

	fmt.Println("OS", runtime.GOOS)
	fmt.Println("Arch", runtime.GOARCH)

	fmt.Println("GO MAX PROCS", runtime.GOMAXPROCS(0))

	go f1()

	fmt.Println("No of go routines after go f1()", runtime.NumGoroutine())

	f2()

	time.Sleep(time.Second * 2)

	fmt.Println("main execution stopped")
}
