package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 go routine started")
	for i := 1; i < 3; i++ {
		fmt.Println("f1: i-", i)
		time.Sleep(time.Second)
	}
	fmt.Println("f1 go routine finished")
	wg.Done()
}

func f2() {
	fmt.Println("f2 go routine started")
	for i := 0; i <= 3; i++ {
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

	var wg sync.WaitGroup

	wg.Add(1)

	go f1(&wg)

	fmt.Println("No of go routines after go f1()", runtime.NumGoroutine())

	f2()

	wg.Wait()

	fmt.Println("main execution stopped")
}
