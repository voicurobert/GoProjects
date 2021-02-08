package main

import (
	"fmt"
	"strings"
)

func main() {

	p := fmt.Println

	result := strings.Contains("i love go programming", "love")

	p(result)

	result = strings.ContainsAny("success", "xys")
	p(result)

	result = strings.ContainsRune("golang", 'g')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("GO"))

	p(strings.ToUpper("go"))

	p("go" == "go")

	p("GO" == "go")

	p(strings.ToLower("GO") == strings.ToLower("go"))

	p(strings.EqualFold("GO", "go"))

	myStr := strings.Repeat("ab", 10)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1)
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr)

	s := strings.Split("a, b, c", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)
	s = strings.Split("a,b,c", "")
	fmt.Printf("%#v\n", s)

	s = []string{"I", "learn", "Golang"}
	myStr = strings.Join(s, " ")
	p(myStr)

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr)
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)

	s1 := strings.TrimSpace("\t Goodby Windows")
	p(s1)

	s2 := strings.Trim("...Hello Go!!!?", ".!?")
	p(s2)
}
