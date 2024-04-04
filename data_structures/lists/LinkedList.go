package lists

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (linkedlist *LinkedList) Add(value int) {
	newNode := &Node{value: value, next: nil}
	aux := linkedlist.head
	if aux == nil {
		linkedlist.head = newNode
	} else {
		for aux.next != nil {
			aux = aux.next
		}
		aux.next = newNode
	}
	linkedlist.size++
}

func (linkedlist *LinkedList) AddOnIndex(value, index int) error {
	if index >= 0 && index <= linkedlist.size {
		newNode := &Node{value: value, next: nil}
		if index == 0 {
			newNode.next = linkedlist.head
			linkedlist.head = newNode
		} else {
			prev := linkedlist.head
			for i := 0; i < index-1; i++ {
				prev = prev.next
			}
			newNode.next = prev.next
			prev.next = newNode
		}
		linkedlist.size++
		return nil
	} else {
		return errors.New("index out of bounds")
	}
}

func (linkedlist *LinkedList) RemoveOnIndex(index int) error {
	if index >= 0 && index < linkedlist.size {
		if index == 0 {
			linkedlist.head = linkedlist.head.next
		} else {
			prev := linkedlist.head
			for i := 0; i < index-1; i++ {
				prev = prev.next
			}
			prev.next = prev.next.next
		}
		linkedlist.size--
		return nil
	}
	return errors.New("index out of bounds")
}

func (linkedlist *LinkedList) Get(index int) (int, error) {
	if index >= 0 && index < linkedlist.size {
		aux := linkedlist.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		return aux.value, nil
	}
	return 0, errors.New("index out of bounds")
}

func (linkedlist *LinkedList) Set(value, index int) error {
	if index >= 0 && index < linkedlist.size {
		aux := linkedlist.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		aux.value = value
		return nil
	}
	return errors.New("index out of bounds")
}

func (linkedlist *LinkedList) Size() int {
	return linkedlist.size
}

func (linkedlist *LinkedList) PrintList() {
	aux := linkedlist.head
	for aux != nil {
		fmt.Print(aux.value, " ")
		aux = aux.next
	}
}

func (linkedlist *LinkedList) Reverse() {
	if linkedlist.size > 0 {
		var prev *Node = nil
		curr := linkedlist.head
		var next *Node = nil
		for curr != nil {
			next = curr.next
			curr.next = prev
			prev = curr
			curr = next
		}
	}
}
