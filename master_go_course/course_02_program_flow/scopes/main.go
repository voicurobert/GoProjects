package main

import (
	"fmt"

	f "fmt" // permitted
)

const done = false // package scoped

var b int = 10 // package scoped

func main() {
	var task = "Running" // local scoped
	fmt.Println(task, done)

	const done = true // permitted - shadows done constant from package scope

	f.Println(done)
}

func f1() {
	fmt.Println("done ", done)

}
