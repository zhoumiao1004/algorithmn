package main

import "fmt"

// 42.接雨水
// https://leetcode.cn/problems/trapping-rain-water/description/
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1] 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）
// 思路1: 单调栈
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
				left := st[len(st)-1] // 栈顶元素是mid左边第一个大于的位置,注意这里不需要出栈
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

// 思路2: 双指针解法
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

// 11. 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water/description/
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 思路: 双指针
func maxArea(height []int) int {
	result := 0
	left, right := 0, len(height)-1
	for left < right {
		w := right - left
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		result = max(result, w*h)
	}
	return result
}

// 84. 柱状图中最大的矩形
// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
// 输入：heights = [2,1,5,6,2,3] 输出：10
// 解释：最大的矩形为图中红色区域，面积为 10
func largestRectangleArea(heights []int) int {
	result := 0
	// 首尾补零
	tmp := make([]int, len(heights)+2)
	for i := 0; i < len(heights); i++ {
		tmp[i+1] = heights[i]
	}
	heights = tmp
	var st []int
	for i := 0; i < len(heights); i++ {
		for len(st) > 0 && heights[i] < heights[st[len(st)-1]] {
			mid := st[len(st)-1] // 取栈顶：mid
			st = st[:len(st)-1]  // 出栈
			if len(st) > 0 {
				left := st[len(st)-1] // 取栈顶：左边第一个小于mid的元素
				right := i
				h := heights[mid]
				w := right - left - 1
				result = max(result, h*w)
			}
		}
		st = append(st, i)
	}

	return result
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3})) // 10
	fmt.Println(largestRectangleArea([]int{1}))                // 1
}
