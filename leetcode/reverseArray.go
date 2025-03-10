package leetcode

func ReverseSlice(array []int) []int {
	var (
		reversed = []int{}
	)

	iterator := len(array) - 1
	for iterator != -1 {
		reversed = append(reversed, array[iterator])
		iterator -= 1
	}

	return reversed
}

func ReverseSliceSwap(array []int) {
	left, right := 0, len(array)-1

	for left < right {
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}
}
