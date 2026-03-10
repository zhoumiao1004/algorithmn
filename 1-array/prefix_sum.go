package main

import "fmt"

// 303. 区域和检索 - 数组不可变
// 给定一个整数数组  nums，处理以下类型的多个查询:
// 计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
// 实现 NumArray 类：
// NumArray(int[] nums) 使用数组 nums 初始化对象
// int sumRange(int left, int right) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，包含 left 和 right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )
// 输入：
// ["NumArray", "sumRange", "sumRange", "sumRange"]
// [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
// 输出：
// [null, 1, -1, -3]
// nums: 	[3, 5, 2, -2, 4, 1]
// preSum:  [0, 3, 8, 10, 8, 12, 13]
type NumArray struct {
	PreSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{PreSum: preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.PreSum[right+1] - this.PreSum[left]
}

// 304. 二维区域和检索 - 矩阵不可变
// https://leetcode.cn/problems/range-sum-query-2d-immutable/description/
// 给定一个二维矩阵 matrix，以下类型的多个请求：
// 计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
// 实现 NumMatrix 类：
// NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
// int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。
// 输入:
// ["NumMatrix","sumRegion","sumRegion","sumRegion"]
// [[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
// 输出:
// [null, 8, 11, 12]
type NumMatrix struct {
	// preSum[i][j] 记录矩阵 [0, 0, i-1, j-1] 的元素和
	preSum [][]int
}

func Constructor2(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])
	if n == 0 {
		return NumMatrix{}
	}
	// 构造前缀和矩阵
	preSum := make([][]int, m+1)
	for i := range preSum {
		preSum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 计算每个矩阵 [0, 0, i, j] 的元素和
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{preSum: preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 目标矩阵之和由四个相邻矩阵运算获得
	// return this.preSum[x2+1][y2+1] - this.preSum[x1][y2+1] - this.preSum[x2+1][y1] + this.preSum[x1][y1]
	return 0
}

// 724. 寻找数组的中心下标
// https://leetcode.cn/problems/find-pivot-index/description/
// 给你一个整数数组 nums ，请计算数组的 中心下标 。
// 数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
// 如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
// 如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。
// 输入：nums = [1, 7, 3, 6, 5, 6]
// 输出：3
// 左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
// 右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。
func pivotIndex(nums []int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	leftSum, rightSum := 0, 0
	for i := 0; i < len(nums); i++ {
		leftSum += nums[i]
		rightSum = s - leftSum + nums[i]
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

func pivotIndex2(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	for i := 1; i < len(preSum); i++ {
		leftSum := preSum[i] - preSum[0]
		rightSum := preSum[n] - preSum[i]
		if leftSum == rightSum {
			return i - 1 // 相对于nums，preSum有一位索引偏移
		}
	}
	return -1
}

// 238. 除了自身以外数组的乘积
// https://leetcode.cn/problems/product-of-array-except-self/description/
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除了 nums[i] 之外其余各元素的乘积 。
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
// prefix: [1,   2,  6, 24]
// suffix: [24, 24, 12,  4]
func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}
	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}
	result := make([]int, n)
	result[0] = suffix[1]
	result[n-1] = prefix[n-2]
	for i := 1; i < n-1; i++ {
		result[i] = prefix[i-1] * suffix[i+1]
	}
	return result
}

// 1352. 最后 K 个数的乘积
// https://leetcode.cn/problems/product-of-the-last-k-numbers/
// 设计一个算法，该算法接受一个整数流并检索该流中最后 k 个整数的乘积。
// 实现 ProductOfNumbers 类：
// ProductOfNumbers() 用一个空的流初始化对象。
// void add(int num) 将数字 num 添加到当前数字列表的最后面。
// int getProduct(int k) 返回当前数字列表中，最后 k 个数字的乘积。你可以假设当前列表中始终 至少 包含 k 个数字。
// 题目数据保证：任何时候，任一连续数字序列的乘积都在 32 位整数范围内，不会溢出。
// 输入：
// ["ProductOfNumbers","add","add","add","add","add","getProduct","getProduct","getProduct","add","getProduct"]
// [[],[3],[0],[2],[5],[4],[2],[3],[4],[8],[2]]
// 输出：
// [null,null,null,null,null,null,20,40,0,null,32]
type ProductOfNumbers struct {
	prefix []int
}

func Constructor3() ProductOfNumbers {
	return ProductOfNumbers{
		prefix: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if len(this.prefix) == 0 {
		this.prefix = []int{1}
	} else {
		this.prefix = append(this.prefix, this.prefix[len(this.prefix)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.prefix)
	if k > n-1 {
		return 0
	}
	return this.prefix[n-1] / this.prefix[n-k-1]
}

// 525. 连续数组
// https://leetcode.cn/problems/contiguous-array/description/
// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
// 输入：nums = [0,1]
// 输出：2
// 说明：[0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
// 输入：nums = [0,1,1,1,1,1,0,0,0]
// 输出：6
// 解释：[1,1,1,0,0,0] 是具有相同数量 0 和 1 的最长连续子数组。
func findMaxLength(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i < len(preSum); i++ {
		if nums[i-1] == 0 {
			preSum[i] = preSum[i-1] - 1
		} else {
			preSum[i] = preSum[i-1] + 1
		}
	}
	result := 0
	indexMap := make(map[int]int)
	for i := 0; i < len(preSum); i++ {
		index, ok := indexMap[preSum[i]]
		if !ok {
			indexMap[preSum[i]] = i
		} else {
			result = max(result, i-index)
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
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	valToIndex := make(map[int]int)
	for i := 0; i < len(preSum); i++ {
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
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	result := 0
	m := make(map[int]int)
	for i := 0; i < len(preSum); i++ {
		if cnt, ok := m[preSum[i]-k]; ok {
			result += cnt
		}
		m[preSum[i]]++
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
func longestWPI(hours []int) int {
	n := len(hours)
	preSum := make([]int, n+1)
	for i := 1; i < len(preSum); i++ {
		if hours[i-1] > 8 {
			preSum[i] = preSum[i-1] + 1
		} else {
			preSum[i] = preSum[i-1] - 1
		}
	}
	result := 0
	valToIndex := make(map[int]int)
	for i := 1; i < len(preSum); i++ {
		if preSum[i] > 0 {
			result = max(result, i)
		} else {
			// preSum[i] - x == 1
			index, ok := valToIndex[preSum[i]-1]
			if ok {
				result = max(result, i-index)
			}
		}
		// 注意：由于求最长，所以只要最早出现的index，不要覆盖
		if _, ok := valToIndex[preSum[i]]; !ok {
			valToIndex[preSum[i]] = i
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
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	result := 0
	m := make(map[int]int)
	fmt.Println("preSum=", preSum)
	for _, val := range preSum {
		// preSum[i] - preSum[j] % k == 0 => preSum[i] % k == preSum[j] % k
		r := val % k
		if r < 0 {
			r += k
		}
		cnt, ok := m[r]
		if ok {
			result += cnt
		}
		m[r]++
	}
	fmt.Println("m=", m)
	return result
}

func main() {
	c := Constructor([]int{3, 5, 2, -2, 4, 1})
	fmt.Println(c.SumRange(1, 4))
	fmt.Println(subarraysDivByK([]int{-1, 2, 9}, 2))     // 2
	fmt.Println(subarraysDivByK([]int{2, -2, 2, -4}, 6)) // 2
}
