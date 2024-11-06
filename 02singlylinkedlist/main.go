package main

import "fmt"

// 链表中的结点
type Node struct {
	Val  int
	Next *Node
}

// 单向链表
type LinkedList struct {
	header *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{header: nil}
}

// 判断链表是否为空
func (l *LinkedList) Isempty() bool {
	return l.header == nil
}

// 打印链表
func (l *LinkedList) Print() {
	if l.Isempty() {
		fmt.Println("空链表")
		return
	}
	current := l.header
	for current != nil {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
}

// 返回链表长度
func (l *LinkedList) Length() int {
	length := 0
	current := l.header
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

// 在链表头插入元素
func (l *LinkedList) AddFromHead(val int) {
	if l.Isempty() {
		l.header = &Node{Val: val}
		return
	}
	node := &Node{Val: val, Next: l.header}
	l.header = node
}

// 在链表尾部插入元素
func (l *LinkedList) AddFromTail(val int) {
	node := &Node{
		Val: val,
	}
	if l.Isempty() {
		l.header = node
		return
	}
	current := l.header
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

// 在任意位置i处插入  i从0开始
func (l *LinkedList) Insert(val int, i int) {
	if i <= 0 {
		l.AddFromHead(val)
		return
	}
	if i >= l.Length() {
		l.AddFromTail(val)
		return
	}
	//找到要插入位置i的前一个节点
	preNode := l.header
	for i := 0; i < i-1; i++ {
		preNode = preNode.Next
	}
	node := &Node{
		Val:  val,
		Next: preNode.Next,
	}
	preNode.Next = node

}

//删除开头结点

//删除结尾结点

//删除任意位置i的结点

//查找链表中的某值

func main() {
	link := NewLinkedList()
	link.Insert(1, 0)
	link.Insert(2, 3)
	link.Insert(3, 2)
	link.Insert(10, 1)
	link.Print()

}
