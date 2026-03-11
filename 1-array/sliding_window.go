package main

import (
	"fmt"
	"math"
)

/* 滑动窗口模板
// 索引区间 [left, right) 是窗口
int left = 0, right = 0;
while (right < nums.size()) {
    // 增大窗口
    window.addLast(nums[right]);
    right++;

    while (window needs shrink) {
        // 缩小窗口
        window.removeFirst(nums[left]);
        left++;
    }
}*/

// 76. 最小覆盖子串
// https://leetcode.cn/problems/minimum-window-substring/
// 给定两个字符串 s 和 t，长度分别是 m 和 n，返回 s 中的 最短窗口 子串，使得该子串包含 t 中的每一个字符（包括重复字符）。如果没有这样的子串，返回空字符串 ""。
// 测试用例保证答案唯一。
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。、
func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	window := make(map[byte]int)
	left, right := 0, 0
	valid := 0
	start, length := 0, math.MaxInt
	for right < len(s) {
		c := s[right]
		right++ // 扩大窗口
		// 窗口内数据更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// fmt.Println(window)
		// 判断左侧窗口是否需要收缩
		for valid == len(need) {
			// 更新结果
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			// 窗口内数据更新
			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					valid--
				}
			}
			window[d]--
		}
	}
	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}

// 567. 字符串的排列
// https://leetcode.cn/problems/permutation-in-string/description/
// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的 排列。如果是，返回 true ；否则，返回 false 。
// 换句话说，s1 的排列之一是 s2 的 子串 。
// 输入：s1 = "ab" s2 = "eidbaooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("ba").
func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	window := make(map[byte]int)
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		// 窗口内数据更新
		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}
		// 判断左侧窗口是否要收缩
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			// 窗口内数据更新
			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

// hash table
func checkInclusion2(s1 string, s2 string) bool {
	len1, len2 := len(s1), len(s2)
	if len1 > len2 {
		return false
	}
	var arr1, arr2 [26]int
	for i := 0; i < len1; i++ {
		arr1[s1[i]-'a']++
		arr2[s2[i]-'a']++
	}
	l, r := 0, len1-1
	for r < len2 {
		if arr1 == arr2 {
			return true
		}
		r++
		if r < len2 {
			arr2[s2[r]-'a']++
			arr2[s2[l]-'a']--
			l++
		}
	}
	return false
}

// 438. 找到字符串中所有字母异位词
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
func findAnagrams(s string, p string) []int {
	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	window := make(map[byte]int)
	left, right := 0, 0
	valid := 0
	var results []int
	for right < len(s) {
		c := s[right]
		right++
		// 窗口内数据更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 判断左侧窗口是否要收缩
		for right-left >= len(p) {
			if valid == len(need) {
				results = append(results, left)
			}
			d := s[left]
			left++
			// 窗口内疏忽更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return results
}

// 3. 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
func lengthOfLongestSubstring(s string) int {
	result := 0
	window := make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		right++
		// 窗口内数据更新
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			// 窗口内数据更新
			window[d]--
		}
		result = max(result, right-left)
	}
	return result
}

func lengthOfLongestSubstring2(s string) int {
	result := 0
	m := make(map[byte]int)
	l := 0
	for r := 0; r < len(s); r++ {
		m[s[r]]++
		for m[s[r]] > 1 {
			m[s[l]]--
			l++
		}
		result = max(result, r-l+1)
	}
	return result
}

// 1658. 将 x 减到 0 的最小操作数
// https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/description/
// 给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数组以供接下来的操作使用。
// 如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。
// 输入：nums = [1,1,4,2,3], x = 5
// 输出：2
// 解释：最佳解决方案是移除后两个元素，将 x 减到 0 。
func minOperations(nums []int, x int) int {
	n := len(nums)
	s := 0
	for i := 0; i < n; i++ {
		s += nums[i]
	}
	target := s - x
	left, right := 0, 0
	sum := 0
	maxLen := -1
	for right < n {
		sum += nums[right]
		right++
		for left < right && sum > target {
			sum -= nums[left]
			left++
		}
		if sum == target {
			maxLen = max(maxLen, right-left)
		}
	}
	if maxLen == -1 {
		return -1
	}
	return n - maxLen
}

// 713. 乘积小于 K 的子数组
// https://leetcode.cn/problems/subarray-product-less-than-k/
// 给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。
// 输入：nums = [10,5,2,6], k = 100
// 输出：8
// 解释：8 个乘积小于 100 的子数组分别为：[10]、[5]、[2]、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
// 需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。
func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)
	left, right := 0, 0
	s := 1
	result := 0
	for right < n {
		s *= nums[right]
		right++
		for left < right && s >= k {
			s /= nums[left]
			left++
		}
		// 计算以right结尾的子数组个数
		result += right - left
	}
	return result
}

// 1004. 最大连续1的个数 III
// https://leetcode.cn/problems/max-consecutive-ones-iii/description/
// 给定一个二进制数组 nums 和一个整数 k，假设最多可以翻转 k 个 0 ，则返回执行操作后 数组中连续 1 的最大个数 。
// 输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
// 输出：6
// 解释：[1,1,1,0,0,1,1,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 6。
func longestOnes(nums []int, k int) int {
	n := len(nums)
	var hash [2]int
	left, right := 0, 0
	result := 0
	for right < n {
		val := nums[right]
		right++
		hash[val]++
		for hash[0] > k {
			d := nums[left]
			left++
			hash[d]--
		}
		result = max(result, right-left)
	}
	return result
}

