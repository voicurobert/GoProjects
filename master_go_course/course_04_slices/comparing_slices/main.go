package main

import "fmt"

func main() {
	var n []int
	m := []int{1, 2, 3}

	n = []int{1, 2, 3}

	var eq bool = true
	for i, v := range n {
		if v != m[i] {
			eq = false
			break
		}
	}
	fmt.Println(eq)
}
