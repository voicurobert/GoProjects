package main

import (
	"fmt"
)

/*
Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}

Using copy() function create a copy of the slice. Prove that the slices are not
connected by modifying one slice and notice that the other slice is not modified.
*/

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}

	otherFriends := make([]string, len(friends))
	x := copy(otherFriends, friends)
	fmt.Println(x)
	fmt.Println(otherFriends)
	otherFriends[0] = "aaa"
	fmt.Println(otherFriends)
	fmt.Println(friends)
}
