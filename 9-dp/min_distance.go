package main

import "fmt"

/*
72. 编辑距离
https://leetcode.cn/problems/edit-distance/
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：插入、删除、替换
输入：word1 = "horse", word2 = "ros" 输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// dp[i][j]含义：word1[0...i-1] 和 word2[0...j-1] 的最少操作数
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // 啥都不做
			} else {
				dp[i][j] = min(
					dp[i-1][j-1]+1,
					dp[i-1][j]+1,
					dp[i][j-1]+1,
				)
			}
		}
	}
	return dp[m][n]
}

// 暴力解法
func minDistanceForce(word1 string, word2 string) int {
	// dp(i, j) 返回 s1[0...i] 和 s2【0...j] 的最小编辑距离
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}

		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		} else {
			return min(
				dp(i, j-1)+1,
				dp(i-1, j)+1,
				dp(i-1, j-1)+1,
			)
		}
	}
	return dp(len(word1)-1, len(word2)-1)
}

// 带备忘录递归解法
func minDistanceMemo(word1 string, word2 string) int {
	// dp(i, j) 返回 s1[0...i] 和 s2【0...j] 的最小编辑距离
	m, n := len(word1), len(word2)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		// 查备忘录
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		if word1[i] == word2[j] {
			memo[i][j] = dp(i-1, j-1) // 啥都不做
		} else {
			memo[i][j] = min(
				dp(i, j-1)+1,
				dp(i-1, j)+1,
				dp(i-1, j-1)+1,
			)
		}
		return memo[i][j]
	}
	return dp(len(word1)-1, len(word2)-1)
}

func main() {
	fmt.Println(numDistinct("babgbag", "bag"))
	// fmt.Println(minDistance("horse", "ros"))
}
