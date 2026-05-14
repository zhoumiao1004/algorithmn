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
	var res [][]int
	var path []int
	var backtrack func(start int)

	backtrack = func(start int) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}

// 思路2:球的视角选盒(桶)
func subset(nums []int) [][]int {
	var res [][]int
	var path []int
	var backtrack func(nums []int, i int)

	backtrack = func(nums []int, i int) {
		if i == len(nums) {
			res = append(res, append([]int{}, path...))
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
	return res
}

// 90.子集II
// https://leetcode.cn/problems/subsets-ii/
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var path []int
	sort.Ints(nums)
	var backtrack func(start int)

	backtrack = func(start int) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i-1] == nums[i] {
				continue // 树层去重
			}
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
