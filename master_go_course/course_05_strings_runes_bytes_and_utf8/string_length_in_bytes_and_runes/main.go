package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "golang"
	fmt.Println(len(s1))

	name := "Codru§a"
	fmt.Println(len(name))

	n := utf8.RuneCountInString(name)
	fmt.Println(n)
}