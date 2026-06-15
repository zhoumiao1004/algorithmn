package main

// 1143.最长公共子序列(Longest Common Subsequence，简称 LCS)
// https://leetcode.cn/problems/longest-common-subsequence/description/
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
// 思路1: 自顶向下带memo的递归
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	var dp func(i, j int) int // 定义：dp函数返回 s1[i...] 和 s2[j...] 的公共子序列长度
	dp = func(i, j int) int {
		if i == m || j == n {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if text1[i] == text2[j] {
			memo[i][j] = 1 + dp(i+1, j+1)
		} else {
			memo[i][j] = max(dp(i+1, j), dp(i, j+1))
		}
		return memo[i][j]
	}
	return dp(0, 0)
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(i, j int) int

	dp = func(i, j int) int {
		if i == -1 || j == -1 {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if text1[i] == text2[j] {
			memo[i][j] = dp(i-1, j-1) + 1
		} else {
			memo[i][j] = max(dp(i-1, j), dp(i, j-1))
		}
		return memo[i][j]
	}
	return dp(m-1, n-1)
}

// 思路2: 自底向上迭代解法
func longestCommonSubsequence3(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// 1035.不相交的线
// 在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
// 现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足：
// nums1[i] == nums2[j]
// 且绘制的直线不与任何其他连线（非水平线）相交。
// 请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
// 以这种方法绘制线条，并返回可以绘制的最大连线数。
// 输入：nums1 = [1,4,2], nums2 = [1,2,4] 输出：2
// 解释：可以画出两条不交叉的线
// 但无法画出第三条不相交的直线，因为从 nums1[1]=4 到 nums2[2]=4 的直线将与从 nums1[2]=2 到 nums2[1]=2 的直线相交。
// 思路1: 就是求lcs
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1) // dp[i][j]含义：以下标i-1结尾的nums1和下标j-1结尾的nums2最长公共子序列长度
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// 583. 两个字符串的删除操作
// https://leetcode.cn/problems/delete-operation-for-two-strings/description/
// 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
// 每步 可以删除任意一个字符串中的一个字符。
// 输入: word1 = "sea", word2 = "eat" 输出: 2
// 解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1) // dp[i][j]含义：下标为i-1的word1和下标为j-1的word2需要删除的最小次数
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
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

// 思路2: lcs
func minDistance1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return len(word1) + len(word2) - 2*dp[m][n]
}

// 712. 两个字符串的最小ASCII删除和
// https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings/description/
// 给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。
// 输入: s1 = "sea", s2 = "eat"
// 输出: 231
// 解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
// 在 "eat" 中删除 "t" 并将 116 加入总和。
// 结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
// 思路1: 自顶向下递归解法
func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(i, j int) int

	dp = func(i, j int) int {
		if i == -1 && j == -1 {
			return 0
		} else if i == -1 {
			return dp(i, j-1) + int(s2[j])
		} else if j == -1 {
			return dp(i-1, j) + int(s1[i])
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if s1[i] == s2[j] {
			memo[i][j] = dp(i-1, j-1)
		} else {
			memo[i][j] = min(dp(i-1, j)+int(s1[i]), dp(i, j-1)+int(s2[j]))
		}
		return memo[i][j]
	}

	return dp(m-1, n-1)
}

// 思路2: 自底向上迭代解法
func minimumDeleteSum2(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}
	return dp[m][n]
}
