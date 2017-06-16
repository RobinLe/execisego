package algorithm

import "fmt"

func (t *BTree) PreOrder() {
	if t == nil {
		return
	}
	fmt.Println(t.Data)
	t.LeftChild.PreOrder()
	t.RightChild.PreOrder()
}

func PreOrder2(tree *BTree) {
	// for stack is not empty
	for tree != nil {
		fmt.Println(tree.Data)
		// stack.push(tree.RightChild)
		tree = tree.LeftChild
	}
}

func (t *BTree) InOrder() {
	if t == nil {
		return
	}
	t.LeftChild.InOrder()
	fmt.Println(t.Data)
	t.RightChild.InOrder()
}

func (t *BTree) PostOrder() {
	if t == nil {
		return
	}
	t.LeftChild.PostOrder()
	t.RightChild.PostOrder()
	fmt.Println(t.Data)
}

func (t *BTree) Height() int {
	if t == nil {
		return 0
	}
	lHeight := t.LeftChild.Height()
	rHeight := t.RightChild.Height()
	if lHeight > rHeight {
		return lHeight + 1
	}
	return rHeight + 1
}
