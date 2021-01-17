package main

import "fmt"

func main() {
	price, inStock := 100, true
	if price > 80 {
		fmt.Println("Too Expensive!")
	}

	if price <= 100 && inStock {
		fmt.Println("Buy it")
	}

	if price < 100 {
		fmt.Println("It is cheap")
	} else if price == 100 {
		fmt.Println("Edge")
	} else {
		fmt.Println("Too Expensive!")
	}

	age := 7

	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote retur in %d years", 18-age)
	} else if age == 18 {
		fmt.Println("First vote")
	} else if age > 18 && age <= 100 {
		fmt.Println("Please vote")
	} else {
		fmt.Println("Invalid age")
	}
}
