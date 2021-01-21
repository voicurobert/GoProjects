package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Create a Go program that reads some numbers from the command line and then
calculates the sum and the product of all the numbers given at command line.

The user should give between 2 and 10 numbers.
*/

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please add more than 1 argument!")
	}
	nums := args[1:]

	sum := 0.0
	product := 1.1
	for _, v := range nums {
		nr, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}
		sum += nr
		product *= nr
	}

	fmt.Println(sum, product)
}
