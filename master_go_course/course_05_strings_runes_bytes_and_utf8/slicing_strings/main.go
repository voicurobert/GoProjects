package main

import "fmt"

func main() {
	s1 := "i love go lang"
	fmt.Println(s1[2:5])

	s2 := "§±~`"

	fmt.Println(s2[0:2])

	rs := []rune(s2)
	fmt.Printf("%T \n", rs)
	fmt.Println(string(rs[0:2]))
}
