package main

import "github.com/FabioSebs/InterviewReadyGo/datastructures"

func main() {
	dll := datastructures.NewDLL()
	dll.AddToList(1)
	dll.AddToList(2)
	dll.AddToList(3)
	dll.PrintListBackwards()
}
