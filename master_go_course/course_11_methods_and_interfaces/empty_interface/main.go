package main

import "fmt"

type emptyInterface interface {
}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}

	empty = 5

	fmt.Println(empty)

	empty = []int{2, 4, 5}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int)))

	you := person{}
	you.info = "Your name"

	fmt.Println(you.info)
}
