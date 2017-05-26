package algorithm

import (
	"fmt"
)

// Stack last in first out LIFO struct
type Stack struct {
	Elements []*Element
	Top      int
	Length   int
}

// NewStack create new stack
func NewStack(length int) *Stack {
	return &Stack{Length: length}
}

func (s *Stack) isFull() bool {
	if s.Top == s.Length {
		fmt.Println("Stack is full")
		return true
	}
	return false
}

func (s *Stack) isEmpty() bool {
	if s.Top == 0 {
		fmt.Println("Stack is empty")
		return true
	}
	return false
}

func (s *Stack) push(n *Element) {
	if s.isFull() {
		return
	}
	fmt.Println("Push ", n.Value)
	s.Elements = append(s.Elements[:s.Top], n)
	s.Top++
}

func (s *Stack) pop() *Element {
	if s.isEmpty() {
		return nil
	}
	s.Top--
	fmt.Println("Pop ", s.Elements[s.Top].Value)
	return s.Elements[s.Top]
}

// peek return the top element of stack
func (s *Stack) peek() *Element {
	return s.Elements[s.Top-1]
}

func StackTest() {
	s := NewStack(3)
	s.push(&Element{1})
	s.push(&Element{2})
	s.push(&Element{3})
	s.push(&Element{4})
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.push(&Element{4})
	s.push(&Element{5})
	s.push(&Element{6})
	s.pop()
	fmt.Println("peek ", s.peek())
	s.push(&Element{6})
	fmt.Println("peek ", s.peek())
}