// 424. 替换后的最长重复字符
// https://leetcode.cn/problems/longest-repeating-character-replacement/description/
// 给你一个字符串 s 和一个整数 k 。你可以选择字符串中的任一字符，并将其更改为任何其他大写英文字符。该操作最多可执行 k 次。
// 在执行上述操作后，返回 包含相同字母的最长子字符串的长度。
// 输入：s = "ABAB", k = 2
// 输出：4
// 解释：用两个'A'替换为两个'B',反之亦然。
// 输入：s = "AABABBA", k = 1
// 输出：4
func characterReplacement(s string, k int) int {
	var hash [26]int
	left, right := 0, 0
	maxCnt := 0
	result := 0
	for right < len(s) {
		// 扩大窗口
		c := s[right]
		right++
		hash[c-'A']++
		maxCnt = max(maxCnt, hash[c-'A'])

		for right-left-maxCnt > k {
			// 左侧窗口要收缩
			d := s[left]
			hash[d-'A']--
			left++
		}
		result = max(result, right-left)
	}
	return result
}

// 219. 存在重复元素 II
// https://leetcode.cn/problems/contains-duplicate-ii/submissions/705551348/
// 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
// 输入：nums = [1,2,3,1], k = 3
// 输出：true
func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	window := make(map[int]int)
	right := 0
	for right < n {
		val := nums[right]
		right++
		index, ok := window[val]
		if ok && right-index <= k {
			return true
		}
		window[val] = right
	}
	return false
}

func containsNearbyDuplicate1(nums []int, k int) bool {
	left, right := 0, 0
	window := make(map[int]bool)
	// 滑动窗口算法框架，维护一个大小为 k 的窗口
	for right < len(nums) {
		// 扩大窗口
		if window[nums[right]] {
			return true
		}
		window[nums[right]] = true
		right++

		if right-left > k {
			// 当窗口的大小大于 k 时，缩小窗口
			delete(window, nums[left])
			left++
		}
	}
	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		index, ok := m[nums[i]]
		if ok && i-index <= k {
			return true
		}
		m[nums[i]] = i
	}
	return false
}

/*
220. 存在重复元素 III
https://leetcode.cn/problems/contains-duplicate-iii/
给你一个整数数组 nums 和两个整数 indexDiff 和 valueDiff 。
找出满足下述条件的下标对 (i, j)：
i != j,
abs(i - j) <= indexDiff
abs(nums[i] - nums[j]) <= valueDiff
如果存在，返回 true ；否则，返回 false 。
示例 1：
输入：nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
输出：true
*/
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if indexDiff <= 0 || valueDiff < 0 {
		return false
	}

	getID := func(x, w int) int {
		if x >= 0 {
			return x / w
		}
		return (x+1)/w - 1
	}

	window := make(map[int]int)
	w := valueDiff + 1

	for i := 0; i < len(nums); i++ {
		m := getID(nums[i], w)

		// 为了防止 i == j，所以在扩大窗口之前先判断是否有符合题意的索引对 (i, j)
		// 查找略大于 nums[right] 的那个元素
		if _, ok := window[m]; ok {
			return true
		}
		// 查找略小于 nums[right] 的那个元素
		if v, ok := window[m-1]; ok && math.Abs(float64(nums[i]-v)) < float64(w) {
			return true
		}
		if v, ok := window[m+1]; ok && math.Abs(float64(nums[i]-v)) < float64(w) {
			return true
		}

		// 扩大窗口
		window[m] = nums[i]

		if i >= indexDiff {
			// 缩小窗口
			delete(window, getID(nums[i-indexDiff], w))
		}
	}

	return false
}

/*
209. 长度最小的子数组
https://leetcode.cn/problems/minimum-size-subarray-sum/description/
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
*/
func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt
	s := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		s += nums[right]
		for s >= target {
			result = min(result, right-left+1)
			s -= nums[left]
			left++
		}
	}
	if result == math.MaxInt {
		return 0
	}
	return result
}

/*
395. 至少有 K 个重复字符的最长子串
https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/
给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
如果不存在这样的子字符串，则返回 0。s 仅由小写英文字母组成
输入：s = "aaabb", k = 3
输出：3
解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。
*/
func longestSubstring(s string, k int) int {
	result := 0
	for i := 1; i<=26; i++ {
		r := longestKLetterSubstr(s, k, i)
		result = max(result, r)
	}
	return result
}

// 寻找s中含有count种字符，且每种字符出现次数都大于k的子串长度
func longestKLetterSubstr(s string, k, count int) int {
	result := 0
	left, right := 0, 0
	var window [26]int
	uniqueCount := 0
	validCount := 0
	for right < len(s) {
		c := s[right]
		if window[c-'a'] == 0 {
			uniqueCount++
		}
		window[c-'a']++
		if window[c-'a'] == k {
			validCount++
		}
		right++

		for uniqueCount > count {
			d := s[left]
			if window[d-'a'] == k {
				validCount--
			}
			window[d-'a']--
			if window[d-'a'] == 0 {
				uniqueCount--
			}
			left++
		}
		if validCount == count {
			result = max(result, right-left)
		}
	}
	return result
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3)) // false
	fmt.Println(containsNearbyAlmostDuplicate([]int{-2, 3}, 2, 5))            // true
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 2, 3, 4, 5}, 3, 0)) // true
}
