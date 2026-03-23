package main

import (
	"encoding/base32"
	"sort"
)

// 46.全排列
// https://leetcode.cn/problems/permutations/description/
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 思路1: 盒视角
func permute(nums []int) [][]int {
	var results [][]int
	var path []int
	used := make([]bool, len(nums))
	var dfs func(nums []int)
	dfs = func(nums []int) {
		if len(path) == len(nums) {
			results = append(results, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(nums)
	return results
}

// 使用swap更高效
func permute2(nums []int) [][]int {
	var results [][]int
	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		if start == len(nums) {
			results = append(results, append([]int{}, nums...))
			return
		}
		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			backtrack(nums, start+1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
	backtrack(nums, 0)
	return results
}

// 思路2: 球视角，元素选索引
func permute3(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	count := 0
	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		if count == len(nums) {
			result = append(result, append([]int{}, nums...))
			return
		}
		originalIndex := -1
		swapIndex := -1
		for i := 0; i<len(nums); i++ {
			if used[i] {
				continue
			}
			if originalIndex == -1 {
				originalIndex = i
			}
			swapIndex = i
			// 做选择，元素 nums[originalIndex] 选择 swapIndex 位置
			nums[originalIndex], nums[swapIndex] =  nums[swapIndex], nums[originalIndex]
			used[swapIndex] = true
			count++
			backtrack(nums)
			// 撤销选择
			count--
			used[swapIndex] = false
			nums[originalIndex], nums[swapIndex] =  nums[swapIndex], nums[originalIndex]
		}
	}
	backtrack(nums)
	return result
}

// 47. 全排列 II
// https://leetcode.cn/problems/permutations-ii/description/
// LCR 084. 全排列 II https://leetcode.cn/problems/7p8L0Z/description/
// 输入：nums = [1,1,2]
// 输出：[[1,1,2],[1,2,1],[2,1,1]]
func permuteUnique(nums []int) [][]int {
	var results [][]int
	var path []int
	sort.Ints(nums)
	used := make([]bool, len(nums))
	var dfs func(nums []int)
	dfs = func(nums []int) {
		if len(path) == len(nums) {
			results = append(results, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue // 树层去重复
			}
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(nums)
	return results
}
