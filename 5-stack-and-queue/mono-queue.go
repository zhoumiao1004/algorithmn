package main

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
// 思路1: 封装单调队列的方式解题
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

// 代码模拟单调队列过程，不易理解
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
