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
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{header: nil}
}

//判断是否为空

func (d *DoublyLinkedList) IsEmpty() bool {
	return d.header == nil
}

//返回双向链表的长度

func (d *DoublyLinkedList) Length() int {
	count := 0
	current := d.header
	for current != nil {
		current = current.Next
		count++
	}
	return count
}

// 从头部插入
func (d *DoublyLinkedList) AddFromHead(val int) {
	node := &Node{Val: val}
	if d.IsEmpty() {
		d.header = node
		return
	}
	//链表不为空
	d.header.Prev = node
	node.Next = d.header
	d.header = node
}

//从尾部插入

func (d *DoublyLinkedList) AddFromTail(val int) {
	node := &Node{Val: val}
	if d.IsEmpty() {
		d.header = node
		return
	}
	//找到尾结点
	current := d.header
	for current.Next != nil {
		current = current.Next
	}
	node.Prev = current
	current.Next = node
}

// 双端链表任意位置插入
func (d *DoublyLinkedList) Insert(val int, index int) {
	if index <= 0 {
		d.AddFromHead(val)
		return
	}
	if index >= d.Length() {
		d.AddFromTail(val)
		return
	}
	//找到i位置，因为是双端链表
	node := &Node{Val: val}
	current := d.header
	for i := 0; i < index; i++ {
		current = current.Next
	}
	node.Next = current
	node.Prev = current.Prev
	current.Prev.Next = node
	current.Prev = node

}

//删除头节点

func (d *DoublyLinkedList) RemoveOfHead() {
	if d.IsEmpty() {
		fmt.Println("链表为空")
		return
	}
	//只有1个结点的情况
	if d.header.Next == nil {
		d.header = nil
		return
	}

	d.header.Next.Prev = nil
	d.header = d.header.Next

}

// 删除尾结点
func (d *DoublyLinkedList) RemoveOfTail() {
	if d.IsEmpty() {
		fmt.Println("链表为空")
		return
	}
	//只有1个结点的情况
	if d.header.Next == nil {
		d.header = nil
		return
	}
	current := d.header
	for current.Next != nil {
		current = current.Next
	}
	current.Prev.Next = nil

}

//删除任意结点

func (d *DoublyLinkedList) Remove(index int) {
	if index <= 0 {
		d.RemoveOfHead()
		return
	}
	if index >= d.Length()-1 {
		d.RemoveOfTail()
		return
	}

	current := d.header
	for i := 0; i < index; i++ {
		current = current.Next
	}
	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev

}

//是否包含值

func (d *DoublyLinkedList) Find(value int) bool {
	if d.IsEmpty() {
		return false
	}
	current := d.header
	flag := false
	for current != nil {
		if current.Val == value {
			flag = true
			break
		}
		current = current.Next
	}

	return flag

}

//遍历双端链表

func (d *DoublyLinkedList) Print() {

	if d.IsEmpty() {
		fmt.Println("链表为空")
		return
	}
	//从头节点开始
	current := d.header
	for current.Next != nil {
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Printf("%d", current.Val)
}

func main() {
	l := NewDoublyLinkedList()
	l.AddFromHead(1)
	l.AddFromHead(2)
	l.AddFromHead(3)
	l.Print()
}
