package main

import "fmt"

func addOne(x *int) {
	*x = *x + 1 // *x returns the value the pointer points to
}

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b *badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func main() {

	x := 5
	fmt.Println(x)

	// var xPtr2 *int = &x
	xPtr := &x // give me the memory address of this value &
	fmt.Println(xPtr)

	addOne(xPtr)
	fmt.Println(x)

	p3 := position{2, 5}
	fmt.Println(p3)

	b := badGuy{"Jabba Hut", 100, p3}

	whereIsBadGuy(&b)
}
