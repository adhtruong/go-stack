package stack

import (
	"testing"
)

func TestStackEmpty(t *testing.T) {
	stack := NewStack()
	if !stack.IsEmpty() {
		t.Errorf("Require stack empty")
	}
}

func TestStackAddAndRemove(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	if stack.Size != 1 {
		t.Errorf("Require stack size 1")
	}
	element := stack.Pop()
	if element != 1 {
		t.Errorf("Popped element does not match")
	}
	if stack.Size != 0 {
		t.Errorf("Require stack size 0")
	}
}

func TestPopOfEmpty(t *testing.T) {
	stack := NewStack()
	if stack.Pop() != nil {
		t.Errorf("Require nil from empty stack")
	}
}
