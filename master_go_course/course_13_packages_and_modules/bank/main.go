package main

import (
	"fmt"

	"github.com/voicurobert/GoProjects/master_go_course/course_13_packages_and_modules/numbers"
)

func main() {
	fmt.Printf("%d is even: %t \n", 40, numbers.Even(40))
	fmt.Printf("%d is prime: %t \n", 13, numbers.IsPrime(13))
	fmt.Printf("%d is prime: %t \n", 100, numbers.IsPrime(100))
}
