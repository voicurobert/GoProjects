package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("This is f1 function")
}

func f2(a int, b int) {
	fmt.Println("Sum:", a+b)
}

func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 {
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

func sum(a, b int) (s int) {
	s = a + b
	return
}

func main() {
	f1()
	f2(5, 7)
	f3(4, 5, 6, 7.0, 8.0, "String")
	p := f4(5.1)
	fmt.Println(p)

	a, b := f5(10, 20)
	fmt.Println(a, b)
	s := sum(3, 5)
	fmt.Println(s)
}
