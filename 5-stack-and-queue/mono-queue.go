package main

import "math"

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
// 思路1: 单调队列
type MonoQueue struct {
	deque []int
}

func (m *MonoQueue) Front() int  { return m.deque[0] }
func (m *MonoQueue) Back() int   { return m.deque[len(m.deque)-1] }
func (m *MonoQueue) Empty() bool { return len(m.deque) == 0 }
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

// 思路1: 滑动窗口
func maxSlidingWindow(nums []int, k int) []int {
	window := NewMonotonicQueue()
	var res []int
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			window.Push(nums[i])
		} else {
			window.Push(nums[i])
			res = append(res, window.Max())
			window.Pop()
		}
	}
	return res
}

func maxSlidingWindow2(nums []int, k int) []int {
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

// 代码模拟单调队列过程，不易理解
func maxSlidingWindow3(nums []int, k int) []int {
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

// 1438. 绝对差不超过限制的最长连续子数组
// https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
// 给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit。
// 输入：nums = [8,2,4,7], limit = 4
// 输出：2
// 解释：所有子数组如下：
// [8] 最大绝对差 |8-8| = 0 <= 4.
// [8,2] 最大绝对差 |8-2| = 6 > 4.
// [8,2,4] 最大绝对差 |8-2| = 6 > 4.
// [8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
// [2] 最大绝对差 |2-2| = 0 <= 4.
// [2,4] 最大绝对差 |2-4| = 2 <= 4.
// [2,4,7] 最大绝对差 |2-7| = 5 > 4.
// [4] 最大绝对差 |4-4| = 0 <= 4.
// [4,7] 最大绝对差 |4-7| = 3 <= 4.
// [7] 最大绝对差 |7-7| = 0 <= 4.
// 因此，满足题意的最长子数组的长度为 2 。
type MonotonicQueue struct {
	q    []int
	minq []int
	maxq []int
}

func NewMonotonicQueue() MonotonicQueue {
	return MonotonicQueue{
		q:    []int{},
		minq: []int{},
		maxq: []int{},
	}
}

func (mq *MonotonicQueue) Min() int      { return mq.minq[0] }
func (mq *MonotonicQueue) Max() int      { return mq.maxq[0] }
func (mq *MonotonicQueue) Size() int     { return len(mq.q) }
func (mq *MonotonicQueue) IsEmpty() bool { return len(mq.q) == 0 }
func (mq *MonotonicQueue) Push(elem int) {
	mq.q = append(mq.q, elem)

	for len(mq.maxq) > 0 && mq.maxq[len(mq.maxq)-1] < elem {
		mq.maxq = mq.maxq[:len(mq.maxq)-1]
	}
	mq.maxq = append(mq.maxq, elem)

	for len(mq.minq) > 0 && mq.minq[len(mq.minq)-1] > elem {
		mq.minq = mq.minq[:len(mq.minq)-1]
	}
	mq.minq = append(mq.minq, elem)
}

func (mq *MonotonicQueue) Pop() int {
	val := mq.q[0]
	mq.q = mq.q[1:] // 注: 别忘了出队
	if val == mq.minq[0] {
		mq.minq = mq.minq[1:]
	}
	if val == mq.maxq[0] {
		mq.maxq = mq.maxq[1:]
	}
	return val
}

func longestSubarray(nums []int, limit int) int {
	left, right := 0, 0
	windowSize, res := 0, 0
	window := NewMonotonicQueue()
	for right < len(nums) {
		window.Push(nums[right])
		right++
		windowSize++
		for window.Max()-window.Min() > limit {
			window.Pop()
			left++
			windowSize--
		}
		res = max(res, windowSize)
	}
	return res
}

// 862. 和至少为 K 的最短子数组
// https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/description/
// 给你一个整数数组 nums 和一个整数 k ，找出 nums 中和至少为 k 的 最短非空子数组 ，并返回该子数组的长度。如果不存在这样的 子数组 ，返回 -1 。
// 子数组 是数组中 连续 的一部分。
// 输入：nums = [1], k = 1
// 输出：1
func shortestSubarray(nums []int, k int) int {
	result := math.MaxInt
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	window := NewMonotonicQueue()
	left, right := 0, 0
	for right <= n {
		window.Push(preSum[right])
		right++
		for right <= n && !window.IsEmpty() && preSum[right]-window.Min() >= k {
			result = min(result, right-left)
			window.Pop()
			left++
		}
	}
	if result == math.MaxInt {
		return -1
	}
	return result
}

// 918. 环形子数组的最大和
// https://leetcode.cn/problems/maximum-sum-circular-subarray/
// 给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
// 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个元素是 nums[(i - 1 + n) % n] 。
// 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。
// 输入：nums = [1,-2,3,-2]
// 输出：3
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	preSum := make([]int, 2*n+1) // 计算环状 nums 的前缀和
	preSum[0] = 0
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[(i-1)%n]
	}

	maxSum := math.MinInt32
	window := NewMonotonicQueue() // 维护一个滑动窗口，以便根据窗口中的最小值计算最大子数组和
	window.Push(0)
	for i := 1; i < len(preSum); i++ {
		maxSum = max(maxSum, preSum[i]-window.Min())
		if window.Size() == n {
			window.Pop()
		}
		window.Push(preSum[i])
	}

	return maxSum
}
