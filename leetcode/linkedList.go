package leetcode

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value int
	Next  *Node
}

func NewLinkedList(headValue int) LinkedList {
	// make node and set it to head and tail
	node := &Node{
		Value: headValue,
		Next:  nil,
	}

	return LinkedList{
		Head: node,
		Tail: node,
	}
}

func (ll *LinkedList) AddToList(newValue int) {
	// make the new node
	newNode := &Node{
		Value: newValue,
		Next:  nil,
	}

	// edge case for empty list
	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
		return
	}

	// update tail (the reason the head's next doesnt need to be updated is bc the tail node is the same address of the head)
	// essentially on the first iteration you are updating the head and tail
	ll.Tail.Next = newNode
	ll.Tail = newNode
}

// * todo
func (ll *LinkedList) ReverseLinkedList() {

}

// * todo
func (ll *LinkedList) FindTarget() {

}

func (ll *LinkedList) PrintList() {
	head := ll.Head
	for head != nil {
		if head.Next == nil {
			fmt.Printf("Value: %d \n", head.Value)
		} else {
			fmt.Printf("Value: %d ----> ", head.Value)
		}
		head = head.Next
	}
}
