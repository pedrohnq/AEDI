package main

import (
	"aedi/data_structures/lists"
	"aedi/data_structures/queues"
	"aedi/data_structures/stacks"
	"fmt"
)

func usingArrayList() {
	fmt.Println("Using ArrayList!")
	arrayList := lists.ArrayList{}
	arrayList.Init(5)
	arrayList.Add(1)
	arrayList.Add(2)
	arrayList.Add(3)
	arrayList.Add(4)
	arrayList.Add(5)
	arrayList.AddOnIndex(10, 3)
	arrayList.RemoveOnIndex(4)
	fmt.Println(arrayList)
	element_in_index_2, _ := arrayList.Get(2)
	fmt.Println(element_in_index_2)
	arrayList.Set(11, 2)
	element_in_index_2, _ = arrayList.Get(2)
	fmt.Println(element_in_index_2)
	fmt.Println(arrayList.Size())
	fmt.Println(arrayList)
	arrayList.Reverse()
	fmt.Println(arrayList)
	fmt.Println("Finished using ArrayList!")
	fmt.Println("--------------------------")
}

func usingLinkedList() {
	fmt.Println("Using LinkedList!")
	linkedlist := lists.LinkedList{}
	linkedlist.Add(1)
	linkedlist.Add(2)
	linkedlist.Add(3)
	linkedlist.Add(4)
	linkedlist.Add(5)
	linkedlist.AddOnIndex(10, 3)
	linkedlist.RemoveOnIndex(4)
	linkedlist.Set(11, 2)
	element_in_index_2, _ := linkedlist.Get(2)
	fmt.Println(element_in_index_2)
	fmt.Println(linkedlist.Size())
	linkedlist.PrintList()
	linkedlist.Reverse()
	linkedlist.PrintList()
	fmt.Println("Finished using LinkedList!")
	fmt.Println("--------------------------")
}

func usingDoublyLinkedList() {
	fmt.Println("Using DoublyLinkedList!")
	doublylinkedlist := lists.DoublyLinkedList{}
	fmt.Println("Adding elements to DoublyLinkedList!")
	doublylinkedlist.Add(1)
	doublylinkedlist.Add(2)
	doublylinkedlist.Add(3)
	doublylinkedlist.Add(4)
	doublylinkedlist.Add(5)
	doublylinkedlist.PrintList()

	fmt.Println("Adding element 10 to index 3 in DoublyLinkedList!")
	doublylinkedlist.AddOnIndex(10, 3)
	doublylinkedlist.PrintList()

	fmt.Println("Removing element in index 4 in DoublyLinkedList!")
	doublylinkedlist.RemoveOnIndex(4)
	doublylinkedlist.PrintList()

	fmt.Println("Setting element 11 in index 2 in DoublyLinkedList!")
	doublylinkedlist.Set(11, 2)
	doublylinkedlist.PrintList()

	fmt.Println("Getting element in index 2 in DoublyLinkedList!")
	element_in_index_2, _ := doublylinkedlist.Get(2)
	fmt.Println(element_in_index_2)

	fmt.Println("Inverting DoublyLinkedList!")
	doublylinkedlist.PrintList()
	doublylinkedlist.Reverse()
	doublylinkedlist.PrintList()
	fmt.Println("Finished using DoublyLinkedList!")
	fmt.Println("--------------------------")
}

func balancedParenthesesArrayStack(str string) bool {
	stack := stacks.StackArrayList{}

	if string(str[0]) != "(" {
		fmt.Println("A string ", str, " is not starting with '('")
		return false
	}

	for i := 0; i < len(str); i++ {
		if string(str[i]) == "(" {
			stack.Push(1)
		} else if string(str[i]) == ")" {
			stack.Pop()
		}
	}
	if stack.Size() == 0 {
		fmt.Println("The string ", str, " is balanced")
		return true
	} else {
		fmt.Println("The string ", str, " is not balanced")
		return false
	}
}

func balancedParenthesesStackLinkedList(str string) bool {
	stack := stacks.StackLinkedList{}

	if string(str[0]) != "(" {
		fmt.Println("A string ", str, " is not starting with '('")
		return false
	}

	for i := 0; i < len(str); i++ {
		if string(str[i]) == "(" {
			stack.Push(1)
		} else if string(str[i]) == ")" {
			stack.Pop()
		}
	}
	if stack.Size() == 0 {
		fmt.Println("The string ", str, " is balanced")
		return true
	} else {
		fmt.Println("The string ", str, " is not balanced")
		return false
	}
}

func usingStackArrayList() {
	fmt.Println("Using StackArrayList!")
	str1 := "(()(())())"
	str2 := "(()(())("
	str3 := ")(())())("
	balancedParenthesesArrayStack(str1)
	balancedParenthesesArrayStack(str2)
	balancedParenthesesArrayStack(str3)
	fmt.Println("Finished using StackArrayList!")
}

func usingStackLinkedList() {
	fmt.Println("Using StackLinkedList!")
	str1 := "(()(())())"
	str2 := "(()(())("
	str3 := ")(())())("
	balancedParenthesesStackLinkedList(str1)
	balancedParenthesesStackLinkedList(str2)
	balancedParenthesesStackLinkedList(str3)
	fmt.Println("Finished using StackLinkedList!")
}

func usingQueueArrayList() {
	fmt.Println("Using QueueArrayList!")
	queue := queues.QueueArrayList{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	fmt.Println("Números na fila: ", queue.Size())
	next, _ := queue.Peek()
	fmt.Println("Próximo número: ", next)
	out_list, _ := queue.Dequeue()
	fmt.Println("Saiu da fila: ", out_list)
	fmt.Println("Números na fila: ", queue.Size())
	next, _ = queue.Peek()
	fmt.Println("Próximo número: ", next)
	out_list, _ = queue.Dequeue()
	fmt.Println("Saiu da fila: ", out_list)
	fmt.Println("Finished using QueueArrayList!")
}

func main() {
	usingArrayList()
	usingLinkedList()
	usingDoublyLinkedList()
	usingStackArrayList()
	usingStackLinkedList()
	usingQueueArrayList()
}
