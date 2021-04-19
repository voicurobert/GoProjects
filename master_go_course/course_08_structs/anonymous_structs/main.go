package main

import "fmt"

func main() {

	// anonymous strct
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}

	fmt.Printf("%#v\n", diana)
	fmt.Printf("Diana age %d\n", diana.age)

	// anonymous fields
	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1983 by George Orwell", 10.2, false}

	fmt.Printf("%#v\n", b1)

	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"John", 40000, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e)
}
