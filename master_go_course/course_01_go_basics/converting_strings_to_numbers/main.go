package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99) // this is a rune
	fmt.Println(s)

	s2 := fmt.Sprintf("%d", 44)
	fmt.Println(s2)

	fmt.Println(string(34324))

	// string to numbers
	s1 := "3.123"
	fmt.Printf("%T \n", s1)

	f1, _ := strconv.ParseFloat(s1, 64)
	fmt.Println(f1)

	i, _ := strconv.Atoi("-50")
	j := strconv.Itoa(20)
	fmt.Printf("i type is %T, i value is %v \n", i, i)
	fmt.Printf("j type is %T, j value is %q \n", j, j)

}
