package linkedlist

import "fmt"

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
	len  int
}

type LinkedListInterface interface {
	Append(int)
	ShowValues()
	Length()
	RemoveByValue(int)
	Index(int)
}

func InitializeLinkedList() LinkedListInterface {
	return &linkedList{len: 0}
}

func (list *linkedList) Append(value int) {
	newNode := &node{value: value}

	if list.head == nil {
		list.head = newNode
		list.len += 1
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}

		current.next = newNode
		list.len += 1
	}
}

func (list *linkedList) ShowValues() {
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

func (list *linkedList) RemoveByValue(value int) {
	if list.head == nil {
		fmt.Println("List is empty")
		return

	} else if list.head.value == value {
		list.head = list.head.next
		list.len = list.len - 1
		return

	} else {
		ancestor := list.head
		pointer := list.head.next
		for pointer != nil {
			if pointer.value == value {
				ancestor.next = pointer.next
				pointer.next = nil
				list.len = list.len - 1
				return
			}
			ancestor = pointer
			pointer = pointer.next
		}
		fmt.Println("Value not found")
	}
}

func (list *linkedList) Length() {
	fmt.Println(list.len)
}

func (list *linkedList) Index(value int) {
	if list.head == nil {
		fmt.Println("List is empty")
		return

	} else {
		index := 0
		current := list.head
		for current != nil {
			if current.value == value {
				fmt.Println(index)
				return
			}
			current = current.next
			index++
		}
	}
}
