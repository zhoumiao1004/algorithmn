package main

import "fmt"

// 1143.最长公共子序列(Longest Common Subsequence，简称 LCS)
// https://leetcode.cn/problems/longest-common-subsequence/description/
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	// dp[i][j]含义：[0,i-1]的text1和[0,j-1]的text2的最长公共子序列长度
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
	// fmt.Println(dp)
	return dp[m][n]
}

// 自顶向下的递归，通过memo保存子问题结果
func longestCommonSubsequence2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	// 定义：dp函数返回 s1[i...] 和 s2[j...] 的公共子序列长度
	var dp func(s1, s2 string, i, j int) int
	dp = func(s1, s2 string, i, j int) int {
		if i == len(s1) || j == len(s2) {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if s1[i] == s2[j] {
			memo[i][j] = 1 + dp(s1, s2, i+1, j+1)
		} else {
			memo[i][j] = max(dp(s1, s2, i+1, j), dp(s1, s2, i, j+1))
		}
		return memo[i][j]
	}
	return dp(text1, text2, 0, 0)
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
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	// dp[i][j]含义：以下标i-1结尾的nums1和下标j-1结尾的nums2最长公共子序列长度
	dp := make([][]int, m+1)
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
// [0, 1, 2, 3]
// [1, 2, 3, 4]
// [2, 1, 2, 3]
// [3, 2, 2, 2]
func minDistance(word1 string, word2 string) int {
	// dp[i][j]含义：下标为i-1的word1和下标为j-1的word2需要删除的最小次数
	// 递推公式
	// 1.相同 dp[i][j] = dp[i-1][j-1]
	// 2.不同 dp[i][j] = min(dp[i-1][j] + 1, dp[i][j-1] + 1)
	m, n := len(word1), len(word2)
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
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

// 方法2:
// 1.求最长公共子序列长度
// 2.len(word1) + len(word2) - 2*dp[m][n]
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
	fmt.Println(dp)
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
func minimumDeleteSum(s1 string, s2 string) int {
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

// 1312. 让字符串成为回文串的最少插入次数
// https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/description/
// 给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。
// 请你返回让 s 成为回文串的 最少操作次数 。
// 「回文串」是正读和反读都相同的字符串。
// 输入：s = "zzazz"
// 输出：0
// 输入：s = "mbadm"
// 输出：2
// 解释：字符串可变为 "mbdadbm" 或者 "mdbabdm" 。
func minInsertions(s string) int {
	// dp[i][j]含义：字符串s[i..j]，最少需要插入dp[i][j]次才能成为回文串。因此要求的是dp[0, n-1]
	// 假设已经计算出了子问题 dp[i+1][j-1] 的值，如果s[i] == s[j], s[i..j]需要的次数也是dp[i+1][j-1]; 如果s[i] != s[j] 有几种情况：
	// i左边插入s[j]相同的字符 or j右边插入s[i]相同的字符
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i<n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n-1; i>=0; i-- {
		for j := i+1; j<n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}
