package main

import (
	"fmt"
)

func main() {
	var nums []int
	fmt.Printf("%#v \n", nums)
	fmt.Println(len(nums), cap(nums))

	nums = append(nums, 1, 2)
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 3)
	fmt.Println(len(nums), cap(nums))

	nums = append(nums, 4)
	fmt.Println(len(nums), cap(nums))

	nums = append(nums, 6)
	fmt.Println(len(nums), cap(nums))

	nums = append(nums[0:4], 200, 300, 400, 500, 600)
	fmt.Println(len(nums), cap(nums))

	letters := []string{"a", "b", "c", "d", "e", "f"}
	letters = append(letters[:1], "x", "y")
	fmt.Println(letters, len(letters), cap(letters))

	fmt.Println(letters[4]) // index out of range
	fmt.Println(letters[3:6])
}
