package main

import "fmt"

/*
392. 判断子序列
https://leetcode.cn/problems/is-subsequence/
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
进阶：
如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
输入：s = "abc", t = "ahbgdc"
输出：true
*/
// 方法1: dp
// 方法2: 双指针
func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	left, right := 0, 0
	for right < len(t) {
		if t[right] == s[left] {
			left++
			if left == len(s) {
				return true
			}
		}
		right++
	}
	return false
}

/*
792. 匹配子序列的单词数
https://leetcode.cn/problems/number-of-matching-subsequences/description/
给定字符串 s 和字符串数组 words, 返回  words[i] 中是s的子序列的单词个数 。
字符串的 子序列 是从原始字符串中生成的新字符串，可以从中删去一些字符(可以是none)，而不改变其余字符的相对顺序。
例如， “ace” 是 “abcde” 的子序列。
输入: s = "abcde", words = ["a","bb","acd","ace"]
输出: 3
解释: 有三个是 s 的子序列的单词: "a", "acd", "ace"。
*/
func numMatchingSubseq(s string, words []string) int {
	// 对 s 进行预处理，记录 char -> 该 char 的索引列表
	charToIndexes := make([][]int, 26)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if charToIndexes[c-'a'] == nil {
			charToIndexes[c-'a'] = []int{}
		}
		charToIndexes[c-'a'] = append(charToIndexes[c-'a'], i)
	}

	res := 0
	for _, word := range words {
		// 字符串 word 上的指针 i
		i := 0
		// 字符串 s 上的指针 j
		j := 0
		// 现在判断 word 是否是 s 的子序列
		// 借助 charToIndexes 查找 word 中每个字符在 s 中的索引
		for i < len(word) {
			c := word[i]
			// 整个 s 压根儿没有字符 word[i]
			if charToIndexes[c-'a'] == nil {
				break
			}
			// 二分搜索大于等于 j 的最小索引
			// 即在 s[j..] 中搜索等于 word[i] 的最小索引
			pos := leftBound(charToIndexes[c-'a'], j)
			if pos == len(charToIndexes[c-'a']) {
				break
			}
			j = charToIndexes[c-'a'][pos]
			// 如果找到，即 word[i] == s[j]，继续往后匹配
			j++
			i++
		}
		// 如果 word 完成匹配，则是 s 的子序列
		if i == len(word) {
			res++
		}
	}

	return res
}

// 查找左侧边界的二分查找
func leftBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if target > arr[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

/*
658. 找到 K 个最接近的元素
https://leetcode.cn/problems/find-k-closest-elements/description/
给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
整数 a 比整数 b 更接近 x 需要满足：
|a - x| < |b - x| 或者
|a - x| == |b - x| 且 a < b
输入：arr = [1,2,3,4,5], k = 4, x = 3
输出：[1,2,3,4]
*/
func findClosestElements(arr []int, k int, x int) []int {
	var results []int
	p := leftBound(arr, x)
	left, right := p-1, p
	// 拓展区间，直到区间内包含k个数
	for right-left-1 < k {
		if left == -1 {
			right++
		} else if right == len(arr) {
			left--
		} else if x-arr[left] > arr[right]-x {
			right++
		} else {
			left--
		}
	}
	for i := left + 1; i < right; i++ {
		results = append(results, arr[i])
	}
	return results
}

// 162. 寻找峰值
// https://leetcode.cn/problems/find-peak-element/description/
// 峰值元素是指其值严格大于左右相邻值的元素。
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
// 你可以假设 nums[-1] = nums[n] = -∞ 。
// 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
// 输入：nums = [1,2,3,1]
// 输出：2
// 解释：3 是峰值元素，你的函数应该返回其索引 2。
// 输入：nums = [1,2,1,3,5,6,4]
// 输出：1 或 5
// 解释：你的函数可以返回索引 1，其峰值元素为 2；
// 或者返回索引 5， 其峰值元素为 6。
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-2
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1 // 上坡，峰值在右边，收缩左边界
		} else if nums[mid] > nums[mid+1] {
			right = mid - 1 // 下坡，峰值在左边，收缩右边界
		} else if nums[mid] == nums[mid+1] {
			right = mid - 1
		}
	}
	return left
}

