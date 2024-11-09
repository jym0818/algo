package main

import "fmt"

//双向循环链表

type Node struct {
	Val        int
	Prev, Next *Node
}

type DoublyLoopLinkedList struct {
	header *Node
}

func NewLinkedList() *DoublyLoopLinkedList {
	return &DoublyLoopLinkedList{header: nil}

}

//是否为空

func (l *DoublyLoopLinkedList) IsEmpty() bool {
	return l.header == nil
}

//链表的长度

func (l *DoublyLoopLinkedList) Length() int {
	if l.IsEmpty() {
		return 0
	}
	current := l.header
	count := 1
	for current.Next != l.header {
		count++
		current = current.Next
	}
	return count
}

// 链表头插入
func (l *DoublyLoopLinkedList) AddFromHead(val int) {
	node := &Node{Val: val}
	if l.IsEmpty() {
		l.header = node
		node.Next = l.header
		node.Prev = l.header
		return
	}

	//找到最后一个结点 ，让它指向新插入的头节点
	current := l.header
	for current.Next != l.header {
		current = current.Next
	}

	current.Next = node
	node.Prev = current
	node.Next = l.header
	l.header.Prev = node
	l.header = node

}

// 双向循环链表尾部插入
func (l *DoublyLoopLinkedList) AddFromTail(val int) {
	node := &Node{Val: val}
	if l.IsEmpty() {
		l.header = node
		node.Next = l.header
		node.Prev = l.header
		return
	}
	//找到最后一个结点
	current := l.header
	for current.Next != l.header {
		current = current.Next
	}
	current.Next = node
	node.Prev = current
	node.Next = l.header
	l.header.Prev = node

}

// 任意位置插入

func (l *DoublyLoopLinkedList) Insert(val int, index int) {

	if index <= 0 {
		l.AddFromHead(val)
		return
	}
	if index >= l.Length() {
		l.AddFromTail(val)
		return
	}
	//找到结点i的位置
	current := l.header
	node := &Node{Val: val}
	for i := 0; i < index; i++ {
		current = current.Next
	}
	node.Prev = current.Prev
	node.Next = current
	current.Prev.Next = node
	current.Prev = node

}

//删除头节点

func (l *DoublyLoopLinkedList) RemoveOfHead() {
	if l.IsEmpty() {
		fmt.Println("链表为空")
		return
	}
	if l.header.Next == l.header {
		l.header = nil
		return
	}
	l.header.Prev.Next = l.header.Next
	l.header.Next.Prev = l.header.Prev
	l.header = l.header.Next
}

// 删除尾节点
func (l *DoublyLoopLinkedList) RemoveOfTail() {
	if l.IsEmpty() {
		fmt.Println("链表为空")
		return
	}
	if l.header.Next == l.header {
		l.header = nil
		return
	}
	lastNode := l.header.Prev
	lastNode.Prev.Next = l.header
	l.header.Prev = lastNode.Prev
}

// 删除任意结点
func (l *DoublyLoopLinkedList) Remove(index int) {
	if index <= 0 {
		l.RemoveOfHead()
		return
	}
	if index >= l.Length()-1 {
		l.RemoveOfTail()
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
	if l.IsEmpty() {
		return false
	}
	current := l.header
	for current.Next != l.header {
		if current.Val == val {
			return true
		}
		current = current.Next
	}
	if current.Val == val {
		return true
	}
	return false
}

//链表遍历

func (l *DoublyLoopLinkedList) Print() {
	if l.IsEmpty() {
		fmt.Println("链表为空")
		return
	}
	current := l.header

	for current.Next != l.header {
		fmt.Println("%d->", current.Val)
		current = current.Next
	}
	fmt.Println("%d->", current.Val)

}
