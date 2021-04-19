package main

import "fmt"

func main() {
	people := [5]string{
		"Hellen",
		"Mark",
		"Brenda",
		"Antoion",
		"Michael",
	}

	friends := [2]string{"Mark", "Marry"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}

	// go to statement

}
