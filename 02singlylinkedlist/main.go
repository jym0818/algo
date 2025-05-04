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
	return &LinkedList{header: &Node{Val: -1, Next: nil}}
}

// 打印链表
func (l *LinkedList) Print() {
	current := l.header.Next
	for current != nil {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

// 返回链表长度
func (l *LinkedList) Length() int {
	length := 0
	current := l.header.Next
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

// 在任意位置i处插入  i从0开始
func (l *LinkedList) Insert(val int, index int) {
	if index < 0 || index > l.Length() {
		fmt.Println("边界错误")
	}
	//找到要插入位置i的前一个节点
	preNode := l.header
	for i := 0; i < index; i++ {
		preNode = preNode.Next
	}
	node := &Node{
		Val:  val,
		Next: preNode.Next,
	}
	preNode.Next = node

}

// 删除任意位置i的结点
func (l *LinkedList) Remove(index int) {
	if index < 0 || index >= l.Length() || l.Length() == 0 {
		return // 处理无效索引和空链表
	}
	preNode := l.header
	for i := 0; i < index; i++ {
		preNode = preNode.Next
	}
	preNode.Next = preNode.Next.Next
}

//查找结点中是否包含某值

func (l *LinkedList) Find(val int) bool {

	current := l.header.Next

	for current != nil {
		if current.Val == val {
			return true
		}
		current = current.Next
	}
	return false

}

func main() {
	l := NewLinkedList()
	l.Print()
	fmt.Println()
	l.Print()
}
