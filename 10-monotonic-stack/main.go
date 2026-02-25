package main

import (
	"fmt"
)

/* 单调栈：通过一个栈保存已经遍历过的元素 */
// 739. 每日温度
// https://leetcode.cn/problems/daily-temperatures/description/
// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 输入: temperatures = [73,74,75,71,69,72,76,73]
// 输出: [1,1,4,2,1,1,0,0]
func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var st []int
	for i := 0; i < len(temperatures); i++ {
		// 比较temperature[i]和栈顶元素大小，如果大于栈顶元素，说明是右边第一个大于栈顶元素的地方
		for len(st) > 0 && temperatures[i] > temperatures[st[len(st)-1]] {
			idx := st[len(st)-1]
			result[idx] = i - idx
			st = st[:len(st)-1] // 弹出栈顶元素
		}
		st = append(st, i)
	}

	return result
}

// 496. 下一个更大元素 I
// https://leetcode.com/problems/next-greater-element-i/description/
// 给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
// 请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
// nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
// 输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
// 输出：[-1,3,-1]
// 解释：nums1 中每个值的下一个更大元素如下所述：
// 4: 不存在下一个更大元素，所以答案是 -1 。
// 1: 下一个更大元素是 3 。
// 2: 不存在下一个更大元素，所以答案是 -1 。
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		result[i] = -1
	}
	// nums1中val和idx的映射
	idxMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		idxMap[nums1[i]] = i
	}
	var st []int
	for i := 0; i < len(nums2); i++ {
		for len(st) > 0 && nums2[i] > st[len(st)-1] {
			idx, ok := idxMap[st[len(st)-1]]
			if ok {
				result[idx] = nums2[i]
			}
			st = st[:len(st)-1]
		}
		st = append(st, nums2[i])
	}

	return result
}

// 503. 下一个更大元素 II
// 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。
// 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1 。
// 输入: nums = [1,2,1] 输出: [2,-1,2]
// 解释: 第一个 1 的下一个更大的数是 2；
// 数字 2 找不到下一个更大的数；
// 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}
	var st []int
	for i := 0; i < 2*n; i++ {
		for len(st) > 0 && nums[i%n] > nums[st[len(st)-1]] {
			result[st[len(st)-1]] = nums[i%n]
			st = st[:len(st)-1]
		}
		st = append(st, i%n)
	}

	return result
}

// 4.接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1] 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）
func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	result := 0
	var st []int
	for i := 0; i < n; i++ {
		for len(st) > 0 && height[i] > height[st[len(st)-1]] {
			mid := st[len(st)-1]
			st = st[:len(st)-1] // 出栈
			if len(st) > 0 {
				left := st[len(st)-1] // 栈顶元素是mid左边第一个大于的位置
				right := i
				w := right - left - 1
				h := min(height[left], height[right]) - height[mid]
				result += w * h
			}
		}
		st = append(st, i) // 入栈
	}
	return result
}

// 双指针解法
func trap2pointer(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	left, right := 0, n-1
	maxLeft, maxRight := nums[0], nums[n-1]
	ans := 0
	for left < right {
		hLeft, hRight := nums[left], nums[right]
		if hLeft < hRight {
			if hLeft > maxLeft {
				maxLeft = hLeft
			} else {
				ans += hLeft
			}
			left++
		} else {
			if hRight > maxRight {
				maxLeft = hLeft
			} else {
				ans += hRight
			}
			right--
		}
	}
	return ans
}

// 84. 柱状图中最大的矩形
// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
// 输入：heights = [2,1,5,6,2,3] 输出：10
// 解释：最大的矩形为图中红色区域，面积为 10
func largestRectangleArea(heights []int) int {
	result := 0
	// 首尾补零
	nums := make([]int, len(heights)+2)
	for i := 0; i < len(heights); i++ {
		nums[i+1] = heights[i]
	}
	var st []int
	for i := 0; i < len(nums); i++ {
		for len(st) > 0 && nums[i] < nums[st[len(st)-1]] {
			mid := st[len(st)-1] // 取栈顶：mid
			st = st[:len(st)-1]  // 出栈
			if len(st) > 0 {
				left := st[len(st)-1] // 取栈顶：左边第一个小于mid的元素
				right := i
				h := nums[mid]
				w := right - left - 1
				result = max(result, h*w)
			}
		}
		st = append(st, i)
	}

	return result
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4})) // 3, -1]
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3})) // 10
	fmt.Println(largestRectangleArea([]int{1}))                // 1

	// people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	// fmt.Println(reconstructQueue(people))
}
