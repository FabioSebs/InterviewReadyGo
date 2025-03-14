package main

import (
	"fmt"

	"github.com/FabioSebs/InterviewReadyGo/leetcode"
)

func main() {
	testArray := []int{1, 1, 2, 3, 4, 5, 5, 6}
	testArray2 := []int{1, 2, 3, 4, 5}

	fmt.Println(leetcode.RemoveDuplicates(testArray))
	fmt.Println(leetcode.RemoveDuplicates(testArray2))
}
