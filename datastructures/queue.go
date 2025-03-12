package datastructures

import "fmt"

type Queue interface {
	Push(val int)
	Pop() int
	PrintQueue()
}

type queue []int

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Push(val int) {
	*q = append(*q, val)
}

func (q *queue) Pop() int {
	switch len(*q) {
	case 0:
		return 0
	case 1:
		pop := (*q)[0]
		*q = []int{}
		return pop
	default:
		pop := (*q)[0]
		*q = (*q)[1:]
		return pop
	}
}

func (q *queue) PrintQueue() {
	for _, val := range *q {
		fmt.Println(val)
	}
}
