package main

import "fmt"

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func main() {

	var p position
	fmt.Println(p)

	p2 := position{4, 2}
	fmt.Println(p2)

	p3 := new(position)
	fmt.Println(p3)

	b := badGuy{"Jabba Hut", 100, p2}
	fmt.Println(b)

	whereIsBadGuy(b)
}
