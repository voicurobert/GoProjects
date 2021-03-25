package main

import "fmt"

// LinkedList is a doubly linked list
type LinkedList struct {
	headNode *Node
}

// Node node struct for doubly linked list
type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

func (linkedList *LinkedList) nodeBetweenValues(firstProperty, secondProperty int) *Node {
	var nodeWith *Node
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

func (linkedList *LinkedList) addToHead(property int) {
	var node = &Node{
		property:     property,
		nextNode:     nil,
		previousNode: nil,
	}
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}
	linkedList.headNode = node
}

// NodeWithValue returns node with property
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (linkedList *LinkedList) addAfter(nodeProperty, property int) {
	var node = &Node{
		property:     property,
		nextNode:     nil,
		previousNode: nil,
	}
	var nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

// LastNode returns the last node in the list
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (linkedList *LinkedList) addToEnd(property int) {
	var node = &Node{
		property:     property,
		nextNode:     nil,
		previousNode: nil,
	}

	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.addToHead(1)
	linkedList.addToHead(3)
	linkedList.addToEnd(5)
	linkedList.addAfter(1, 7)
	fmt.Println(linkedList.headNode.property)
	node := linkedList.nodeBetweenValues(1, 5)
	fmt.Println(node.property)
}
