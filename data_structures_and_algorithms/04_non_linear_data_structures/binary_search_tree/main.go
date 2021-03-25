package main

import "sync"

type treeNode struct {
	key       int
	value     int
	leftNode  *treeNode
	rightNode *treeNode
}

type binarySearchTree struct {
	rootNode *treeNode
	lock     sync.RWMutex
}

func insertTreeNode(rootNode, node *treeNode) {
	if node.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = node
		} else {
			insertTreeNode(rootNode.leftNode, node)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = node
		} else {
			insertTreeNode(rootNode.rightNode, node)
		}
	}
}

func (tree *binarySearchTree) insertElement(key, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	treeNode := &treeNode{
		key:       key,
		value:     value,
		leftNode:  nil,
		rightNode: nil,
	}

	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode)
	}
}

func (tree *binarySearchTree) inOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverseTree(tree.rootNode, function)
}

func inOrderTraverseTree(treeNode *treeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

func main() {

}
