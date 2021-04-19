package main

import "fmt"

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCar(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCarPointer(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func main() {
	c := car{
		brand: "Honda",
		price: 22000,
	}
	changeCar(c, "Honda2", 22233)
	fmt.Println(c)
	c.changeCar("BMW", 30000)
	fmt.Println(c)

	c.changeCarPointer("BMW", 4000)
	fmt.Println(c)

	var yourCar *car
	yourCar = &c
	fmt.Println(*yourCar)

	yourCar.changeCarPointer("VW", 20000)
	fmt.Println(*yourCar)
}
