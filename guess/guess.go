package main

// 1. make the program to print out how many tries it took to win
// 2. see if you can if the user is lying - and quit - when the high bound is reached

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100

	count := 0

	fmt.Println("Please think of a number between", low, "to", high)
	fmt.Println("Press ENTER when ready")

	scanner.Scan()

	for {
		guess := (low + high) / 2
		fmt.Println("I guess the number is ", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low")
		fmt.Println("(c) correct?")

		scanner.Scan()

		response := scanner.Text()

		if response == "a" {
			high = guess - 1
			count++
		} else if response == "b" {
			low = guess + 1
			count++
		} else if response == "c" {
			fmt.Println("I won in", count, "steps!")
			break
		} else {
			fmt.Println("Invalid response, try again.")
		}
	}
}
