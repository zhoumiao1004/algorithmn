package main

import "fmt"

// 977.有序数组的平方
// https://leetcode.cn/problems/squares-of-a-sorted-array/description/
// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
func sortedSquares(nums []int) []int {
	n := len(nums)
	results := make([]int, n)
	left, right := 0, n-1
	k := n - 1
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			results[k] = nums[right] * nums[right]
			right--
		} else {
			results[k] = nums[left] * nums[left]
			left++
		}
		k--
	}
	return results
}

// 344. 反转字符串
// https://leetcode.cn/problems/reverse-string/description/
// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 5. 最长回文子串
// 给你一个字符串 s，找到 s 中最长的 回文 子串。
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 思路1：中心扩散
func longestPalindrome(s string) string {
	var palindrome func(s string, left, right int) string
	palindrome = func(s string, left, right int) string {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return s[left+1 : right]
	}

	result := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(result) {
			result = s1
		}
		if len(s2) > len(result) {
			result = s2
		}
	}
	return result
}

// 思路2：dp
func longestPalindrome2(s string) string {
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

func longestPalindrome3(s string) string {
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

// 75. 颜色分类
// https://leetcode.cn/problems/sort-colors/
// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
// 必须在不使用库内置的 sort 函数的情况下解决这个问题。
func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	p := 0
	for p <= p2 {
		if nums[p] == 0 {
			nums[p0], nums[p] = nums[p], nums[p0]
			p0++
		} else if nums[p] == 2 {
			nums[p], nums[p2] = nums[p2], nums[p]
			p2--
		} else if nums[p] == 1 {
			p++
		}
		// 由于p找到0就会和p0位置的数字换，所以p0一直增加，由于p0之前都是0，所以p需要>=p0
		if p < p0 {
			p = p0
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
	fmt.Println(longestPalindrome("babad"))  // bab
	fmt.Println(longestPalindrome2("babad")) // bab
}
