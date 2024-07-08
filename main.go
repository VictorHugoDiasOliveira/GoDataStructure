package main

import (
	"github.com/VictorHugoDiasOliveira/GoDataStructure/linkedlist"
)

func main() {

	list := linkedlist.InitializeLinkedList()
	list.Append(12)
	list.Append(15)
	list.Append(8)
	list.RemoveByValue(8)
	list.ShowValues()
	list.Length()
}
