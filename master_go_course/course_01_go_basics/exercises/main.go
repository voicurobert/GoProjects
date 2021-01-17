package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	var b float64
	var c bool
	var d string

	x, y, z := 20, 15.5, "Gopher!"

	fmt.Println(a, b, c, d, x, y, z)

	var aa float64 = 7.1

	xx, yy := true, 3.7

	aa, xx = 5.5, false

	_, _, _ = aa, xx, yy

	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)

	const secPerDay = 86400
	const daysYear = 365
	fmt.Printf("Number of seconds in a year %d \n", daysYear*secPerDay)

	// cannot create constant with a slice
	//const m = []int{1, 3, 4, 5, 6, 8}
	//_ = m

	const aaa int = 7
	const bbb float64 = 5.6
	const ccc = float64(aaa) * bbb

	xxx := 8
	_ = xxx
	// const xc int = xxx

	//const noIPv6 = math.Pow(2, 128)

	const (
		Jun = iota + 6
		Jul
		Aug
	)
	fmt.Println(Jun, Jul, Aug)

	x1, y1, z1 := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}
	fmt.Printf("X1 %d\n", x1)
	fmt.Printf("Y1 %f\n", y1)
	fmt.Printf("Z1 %q\n", z1)
	fmt.Printf("score %v\n", score)

	const x2 float64 = 1.422349587101
	fmt.Printf("X2 %.4f", x2)

	shape := "circle"
	radius := 3.2
	const pi2 float64 = 3.14159
	circumference := 2 * pi2 * radius

	fmt.Printf("Shape: %q\n", shape)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
	_ = shape

	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	is := strconv.Itoa(i)
	fmt.Println(is)

	s2i, _ := strconv.Atoi(s2)
	fmt.Println(s2i)

	f2 := fmt.Sprintf("%f", f)
	fmt.Println(f2)

	fs1, _ := strconv.ParseFloat(s1, 64)
	fmt.Println(fs1)

	x3, y3 := 4, 5.1

	z3 := float64(x3) * y3
	fmt.Println(z3)

	a3 := 5
	b3 := 6.2 * float64(a3)
	fmt.Println(b3)

	const (
		distanceToSun = 149000000.6
		speedOfLight  = 299792.458
	)

	time := (distanceToSun * 1000) / speedOfLight
	fmt.Println(time)

	type duration int
	var hour duration

	fmt.Println(hour)
	fmt.Printf("%T\n", hour)
	hour = 3600
	fmt.Println(hour)
	minute := 60
	fmt.Println(hour != duration(minute))
	
	type mile float64
	type kilometer float64
	const m2km = 1.609
	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer
	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)
	fmt.Println(kmBerlinToParis)
}
