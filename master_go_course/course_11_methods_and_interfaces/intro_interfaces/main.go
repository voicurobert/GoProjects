package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func printCircle(c circle) {
	fmt.Println("Shape ", c)
	fmt.Println(c.area())
	fmt.Println(c.perimeter())
}

func printRectangle(r rectangle) {
	fmt.Println("Shape ", r)
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
}

func printShape(s shape) {
	fmt.Printf("Shape %#v\n", s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	c1 := circle{radius: 5.}
	r1 := rectangle{width: 3., height: 5.5}

	printShape(c1)
	printShape(r1)
}
