package main

import (
	"fmt"
	"math"
	"sort"
)

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

// 速度为x，需要f(x)小时吃完
func f(piles []int, x int) int {
	hours := 0
	for i := 0; i < len(piles); i++ {
		hours += piles[i] / x
		if piles[i]%x > 0 {
			hours++
		}
	}
	return hours
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
		if f2(weights, mid) == days {
			// 寻找左边界，所以要收缩右边界
			right = mid - 1
		} else if f2(weights, mid) < days {
			// 速度mid快了
			right = mid - 1
		} else if f2(weights, mid) > days {
			left = mid + 1
		}
	}
	return left
}

// 运载能力为x，需要x天运完货物
func f2(weights []int, x int) int {
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

// 410. 分割数组的最大值
// https://leetcode.cn/problems/split-array-largest-sum/description/
// 给定一个非负整数数组 nums 和一个整数 k ，你需要将这个数组分成 k 个非空的连续子数组，使得这 k 个子数组各自和的最大值 最小。
// 返回分割后最小的和的最大值。
// 子数组 是数组中连续的部分。
// 输入：nums = [7,2,5,10,8], k = 2
// 输出：18
// 解释：一共有四种方法将 nums 分割为 2 个子数组。
// 其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
// 因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
func splitArray(nums []int, k int) int {
	return 0
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

// 209. 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum/description/
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt
	left := 0
	s := 0
	for right := 0; right < len(nums); right++ {
		s += nums[right]
		for s >= target {
			result = min(result, right-left+1)
			s -= nums[left]
			left++
		}
	}
	return result
}

// 1365. 有多少小于当前数字的数字
// https://leetcode.cn/problems/how-many-numbers-are-smaller-than-the-current-number/description/
// 给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。
// 换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。
// 以数组形式返回答案。
// 输入：nums = [8,1,2,2,3]
// 输出：[4,0,1,1,3]
func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)
	m := make(map[int]int)
	for i, val := range tmp {
		if _, ok := m[val]; !ok {
			m[val] = i
		}
	}
	for i, val := range nums {
		result[i] = m[val]
	}
	return result
}

// 941. 有效的山脉数组
// 给定一个整数数组 arr，如果它是有效的山脉数组就返回 true，否则返回 false。
// 让我们回顾一下，如果 arr 满足下述条件，那么它是一个山脉数组：
// arr.length >= 3
// 在 0 < i < arr.length - 1 条件下，存在 i 使得：
// arr[0] < arr[1] < ... arr[i-1] < arr[i]
// arr[i] > arr[i+1] > ... > arr[arr.length - 1]
// 输入：arr = [0,3,2,1]
// 输出：true
func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	incFlag := false
	decFlag := false
	i := 1
	for ; i < n && arr[i-1] < arr[i]; i++ {
		incFlag = true
	}
	for ; i < n && arr[i-1] > arr[i]; i++ {
		decFlag = true
	}
	return i == n && incFlag && decFlag
}

// 1207. 独一无二的出现次数
// https://leetcode.cn/problems/unique-number-of-occurrences/
// 输入：arr = [1,2,2,1,1,3]
// 输出：true
// 解释：在该数组中，1 出现了 3 次，2 出现了 2 次，3 只出现了 1 次。没有两个数的出现次数相同。
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, val := range arr {
		m[val]++
	}
	freq := make(map[int]bool)
	for _, val := range m {
		if freq[val] {
			return false
		}
		freq[val] = true
	}
	return true
}

// 189. 轮转数组
// https://leetcode.cn/problems/rotate-array/description/
// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4]
// 方法1：原地
func rotateInplace(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	for i := 0; i < k%n; i++ {
		val := nums[n-1]
		for j := n - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = val
	}
}

// 方法2：取余
func rotateMod(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	tmp := append([]int{}, nums...)
	for i := 0; i < n; i++ {
		nums[(i+k)%n] = tmp[i]
	}
}

// 922. 按奇偶排序数组 II
// https://leetcode.cn/problems/sort-array-by-parity-ii/description/
// 给定一个非负整数数组 nums，  nums 中一半整数是 奇数 ，一半整数是 偶数 。
// 对数组进行排序，以便当 nums[i] 为奇数时，i 也是 奇数 ；当 nums[i] 为偶数时， i 也是 偶数 。
// 你可以返回 任何满足上述条件的数组作为答案 。
// 输入：nums = [4,2,5,7]
// 输出：[4,5,2,7]
// 解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	even := 0
	odd := 1
	for even < n && odd < n {
		for even < n && nums[even]%2 == 0 {
			even += 2
		}
		for odd < n && nums[odd]%2 == 1 {
			odd += 2
		}
		if even < n && odd < n {
			nums[even], nums[odd] = nums[odd], nums[even]
			even += 2
			odd += 2
		}
	}
	return nums
}

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60}}, 11))
	fmt.Println(searchMatrix([][]int{
		{1}}, 1))
	fmt.Println(searchMatrix([][]int{
		{1}, {3}}, 3))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))       // 1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))       // 2
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))            // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))      // 0
	fmt.Println(findMin([]int{11, 13, 15, 17}))           // 11
	fmt.Println(search2([]int{4, 5, 6, 7, 0, 1, 2}, 0))   // 4
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // 3,4
}
