package main

import "fmt"

func main() {
	friends := map[string]int{"dan": 40, "maria": 25}
	neighbors := friends

	friends["dan"] = 50

	fmt.Println(neighbors)

	people := make(map[string]int)
	for k, v := range friends {
		people[k] = v
	}

	people["Anne"] = 18

	fmt.Println(friends)
	fmt.Println(people)

}