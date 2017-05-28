package algorithm

import (
	"fmt"
	"math/rand"
)

// BinaryTree is a binary tree with integer values.
type BinaryTree struct {
	Value      int
	LeftChild  *BinaryTree
	RightChild *BinaryTree
}

// NewBinaryTree return a new binary tree
func NewBinaryTree(n int) *BinaryTree {
	var t *BinaryTree
	for _, i := range rand.Perm(n) {
		t = insert(t, rand.Intn(100+i))
	}
	// fmt.Println(t.Value)
	return t
}

func insert(tree *BinaryTree, value int) *BinaryTree {
	if tree == nil {
		return &BinaryTree{value, nil, nil}
	}
	if value < tree.Value {
		tree.LeftChild = insert(tree.LeftChild, value)
		return tree
	}
	tree.RightChild = insert(tree.RightChild, value)
	return tree
}

func PreOrder(tree *BinaryTree) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Value)
	PreOrder(tree.LeftChild)
	PreOrder(tree.RightChild)
}

func PreOrder2(tree *BinaryTree) {
	// for stack is not empty
	for tree != nil {
		fmt.Println(tree.Value)
		// stack.push(tree.RightChild)
		tree = tree.LeftChild
	}
}

func InOrder(tree *BinaryTree) {
	if tree == nil {
		return
	}
	InOrder(tree.LeftChild)
	fmt.Println(tree.Value)
	InOrder(tree.RightChild)
}

func PostOrder(tree *BinaryTree) {
	if tree == nil {
		return
	}
	PostOrder(tree.LeftChild)
	PostOrder(tree.RightChild)
	fmt.Println(tree.Value)
}

func TreeTest() {
	tree := NewBinaryTree(3)
	fmt.Println("Pre Order")
	PreOrder(tree)
	fmt.Println("In Order")
	InOrder(tree)
	fmt.Println("Post Order")
	PostOrder(tree)
}
