package main

import "fmt"

// 525. 连续数组
// https://leetcode.cn/problems/contiguous-array/description/
// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
// 输入：nums = [0,1]
// 输出：2
// 说明：[0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
// 输入：nums = [0,1,1,1,1,1,0,0,0]
// 输出：6
// 解释：[1,1,1,0,0,0] 是具有相同数量 0 和 1 的最长连续子数组。
// 思路1: 把0当成-1，和为0的最长子数组
func findMaxLength(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1) // preSum[i] 代表 [0..i-1]的区间和
	for i := 1; i <= n; i++ {
		if nums[i-1] == 0 {
			preSum[i] = preSum[i-1] - 1
		} else {
			preSum[i] = preSum[i-1] + 1
		}
	}

	result := 0
	indexMap := make(map[int]int)
	for i := 0; i <= n; i++ {
		// 查看hashmap中是否已经存在左边界
		index, ok := indexMap[preSum[i]]
		if !ok {
			indexMap[preSum[i]] = i // key不存在：保存左边界
		} else {
			result = max(result, i-index) // key已存在：更新结果 (不能覆盖value，因为要求最大长度)
		}
	}

	return result
}

func findMaxLength2(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1) // preSum[i] 代表 [0..i-1]的区间和
	result := 0
	indexMap := make(map[int]int)
	indexMap[0] = 0 // 对0特殊处理
	for i := 1; i <= n; i++ {
		if nums[i-1] == 0 {
			preSum[i] = preSum[i-1] - 1
		} else {
			preSum[i] = preSum[i-1] + 1
		}
		// 查看hashmap中是否已经存在左边界
		index, ok := indexMap[preSum[i]]
		if !ok {
			indexMap[preSum[i]] = i // key不存在：保存左边界
		} else {
			result = max(result, i-index) // key已存在：更新结果 (不能覆盖value，因为要求最大长度)
		}
	}
	return result
}

// 523. 连续的子数组和
// https://leetcode.cn/problems/continuous-subarray-sum/
// 给你一个整数数组 nums 和一个整数 k ，如果 nums 有一个 好的子数组 返回 true ，否则返回 false：
// 一个 好的子数组 是：
// 长度 至少为 2 ，且
// 子数组元素总和为 k 的倍数。
// 注意：
// 子数组 是数组中 连续 的部分。
// 如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终 视为 k 的一个倍数。
// 输入：nums = [23,2,4,6,7], k = 6
// 输出：true
// 解释：[2,4] 是一个大小为 2 的子数组，并且和为 6 。
// 分析：(preSum[i] - preSum[j]) % k == 0 其实就是 preSum[i] % k == preSum[j] % k。
func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1] // 计算前缀和
	}

	valToIndex := make(map[int]int)
	for i := 0; i <= n; i++ {
		val := preSum[i] % k
		if index, ok := valToIndex[val]; ok {
			if i-index >= 2 {
				return true
			}
		} else {
			valToIndex[val] = i
		}
	}
	return false
}

func checkSubarraySum2(nums []int, k int) bool {
	n := len(nums)
	preSum := make([]int, n+1)
	valToIndex := make(map[int]int)
	valToIndex[0] = 0 // 对0特殊处理
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1] // 计算前缀和
		val := preSum[i] % k
		if index, ok := valToIndex[val]; ok {
			if i-index >= 2 {
				return true
			}
		} else {
			valToIndex[val] = i
		}
	}
	return false
}

// 560. 和为 K 的子数组
// https://leetcode.cn/problems/subarray-sum-equals-k/
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
// 子数组是数组中元素的连续非空序列。
// 输入：nums = [1,1,1], k = 2
// 输出：2
// 输入：nums = [1,2,3], k = 3
// 输出：2
func subarraySum(nums []int, k int) int {
	result := 0
	n := len(nums)
	cntMap := make(map[int]int)
	cntMap[0] = 1 // 对0特殊处理
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		if cnt, ok := cntMap[preSum[i]-k]; ok {
			result += cnt
		}
		cntMap[preSum[i]]++
	}
	return result
}

// 1124. 表现良好的最长时间段
// https://leetcode.cn/problems/longest-well-performing-interval/description/
// 给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
// 我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
// 所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
// 请你返回「表现良好时间段」的最大长度。
// 输入：hours = [9,9,6,0,6,6,9]
// 输出：3
// 解释：最长的表现良好时间段是 [9,9,6]。
// 思路1: 前缀和
func longestWPI(hours []int) int {
	result := 0
	valToIndex := make(map[int]int)
	valToIndex[0] = 0 // 对0特殊处理
	n := len(hours)
	preSum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if hours[i-1] > 8 {
			preSum[i] = preSum[i-1] + 1
		} else {
			preSum[i] = preSum[i-1] - 1
		}
		if preSum[i] > 0 {
			result = max(result, i)
		} else {
			index, ok := valToIndex[preSum[i]-1]
			if ok {
				result = max(result, i-index)
			}
		}
		if _, ok := valToIndex[preSum[i]]; !ok {
			valToIndex[preSum[i]] = i // 注意：由于求最长，所以只要最早出现的index，不要覆盖
		}
	}

	return result
}

// 974. 和可被 K 整除的子数组
// https://leetcode.cn/problems/subarray-sums-divisible-by-k/description/
// 给定一个整数数组 nums 和一个整数 k ，返回其中元素之和可被 k 整除的非空 子数组 的数目。
// 子数组 是数组中 连续 的部分。
// 输入：nums = [4,5,0,-2,-3,1], k = 5
// 输出：7
// 解释：
// 有 7 个子数组满足其元素之和可被 k = 5 整除：
// [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
func subarraysDivByK(nums []int, k int) int {
	result := 0
	n := len(nums)
	cntMap := make(map[int]int)
	cntMap[0] = 1 // 对0特殊处理
	preSum := make([]int, n+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		r := preSum[i] % k // preSum[i] - preSum[j] % k == 0 => preSum[i] % k == preSum[j] % k
		if r < 0 {
			r += k
		}
		cnt, ok := cntMap[r]
		if ok {
			result += cnt
		}
		cntMap[r]++
	}

	return result
}

func main() {
	fmt.Println(subarraysDivByK([]int{-1, 2, 9}, 2))     // 2
	fmt.Println(subarraysDivByK([]int{2, -2, 2, -4}, 6)) // 2
}
