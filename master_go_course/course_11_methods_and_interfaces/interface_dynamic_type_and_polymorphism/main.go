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

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func printShape(s shape) {
	fmt.Printf("Shape %#v\n", s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	/*
		var s shape
		fmt.Printf("%T \n", s)

		ball := circle{radius: 2.5}

		s = ball
		fmt.Printf("%T \n", s)

		printShape(s)

		room := rectangle{width: 300.0, height: 600.0}
		s = room
		fmt.Printf("%T \n", s)
	*/

	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)

	ball, ok := s.(circle)
	if ok {
		fmt.Println(ball.volume())
	}

	s = rectangle{width: 30.0, height: 20.0}

	switch value := s.(type) {
	case circle:
		fmt.Println("has circle type", value)
	case rectangle:
		fmt.Println("has rectangle type", value)
	}
}
