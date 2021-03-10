package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type object interface {
	volume() float64
}

type geometry interface {
	shape
	object
	getColor() string
}

type cube struct {
	edge float64
}

func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)
}

func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v
}

func (c cube) getColor() string {
	return "RED"
}

func main() {
	c := cube{edge: 2}
	a, v := measure(c)

	fmt.Println("Area: volume:", a, v)
}
