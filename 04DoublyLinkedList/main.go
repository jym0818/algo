package main

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

// 双端链表
type DoublyLinkedList struct {
	header *Node
	tailer *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	header := &Node{Val: -1} // 头哨兵
	tailer := &Node{Val: -1} // 尾哨兵
	header.Next = tailer
	tailer.Prev = header
	return &DoublyLinkedList{header: header, tailer: tailer}
}

//返回双向链表的长度

func (l *DoublyLinkedList) Length() int {
	length := 0
	current := l.header.Next
	for current != l.tailer {
		current = current.Next
		length++
	}
	return length
}

// 双端链表任意位置插入
func (l *DoublyLinkedList) Insert(val int, index int) {
	if index < 0 || index > l.Length() {
		return
	}
	//找到i位置，因为是双端链表
	current := l.header.Next
	for i := 0; i < index; i++ {
		current = current.Next
	}
	newNode := &Node{
		Val:  val,
		Prev: current.Prev,
		Next: current,
	}
	current.Prev.Next = newNode
	current.Prev = newNode

}

//删除任意结点

func (l *DoublyLinkedList) Remove(index int) {
	if index < 0 || index >= l.Length() {
		return
	}

	current := l.header.Next
	for i := 0; i < index; i++ {
		current = current.Next
	}
	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev

}

//是否包含值

func (l *DoublyLinkedList) Find(value int) bool {

	current := l.header.Next
	for current != l.tailer {
		if current.Val == value {
			return true
		}
		current = current.Next
	}

	return false

}

//遍历双端链表

func (l *DoublyLinkedList) Print() {

	//从头节点开始
	current := l.header
	for current.Next != l.tailer {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Printf("%d", current.Val)
}

func main() {
	l := NewDoublyLinkedList()

	l.Print()
}
