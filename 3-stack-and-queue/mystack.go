package main

// 225. 用队列实现栈
// https://leetcode.cn/problems/implement-stack-using-queues/description/
type MyStack struct {
	que []int // slice模拟队列，只能左进右出
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.que = append(this.que, x)
}

func (this *MyStack) Pop() int {
	// 要弹出最后一个进入的元素，思路：前n-1个出队再入队，然后队首出队
	n := len(this.que)
	for i := 0; i < n-1; i++ {
		e := this.que[0]
		this.que = this.que[1:]        // 出队
		this.que = append(this.que, e) // 加入队尾
	}
	val := this.que[0]
	this.que = this.que[1:]
	return val
}

func (this *MyStack) Top() int {
	return this.que[len(this.que)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.que) == 0
}
