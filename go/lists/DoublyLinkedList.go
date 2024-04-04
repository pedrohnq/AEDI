package lists

import (
	"errors"
)

type NodeDoubly struct {
	value int
	next  *NodeDoubly
	prev  *NodeDoubly
}

type DoublyLinkedList struct {
	head *NodeDoubly
	tail *NodeDoubly
	size int
}

func (doublylinkedlist *DoublyLinkedList) Add(value int) {
	newNode := &NodeDoubly{value: value, next: nil, prev: nil}
	if doublylinkedlist.head == nil {
		doublylinkedlist.head = newNode
		doublylinkedlist.tail = newNode
	} else {
		doublylinkedlist.tail.next = newNode
		newNode.prev = doublylinkedlist.tail
		doublylinkedlist.tail = newNode
	}
	doublylinkedlist.size++
}

func (doublylinkedlist *DoublyLinkedList) AddOnIndex(value, index int) error {
	if index >= 0 && index <= doublylinkedlist.size {
		newNode := &NodeDoubly{value: value, next: nil, prev: nil}
		if index == 0 {
			newNode.next = doublylinkedlist.head
			doublylinkedlist.head.prev = newNode
			doublylinkedlist.head = newNode
		} else {
			current := doublylinkedlist.head
			for i := 0; i < index-1; i++ {
				current = current.next
			}
			newNode.next = current.next
			newNode.prev = current
			current.next = newNode
			if index == doublylinkedlist.size {
				doublylinkedlist.tail = newNode
			}
		}
		doublylinkedlist.size++
		return nil
	}
	return errors.New("index out of bounds")
}

func (doublylinkedlist *DoublyLinkedList) RemoveOnIndex(index int) error {
	if index >= 0 && index < doublylinkedlist.size {
		if index == 0 {
			nextNode := doublylinkedlist.head.next
			if nextNode != nil {
				nextNode.prev = nil
			}
			doublylinkedlist.head = nextNode
		} else {
			currentNode := doublylinkedlist.head
			for i := 0; i < index-1; i++ {
				currentNode = currentNode.next
			}
			nextNode := currentNode.next
			currentNode.next = nextNode.next
			if nextNode.next != nil {
				nextNode.next.prev = currentNode
			}
			if index == doublylinkedlist.size-1 {
				doublylinkedlist.tail = currentNode
			}
		}
		doublylinkedlist.size--
		return nil
	}
	return errors.New("index out of bounds")
}

func (doublylinkedlist *DoublyLinkedList) Get(index int) (int, error) {
	if index >= 0 && index < doublylinkedlist.size {
		current := doublylinkedlist.head
		if index < doublylinkedlist.size/2 {
			for i := 0; i < index; i++ {
				current = current.next
			}
		} else {
			current := doublylinkedlist.tail
			for i := doublylinkedlist.size - 1; i > index; i-- {
				current = current.prev
			}
		}
		return current.value, nil
	}
	return 0, errors.New("index out of bounds")
}

func (doublylinkedlist *DoublyLinkedList) Set(value, index int) error {
	if index >= 0 && index < doublylinkedlist.size {
		current := doublylinkedlist.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		current.value = value
		return nil
	}
	return errors.New("index out of bounds")
}

func (doublylinkedlist *DoublyLinkedList) Size() int {
	return doublylinkedlist.size
}

func (doublylinkedlist *DoublyLinkedList) PrintList() {
	current := doublylinkedlist.head
	for current != nil {
		print(current.value, " ")
		current = current.next
	}
}

func (doublylinkedlist *DoublyLinkedList) Reverse() {
	current := doublylinkedlist.head
	for current != nil {
		temp := current.next
		current.next = current.prev
		current.prev = temp
		current = temp
	}
	temp := doublylinkedlist.head
	doublylinkedlist.head = doublylinkedlist.tail
	doublylinkedlist.tail = temp
}
