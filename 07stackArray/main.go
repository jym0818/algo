package main

import (
	"fmt"
	"math"
)

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{data: make([]int, 0, 16)}
}

func (s *Stack) Length() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}
func (s *Stack) Pop() int {
	val := s.Peek()
	if s.IsEmpty() {
		return val
	}
	length := s.Length()
	s.data = s.data[0 : length-1]
	return val
}
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		fmt.Println("栈为空")
		return math.MaxInt
	}
	length := s.Length()
	return s.data[length-1]
}
