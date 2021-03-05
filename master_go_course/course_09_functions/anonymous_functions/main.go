package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {

	func(msg string) {
		fmt.Println(msg)
	}("ASD")

	a := increment(10)
	fmt.Printf("%T\n", a)
	a()
	fmt.Println(a())
}
