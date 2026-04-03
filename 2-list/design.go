package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List interface {
	Get(index int) int
	AddAtHead(val int)
	AddAtTail(val int)
	AddAtIndex(index int, val int)
	DeleteAtIndex(index int)
}

func newMyLinkedList() List {
	return &MyLinkedList{
		dummy: &ListNode{},
		Size:  0,
	}
}

// 707.设计链表
// https://leetcode.cn/problems/design-linked-list/description/
type MyLinkedList struct {
	dummy *ListNode
	Size  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummy: &ListNode{},
		Size:  0,
	}
}

func (l *MyLinkedList) String() string {
	var vals []string
	cur := l.dummy.Next
	for cur != nil {
		vals = append(vals, fmt.Sprintf("%d", cur.Val))
		cur = cur.Next
	}
	return strings.Join(vals, "->")
}
func (l *MyLinkedList) Get(index int) int {
	if l == nil || index < 0 || index >= l.Size {
		return -1
	}
	cur := l.dummy.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{Val: val}
	node.Next = l.dummy.Next
	l.dummy.Next = node
	l.Size++
}

func (l *MyLinkedList) AddAtTail(val int) {
	cur := l.dummy
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
	l.Size++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > l.Size { // 注意这里是大于
		return
	}
	cur := l.dummy
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	node := &ListNode{Val: val, Next: cur.Next}
	cur.Next = node
	l.Size++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if l == nil || index < 0 || index >= l.Size {
		return
	}
	cur := l.dummy
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	l.Size--
}

func main() {
	myList := Constructor()
	myList.AddAtTail(3)
	myList.AddAtHead(2)
	myList.AddAtHead(1)
	myList.AddAtTail(4)
	myList.AddAtTail(5)
	fmt.Println(myList.String())
	fmt.Println(myList.Get(1))
	myList.AddAtIndex(2, 9)
	fmt.Println(myList.String())
	myList.DeleteAtIndex(4)
	fmt.Println(myList.String())
}
