package main

import "fmt"

func main() {
	var cities []string
	fmt.Println(cities == nil)
	fmt.Printf("%#v size %d \n", cities, len(cities))

	numbers := []int{2, 3, 5}
	fmt.Println(numbers)

	nums := make([]int, 10)
	fmt.Printf("%#v \n", nums)

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println(myFriend)

	for i, v := range numbers {
		fmt.Println(i, v)
	}

	var n []int
	n = numbers
	fmt.Println(n)
}
