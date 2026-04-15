package main

// 622. 设计循环队列
// 设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。
// 循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。
type MyCircularQueue struct {
	
}

func Constructor(k int) MyCircularQueue {

}

func (this *MyCircularQueue) EnQueue(value int) bool {

}

func (this *MyCircularQueue) DeQueue() bool {

}

func (this *MyCircularQueue) Front() int {

}

func (this *MyCircularQueue) Rear() int {

}

func (this *MyCircularQueue) IsEmpty() bool {

}

func (this *MyCircularQueue) IsFull() bool {

}

func main() {
	circularQueue := Constructor(3)  // 设置长度为 3
	circularQueue.EnQueue(1)   // 返回 true
	circularQueue.EnQueue(2)   // 返回 true
	circularQueue.EnQueue(3)   // 返回 true
	circularQueue.EnQueue(4)   // 返回 false，队列已满
	circularQueue.Rear()   // 返回 3
	circularQueue.IsFull()   // 返回 true
	circularQueue.DeQueue()   // 返回 true
	circularQueue.EnQueue(4)   // 返回 true
	circularQueue.Rear()   // 返回 4
}