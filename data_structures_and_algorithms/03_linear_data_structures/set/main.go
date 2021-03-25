package main

import "fmt"

// Set struct
type Set struct {
	integerMap map[int]bool
}

func (set *Set) new() {
	set.integerMap = make(map[int]bool)
}

func (set *Set) containsElement(element int) bool {
	_, exists := set.integerMap[element]
	return exists
}

func (set *Set) addElement(element int) {
	if !set.containsElement(element) {
		set.integerMap[element] = true
	}
}

func (set *Set) deleteElement(element int) {
	delete(set.integerMap, element)
}

func (set *Set) intersect(another *Set) *Set {
	var newSet = &Set{}
	newSet.new()
	for value := range set.integerMap {
		if another.containsElement(value) {
			newSet.addElement(value)
		}
	}
	return newSet
}

func (set *Set) union(another *Set) *Set {
	var union = &Set{}
	union.new()
	for value := range set.integerMap {
		union.addElement(value)
	}
	for value := range another.integerMap {
		union.addElement(value)
	}
	return union
}

func main() {
	set := &Set{}
	set.new()
	set.addElement(10)
	set.addElement(11)
	fmt.Println(set.containsElement(11))

	anotherSet := &Set{}
	anotherSet.new()
	anotherSet.addElement(33)

	fmt.Println(set.union(anotherSet))
}
