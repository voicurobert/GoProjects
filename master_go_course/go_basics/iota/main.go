package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	const (
		c11 = iota
		c23
		c33
	)

	fmt.Println(c1, c2, c3)
	fmt.Println(c11, c23, c33)

	const (
		North = iota
		East
		South
		West
	)

	fmt.Println(West)

	const (
		a = iota*2 + 1 // 0
		b              // 2
		c              // 4
	)

	fmt.Println(a, b, c)

	// x = -2, y = -4, z = -5

	const (
		x = -(iota + 2) // -2
		_               // -3 but it is skipped
		y               // -4
		z               // -5
	)
	fmt.Println(x, y, z)
}
