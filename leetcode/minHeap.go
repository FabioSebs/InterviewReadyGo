package leetcode

import (
	"container/heap"
	"math"
)

// Define a struct to store the element value, its index in the list, and which list it belongs to
type Element struct {
	val  int
	idx  int
	list int
}

// MinHeap implementation
type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val } // Min-Heap based on value
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func smallestRange(nums [][]int) []int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	maxVal := math.MinInt32 // Track the maximum value in the heap
	rangeStart, rangeEnd := 0, math.MaxInt32

	// Step 1: Initialize heap with the first element of each list
	for i := 0; i < len(nums); i++ {
		heap.Push(minHeap, Element{val: nums[i][0], idx: 0, list: i})
		if nums[i][0] > maxVal {
			maxVal = nums[i][0]
		}
	}

	// Step 2: Extract the min element and insert the next element from the same list
	for minHeap.Len() == len(nums) { // Ensure we always have one element from each list
		minElem := heap.Pop(minHeap).(Element)

		// Update range if the new range is smaller
		if maxVal-minElem.val < rangeEnd-rangeStart {
			rangeStart, rangeEnd = minElem.val, maxVal
		}

		// Move to the next element in the same list
		if minElem.idx+1 < len(nums[minElem.list]) {
			nextElem := nums[minElem.list][minElem.idx+1]
			heap.Push(minHeap, Element{val: nextElem, idx: minElem.idx + 1, list: minElem.list})

			// Update maxVal
			if nextElem > maxVal {
				maxVal = nextElem
			}
		} else {
			break // If any list is exhausted, stop since we can't include all lists anymore
		}
	}

	return []int{rangeStart, rangeEnd}
}
