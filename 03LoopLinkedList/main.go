package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LoopLinkedList struct {
	header *Node
}

func NewLoopLinkedList() *LoopLinkedList {
	header := &Node{Val: 0}
	header.Next = header // 关键闭环操作
	return &LoopLinkedList{header: header}
}

// 链表长度
func (l *LoopLinkedList) Length() int {
	length := 0
	current := l.header.Next
	for current != l.header {
		length++
		current = current.Next
	}
	return length

}

//任意位置插入元素 index从 0开始

func (l *LoopLinkedList) Insert(val int, index int) {
	if index < 0 || index > l.Length() {
		return
	}

	current := l.header
	for i := 0; i < index; i++ {
		current = current.Next
	}
	node := &Node{Val: val, Next: current.Next}
	current.Next = node
}

// 删除任意位置元素
func (l *LoopLinkedList) Remove(index int) {
	if index < 0 || index >= l.Length() || l.Length() == 0 {
		return // 索引越界直接返回
	}

	current := l.header
	for i := 0; i < index; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next // 自动保持闭环

}

//查找元素

func (l *LoopLinkedList) Find(val int) bool {
	current := l.header.Next // 从第一个实际节点开始
	for current != l.header {
		if current.Val == val {
			return true
		}
		current = current.Next
	}
	return false
}

//遍历链表

func (l *LoopLinkedList) Print() {
	if l.header.Next == l.header {
		fmt.Println("空链表")
		return
	}

	current := l.header.Next
	for current != l.header {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Println("头哨兵")
}

func main() {
	l := NewLoopLinkedList()

	l.Print()

}
