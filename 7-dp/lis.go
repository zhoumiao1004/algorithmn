package main

import (
	"fmt"
	"sort"
)

// 300.最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/description/
// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
// 输入：nums = [10,9,2,5,3,7,101,18] 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
// 时间复杂度 O(N*N)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	dp := make([]int, n) // dp[i]含义：以nums[i]这个数结尾的LIS的长度
	for i := 0; i < n; i++ {
		dp[i] = 1 // base case
	}
	res := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			res = max(res, dp[i]) // 顺便计算以各个 nums[i] 结尾的LIS长度最大值
		}
	}
	return res
}

// 354. 俄罗斯套娃信封问题
// https://leetcode.cn/problems/russian-doll-envelopes/description/
func maxEnvelopes(envelopes [][]int) int {
	var lengthOfLIS func(nums []int) int
	lengthOfLIS = func(nums []int) int {
		piles := 0
		n := len(nums)
		top := make([]int, n)
		for _, poker := range nums {
			// 要处理的扑克牌
			left, right := 0, piles
			// 二分查找插入位置
			for left < right {
				mid := (left + right) / 2
				if top[mid] >= poker {
					right = mid
				} else {
					left = mid + 1
				}
			}
			if left == piles {
				piles++
			}
			// 把这张牌放到牌堆顶
			top[left] = poker
		}
		// 牌堆数就是 LIS 长度
		return piles
	}

	n := len(envelopes)
	// 按宽度升序排列，如果宽度一样，则按高度降序排列
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	// 对高度数组寻找 LIS
	height := make([]int, n)
	for i := 0; i < n; i++ {
		height[i] = envelopes[i][1]
	}

	return lengthOfLIS(height)
}

// 674. 最长连续递增序列（子数组）
// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/description/
// 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
// 输入：nums = [1,3,5,4,7] 输出：3
// 解释：最长连续递增序列是 [1,3,5], 长度为3。
// 尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n) // dp[i]含义：以下标i结尾的字符串的最长连续递增序列的长度
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			res = max(res, dp[i])
		}
	}
	return res
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))         // 3
}
