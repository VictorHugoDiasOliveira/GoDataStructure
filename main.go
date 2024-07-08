package main

import (
	"github.com/VictorHugoDiasOliveira/GoDataStructure/doublylinkedlist"
)

func main() {

	// list := linkedlist.InitializeLinkedList()
	// list.Append(12)
	// list.Append(15)
	// list.Append(8)
	// list.RemoveByValue(8)
	// list.ShowValues()
	// list.Length()

	doublylist := doublylinkedlist.InitializeDoublyLinkedList()
	doublylist.Append(18)
	doublylist.Append(19)
	doublylist.Append(5)
	doublylist.Append(12)
	doublylist.ShowForwardValues()
	doublylist.ShowBackwardValues()
}
