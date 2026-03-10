package main

import "fmt"

// 48. 旋转图像
// https://leetcode.cn/problems/rotate-image/description/
// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[[7,4,1],[8,5,2],[9,6,3]]
func rotate(matrix [][]int) {
	n := len(matrix)
	// 轴对称
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 反转每一行
	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}

// 151. 反转字符串中的单词
// https://leetcode.cn/problems/reverse-words-in-a-string/
func reverseWords(s string) string {

}

// 61. 旋转链表
// https://leetcode.cn/problems/rotate-list/
// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[4,5,1,2,3]
func rotateRight(head *ListNode, k int) *ListNode {

}

// 54. 螺旋矩阵
// https://leetcode.cn/problems/spiral-matrix/
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
func spiralOrder(matrix [][]int) []int {
	var result []int
	m := len(matrix)
	n := len(matrix[0])
	upper_bound, lower_bound := 0, m-1
	left_bound, right_bound := 0, n-1
	for len(result) < m*n {
		if upper_bound <= lower_bound {
			for j := left_bound; j <= right_bound; j++ {
				result = append(result, matrix[upper_bound][j])
			}
			upper_bound++
		}
		if left_bound <= right_bound {
			for i := upper_bound; i <= lower_bound; i++ {
				result = append(result, matrix[i][right_bound])
			}
			right_bound--
		}
		if upper_bound <= lower_bound {
			for j := right_bound; j >= left_bound; j-- {
				result = append(result, matrix[lower_bound][j])
			}
			lower_bound--
		}
		if left_bound <= right_bound {
			for i := lower_bound; i >= upper_bound; i-- {
				result = append(result, matrix[i][left_bound])
			}
			left_bound++
		}
	}
	return result
}

// 59. 螺旋矩阵 II
// https://leetcode.cn/problems/spiral-matrix-ii/description/
// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]
func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	startx, starty := 0, 0
	offset := 0
	cnt := 1
	loop := n / 2
	for loop > 0 {
		i, j := startx, starty
		for ; j < n-1-offset; j++ {
			result[i][j] = cnt
			cnt++
		}
		for ; i < n-1-offset; i++ {
			result[i][j] = cnt
			cnt++
		}
		for ; j > offset; j-- {
			result[i][j] = cnt
			cnt++
		}
		for ; i > offset; i-- {
			result[i][j] = cnt
			cnt++
		}
		startx++
		starty++
		offset++
		loop--
	}
	if n%2 == 1 {
		result[n/2][n/2] = cnt
	}
	return result
}

// 74.搜索二维矩阵
// https://leetcode.cn/problems/search-a-2d-matrix/
// 给你一个满足下述两条属性的 m x n 整数矩阵：
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
// 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// 输出：true
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 纵向二分,找左边界
	left, right := 0, m-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid][0] > target {
			right = mid - 1
		} else if matrix[mid][0] < target {
			left = mid + 1
		} else {
			return true
		}
	}
	// 第一列没找到：right指向的第一个小于target的位置
	fmt.Println(left, right)
	if right < 0 {
		right++
	}
	// 横向二分
	row := right
	left, right = 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 240.搜索二维矩阵 II
// https://leetcode.cn/problems/search-a-2d-matrix-ii/
// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
// 输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
// 输出：true
func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := m-1, 0
	for i >= 0 && j < n {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			j++
		} else if matrix[i][j] > target {
			i--
		}
	}
	return false
}
