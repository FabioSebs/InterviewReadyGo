package leetcode

type PrefixHelper struct {
	Letters []rune
	Index   []int
}

func LongestCommonPrefix(strs []string) string {
	var (
		lowestWordCount int    = 1000
		prefixObserver         = map[int]string{}
		prefix          string = ""
	)

	for _, word := range strs {
		if lowestWordCount > len(word)-1 {
			lowestWordCount = len(word) - 1
		}

		for idx, letter := range word {
			prefixObserver[idx] += string(letter)
		}
	}

	for idx := 0; idx <= lowestWordCount; idx++ {
		passed := true
		firstLetter := prefixObserver[idx][0]
		for cnt, ltr := range prefixObserver[idx] {
			if cnt == 0 {
				continue
			}
			if rune(firstLetter) != ltr {
				passed = false
			}
		}
		if passed {
			prefix += string(firstLetter)
		}
	}

	return prefix
}
