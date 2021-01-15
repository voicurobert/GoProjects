package main

import "fmt"

func main() {
	// int
	var i1 int8 = 100
	fmt.Printf("%T \n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T \n", i2)

	// float type

	var f1, f2, f3 float64 = 1.1, -.2, 5. // 1.1, -0.2, 5.0
	fmt.Printf("%T %T %T \n", f1, f2, f3)

	// rune
	var r1 rune = 'f'
	fmt.Printf("%v \n", r1)

	// strings
	var s string = "Hello, Go"
	fmt.Printf("%T \n", s)

	// array
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T, %v \n", numbers, numbers)

	// splice
	var cities = []string{"London", "Tokio", "New York"}
	fmt.Printf("%T, %v \n", cities, cities)

	// map
	var balances = map[string]float64{
		"USD": 233.4,
		"EUR": 334.5,
	}
	fmt.Printf("%T, %v \n", balances, balances)

	// struct
	type person struct {
		name string
		age  int
	}

	var p person
	fmt.Printf("%T, %v \n", p, p)

	// pointer
	var x int = 2
	ptr := &x
	fmt.Printf("%T, %v \n", ptr, ptr)

	fmt.Printf("%T, %v \n", f, f)
}


func f() {

}