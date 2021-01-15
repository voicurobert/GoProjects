package main

import "fmt"

func main() {
	// int age = 10; - c++
	var age int = 30
	fmt.Println("Age:", age)

	var name = "Dan"
	_ = name

	s := "Learning golang"

	fmt.Println("Your name is:", name)
	fmt.Println(s)

	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8
	_, _ = i, j

	// swapping variables
	j, i = i, j
	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)
}
