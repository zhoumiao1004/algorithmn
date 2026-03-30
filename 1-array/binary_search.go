package main

import "fmt"

// 704. 二分查找
// 本质：通过收缩左右边界，缩小搜索范围
// https://leetcode.cn/problems/binary-search/description/
// 输入: nums = [-1,0,3,5,9,12], target = 9 输出: 4
// 解释: 9 出现在 nums 中并且下标为 4
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// LCR 172. 统计目标成绩的出现次数
// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/description/
// 某班级考试成绩按非严格递增顺序记录于整数数组 scores，请返回目标成绩 target 的出现次数。
// 输入: scores = [2, 2, 3, 4, 4, 4, 5, 6, 6, 8], target = 4
// 输出: 3
func countTarget(scores []int, target int) int {
	n := len(scores)
	left := getLeft(scores, target)
	right := getRight(scores, target)
	if left < 0 || right > n-1 {
		return 0
	}
	return right - left + 1
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := getLeft(nums, target)
	right := getRight(nums, target)
	if left < 0 || left > len(nums)-1 || right < 0 || right > len(nums)-1 {
		return []int{-1, -1}
	}
	if nums[left] != target || nums[right] != target {
		return []int{-1, -1}
	}
	return []int{left, right}
}

func getLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}
	return left
}

func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}
	return right
}

// 35. 搜索插入位置
// https://leetcode.cn/problems/search-insert-position/description/
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
// 输入: nums = [1,3,5,6], target = 2
// 输出: 1
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			// 收缩右边界
			right = mid - 1
		}
	}
	// right指向第一个小于的位置，left指向后一个位置
	// return right + 1
	return left
}

// 875. 爱吃香蕉的珂珂
// 珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 h 小时后回来。
// 珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
// 珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
// 返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。
// 输入：piles = [3,6,7,11], h = 8
// 输出：4
// 输入：piles = [30,11,23,4,20], h = 5
// 输出：30
func minEatingSpeed(piles []int, h int) int {
	// 速度为x，需要f(x)小时吃完
	var f func(piles []int, x int) int
	f = func(piles []int, x int) int {
		hours := 0
		for i := 0; i < len(piles); i++ {
			hours += piles[i] / x
			if piles[i]%x > 0 {
				hours++
			}
		}
		return hours
	}
	left, right := 1, 1000000000+1
	for left <= right {
		mid := left + (right-left)/2
		if f(piles, mid) == h {
			// 搜索左侧边界，就要收缩右侧边界
			right = mid - 1
		} else if f(piles, mid) < h {
			// mid速度快了导致需要的时间小于h，需要让f(x）返回大一点，速度mid要降低，所以收缩右边界
			right = mid - 1
		} else if f(piles, mid) > h {
			left = mid + 1
		}
	}
	return left
}

// 1011. 在 D 天内送达包裹的能力
// https://leetcode.cn/problems/capacity-to-ship-packages-within-d-days/description/
// 传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。
// 传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量（weights）的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
// 返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。
// 输入：weights = [1,2,3,4,5,6,7,8,9,10], days = 5
// 输出：15
// 解释：船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
// 第 1 天：1, 2, 3, 4, 5
// 第 2 天：6, 7
// 第 3 天：8
// 第 4 天：9
// 第 5 天：10
func shipWithinDays(weights []int, days int) int {
	// 运载能力为x，需要x天运完货物
	var leastDays func(weights []int, x int) int
	leastDays = func(weights []int, x int) int {
		days := 1
		s := 0
		for i := 0; i < len(weights); i++ {
			if s+weights[i] > x {
				days++
				s = weights[i]
			} else {
				s += weights[i]
			}
		}
		return days
	}
	// left, right := 1, 500*5*10000+1
	left, right := 0, 1
	for _, w := range weights {
		if left < w {
			left = w
		}
		right += w
	}
	for left <= right {
		mid := left + (right-left)/2
		if leastDays(weights, mid) == days {
			// 寻找左边界，所以要收缩右边界
			right = mid - 1
		} else if leastDays(weights, mid) < days {
			// 速度mid快了
			right = mid - 1
		} else if leastDays(weights, mid) > days {
			left = mid + 1
		}
	}
	return left
}

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

// 思路2: 收缩右边界
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

// 540. 有序数组中的单一元素
// https://leetcode.cn/problems/single-element-in-a-sorted-array/description/
// 给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
// 请你找出并返回只出现一次的那个数。
// 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
// 输入: nums = [1,1,2,3,3,4,4,8,8]
// 输出: 2
// 输入: nums =  [3,3,7,7,10,11,11]
// 输出: 10
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if mid%2 == 0 {
			if nums[mid] == nums[mid-1] {
				right = mid
			} else {
				left = mid
			}
		} else if mid%2 == 1 {
			if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return nums[left]
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))       // 1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))       // 2
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))            // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))      // 0
	fmt.Println(findMin([]int{11, 13, 15, 17}))           // 11
	fmt.Println(search2([]int{4, 5, 6, 7, 0, 1, 2}, 0))   // 4
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // 3,4
}
