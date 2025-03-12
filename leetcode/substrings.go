package leetcode

func NumberOfSubstrings(s string) int {
	count := map[byte]int{'a': 0, 'b': 0, 'c': 0} // To track occurrences of 'a', 'b', and 'c'
	left, result := 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]]++ // Expand window by adding right character

		// While the window contains at least one 'a', 'b', and 'c'
		for count['a'] > 0 && count['b'] > 0 && count['c'] > 0 {
			result += len(s) - right // All substrings starting from left to the end are valid
			count[s[left]]--         // Shrink window from left
			left++                   // Move left pointer
		}
	}

	return result
}
