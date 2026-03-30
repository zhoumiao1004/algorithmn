package main

import (
	"fmt"
	"sort"
)

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
	tmp := append([]int{}, nums...)
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
}
