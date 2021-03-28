package stack

import (
	"sync"
)

type element struct {
	data interface{}
	next *element
}

type Stack struct {
	head *element
	Size int
	lock *sync.Mutex
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.lock = &sync.Mutex{}
	return stack
}

func (stack *Stack) Push(data interface{}) {
	stack.lock.Lock()

	element := new(element)
	element.data = data
	element.next = stack.head

	stack.head = element
	stack.Size++

	stack.lock.Unlock()
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}

	stack.lock.Lock()

	r := stack.head.data
	stack.head = stack.head.next
	stack.Size--

	stack.lock.Unlock()

	return r
}

func (stack *Stack) IsEmpty() bool {
	return stack.head == nil
}
