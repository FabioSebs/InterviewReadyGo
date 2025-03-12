package datastructures

import "fmt"

type Stack interface {
	Push(val int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	PrintStack()
}

type stack []int

func NewStack() Stack {
	return &stack{}
}

func (s *stack) Push(val int) {
	*s = append(*s, val)
}

func (s *stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}

	topIndex := len(*s) - 1
	val := (*s)[topIndex]

	*s = (*s)[:topIndex]
	return val, nil
}

func (s *stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) PrintStack() {
	for idx := len(*s) - 1; idx >= 0; idx-- {
		fmt.Println((*s)[idx])
	}
}
