package main

import (
	"math"
	"strings"
)

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

// 115. 不同的子序列
// https://leetcode.cn/problems/distinct-subsequences/description/
// 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数
// 输入：s = "babgbag", t = "bag" 输出：5
// 2种视角：从s的视角；从t的视角
// 思路1: 从s的视角,如果s[0]能匹配t[0],又有两种情况
// 如果s[0] 匹配 t[0], 原问题转化为s[1...]的所有子序列中计算t[1...]出现的次数
// 也可以不让 s[0] 匹配 t[0], 原问题转化为s[1...]的所有子序列中计算t[0...]出现的次数
// 为了给 s[0] 之后的元素匹配的机会，比如 s = "aab", t = "ab"，就有两种匹配方式：a_b 和 _ab。
func numDistinctMemo(s string, t string) int {
	m, n := len(s), len(t)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	var dp func(s, t string, i, j int) int
	dp = func(s, t string, i, j int) int {
		if j == len(t) {
			return 1
		}
		if len(s)-i < len(t)-j {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if s[i] == t[j] {
			memo[i][j] = dp(s, t, i+1, j+1) + dp(s, t, i+1, j)
		} else {
			memo[i][j] = dp(s, t, i+1, j)
		}
		return memo[i][j]
	}
	return dp(s, t, 0, 0)
}

// 自底向上递归dp数组
// [1, 0, 0, 0]
// [1, 1, 0, 0]
// [1, 1, 1, 0]
// [1, 2, 1, 0]
// [1, 2, 1, 1]
// [1, 3, 1, 1]
// [1, 3, 4, 1]
// [1, 3, 4, 5]
func numDistinct(s string, t string) int {
	// dp[i][j]含义：[0,i-1]的s和[0,j-1]的t的个数
	// 1.s[i] == t[j]: dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
	// 2.s[i] != t[j]: dp[i][j] = dp[i-1][j]
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j] // 两边都删除的个数 + 删除s最后一个
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	// fmt.Println(dp)
	return dp[m][n]
}

// 139.单词拆分
// https://leetcode.cn/problems/word-break/
// 输入: s = "leetcode", wordDict = ["leet", "code"] 输出: true
// 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
// 注意：单词放入是有顺序的，所以是排列问题，不能求组合
// 1.遍历的思路，就是用回溯算法解决，回溯最经典的应用就是排列组合问题。时间复杂度2的N次方
func wordBreakBT(s string, wordDict []string) bool {
	// memo := make(map[string]bool)
	found := false
	var path []string
	var backtrack func(wordDict []string, start int)
	backtrack = func(wordDict []string, start int) {
		if found {
			return
		}
		if start == len(wordDict) {
			found = true
			return
		}

		for i := 0; i < len(wordDict); i++ {
			word := wordDict[i]
			if start+len(word) <= len(s) && s[start:start+len(word)] == word {
				// 做选择
				path = append(path, word)
				// 进入下一层回溯树
				backtrack(wordDict, i+len(word))
				// 撤销选择
				path = path[:len(path)-1]
			}

		}
	}
	backtrack(wordDict, 0)
	return found
}

// 2.分解的思路
func wordBreakMemo(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	memo := make([]int, len(s)) // s[i]能否被单词拼出, -1 代表未计算，0 代表无法凑出，1 代表可以凑出
	// 定义：返回s[start...] 子串是否能被单词拼出
	var dp func(s string, start int) bool
	dp = func(s string, start int) bool {
		if start == len(s) {
			return true
		}
		if memo[start] != 0 {
			return memo[start] == 1
		}
		// 遍历 s[start...] 的所有前缀，看看哪些前缀存在 wordDict 中
		for i := 1; start+i <= len(s); i++ {
			prefix := s[start : start+i]
			if wordSet[prefix] && dp(s, start+i) {
				memo[start] = 1
				return true
			}
		}
		// s[1...] 无法被拼出
		memo[start] = 0
		return false
	}
	return dp(s, 0)
}

func wordBreak(s string, wordDict []string) bool {
	// dp[j]含义：[0,j)范围的子串，能否由字典里的单词组成
	// 用集合中的物品，装大小为j的背包
	// if dp[i] = true && [i,j]区间内的字符串在字典中 : dp[j] = true
	// 遍历顺序：求排列，先遍历背包再遍历物品
	wordMap := make(map[string]bool)
	for _, w := range wordDict {
		wordMap[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for j := 1; j <= n; j++ { // 背包
		for i := 0; i <= j; i++ { // 物品
			if dp[i] && wordMap[s[i:j]] {
				dp[j] = true
			}
		}
	}
	// fmt.Println(dp)
	return dp[n]
}

// 140. 单词拆分 II
// https://leetcode.cn/problems/word-break-ii/
// 给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。
// 注意：词典中的同一个单词可能在分段中被重复使用多次。
// 输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
// 输出:["cats and dog","cat sand dog"]
// 1.遍历的思路（回溯算法）
func wordBreakIIBT(s string, wordDict []string) []string {
	var result []string
	var path []string

	var backtrack func(s string, start int)
	backtrack = func(s string, start int) {
		if start == len(s) {
			result = append(result, strings.Join(path, " "))
			return
		}
		for _, word := range wordDict {
			if start+len(word) <= len(s) && s[start:start+len(word)] == word {
				path = append(path, word)
				backtrack(s, start+len(word))
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(s, 0)
	return result
}

// 2.分解的思路（动态规划）
func wordBreakII(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	memo := make([][]string, len(s))

	// dp[start...]能由单词拼成的句子
	var dp func(s string, start int) []string
	dp = func(s string, start int) []string {
		var result []string
		if start == len(s) {
			return []string{""}
		}
		if len(memo[start]) > 0 {
			return memo[start]
		}
		// 遍历 s[start...] 的所有前缀
		for i := 1; start+i <= len(s); i++ {
			prefix := s[start : start+i]
			if wordSet[prefix] {
				for _, sentence := range dp(s, start+i) {
					if sentence == "" {
						result = append(result, prefix)
					} else {
						result = append(result, prefix+" "+sentence)
					}
				}
			}
		}
		return result
	}
	return dp(s, 0)
}