func findPeakElement2(nums []int) int {
	// 取两端都闭的二分搜索
	left, right := 0, len(nums)-1
	// 因为题目必然有解，所以设置 left == right 为结束条件
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			// mid 本身就是峰值或其左侧有一个峰值
			right = mid
		} else {
			// mid 右侧有一个峰值
			left = mid + 1
		}
	}
	return left
}

/*
852. 山脉数组的峰顶索引
https://leetcode.cn/problems/peak-index-in-a-mountain-array/description/
给定一个长度为 n 的整数 山脉 数组 arr ，其中的值递增到一个 峰值元素 然后递减。
返回峰值元素的下标。
你必须设计并实现时间复杂度为 O(log(n)) 的解决方案。
输入：arr = [0,2,1,0]
输出：1
*/
func peakIndexInMountainArray(nums []int) int {
	left, right := 0, len(nums)-2
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1 // 上坡，峰值在右边，收缩左边界
		} else if nums[mid] > nums[mid+1] {
			right = mid - 1 // 下坡，峰值在左边，收缩右边界
		} else if nums[mid] == nums[mid+1] {
			right = mid - 1
		}
	}
	return left
}

// 33.搜索旋转排序数组
// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
// 整数数组 nums 按升序排列，数组中的值 互不相同 。
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 向左旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 下标 3 上向左旋转后可能变为 [4,5,6,7,0,1,2] 。
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4
// 输入：nums = [4,5,6,7,0,1,2], target = 3
// 输出：-1
func search2(nums []int, target int) int {
	// 2次二分：第一次找到最小值下标，第二次在有序数组中找值
	n := len(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < nums[n-1] { // 中点在第二段，最小值在左边，收缩右边界
			right = mid - 1
		} else if nums[mid] > nums[n-1] {
			left = mid + 1
		} else if nums[mid] == nums[n-1] {
			right = mid - 1
		}
	}
	minIndex := left
	if nums[n-1] < target {
		left, right = 0, minIndex
	} else {
		left, right = minIndex, n-1
	}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1
}

/*
81. 搜索旋转排序数组 II
https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/
已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。
给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。
你必须尽可能减少整个操作步骤。
输入：nums = [2,5,6,0,0,1,2], target = 0
输出：true
输入：nums = [2,5,6,0,0,1,2], target = 3
输出：false
*/
func searchRotatedSortedArray(nums []int, target int) bool {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		// 需要在计算 mid 之前收缩左右边界去重
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		for left < right && nums[right] == nums[right-1] {
			right--
		}
		// 1.找到最小值下标
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < nums[n-1] {
			// mid在第二段,
			right = mid - 1
		} else if nums[mid] > nums[n-1] {
			// mid在第一段，最小值在后边，收缩左边界
			left = mid + 1
		} else if nums[mid] == nums[n-1] {
			// 最小值在左边，收缩右边界
			right = mid - 1
		}
	}
	// 2.left是最小值下标，在区间使用二分
	l, r := 0, left-1
	if target <= nums[n-1] {
		l, r = left, n-1
	}
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return true
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}

// 153. 寻找旋转排序数组中的最小值
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/
// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
// 若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
// 若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
// 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
// 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
// 输入：nums = [3,4,5,1,2]
// 输出：1
// 解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。
// 输入：nums = [4,5,6,7,0,1,2]
// 输出：0
// 解释：原数组为 [0,1,2,4,5,6,7] ，旋转 4 次得到输入数组。
func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < nums[n-1] {
			// 中点在第二段，最小值在中点左边，收缩右边界
			right = mid - 1
		} else if nums[mid] > nums[n-1] {
			left = mid + 1
		} else if nums[mid] == nums[n-1] {
			right = mid - 1 // 最小值在左侧，继续收缩右边界
		}
	}
	return nums[left]
}

func main() {
	fmt.Println(matrixReshape([][]int{[]int{1, 2}, []int{3, 4}}, 4, 1))
}
