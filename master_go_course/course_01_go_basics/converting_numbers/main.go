package main

import "fmt"

func main() {
	var x = 3   // int
	var y = 3.1 // float

	// x = x * int(y)
	x = int(float64(x) * y)
	fmt.Println(x)

	y = float64(x) * y
	fmt.Println(y)

	var a = 5
	fmt.Printf("%T \n", a)

	var b int64 = 2
	_ = b
}
