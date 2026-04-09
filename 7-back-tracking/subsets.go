package main

import (
	"fmt"
	"sort"
)

// 78. 子集
// https://leetcode.cn/problems/subsets/description/
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 思路1:盒(桶)的视角选球
func subsets(nums []int) [][]int {
	var results [][]int
	var path []int
	var backtrack func(nums []int, start int)
	
	backtrack = func(nums []int, start int) {
		results = append(results, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	
	backtrack(nums, 0)
	return results
}

// 思路2:球的视角选盒(桶)
func subset(nums []int) [][]int {
	var results [][]int
	var path []int
	var backtrack func(nums []int, i int)
	
	backtrack = func(nums []int, i int) {
		if i == len(nums) {
			results = append(results, append([]int{}, path...))
			return
		}
		// 第一种选择：球在盒中
		path = append(path, nums[i])
		backtrack(nums, i+1)
		path = path[:len(path)-1] // 撤销选择
		// 第二种选择：球不在盒中
		backtrack(nums, i+1)
	}

	backtrack(nums, 0)
	return results
}

// 90.子集II
// https://leetcode.cn/problems/subsets-ii/
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
func subsetsWithDup(nums []int) [][]int {
	var results [][]int
	var path []int
	sort.Ints(nums)
	used := make([]bool, len(nums))
	var backtrack func(nums []int, startIndex int)
	
	backtrack = func(nums []int, startIndes int) {
		results = append(results, append([]int{}, path...))
		for i := startIndes; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue // 树层去重
			}
			path = append(path, nums[i])
			used[i] = true
			backtrack(nums, i+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack(nums, 0)
	return results
}

// 491. 非递减子序列
// https://leetcode.cn/problems/non-decreasing-subsequences/description/
// 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
// 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
// 输入：nums = [4,6,7,7]
// 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
// 注：不能排序
func findSubsequences(nums []int) [][]int {
	var results [][]int
	var path []int
	var backtrack func(nums []int, startIndex int)

	backtrack = func(nums []int, startIndex int) {
		if len(path) > 1 {
			results = append(results, append([]int{}, path...)) // 注意这里没有return，因为找到一个子序列，还能继续往树的下一层找更长的子序列
		}
		uset := make(map[int]bool) // 同层去重
		for i := startIndex; i < len(nums); i++ {
			if uset[nums[i]] {
				continue
			}
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			path = append(path, nums[i])
			uset[nums[i]] = true
			backtrack(nums, i+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(nums, 0)
	return results
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
