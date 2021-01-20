package main

import (
	"fmt"
	"time"
)

func main() {

	// switch add break automatically after each case statement
	language := "GO"
	switch language {
	case "python":
		fmt.Println("Pythin")
	case "golang", "GO":
		fmt.Println("Golang")
	default:
		fmt.Println("Default")
	}

	n := 5
	switch true { // compare only boolean expression in switch cases, default will not be execute
	case n%2 == 0:
		fmt.Println("even number")
	case n%2 != 0:
		fmt.Println("odd number")
	}

	hour := time.Now().Hour()
	fmt.Println(hour)
	switch { // missing expression means true
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
