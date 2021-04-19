package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	for i, v := range s {
		fmt.Println(i, v)
	}
}
