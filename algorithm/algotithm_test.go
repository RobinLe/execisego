package algorithm

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(&Element{1})
	stack.Push(&Element{2})
	stack.Pop()
	stack.Push(&Element{3})
	stack.Pop()
	stack.Pop()
	stack.Push(&Element{4})
	stack.Pop()
}
