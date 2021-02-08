package main

import "fmt"

func main() {
	s1 := "Hi there Go!"
	fmt.Println(s1)
	fmt.Println("He says: \"Hello\"")

	fmt.Println(`He says: "Hello"`)
	// `` raw strings

	s2 := `I like go`
	fmt.Println(s2)

	fmt.Println("Price: 100 \nBrand: Nike")

	fmt.Println(`C:\Users\RVO`)
	fmt.Println("C:\\Users\\RVO")

	var s3 = "I love " + "Go"
	fmt.Println(s3 + "!")

	fmt.Println("Element at index 0:", s3[0])

	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3)
}
