package main

import "fmt"

// 895. 最大频率栈
// https://leetcode.cn/problems/maximum-frequency-stack/description/
// 设计一个类似堆栈的数据结构，将元素推入堆栈，并从堆栈中弹出出现频率最高的元素。
// 实现 FreqStack 类:
// FreqStack() 构造一个空的堆栈。
// void push(int val) 将一个整数 val 压入栈顶。
// int pop() 删除并返回堆栈中出现频率最高的元素。
// 如果出现频率最高的元素不只一个，则移除并返回最接近栈顶的元素。
type FreqStack struct {
	maxFreq    int
	valToFreq  map[int]int
	freqToVals map[int][]int
}

func Constructor() FreqStack {
	return FreqStack{
		valToFreq:  make(map[int]int),
		freqToVals: make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {
	this.valToFreq[val]++
	freq := this.valToFreq[val]
	this.freqToVals[freq] = append(this.freqToVals[freq], val)
	this.maxFreq = max(this.maxFreq, freq)
}

func (this *FreqStack) Pop() int {
	vals := this.freqToVals[this.maxFreq]
	v := vals[len(vals)-1] // 取出频率最高的元素中离栈顶最近的一个
	this.freqToVals[this.maxFreq] = vals[:len(vals)-1]
	this.valToFreq[v]--
	if len(this.freqToVals[this.maxFreq]) == 0 {
		delete(this.freqToVals, this.maxFreq)
		this.maxFreq--
	}
	return v
}

func main() {
	freqStack := Constructor()
	freqStack.Push(5)            //堆栈为 [5]
	freqStack.Push(7)            //堆栈是 [5,7]
	freqStack.Push(5)            //堆栈是 [5,7,5]
	freqStack.Push(7)            //堆栈是 [5,7,5,7]
	freqStack.Push(4)            //堆栈是 [5,7,5,7,4]
	freqStack.Push(5)            //堆栈是 [5,7,5,7,4,5]
	fmt.Println(freqStack.Pop()) //返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
	fmt.Println(freqStack.Pop()) //返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
	fmt.Println(freqStack.Pop()) //返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
	fmt.Println(freqStack.Pop()) //返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]。
}
