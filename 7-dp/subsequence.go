package main

import (
	"fmt"
	"math"
)

// 392. 判断子序列
// https://leetcode.cn/problems/is-subsequence/description/
// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
// 输入：s = "abc", t = "ahbgdc" 输出：true
// [0, 0, 0, 0, 0, 0, 0]
// [0, 1, 1, 1, 1, 1, 1]
// [0, 0, 0, 2, 2, 2, 2]
// [0, 0, 0, 0, 0, 0, 3]
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// dp[i][j] = dp[i][j-1] // 求的是s是不是t的子序列，t模拟删除最后一个字符
				dp[i][j] = max(dp[i][j-1], dp[i-1][j]) // lcs
			}
		}
	}
	return dp[m][n] == len(s)
}

// 思路2: 双指针
func isSubsequence2(s string, t string) bool {
	i := 0
	for j := 0; j < len(t); j++ {
		if i < len(s) && t[j] == s[i] {
			i++
		}
	}
	return i == len(s)
}

// 647. 回文子串
// https://leetcode.cn/problems/palindromic-substrings/description/
// 输入：s = "abc" 输出：3
// 解释：三个回文子串: "a", "b", "c"
// 输入：s = "aaa" 输出：6
// 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	res := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 || dp[i+1][j-1] {
					dp[i][j] = true
					res++
				}
			}
		}
	}
	return res
}

func countSubstrings2(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	res := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		res++
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				if j == i+1 || dp[i+1][j-1] {
					dp[i][j] = true
					res++
				}
			}
		}
	}
	return res
}

// 516.最长回文子序列
// https://leetcode.cn/problems/longest-palindromic-subsequence/description/
// 输入：s = "bbbab" 输出：4
// 解释：一个可能的最长回文子序列为 "bbbb" 。
// 输入：s = "cbbd" 输出：2
// 解释：一个可能的最长回文子序列为 "bb" 。
// dp[i][j]：字符串s在[i, j]范围内最长的回文子序列的长度为dp[i][j]。
// 递推公式：如果s[i]与s[j]不相同，说明s[i]和s[j]的同时加入 并不能增加[i,j]区间回文子序列的长度，那么分别加入s[i]、s[j]看看哪一个可以组成最长的回文子序列。
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
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
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}

func minInsertions2(s string) int {
	// 计算把 s 变成回文串的最少插入次数
	return len(s) - longestPalindromeSubseq(s)
}

// 132. 分割回文串 II
// https://leetcode.cn/problems/palindrome-partitioning-ii/description/
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串。
// 返回符合要求的 最少分割次数 。
// 输入：s = "aab"
// 输出：1
// 解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
func minCut(s string) int {
	// 预处理：先统计左闭右闭子串s[i:j]是不是回文串
	isValid := make([][]bool, len(s))
	for i := 0; i < len(isValid); i++ {
		isValid[i] = make([]bool, len(s))
		isValid[i][i] = true
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 || isValid[i+1][j-1] {
					isValid[i][j] = true
				}
			}
		}
	}
	// 1.dp[i]含义：切割字符串s[0:i]为多个回文串，最少分割次数
	// 2.递推公式：dp[i] = min(dp[i], dp[j] + 1)
	// 3.初始化：求最少，所以初始化为MaxInt
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < len(s); i++ {
		if isValid[0][i] {
			dp[i] = 0 // 0到i的子串已经是回文串了，不需要切割
			continue
		}
		// 0到i的子串不是回文串，需要切割，使用j来0到i之间尝试
		for j := 0; j < i; j++ {
			if isValid[j+1][i] { // 从j+1到i-1是回文串，可以在j后面切一刀分割
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(s)-1]
}

// 673. 最长递增子序列的个数
// https://leetcode.cn/problems/number-of-longest-increasing-subsequence/description/
// 给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。
// 注意 这个数列必须是 严格 递增的。
// 输入: [1,3,5,4,7]
// 输出: 2
// 解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	// dp[i]含义：以i结尾的nums数组最长递增子序列长度
	// count[i]：以i结尾的nums数组最长递增子序列的个数
	dp := make([]int, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
	}
	maxCount := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
				dp[i] = max(dp[i], dp[j]+1)
			}
			maxCount = max(maxCount, dp[i])
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if dp[i] == maxCount {
			res += count[i]
		}
	}
	return res
}

func findNumberOfLIS2(nums []int) int {
	// dp[i]含义：以i结尾的nums数组最长递增子序列长度
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}
	maxCount := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i][0] = max(dp[i][0], dp[j][0]+1)
			}
			if dp[i][0] == maxCount {
				dp[i][1]++
			} else if dp[i][0] > maxCount {
				dp[i][1] = 1
				maxCount = dp[i][0]
			}
		}
	}
	// fmt.Println(dp)
	// fmt.Println(maxCount)
	for i := 0; i < n; i++ {
		if dp[i][0] == maxCount {
			return dp[i][1]
		}
	}
	return 1
}

func main() {
	fmt.Println(countSubstrings("abc"))                 // 3
	fmt.Println(findNumberOfLIS2([]int{1, 3, 5, 4, 7})) // 2
	fmt.Println(findNumberOfLIS2([]int{2, 2, 2, 2, 2})) // 5
	//输入：s = "abc", t = "ahbgdc"
	//输出：true
	//[0 0 0 0 0 0 0]
	//[0 1 1 1 1 1 1]
	//[0 0 0 2 2 2 2]
	//[0 0 0 0 0 0 3]
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums := []int{-1}
	fmt.Println(maxSubArray(nums))                        // 6 连续子数组 [4,-1,2,1] 的和最大，为 6 。
	fmt.Println(longestCommonSubsequence("abced", "ace")) // 3
}
