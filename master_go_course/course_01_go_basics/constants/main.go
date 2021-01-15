package main

import (
	"fmt"
)

func main() {
	const days int = 7
	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234 // in hours
	fmt.Printf("Duration is seconds: %v \n", duration*secondsInHour)

	x, y := 5, 1

	fmt.Println(x / y)

	const a = 5
	const b = 0
	//fmt.Println(a/b)

	const n, m int = 4, 5
	const n1, m1 = 7, 8

	const (
		min1 = -500
		min2
		min3
	)
	fmt.Println(min1, min2, min3)

	// constants rules
	// 1. you cannot change a constant
	// 2. you cannot initialize a constant at runtime, it belongs to compile time
	//const power = math.Pow(2, 3)
	// 3. cannot use a variable to initialize a constant
	// 4. you can use len() func to declare a const. Len is available at compile time
	const l1 = len("hello")

	// typed vs untyped constant
}
