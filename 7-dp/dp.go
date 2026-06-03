package main

import "math"

// 931.下降路径最小和
// https://leetcode.com/problems/minimum-falling-path-sum/
// 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
// 输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
// 输出：13
// 思路1: 自顶向下带备忘录的递归解法
func minFallingPathSum(matrix [][]int) int {
	var dp func(i, j int) int // dp函数含义：从0行下落，落到matrix[i][j]的最小路径和
	n := len(matrix)
	memos := make([][]int, n)
	for i := 0; i < n; i++ {
		memos[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memos[i][j] = 10001
		}
	}

	dp = func(i int, j int) int {
		n := len(matrix)
		if j < 0 || j >= n {
			return math.MaxInt
		}
		if i == 0 {
			return matrix[i][j]
		}
		if memos[i][j] != 10001 {
			return memos[i][j]
		}
		memos[i][j] = matrix[i][j] + min(
			dp(i-1, j-1),
			dp(i-1, j),
			dp(i-1, j+1),
		)
		return memos[i][j]
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		res = min(res, dp(n-1, i))
	}
	return res
}

// 思路2: 自底向上迭代解法
func minFallingPathSum2(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n) // dp[i][j]含义：下降路径最小和
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = matrix[0][j] // 初始化第一行
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

	res := math.MaxInt
	for j := 0; j < n; j++ {
		res = min(res, dp[n-1][j]) // 结果在最后一行
	}
	return res
}

// 115. 不同的子序列
// https://leetcode.cn/problems/distinct-subsequences/description/
// 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数
// 输入：s = "babgbag", t = "bag" 输出：5
// 2种视角：s的视角 or t的视角
// 思路1: 从s的视角,如果s[0]能匹配t[0],又有两种情况
// 如果s[0] 匹配 t[0], 原问题转化为s[1...]的所有子序列中计算t[1...]出现的次数
// 也可以不让 s[0] 匹配 t[0], 原问题转化为s[1...]的所有子序列中计算t[0...]出现的次数
// 为了给 s[0] 之后的元素匹配的机会，比如 s = "aab", t = "ab"，就有两种匹配方式：a_b 和 _ab。
// 视角1: 从t的视角穷举。时间复杂度：状态的个数O(M*N) * 函数本身O(M)
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1 // 未计算初始化为-1
		}
	}

	var dp func(i, j int) int // s[i...] 中出现 t[j...] 的次数

	dp = func(i, j int) int {
		if j == len(t) {
			return 1 // t已经全部匹配完
		}
		if len(s)-i < len(t)-j {
			return 0 // s[i...] 比 t[j...] 还短，必然没有匹配的子序列
		}
		if memo[i][j] != -1 {
			return memo[i][j] // 计算过
		}
		res := 0
		for k := i; k < len(s); k++ {
			if s[k] == t[j] {
				res += dp(k+1, j+1)
			}
		}
		memo[i][j] = res
		return res
	}

	return dp(0, 0)
}

// 视角2: 从s的视角穷举，时间复杂度：状态的个数O(M*N)
func numDistinct2(s string, t string) int {
	m, n := len(s), len(t)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1 // 未计算初始化为-1
		}
	}

	var dp func(i, j int) int // s[i...] 中出现 t[j...] 的次数

	dp = func(i, j int) int {
		if j == len(t) {
			return 1 // t已经全部匹配完
		}
		if len(s)-i < len(t)-j {
			return 0 // s[i...] 比 t[j...] 还短，必然没有匹配的子序列
		}
		if memo[i][j] != -1 {
			return memo[i][j] // 计算过
		}
		if s[i] == t[j] {
			memo[i][j] = dp(i+1, j+1) + dp(i+1, j) // 明明可以匹配，为什么不让他俩匹配呢？主要是为了给s[i]之后的元素匹配的机会
		} else {
			memo[i][j] = dp(i+1, j)
		}
		return memo[i][j]
	}

	return dp(0, 0)
}

// 思路2: 自底向上递归dp数组
func numDistinct3(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1) // dp[i][j]含义：[0,i-1]的s和[0,j-1]的t的个数
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1 // base case
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
	return dp[m][n]
}
