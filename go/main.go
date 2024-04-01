package main

import (
	"aedi/go/lists"
	"fmt"
)

func main() {
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
	arrayList.Invert()
	fmt.Println(arrayList)
	fmt.Println("Finished using ArrayList!")
	fmt.Println("--------------------------")

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
	element_in_index_2, _ = linkedlist.Get(2)
	fmt.Println(element_in_index_2)

	linkedlist.PrintList()
	linkedlist.Invert()
	linkedlist.PrintList()
	fmt.Println("Finished using LinkedList!")
	fmt.Println("--------------------------")

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
	element_in_index_2, _ = doublylinkedlist.Get(2)
	fmt.Println(element_in_index_2)

	fmt.Println("Inverting DoublyLinkedList!")
	doublylinkedlist.PrintList()
	doublylinkedlist.Invert()
	doublylinkedlist.PrintList()
	fmt.Println("Finished using DoublyLinkedList!")
	fmt.Println("--------------------------")
}
