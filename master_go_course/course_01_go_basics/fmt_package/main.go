package main

import "fmt"

func main() {
	name := "Andrei"

	fmt.Println("I am", name)

	a, b := 4, 6
	fmt.Println("Sum:", a+b, "Mean value:", (a+b)/2)

	fmt.Printf("Your age is %d \n", 21)

	fmt.Printf("X is %d, Y is %f \n", 5, 6.8)

	fmt.Printf("He says \"Hello Go\" \n")

	figure := "Circle"
	radius := 5
	pi := 3.14159
	_, _ = figure, pi
	fmt.Printf("Radius is %d \n", radius)
	fmt.Printf("Radius is %+d \n", radius)

	fmt.Printf("Pi const is %f \n", pi)

	fmt.Printf("The diameter of a %s with a Radius of %d is %f \n", figure, radius, float64(radius)*2*pi)

	// %q for quoted string adds "" to the string
	fmt.Printf("This is a %q \n", figure)

	// %v > replaced by any type
	fmt.Printf("The diameter of a %v with a Radius of %v is %v \n", figure, radius, float64(radius)*2*pi)

	// %T returns the type of a variable
	fmt.Printf("Figure is of type %T \n", figure)

	// %t formats a boolean value
	closed := true
	fmt.Printf("closed?: %t \n", closed)

	// %b converts a int to base 2
	fmt.Printf("%08b \n", 55)
	// %08b prints the bites in 8

	// %x hexa
	fmt.Printf("%x \n", 100)

	x := 3.4
	y := 6.9
	fmt.Printf("X * Y = %.3f\n", x * y)

}
