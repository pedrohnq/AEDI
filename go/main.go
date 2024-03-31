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
}
