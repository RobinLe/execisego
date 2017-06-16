package algorithm

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// NewBTree return a new binary tree
func NewBTree(n int) *BTree {
	var tree *BTree
	for _, i := range rand.Perm(n) {
		tree = insert(tree, rand.Intn(100+i))
	}
	return tree
}

// insert node to btree
func insert(tree *BTree, data int) *BTree {
	if tree == nil {
		return &BTree{data, nil, nil}
	}
	if rand.Intn(20)%2 == 0 {
		tree.LeftChild = insert(tree.LeftChild, data)
		return tree
	}
	tree.RightChild = insert(tree.RightChild, data)
	return tree
}

// IsBTS binary tree is search tree or not
// inorder traversal and keep track the previous node vale
func IsBTS(tree *BTree, prev **BTree) bool {
	if tree != nil {
		if !IsBTS(tree.LeftChild, prev) {
			return false
		}
		// compare previous node data with current data
		if *prev != nil && (*prev).Data >= tree.Data {
			return false
		}
		// save node data and keep track
		*prev = tree
		return IsBTS(tree.RightChild, prev)
	}
	return true
}

// PrintBTree print btree
func PrintBTree(tree *BTree) {
	q := NewQueue(20)
	q.enQueue(tree)
	depth := tree.Height()
	for !q.IsEmpty() {
		space := strings.Repeat(" ", int(math.Pow(2.0, float64(depth-1))))
		fmt.Print(space)
		depth--
		count := q.Rear - q.Front
		for i := 0; i < count; i++ {
			tree := q.deQueue()
			fmt.Print(tree.Data)
			if tree.LeftChild != nil {
				q.enQueue(tree.LeftChild)
				fmt.Print("/")
			}
			if tree.RightChild != nil {
				q.enQueue(tree.RightChild)
				fmt.Print("\\")
			}
			fmt.Print(space)
		}
		fmt.Println()
	}
}
