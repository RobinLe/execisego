package algorithm

// Element base element of data struct
type Element struct {
	Value int
}

// Stack last in first out LIFO struct
type Stack struct {
	Elements []*Element
	Top      int
	Length   int
}

// Queue first in first out FIFO
type Queue struct {
	Elements []*BTree
	Rear     int
	Front    int
	Length   int
}

// BTree struct of binary tree
type BTree struct {
	Data       int
	LeftChild  *BTree
	RightChild *BTree
}
