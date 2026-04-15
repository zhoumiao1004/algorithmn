package main

import "fmt"

// 303. 区域和检索 - 数组不可变
// https://leetcode.cn/problems/range-sum-query-immutable/description/
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
	preSum [][]int // 记录矩阵 [0, 0, i-1, j-1] 的元素和
}

func Constructor2(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	if m == 0 || n == 0 {
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
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}

// 1314. 矩阵区域和
// https://leetcode.cn/problems/matrix-block-sum/description/
// 给你一个 m x n 的矩阵 mat 和一个整数 k ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素 mat[r][c] 的和：
// i - k <= r <= i + k,
// j - k <= c <= j + k 且
// (r, c) 在矩阵内。
// 输入：mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
// 输出：[[12,21,16],[27,45,33],[24,39,28]]
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	numMatrix := Constructor2(mat)
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x1 := max(i-k, 0)
			y1 := max(j-k, 0)
			x2 := min(i+k, m-1)
			y2 := min(j+k, n-1)
			res[i][j] = numMatrix.SumRegion(x1, y1, x2, y2)
		}
	}
	return res
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

// 思路2: 前缀和
func pivotIndex2(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	for i := 1; i < len(preSum); i++ {
		leftSum := preSum[i] - preSum[0]    // i之前元素之和（不包括i
		rightSum := preSum[n] - preSum[i+1] // i之后元素之和（不包括i）
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

func main() {
	c := Constructor([]int{3, 5, 2, -2, 4, 1})
	fmt.Println(c.SumRange(1, 4))
}
