package main

import "fmt"

func main() {
	const a float64 = 5.5 // typed constant
	const b = 6.7         // untyped constant

	const c float64 = a * b
	const str = "Hello" + "Go!"

	const d = 5 > 10

	fmt.Println(d)

	// const x int = 5
	// const y float64 = 5.1 * x

	const x = 5
	const y = 2.2 * x // this is allowed because x is an untyped const

	var i int = x     // x changes to int
	var j float64 = y // var j float64 = float64(y)
	fmt.Println(i, j)
	var p byte = x
	fmt.Println(p)

}
