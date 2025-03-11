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

func (ll *LinkedList) ReverseLinkedList() {
	var prev *Node     // Pointer to track the previous node (initially nil)
	current := ll.Head // Start from the head of the linked list

	// Loop through the linked list
	for current != nil {
		next := current.Next // Store the next node before breaking the link

		current.Next = prev // Reverse the pointer (make current point to previous node)

		prev = current // Move prev forward (prev becomes the current node)
		current = next // Move current forward (to the stored next node)
	}

	// Swap head and tail because the list is now reversed
	ll.Tail, ll.Head = ll.Head, prev
}

func (ll *LinkedList) FindTarget(target int) {
	var (
		head         = ll.Head
		position int = 1
	)

	for head != nil {
		if target == head.Value {
			fmt.Printf("Found Target ! Position : %d", target)
		}
		position += 1
		head = head.Next
	}
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
