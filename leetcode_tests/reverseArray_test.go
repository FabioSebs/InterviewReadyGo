package leetcodetests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/FabioSebs/InterviewReadyGo/leetcode"
)

func TestReverseArray(t *testing.T) {
	var (
		tests = []struct {
			array  []int
			answer []int
		}{
			{
				array:  []int{1, 2, 3},
				answer: []int{3, 2, 1},
			},
			{
				array:  []int{5, 3, 2},
				answer: []int{2, 3, 5},
			},
			{
				array:  []int{10, 9, 8},
				answer: []int{8, 9, 10},
			},
		}
	)

	for testNo, testCase := range tests {
		testCase := testCase
		t.Run(fmt.Sprintf("test case reverse slice %d", testNo), func(t *testing.T) {
			t.Parallel()

			answer := leetcode.ReverseSlice(testCase.array)

			if !reflect.DeepEqual(answer, testCase.answer) {
				t.Errorf("got %v, want %v", answer, testCase.answer)
			}
		})
	}

	for testNo, testCase := range tests {
		testCase := testCase
		t.Run(fmt.Sprintf("test case reverse slice swap %d", testNo), func(t *testing.T) {
			t.Parallel()

			answer := testCase.array
			leetcode.ReverseSliceSwap(answer)

			if !reflect.DeepEqual(answer, testCase.answer) {
				t.Errorf("got %v, want %v", answer, testCase.answer)
			}
		})
	}
}
