package main

import (
	"fmt"
	"sort"
)

// 78. 子集
// https://leetcode.cn/problems/subsets/description/
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 可以取0-3个
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
	var dfs func(nums []int, startIndex int)
	dfs = func(nums []int, startIndes int) {
		results = append(results, append([]int{}, path...))
		for i := startIndes; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue // 树层去重
			}
			path = append(path, nums[i])
			used[i] = true
			dfs(nums, i+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(nums, 0)
	return results
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
