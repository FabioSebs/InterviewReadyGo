package main

import (
	"fmt"

	"github.com/FabioSebs/InterviewReadyGo/leetcode"
)

func main() {
	testArray := []string{"flower", "flow", "flight"}

	fmt.Println(leetcode.LongestCommonPrefix(testArray))
}
