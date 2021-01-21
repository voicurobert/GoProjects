package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}

	s3, s4 := s1[0:2], s1[1:3]

	fmt.Println("s1: ", s1)
	fmt.Println("s3: ", s3)
	fmt.Println("s4: ", s4)

	s3[1] = 600
	fmt.Println("s1: ", s1)
	fmt.Println("s3: ", s3)
	fmt.Println("s4: ", s4)

	arr1 := [4]int{10, 20, 30, 40}

	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2
	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Println(slice2)

	cars := []string{"ford", "honda", "audi", "range rover"}
	newCars := []string{}
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "nissan"
	fmt.Println(cars)
	fmt.Println(newCars)

	a1 := []int{10, 20, 30, 40, 50}
	newSlice := a1[0:3]
	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = a1[2:5]
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p \n", &a1)
	fmt.Printf("%p \n", &newSlice)
	newSlice[0] = 1000
	fmt.Println("a1: ", a1)

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("array size %d \n", unsafe.Sizeof(a))
	fmt.Printf("slice size %d \n", unsafe.Sizeof(s))
}
