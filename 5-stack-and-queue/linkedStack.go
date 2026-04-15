package main

import (
	"container/list"
	"fmt"
)

type MyLinkedStack struct {
	list *list.List
}

func NewMyLinkedStack() *MyLinkedStack {
	return &MyLinkedStack{list: list.New()}
}

// 向栈顶加入元素，时间复杂度 O(1)
func (s *MyLinkedStack) Push(e interface{}) {
	s.list.PushBack(e)
}

// 从栈顶弹出元素，时间复杂度 O(1)
func (s *MyLinkedStack) Pop() interface{} {
	element := s.list.Back()
	if element != nil {
		s.list.Remove(element)
		return element.Value
	}
	return nil
}

// 查看栈顶元素，时间复杂度 O(1)
func (s *MyLinkedStack) Peek() interface{} {
	element := s.list.Back()
	if element != nil {
		return element.Value
	}
	return nil
}

// 返回栈中的元素个数，时间复杂度 O(1)
func (s *MyLinkedStack) Size() int {
	return s.list.Len()
}

func main() {
	stack := NewMyLinkedStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())  // 3
	fmt.Println(stack.Pop())  // 2
	fmt.Println(stack.Peek()) // 1
}
