package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John",
		salary: 4000,
		contactInfo: Contact{
			email:   "john@bla.com",
			address: "Address",
			phone:   334551213,
		},
	}

	fmt.Printf("%#v\n", john)

	fmt.Printf("Email %s\n", john.contactInfo.email)

	john.contactInfo.email = "john2@bla.com"

	fmt.Printf("Email %s\n", john.contactInfo.email)

	myContact := Contact{email: "andrei@domain.com", phone: 31231312, address: "Adresa"}

	fmt.Println(myContact)

}
