package stacks

import "aedi/data_structures/lists"

type StackLinkedList struct {
	data lists.LinkedList
	size int
}

func (stack *StackLinkedList) Push(value int) {
	stack.data.Add(value)
}

func (stack *StackLinkedList) Pop() {
	stack.data.RemoveOnIndex(stack.size - 1)
}

func (stack *StackLinkedList) Peek() (int, error) {
	return stack.data.Get(stack.size - 1)
}

func (stack *StackLinkedList) Size() int {
	return stack.size
}
