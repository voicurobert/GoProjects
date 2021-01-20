package main

import "fmt"

func main() {
	grades := [3]int{
		1: 10,
		0: 4,
		2: 7,
	}
	fmt.Printf("%#v \n", grades)

	accounts := [3]int{2: 50}
	fmt.Printf("%#v \n", accounts)

	names := [...]string{
		5:  "dan",
		3:  "sss",
		10: "aaa",
	}
	fmt.Printf("%#v %d\n", names, len(names))

	cities := [...]string{
		5:        "Paris",
		"London", // will have index 6
		1:        "New york",
	}
	fmt.Printf("%#v %d\n", cities, len(cities))

	weekend := [7]bool{5: true, 6: true}
	fmt.Printf("%#v %d\n", weekend, len(weekend))
}
