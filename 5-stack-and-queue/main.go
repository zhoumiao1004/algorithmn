package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
)

// 232.用栈实现队列
// https://leetcode.cn/problems/implement-queue-using-stacks/description/
type MyQueue struct {
	in  []int
	out []int
}

func Constructor1() MyQueue {
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

// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses/description/
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 1.左括号必须用相同类型的右括号闭合。
// 2.左括号必须以正确的顺序闭合。
// 3.每个右括号都有一个对应的相同类型的左括号。
// 输入：s = "()" 输出：true
func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	var st []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			st = append(st, c)
			continue
		}
		if len(st) == 0 {
			return false
		}
		if m[c] != st[len(st)-1] {
			return false
		}
		st = st[:len(st)-1]
	}
	return len(st) == 0
}

// 1047. 删除字符串中的所有相邻重复项
// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
// 输入："abbaca"
// 输出："ca"
func removeDuplicates(s string) string {
	var st []byte
	for i := 0; i < len(s); i++ {
		if len(st) > 0 && s[i] == st[len(st)-1] {
			st = st[:len(st)-1] // 出栈
		} else {
			st = append(st, s[i])
		}
	}
	return string(st)
}

// 150. 逆波兰表达式求值
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
// 有效的算符为 '+'、'-'、'*' 和 '/' 。
// 输入：tokens = ["2","1","+","3","*"]
// 输出：9
// 解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
func evalRPN(tokens []string) int {
	r := 0
	var st []string
	for _, s := range tokens {
		if s == "+" || s == "-" || s == "*" || s == "/" {
			a := st[len(st)-1]
			st = st[:len(st)-1] // 出栈
			b := st[len(st)-1]
			st = st[:len(st)-1] // 出栈
			v2, _ := strconv.Atoi(a)
			v1, _ := strconv.Atoi(b)
			if s == "+" {
				r = v1 + v2
			} else if s == "-" {
				r = v1 - v2
			} else if s == "*" {
				r = v1 * v2
			} else if s == "/" {
				r = v1 / v2
			}
			st = append(st, fmt.Sprintf("%d", r))
		} else {
			st = append(st, s)
		}
	}
	if len(st) > 0 {
		r, _ = strconv.Atoi(st[0])
	}
	return r
}

// 239. 滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum/description/
// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回 滑动窗口中的最大值 。
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//
//	1 [3  -1  -3] 5  3  6  7       3
//	1  3 [-1  -3  5] 3  6  7       5
//	1  3  -1 [-3  5  3] 6  7       5
//	1  3  -1  -3 [5  3  6] 7       6
//	1  3  -1  -3  5 [3  6  7]      7
//
// 封装单调队列的方式解题
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums) - k + 1
	result := make([]int, n)
	q := MonoQueue{}
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	result[0] = q.Front()
	for i := k; i < len(nums); i++ {
		q.Pop(nums[i-k]) // 移除最前面的元素
		q.Push(nums[i])  // 添加最后面的元素
		result[i-k+1] = q.Front()
	}
	return result
}

type MonoQueue struct {
	deque []int
}

func (m *MonoQueue) Front() int {
	return m.deque[0]
}

func (m *MonoQueue) Back() int {
	return m.deque[len(m.deque)-1]
}

func (m *MonoQueue) Empty() bool {
	return len(m.deque) == 0
}

func (m *MonoQueue) Push(val int) {
	// 从后往前把小于val的元素都弹出
	for !m.Empty() && val > m.Back() {
		m.deque = m.deque[:len(m.deque)-1]
	}
	m.deque = append(m.deque, val)
}

func (m *MonoQueue) Pop(val int) {
	// 由于小的队尾的元素已经在push的时候被卷走了，只需要判断pop的是不是队首的最大元素
	if !m.Empty() && val == m.Front() {
		m.deque = m.deque[1:]
	}
}

func maxSlidingWindow2(nums []int, k int) []int {
	var q []int
	n := len(nums)
	if n == 0 || n < k {
		return []int{}
	}
	result := make([]int, n-k+1)
	for i := 0; i < n; i++ {
		// 遗弃的是最大值
		if i >= k && nums[i-k] == q[0] {
			q = q[1:]
		}
		for len(q) > 0 && nums[i] > q[0] {
			q = q[1:]
		}
		q = append(q, nums[i])
		if i >= k {
			result[i-k] = q[0]
		}
	}
	return result
}

// 347. 前 K 个高频元素
// https://leetcode.cn/problems/top-k-frequent-elements/description/
// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
// 输入：nums = [1,1,1,2,2,3], k = 2
// 输出：[1,2]
// 方法1:排序O(nlogn)
func topKFrequent(nums []int, k int) []int {
	var results []int
	cntMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cntMap[nums[i]]++
	}
	for key := range cntMap {
		results = append(results, key)
	}

	sort.Slice(results, func(i, j int) bool {
		return cntMap[results[i]] > cntMap[results[j]]
	})
	return results[:k]
}

// 方法2:小顶堆
// 时间复杂度: O(nlogk)
// 空间复杂度: O(n)
func topKFrequent2(nums []int, k int) []int {
	m := make(map[int]int)
	//记录每个元素出现的次数
	for _, val := range nums {
		m[val]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, val := range m {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, k)
	//按顺序返回堆中的元素
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).([2]int)[0]
	}
	return result
}

// 构建小顶堆
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1] // 小顶堆
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	obj := Constructor1()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Empty())
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(topKFrequent2([]int{1, 1, 1, 2, 2, 3}, 2))
}
