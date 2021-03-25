package main

import (
	"container/heap"
	"fmt"
)

/*
A heap is a data structure that is based on the heap property. The heap data structure is used in selection, graph, and k-way merge algorithms.
Operations such as finding, merging, insertion, key changes, and deleting are performed on heaps. Heaps are part of the container/heap package in Go.
According to the heap order (maximum heap) property, the value stored at each node is greater than or equal to its children.
*/

// IntegerHeap type
type IntegerHeap []int

// Len gets the length of IntegerHeap
func (iHeap IntegerHeap) Len() int {

	return len(iHeap)
}

// Less checks if element of i index is less that j index
func (iHeap IntegerHeap) Less(i, j int) bool {
	return iHeap[i] < iHeap[j]
}

// Swap swaps the element of i to j index
func (iHeap IntegerHeap) Swap(i, j int) {
	iHeap[i], iHeap[j] = iHeap[j], iHeap[i]
}

// Push added the item heapInterface
func (iHeap *IntegerHeap) Push(heapInterface interface{}) {
	*iHeap = append(*iHeap, heapInterface.(int))
}

// Pop gets the item from the heap
func (iHeap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iHeap
	n = len(previous)
	x1 = previous[n-1]
	*iHeap = previous[0 : n-1]
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("Mininum: %d \n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
