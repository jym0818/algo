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
	return &LoopLinkedList{
		header: nil,
	}
}

// 判断环形链表是否为空
func (l *LoopLinkedList) IsEmpty() bool {
	return l.header == nil
}

//链表长度

func (l *LoopLinkedList) Length() int {
	if l.IsEmpty() {
		return 0
	}
	length := 1
	current := l.header
	for current.Next != l.header {
		current = current.Next
		length++
	}
	return length

}

// 头部插入元素
func (l *LoopLinkedList) AddFromHead(val int) {
	node := &Node{
		Val: val,
	}
	if l.IsEmpty() {
		l.header = node
		node.Next = l.header
		return
	}
	//找到最后一个node，让他指向新的元素
	current := l.header
	for current.Next != l.header {
		current = current.Next
	}
	current.Next = node
	node.Next = l.header
	l.header = node

}

//尾部插入元素

func (l *LoopLinkedList) AddFromTail(val int) {
	node := &Node{Val: val}
	if l.IsEmpty() {
		l.header = node
		node.Next = l.header
		return
	}

	//找到尾结点
	current := l.header
	for current.Next != l.header {
		current = current.Next
	}
	current.Next = node
	node.Next = l.header

}

//任意位置插入元素 index从 0开始

func (l *LoopLinkedList) Insert(val int, index int) {
	if index <= 0 {
		l.AddFromHead(val)
		return
	}
	if index >= l.Length() {
		l.AddFromTail(val)
		return
	}
	current := l.header
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	node := &Node{Val: val}
	node.Next = current.Next
	current.Next = node

}

// 删除头结点
func (l *LoopLinkedList) ReomveOfHead() {
	if l.IsEmpty() {
		fmt.Println("链表为空")
		return
	}
	//让头节点指向下个结点并且让最后一个结点也指向下个结点
	current := l.header
	//只有1个结点
	if current.Next == l.header {
		l.header = nil
		return
	}
	for current.Next != l.header {
		current = current.Next
	}

	current.Next = l.header.Next
	l.header = l.header.Next

}

// 删除尾部元素
func (l *LoopLinkedList) RemoveOfTail() {
	if l.IsEmpty() {
		fmt.Println("链表为空")
		return
	}

	//找到倒数第二个结点 ，让它指向头节点
	current := l.header
	//如果链表只有1个元素
	if current.Next == l.header {
		l.header = nil
		return
	}
	for current.Next.Next != l.header {
		current = current.Next
	}
	current.Next = l.header
}

// 删除任意位置元素
func (l *LoopLinkedList) Remove(index int) {
	if l.IsEmpty() {
		fmt.Println("Loop Linked List is Empty")
		return
	}
	if index <= 0 {
		l.ReomveOfHead()
		return
	}
	if index >= l.Length()-1 {
		l.RemoveOfTail()
		return
	}
	current := l.header
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next

}

//查找元素

func (l *LoopLinkedList) Find(val int) bool {
	if l.IsEmpty() {
		return false
	}
	current := l.header
	flag := false

	for current.Next != l.header {
		if current.Val == val {
			flag = true
			break
		}
		current = current.Next
	}
	if current.Val == val {
		flag = true
	}
	return flag
}

//遍历链表

func (l *LoopLinkedList) Print() {
	if l.IsEmpty() {
		fmt.Println("链表为空")
		return
	}
	current := l.header

	for current.Next != l.header {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Printf("%d->", current.Val)

}

func main() {
	l := NewLoopLinkedList()
	l.Insert(1, 0)
	l.Insert(2, -1)
	l.Insert(3, 4)
	l.Insert(4, 6)
	l.Insert(10, 2)

	l.Print()

}
