package main

import (
	"fmt"
	"math"
)

// 53. 最大子数组和
// https://leetcode.cn/problems/maximum-subarray/description/
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
// 思路1: dp
func maxSubArray(nums []int) int {
	n := len(nums)
	res := nums[0]
	dp := make([]int, n) // dp[i]定义：以nums[i]结尾的最大子数组和
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}

// 思路2: 贪心，和为负数就放弃
func maxSubArrayGreedy(nums []int) int {
	res := math.MinInt
	s := 0
	for i := 0; i < len(nums); i++ {
		s = max(s+nums[i], nums[i])
		res = max(res, s)
	}
	return res
}

// 思路3: 滑动窗口
func maxSubArraySlideWindow(nums []int) int {
	left, right := 0, 0
	windowSum := 0
	maxSum := math.MinInt32
	for right < len(nums) {
		windowSum += nums[right]
		right++
		// 更新答案
		maxSum = max(maxSum, windowSum)
		// 判断窗口是否要收缩
		for windowSum < 0 {
			windowSum -= nums[left]
			left++
		}
	}
	return maxSum
}

// 思路4: 前缀和思路：以nums[i]结尾的最大子数组和 = preSum[i+1] - min(preSum[0...i])
func maxSubArrayPreSum(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1] // 计算前缀和
	}

	minVal := math.MaxInt
	res := math.MinInt

	for i := 0; i < n; i++ {
		res = max(res, preSum[i]-minVal)
		minVal = min(minVal, preSum[i])
	}
	return res
}

func main() {
	fmt.Println(maxSubArray([]int{-1, 0, -1}))
}
