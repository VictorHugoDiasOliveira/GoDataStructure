package doublylinkedlist

import "fmt"

type node struct {
	value    int
	previous *node
	next     *node
}

type doublyLinkedList struct {
	head *node
	tail *node
	len  int
}

type DoublyLinkedList interface {
	Append(int)
	Length()
	ShowForwardValues()
	ShowBackwardValues()
}

func InitializeDoublyLinkedList() DoublyLinkedList {
	return &doublyLinkedList{len: 0}
}

func (list *doublyLinkedList) Length() {
	fmt.Println(list.len)
}

func (list *doublyLinkedList) Append(value int) {
	newNode := &node{value: value}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		list.len++
	} else {
		newNode.previous = list.tail
		list.tail.next = newNode
		list.tail = newNode
		list.len++
	}
}

func (list *doublyLinkedList) ShowForwardValues() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	current := list.head
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.next
	}
	fmt.Println()
}

func (list *doublyLinkedList) ShowBackwardValues() {
	if list.tail == nil {
		fmt.Println("List is empty")
		return
	}

	current := list.tail
	for current != nil {
		fmt.Printf("%d ", current.value)
		current = current.previous
	}
	fmt.Println()
}
