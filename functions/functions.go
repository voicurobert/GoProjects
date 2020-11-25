package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

func beSociable(name string) {
	sayHello(name)
	fmt.Println("How's the weather?")
	sayGoodbye(name)
}

func addOne(x int) int {
	return x + 1
}

func sayHelloABunch() {
	fmt.Println("HELLO")
	sayHelloABunch()
}

func main() {
	beSociable("Bill")
	beSociable("Bob")

	x := 5
	x = addOne(x)

	x = addOne(addOne(x))
	fmt.Println(x)
}
