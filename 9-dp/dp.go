package main

import "math"

// 931.下降路径最小和
// https://leetcode.com/problems/minimum-falling-path-sum/
// 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
// 输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
// 输出：13
// 暴力思路：定义dp函数
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	result := math.MaxInt
	memos := make([][]int, n)
	for i := 0; i < n; i++ {
		memos[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			memos[i][j] = math.MaxInt
		}
	}
	// dp函数含义：从0行下落，落到matrix[i][j]的最小路径和
	var dp func(matrix [][]int, i, j int) int
	dp = func(matrix [][]int, i int, j int) int {
		n := len(matrix)
		if i < 0 || i >= n || j < 0 || j >= n {
			return math.MaxInt
		}
		if i == 0 {
			return matrix[i][j]
		}
		if memos[i][j] != math.MaxInt {
			return memos[i][j]
		}
		// 可能由上一层的3个位置得到
		memos[i][j] = matrix[i][j] + min(
			dp(matrix, i-1, j-1),
			dp(matrix, i-1, j),
			dp(matrix, i-1, j+1),
		)
		return memos[i][j]
	}

	// 终点可能出现在最后一行的任意一列
	for i := 0; i < n; i++ {
		result = min(result, dp(matrix, n-1, i))
	}
	return result
}

// 自底向上迭代：dp数组
func minFallingPathSum2(matrix [][]int) int {
	result := math.MaxInt
	n := len(matrix)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt
		}
	}

	for j := 0; j < n; j++ {
		dp[0][j] = matrix[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			minVal := dp[i-1][j]
			if j > 0 {
				minVal = min(minVal, dp[i-1][j-1])
			}
			if j < n-1 {
				minVal = min(minVal, dp[i-1][j+1])
			}
			dp[i][j] = matrix[i][j] + minVal
		}
	}
	for j := 0; j < n; j++ {
		result = min(result, dp[n-1][j])
	}
	return result
}
