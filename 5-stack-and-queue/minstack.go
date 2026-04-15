package main

import (
	"fmt"
	"math"
)

// 155. 最小栈
// https://leetcode.cn/problems/min-stack/description/
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// 实现 MinStack 类:
// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。
type MinStack struct {
	st    []int
	minSt []int // 记录入栈时的最小值
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.st = append(this.st, val)
	this.minSt = append(this.minSt, min(val, this.GetMin()))
}

func (this *MinStack) Pop() {
	this.st = this.st[:len(this.st)-1]
	this.minSt = this.minSt[:len(this.minSt)-1]
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minSt) == 0 {
		return math.MaxInt
	}
	return this.minSt[len(this.minSt)-1]
}

type MinStack2 struct {
	st    []int
	minSt []int // 记录入栈时的最小值
}

func Constructor2() MinStack2 {
	return MinStack2{}
}

func (this *MinStack2) Push(val int) {
	this.st = append(this.st, val)
	if len(this.st) == 0 || val <= this.GetMin() {
		this.minSt = append(this.minSt, val)
	}
}

func (this *MinStack2) Pop() {
	val := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]
	if val == this.minSt[len(this.minSt)-1] {
		this.minSt = this.minSt[:len(this.minSt)-1] // 出栈的是最小值
	}
}

func (this *MinStack2) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack2) GetMin() int {
	return this.minSt[len(this.minSt)-1]
}

func main() {
	minStack := Constructor2()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // --> 返回 -3.
	minStack.Pop()
	fmt.Println(minStack.Top())    // --> 返回 0.
	fmt.Println(minStack.GetMin()) // --> 返回 -2.
}
