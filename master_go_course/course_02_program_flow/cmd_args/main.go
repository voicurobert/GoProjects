package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Args", os.Args)
	fmt.Println("Path", os.Args[0])
	fmt.Println("1st argument", os.Args[1])
	fmt.Println("size", len(os.Args))
}