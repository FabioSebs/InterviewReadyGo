package main

import "github.com/FabioSebs/InterviewReadyGo/leetcode"

func main() {
	myLinkedList := leetcode.NewLinkedList(1)
	myLinkedList.AddToList(2)
	myLinkedList.AddToList(3)
	myLinkedList.PrintList()
	myLinkedList.PrintList()
}
