package algorithm

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	StackTest()
}

func TestQueue(t *testing.T) {
	q := NewQueue(3)
	q.deQueue()
	q.enQueue(&BTree{1, nil, nil})
	q.enQueue(&BTree{2, nil, nil})
	q.enQueue(&BTree{3, nil, nil})
	q.enQueue(&BTree{4, nil, nil})
	q.deQueue()
	q.deQueue()
	q.deQueue()
	q.deQueue()
	q.enQueue(&BTree{5, nil, nil})
	q.enQueue(&BTree{6, nil, nil})
	q.deQueue()
	q.deQueue()
}

func TestTreeTraversal(t *testing.T) {
	tree := NewBTree(5)
	PrintBTree(tree)
	fmt.Println("Pre Order")
	tree.PreOrder()
	fmt.Println("In Order")
	tree.InOrder()
	fmt.Println("Post Order")
	tree.PostOrder()
}

func TestTreePrint(t *testing.T) {
	for i := 1; i < 10; i++ {
		tree := NewBTree(i)
		var prev *BTree
		PrintBTree(tree)
		fmt.Println(IsBTS(tree, &prev))
		fmt.Println()
	}
}
