package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%#v \n", numbers)

	numbers[0] = 7
	fmt.Printf("%#v \n", numbers)

	// numbers[5] = 100 index of out bound

	for i, v := range numbers {
		fmt.Println(i, v)
	}

	fmt.Println(strings.Repeat("#", 20))

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	balances := [2][3]int{
		{5, 6, 7},
		{3, 8, 0},
	}
	fmt.Printf("%#v \n", balances)

	m := [3]int{1, 2, 3}
	n := m // two different arrays
	fmt.Println("n is equal to m: ", n == m)

	m[0] = -1
	fmt.Println("n is equal to m: ", n == m)
}
