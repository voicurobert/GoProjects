package main

import "fmt"

func main() {
	var employees map[string]string
	fmt.Printf("%#v \n", employees)

	fmt.Printf("No of pairs: %d\n", len(employees))

	fmt.Printf("The value for key %q is %q \n", "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("Value %#v \n", accounts["asd"])

	var m1 map[[5]int]string

	_ = m1 // can use arrays as keys for maps, arrays are comparable

	// employees["dan"] = "Programmer" employees is a nil map

	people := map[string]float64{}
	people["John"] = 21.4
	people["Marry"] = 10

	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 32,
		"EUR": 100,
	}
	fmt.Printf("%#v \n", balances)

	m := map[int]int{1: 10, 2: 12}
	_ = m

	balances["USD"] = 110.6
	balances["GBP"] = 100

	fmt.Printf("%#v \n", balances)

	v, ok := balances["RON"]
	fmt.Println(ok, v)

	for k, v := range balances {
		fmt.Printf("key %#v, value %#v \n", k, v)
	}

	delete(balances, "USD")

	fmt.Printf("%#v \n", balances)
}
