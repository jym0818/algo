package main

import "fmt"

//双向循环链表

type Node struct {
	Val        int
	Prev, Next *Node
}

type DoublyLoopLinkedList struct {
	header *Node
	tailer *Node
}

func NewLinkedList() *DoublyLoopLinkedList {
	head := &Node{Val: 0}
	tail := &Node{Val: 0}
	head.Next = tail
	head.Prev = tail
	tail.Prev = head
	tail.Next = head
	return &DoublyLoopLinkedList{header: head, tailer: tail}

}

//链表的长度

func (l *DoublyLoopLinkedList) Length() int {
	length := 0
	current := l.header.Next
	for current != l.tailer {
		length++
		current = current.Next
	}
	return length
}

// 任意位置插入

func (l *DoublyLoopLinkedList) Insert(val int, index int) {
	if index < 0 || index > l.Length() {
		return
	}
	//找到结点i的位置
	current := l.header.Next
	for i := 0; i < index; i++ {
		current = current.Next
	}
	node := &Node{Val: val}

	node.Prev = current.Prev
	node.Next = current
	current.Prev.Next = node
	current.Prev = node

}

// 删除任意结点
func (l *DoublyLoopLinkedList) Remove(index int) {
	if index < 0 || index >= l.Length() || l.Length() == 0 {
		return
	}
	//找到这个结点
	current := l.header

	for i := 0; i < index; i++ {
		current = current.Next
	}
	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev

}

// 是否包含值
func (l *DoublyLoopLinkedList) Find(val int) bool {

	current := l.header.Next
	for current != l.tailer {
		if current.Val == val {
			return true
		}
		current = current.Next
	}

	return false
}

//链表遍历

func (l *DoublyLoopLinkedList) Print() {

	current := l.header.Next

	for current != l.tailer {
		fmt.Println("%d->", current.Val)
		current = current.Next
	}
	fmt.Println("%d->", current.Val)

}
