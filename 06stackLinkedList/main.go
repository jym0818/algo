package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	top    *Node
	length int
}

// 构造函数
func NewStack() *Stack {
	return &Stack{top: nil, length: 0}
}

// 返回栈的长度
func (s *Stack) len() int {
	return s.length
}

// 入栈
func (s *Stack) Push(val int) {
	node := &Node{Val: val}
	node.Next = s.top
	s.top = node
	s.length++
}

// 出栈
func (s *Stack) Pop() int {
	if s.len() == 0 {
		fmt.Println("栈为空")
		return math.MaxInt
	}
	val := s.top.Val
	s.top = s.top.Next
	s.length--
	return val
}

//访问栈顶元素

func (s *Stack) Peek() int {
	if s.len() == 0 {
		fmt.Println("栈为空")
		return math.MaxInt
	}
	return s.top.Val
}
