package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var1, var2 := 'a', 'b' // rune literals
	fmt.Printf("Type %T, value %d\n", var1, var1)

	fmt.Printf("Type %T, value %d\n", var2, var2)

	str := "†ara"
	fmt.Println(len(str))

	fmt.Println("byte", str[3])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Println()

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}

	fmt.Println()

	for _, r := range str {
		fmt.Printf("%c", r)
	}
}
