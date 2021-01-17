package main

import "fmt"

func main() {
	a, b := 4, 2

	r := (a + b) / (a - b) * 2
	fmt.Println(r)

	r = 9 % a
	fmt.Println(r)

	aa, bb := 2, 3
	// increment assignment
	aa += bb
	fmt.Println(aa)

	// decrement assignment
	bb -= 2
	fmt.Println(bb)

	// multiplication assignment
	bb *= 10
	fmt.Println(bb)

	// division assignment
	bb /= 1
	fmt.Println(bb)

	// modulus assignment
	aa %= 3
	fmt.Println(aa)

	// increment statement
	x := 1
	x += 1
	x++
	fmt.Println(x)

	// && logical and
	c, d := 5, 10
	fmt.Println(c > 1 && d == 10)


}
