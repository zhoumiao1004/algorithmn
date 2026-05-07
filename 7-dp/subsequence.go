package main

import (
	"fmt"
	"math"
	"sort"
)

// 300.最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/description/
// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
// 输入：nums = [10,9,2,5,3,7,101,18] 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
// 时间复杂度 O(N*N)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	// dp[i]含义：以nums[i]这个数结尾的LIS的长度
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 // base case
	}
	res := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			res = max(res, dp[i])
		}
	}
	return res
}

// 354. 俄罗斯套娃信封问题
// https://leetcode.cn/problems/russian-doll-envelopes/description/
func maxEnvelopes(envelopes [][]int) int {

	var lengthOfLIS func(nums []int) int
	lengthOfLIS = func(nums []int) int {
		piles := 0
		n := len(nums)
		top := make([]int, n)
		for _, poker := range nums {
			// 要处理的扑克牌
			left, right := 0, piles
			// 二分查找插入位置
			for left < right {
				mid := (left + right) / 2
				if top[mid] >= poker {
					right = mid
				} else {
					left = mid + 1
				}
			}
			if left == piles {
				piles++
			}
			// 把这张牌放到牌堆顶
			top[left] = poker
		}
		// 牌堆数就是 LIS 长度
		return piles
	}

	n := len(envelopes)
	// 按宽度升序排列，如果宽度一样，则按高度降序排列
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	// 对高度数组寻找 LIS
	height := make([]int, n)
	for i := 0; i < n; i++ {
		height[i] = envelopes[i][1]
	}

	return lengthOfLIS(height)
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
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			res = max(res, dp[i])
		}
	}
	return res
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

// 思路3:前缀积prefix
// TODO

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

// 115. 不同的子序列
// https://leetcode.cn/problems/distinct-subsequences/description/
// 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数
// 输入：s = "babgbag", t = "bag" 输出：5
// 2种视角：从s的视角；从t的视角
// 思路1: 从s的视角,如果s[0]能匹配t[0],又有两种情况
// 如果s[0] 匹配 t[0], 原问题转化为s[1...]的所有子序列中计算t[1...]出现的次数
// 也可以不让 s[0] 匹配 t[0], 原问题转化为s[1...]的所有子序列中计算t[0...]出现的次数
// 为了给 s[0] 之后的元素匹配的机会，比如 s = "aab", t = "ab"，就有两种匹配方式：a_b 和 _ab。
// 思路1: 带memo的递归解法
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

// 思路2: 自底向上递归dp数组
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

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))         // 3
	fmt.Println(countSubstrings("abc"))                         // 3
	fmt.Println(findNumberOfLIS2([]int{1, 3, 5, 4, 7}))         // 2
	fmt.Println(findNumberOfLIS2([]int{2, 2, 2, 2, 2}))         // 5
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
