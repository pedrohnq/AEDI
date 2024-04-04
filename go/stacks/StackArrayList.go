package stacks

import (
	"aedi/go/lists"
)

type StackArrayList struct {
	data lists.ArrayList
	size int
}

func (stack *StackArrayList) Push(value int) {
	if stack.size == 0 {
		stack.data.Init(1)
	}
	stack.data.Add(value)
	stack.size++
}

func (stack *StackArrayList) Pop() {
	stack.data.RemoveOnIndex(stack.size - 1)
	stack.size--
}

func (stack *StackArrayList) Peek() (int, error) {
	return stack.data.Get(stack.size - 1)
}

func (stack *StackArrayList) Size() int {
	return stack.size
}
