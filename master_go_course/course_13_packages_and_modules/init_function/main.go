package main

import "fmt"

// declaring array
var numbers [10]int

func init() {
	fmt.Println("In init function")
	numbers[0] = 1
	numbers[1] = 2
}

func init() {
	fmt.Println("In init 2 function")
}

func main() {
	fmt.Println("In main function")
	fmt.Printf("%#v \n", numbers)
}
