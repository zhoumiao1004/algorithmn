package main

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

// 167. 两数之和 II - 输入有序数组
// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
// 给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
// 以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
// 你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
// 你所设计的解决方案必须只使用常量级的额外空间。
// 输入：numbers = [2,7,11,15], target = 9
// 输出：[1,2]
// 解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{-1, -1}
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
