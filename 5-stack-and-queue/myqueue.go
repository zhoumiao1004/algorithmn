package main

import "fmt"

// 232.用栈实现队列
// https://leetcode.cn/problems/implement-queue-using-stacks/description/
type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			val := this.in[len(this.in)-1]
			this.in = this.in[:len(this.in)-1] // 出栈
			this.out = append(this.out, val)   // 入栈
		}
	}
	if len(this.out) == 0 {
		return 0
	}
	val := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]

	return val
}

func (this *MyQueue) Peek() int {
	v := this.Pop()
	this.out = append(this.out, v) // this.out是队列front，放回去
	return v
}

func (this *MyQueue) Empty() bool {
	return len(this.in)+len(this.out) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Empty())
}
