package main

import (
	"container/list"
	"fmt"
)

type MyLinkedQueue struct {
	list *list.List
}

func NewMyLinkedQueue() *MyLinkedQueue {
	return &MyLinkedQueue{list: list.New()}
}

// 向队尾插入元素，时间复杂度 O(1)
func (q *MyLinkedQueue) Push(e interface{}) {
	q.list.PushBack(e)
}

// 从队头删除元素，时间复杂度 O(1)
func (q *MyLinkedQueue) Pop() interface{} {
	front := q.list.Front()
	if front != nil {
		return q.list.Remove(front)
	}
	return nil
}

// 查看队头元素，时间复杂度 O(1)
func (q *MyLinkedQueue) Peek() interface{} {
	front := q.list.Front()
	if front != nil {
		return front.Value
	}
	return nil
}

func (q *MyLinkedQueue) Size() int {
	return q.list.Len()
}

func main() {
	queue := NewMyLinkedQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Peek()) // 1
	fmt.Println(queue.Pop())  // 1
	fmt.Println(queue.Pop())  // 2
	fmt.Println(queue.Peek()) // 3
}
