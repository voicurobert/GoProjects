package main

import (
	"fmt"
)

// Node struct
type Node struct {
	property int
	NextNode *Node
}

// LinkedList struct
type LinkedList struct {
	HeadNode *Node
}

// AddToHead add node to linkedList as head node
func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property

	nextNode := linkedList.HeadNode

	if nextNode != nil {
		node.NextNode = nextNode
	}

	linkedList.HeadNode = node
}

// IterateList iterates over linkedList
func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.HeadNode; node != nil; node = node.NextNode {
		fmt.Println(node.property)
	}
}

// LastNode returns the last node in the list
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.NextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd add node to the end of the list
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{
		property: property,
		NextNode: nil,
	}
	var lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.NextNode = node
	}
}

// NodeWithValue returns node with property
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter add a new node with property after node with nodeProperty
func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	var node = &Node{
		property: property,
		NextNode: nil,
	}
	var nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.NextNode = nodeWith.NextNode
		nodeWith.NextNode = node
	}
}

func main() {

	var linkedList LinkedList
	linkedList = LinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)

	linkedList.IterateList()
}
