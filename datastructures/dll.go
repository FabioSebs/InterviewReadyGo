package datastructures

import "fmt"

// doubly linked in
type DoublyLinkedList interface {
	AddToList(val int)
	PrintList()
	PrintListBackwards()
}

type DLL struct {
	Head *DLLNode
	Tail *DLLNode
}

type DLLNode struct {
	Value *int
	Next  *DLLNode
	Prev  *DLLNode
}

func NewDLL() DoublyLinkedList {
	var (
		newNode = DLLNode{
			Value: nil,
			Next:  nil,
			Prev:  nil,
		}
	)

	return &DLL{
		Head: &newNode,
		Tail: &newNode,
	}
}

func (dl *DLL) AddToList(val int) {
	var (
		newNode = DLLNode{
			Value: &val,
			Next:  nil,
			Prev:  nil,
		}
	)

	// first item to add
	if dl.Head.Value == nil {
		dl.Head = &newNode
		dl.Tail = &newNode
		return
	}

	// onwards
	secondToLast := dl.Tail
	dl.Tail.Next = &newNode
	dl.Tail = &newNode
	dl.Tail.Prev = secondToLast
}

func (dl *DLL) PrintList() {
	head := dl.Head
	for head != nil {
		fmt.Println(*head.Value)
		head = head.Next
	}
}

func (dl *DLL) PrintListBackwards() {
	tail := dl.Tail
	for tail != nil {
		fmt.Println(*tail.Value)
		tail = tail.Prev
	}
}
