package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "bla"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.4
	*name = "bla"
	*sold = false
}

// Product name
type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Beer"
}

func changeProductByPointer(p *Product) {
	p.price = 300
	p.productName = "Beer"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["A"] = 10
	m["B"] = 30
}

func main() {
	q, p, n, s := 5, 300.0, "xxx", true
	fmt.Println("Before", q, p, n, s)
	changeValues(q, p, n, s)
	fmt.Println("After", q, p, n, s)
	fmt.Println("Before pointers", q, p, n, s)
	changeValuesByPointer(&q, &p, &n, &s)
	fmt.Println("After pointers", q, p, n, s)

	gift := Product{
		price:       100,
		productName: "Cola",
	}
	fmt.Println("Before changeProduct", gift)
	changeProduct(gift)
	fmt.Println("After changeProduct", gift)

	fmt.Println("Before changeProduct", gift)
	changeProductByPointer(&gift)
	fmt.Println("After changeProduct", gift)

	prices := []int{10, 20, 30}
	changeSlice(prices)
	fmt.Println(prices)

	m := map[string]int{"A": 1000, "B": 2000}
	changeMap(m)
	fmt.Println(m)

}
