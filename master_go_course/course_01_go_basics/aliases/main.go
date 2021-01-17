package main

import "fmt"

func main() {
	var a uint8 = 10
	var b byte

	b = a
	_ = b

	fmt.Println(b)

	type second = uint // alias for uint

	var hour second = 3600
	fmt.Printf("Minutes in an hour %d \n", hour/60)
}
