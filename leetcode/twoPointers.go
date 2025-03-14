package leetcode

// understanding two pointers
// - left pointer
// - right pointer
// - move the two points inwards

// 1. Finding a pair that sums to a target in a sorted array
func SortedExercise(nums []int, target int) []int {
	// create pointers
	left, right := 0, len(nums)-1

	// create the loop so that the pointers meet
	for left < right {
		sum := nums[left] - nums[right]

		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	// no pairs found
	return []int{}
}

// 2. Checking if a string is a Palindrome
func IsPalindrome(s string) bool {
	// create pointers
	left, right := 0, len(s)-1

	// create the loop so that the pointers meet
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// 3. Removing duplicates from a sorted array
func RemoveDuplicates(nums []int) []int {
	var (
		newArray []int
	)

	valueMap := map[int]int{}

	// create pointers
	left, right := 0, len(nums)-1

	// create the loop so that the pointers meet
	for left < right {
		valueMap[nums[left]] += 1
		valueMap[nums[right]] += 1
		left++
		right--
	}

	for key, _ := range valueMap {
		newArray = append(newArray, key)
	}

	return BubbleSort(newArray)
}

func BubbleSort(nums []int) []int {
	n := len(nums)

	// typical for loop
	for i := 0; i < n-1; i++ {
		// second for loop for checking comparison
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				// swap adjacent elements if they are in the wrong order
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
