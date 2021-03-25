package main

import (
	"fmt"
	"strconv"
)

type element struct {
	elementValue int
}

type stack struct {
	elements []*element
	count    int
}

func (e *element) String() string {
	return strconv.Itoa(e.elementValue)
}

func (s *stack) new() {
	s.elements = make([]*element, 0)
}

func (s *stack) push(e *element) {
	s.elements = append(s.elements[:s.count], e)
	s.count++
}

func (s *stack) pop() *element {
	if s.count == 0 {
		return nil
	}

	length := len(s.elements)
	e := s.elements[length-1]
	if length > 1 {
		s.elements = s.elements[:length-1]
	} else {
		s.elements = s.elements[0:]
	}
	s.count = len(s.elements)
	return e
}

func main() {
	s := &stack{}
	s.new()
	e1 := &element{elementValue: 3}
	e2 := &element{elementValue: 1}
	e3 := &element{elementValue: 2}
	e4 := &element{elementValue: 7}

	s.push(e1)
	s.push(e2)
	s.push(e3)
	s.push(e4)
	fmt.Printf("%#v", s)
}
