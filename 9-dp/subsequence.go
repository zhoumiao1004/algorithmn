package main

import (
	"fmt"
	"math"
)

// 300.最长递增子序列
// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
// 输入：nums = [10,9,2,5,3,7,101,18] 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
// 时间复杂度 O(N*N)
func lengthOfLIS(nums []int) int {
	// dp[i]含义：以nums[i]这个数结尾的最长递增子序列的长度
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 // base case
	}
	result := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				result = max(result, dp[i]) // 顺便求所有nums[i]结尾的LIS长度最大值
			}
		}
	}
	return result
}

// 674. 最长连续递增序列（子数组）
// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/description/
// 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
// 输入：nums = [1,3,5,4,7] 输出：3
// 解释：最长连续递增序列是 [1,3,5], 长度为3。
// 尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
func findLengthOfLCIS(nums []int) int {
	// dp[i]含义：以下标i结尾的字符串的最长连续递增序列的长度
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	result := 0
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			result = max(result, dp[i])
		}
	}
	return result
}

// 718. 最长重复子数组
// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/description/
// 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
// 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7] 输出：3
// 解释：长度最长的公共子数组是 [3,2,1] 。
func findLength(nums1 []int, nums2 []int) int {
	// dp[i][j]含义：nums1下标以i-1结尾，nums2以j-1结尾的数组的最长重复字数组长度
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				result = max(result, dp[i][j])
			}
		}
	}
	return result
}

// 1143.最长公共子序列
// https://leetcode.cn/problems/longest-common-subsequence/description/
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度
/* text1 = "abcde", text2 = "ace"
		a	c	e
	0	0	0	0
a	0	1	1	1
b	0	1	1	1
c	0	1	2	2
d	0	1	2	2
e	0	1	2	3
*/
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

// 53. 最大子数组和
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
func maxSubArray(nums []int) int {
	// dp[i]含义：以nums[i]结尾的最大(连续)子数组的和
	// 递推公式：dp[i] = max(dp[i-1] + nums[i], nums[i])
	n := len(nums)
	result := nums[0]
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(result, dp[i])
	}
	// fmt.Println(dp)
	return result
}

// 贪心：和为负数就放弃
func maxSubArrayGreedy(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	result := math.MinInt
	s := 0
	for i := 0; i < n; i++ {
		s = max(s+nums[i], nums[i])
		result = max(result, s)
	}
	return result
}

// 152.乘积最大子数组
// 输入: nums = [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
func maxProduct(nums []int) int {
	// dp[i]含义：以i结尾的nums子数组的最大乘积
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][0] = nums[0] // 最小乘积
	dp[0][1] = nums[0] // 最大乘积
	result := nums[0]
	for i := 1; i < n; i++ {
		a, b := dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]
		dp[i][0] = min(min(a, b), nums[i])
		dp[i][1] = max(max(a, b), nums[i])
		result = max(result, dp[i][1])
	}
	return result
}

// 方法2:贪心
func maxProduct2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	result := nums[0]
	preMin, preMax := nums[0], nums[0]
	for i := 1; i < n; i++ {
		a := preMin * nums[i]
		b := preMax * nums[i]
		preMin = min(nums[i], min(a, b))
		preMax = max(nums[i], max(a, b))
		result = max(result, preMax)
	}
	return result
}

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
	// dp[i][j]含义：[0,i-1]的s和[0,j-1]的t，相同子序列长度
	// 递推公式： if s[i-1] == s[j-1] : dp[i][j] = dp[i-1][j-1] + 1 else：dp[i][j] = dp[i][j-1]
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1] // 求的是s是不是t的子序列，t模拟删除最后一个字符
			}
		}
	}
	// fmt.Println(dp)
	return dp[m][n] == len(s)
}

// 647. 回文子串
// https://leetcode.cn/problems/palindromic-substrings/description/
// 输入：s = "abc" 输出：3
// 解释：三个回文子串: "a", "b", "c"
// [true false false]
// [false true false]
// [false false true]
// 输入：s = "aaa" 输出：6
// 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
func countSubstrings(s string) int {
	// dp[i][j]含义：[i,j]范围内的回文子串个数
	// 递推公式
	// 1.相同 s[i] == s[j]
	//  a.j-i <=1 dp[i][j] = true 回文个数+1
	//  b.        if dp[i+1][j-1] == true => dp[i][j] = true 回文个数+1
	// 2.不同 false
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	result := 0
	for i := n - 1; i >= 0; i-- { // 从下往上
		for j := i; j < n; j++ { // 从左往右
			if s[i] == s[j] {
				if j-i <= 1 || dp[i+1][j-1] {
					dp[i][j] = true
					result++
				}
			}
		}
	}
	// fmt.Println(dp)
	return result
}

func countSubstrings2(s string) int {
	// 1.dp[i][j]含义：左闭右闭区间s[i:j]是不是回文串
	// 2.递推公式：
	// if s[i] == s[j]:
	//   if j-i<=1 || dp[i+1][j-1]: dp[i][j] = true
	// 3.初始化：dp[i][i] = true
	// 4.遍历顺序：i从下到上，j从左到右
	n := len(s)
	dp := make([][]bool, n)
	result := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		result++
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				if j == i+1 || dp[i+1][j-1] {
					dp[i][j] = true
					result++
				}
			}
		}
	}
	return result
}

// 5.最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/description/
// 给你一个字符串 s，找到 s 中最长的 回文 子串。
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
func longestPalindrome(s string) string {
	// dp[i][j]含义：s[i][j]是不是回文串
	// 递推公式：
	// if s[i] == s[j]:
	//   if j - i <= 1: dp[i][j] = true
	//   else dp[i][j] = dp[i+1][j-1]
	//
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		// dp[i][i] = true
	}
	maxLen := 1
	left := 0
	right := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] && j-i+1 > maxLen {
					maxLen = j - i + 1
					left = i
					right = j
				}
			}
		}
	}
	return s[left : right+1]
}

func longestPalindrome2(s string) string {
	// 1.dp[i][j]含义：左闭右闭区间[i,j]范围内的字符串s是不是回文串
	// 2.递推公式：
	// if s[i] == s[j]:
	//   if j-i==1: dp[i][j] = true
	//   else dp[ij][j] = dp[i+1][j-1]
	// 3.初始化：
	// 4.遍历顺序：i从下往上，j从左往右
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	maxLen := 0
	left, right := 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 || dp[i+1][j-1] {
					dp[i][j] = true
					if j-i+1 > maxLen {
						left, right = i, j
						maxLen = j - i + 1
					}
				}
			}
		}
	}
	return s[left : right+1]
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
	// dp[i][j]含义：下标i到j范围内的字符串内最长回文子序列长度
	// 递推公式：if s[i] == s[j]: dp[i][j] = dp[i-1][j-1] + 2
	//         else: dp[i][j] = max(dp[i-1][j], dp[i][j-1])
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
	result := 0
	for i := 0; i < n; i++ {
		if dp[i] == maxCount {
			result += count[i]
		}
	}
	return result
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
	fmt.Println(dp)
	fmt.Println(maxCount)
	for i := 0; i < n; i++ {
		if dp[i][0] == maxCount {
			return dp[i][1]
		}
	}
	return 1
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))         // 3
	fmt.Println(countSubstrings("abc"))                         // 3
	fmt.Println(findNumberOfLIS2([]int{1, 3, 5, 4, 7}))         // 2
	fmt.Println(findNumberOfLIS2([]int{2, 2, 2, 2, 2}))         // 5
	fmt.Println(longestPalindrome("babad"))                     // bab
	fmt.Println(longestPalindrome2("babad"))                    // bab
}
