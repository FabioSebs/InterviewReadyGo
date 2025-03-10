package leetcode

// take an array and find the indices that add up to be the desired target
func TwoSum(array []int, target int) []int {
	var (
		possibleAnswers map[int]int = map[int]int{}
	)

	for i, num := range array {
		complement := target - num
		if mapIdx, found := possibleAnswers[complement]; found {
			return []int{mapIdx, i}
		}
		possibleAnswers[num] = i
	}
	return []int{}
}
