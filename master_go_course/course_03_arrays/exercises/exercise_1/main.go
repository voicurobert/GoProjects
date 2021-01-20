package main

import "fmt"

func main() {
	var cities [2]string
	fmt.Printf("%#v\n", cities)

	grades := [3]float64{2.0, 4.6}
	fmt.Printf("%#v\n", grades)

	taskDone := [...]bool{true, true, true}
	fmt.Printf("%#v\n", taskDone)

	count := 0
	nums := [...]int{30, -1, -6, 90, -6}
	for _, v := range nums {
		if v > 0 && v%2 == 0 {
			count++
		}
	}
	fmt.Println(count)
}
