package algorithm

import (
	"fmt"
)

type Element struct {
	Value int
}

type Stack struct {
	Elements []*Element
	Length   int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Element) {
	fmt.Println("Push ", n.Value)
	s.Elements = append(s.Elements[:s.Length], n)
	s.Length++
}

func (s *Stack) Pop() *Element {
	if s.Length == 0 {
		return nil
	}
	s.Length--
	fmt.Println("Pop ", s.Elements[s.Length].Value)
	return s.Elements[s.Length]
}
