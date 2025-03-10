package leetcodetests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/FabioSebs/InterviewReadyGo/leetcode"
)

func TestTwoSum(t *testing.T) {
	var (
		tests = []struct {
			array  []int
			target int
			answer []int
		}{
			{
				array:  []int{1, 2, 3, 4, 5},
				target: 3,
				answer: []int{0, 1},
			},
			{
				array:  []int{11, 4, 6, 8, 9},
				target: 20,
				answer: []int{0, 4},
			},
			{
				array:  []int{9, 4, 6, 2, 1, 5},
				target: 8,
				answer: []int{2, 3},
			},
		}
	)

	for idx, testCase := range tests {
		t.Run(fmt.Sprintf("test case %d", idx+1), func(t *testing.T) {
			answer := leetcode.TwoSum(testCase.array, testCase.target)

			if !reflect.DeepEqual(answer, testCase.answer) {
				t.Errorf("got %v, want %v", answer, testCase.answer)
			}
		})
	}
}
