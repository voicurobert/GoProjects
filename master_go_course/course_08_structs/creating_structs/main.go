package main

import "fmt"

func main() {
	title1, author1, year1 := "The Devine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespear", 1606

	fmt.Println("Book1:", title1, author1, year1)
	fmt.Println("Book2:", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}

	myBook := book{"The Devine comedy", "Dante Aligheri", 1320} // not the recomendet way to create structs
	fmt.Println("My Book", myBook)

	myBook2 := book{
		title:  "Macbeth",
		author: "William",
		year:   1606,
	}

	fmt.Println(myBook2)

	aBook := book{title: "Bla"}
	fmt.Printf("%#v\n", aBook)

}
