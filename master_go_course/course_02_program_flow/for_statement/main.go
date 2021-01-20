package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop using for loop
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// infinite loop
	/*
		sum := 0
		for {
			sum++
		}
	*/

	// continue statement
	fmt.Println("################################")
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13 \n", i)
			count++
		}
		if count == 10 {
			break
		}
	}

}
