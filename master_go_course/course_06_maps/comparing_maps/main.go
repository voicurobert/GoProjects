package main

import "fmt"

func main() {
	a := map[string]string{"a": "x"}
	b := map[string]string{"b": "y"}

	// fmt.Println(a == b)
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == s2)
}
