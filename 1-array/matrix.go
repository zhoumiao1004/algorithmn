package main

import (
	"fmt"
	"sort"
)

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
	// 倒数第k个节点作为头结点
	if head == nil {
		return nil
	}
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	n := k % length
	if n == 0 {
		return head
	}
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	newHead := slow.Next
	fast.Next = head
	slow.Next = nil
	return newHead
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

/*
566. 重塑矩阵
在 MATLAB 中，有一个非常有用的函数 reshape ，它可以将一个 m x n 矩阵重塑为另一个大小不同（r x c）的新矩阵，但保留其原始数据。
给你一个由二维数组 mat 表示的 m x n 矩阵，以及两个正整数 r 和 c ，分别表示想要的重构的矩阵的行数和列数。
重构后的矩阵需要将原始矩阵的所有元素以相同的 行遍历顺序 填充。
如果具有给定参数的 reshape 操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。
输入：mat = [[1,2],[3,4]], r = 1, c = 4
输出：[[1,2,3,4]]
*/
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}
	tmp := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp[i*n+j] = mat[i][j]
		}
	}
	fmt.Println(tmp)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			result[i][j] = tmp[i*c+j]
		}
	}
	return result
}

func matrixReshape2(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			// index := i*c+j ==> x*n+y
			index := i*c + j
			result[i][j] = mat[index/n][index%n]
		}
	}
	return result
}

func matrixReshape3(mat [][]int, r int, c int) [][]int {
	n := len(mat[0])
	// 如果想成功 reshape，元素个数应该相同
	if r*c != len(mat)*n {
		return mat
	}

	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}

	for i := 0; i < len(mat)*n; i++ {
		set(&res, i, get(&mat, i, n))
	}
	return res
}

// 通过一维坐标访问二维数组中的元素
func get(matrix *[][]int, index int, n int) int {
	// 计算二维中的横纵坐标
	i := index / n
	j := index % n
	return (*matrix)[i][j]
}

// 通过一维坐标设置二维数组中的元素
func set(matrix *[][]int, index int, value int) {
	n := len((*matrix)[0])
	// 计算二维中的横纵坐标
	i := index / n
	j := index % n
	(*matrix)[i][j] = value
}

// 1329. 将矩阵按对角线排序
// https://leetcode.cn/problems/sort-the-matrix-diagonally/
// 输入：mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
// 输出：[[1,1,1,1],[1,2,2,2],[1,2,3,3]]
func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	diaMap := make(map[int][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := i - j
			diaMap[k] = append(diaMap[k], mat[i][j])
		}
	}
	for _, v := range diaMap {
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
	}
	// 结果回填到矩阵
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr := diaMap[i-j]
			mat[i][j] = arr[len(arr)-1]
			diaMap[i-j] = arr[:len(arr)-1]
		}
	}
	return mat
}

// 1260. 二维网格迁移
// https://leetcode.cn/problems/shift-2d-grid/description/
// 给你一个 m 行 n 列的二维网格 grid 和一个整数 k。你需要将 grid 迁移 k 次。
// 每次「迁移」操作将会引发下述活动：
// 位于 grid[i][j]（j < n - 1）的元素将会移动到 grid[i][j + 1]。
// 位于 grid[i][n - 1] 的元素将会移动到 grid[i + 1][0]。
// 位于 grid[m - 1][n - 1] 的元素将会移动到 grid[0][0]。
// 请你返回 k 次迁移操作后最终得到的 二维网格。
// 输入：grid = [[1,2,3],[4,5,6],[7,8,9]], k = 1
// 输出：[[9,1,2],[3,4,5],[6,7,8]]
// 1.除最后一列向右移1位 2.最后一列一到第一列 3.右下角移到左上角
func shiftGrid(grid [][]int, k int) [][]int {

}

// 867. 转置矩阵
// https://leetcode.cn/problems/transpose-matrix/
// 给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。
// 矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[[1,4,7],[2,5,8],[3,6,9]]
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	results := make([][]int, n)
	for i := 0; i < n; i++ {
		results[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			results[j][i] = matrix[i][j]
		}
	}
	return results
}

// 14. 最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	left := 0 // 相同的列
	m, n := len(strs), len(strs[0])
	for j := 0; j < n; j++ {
		// 第j列，对比每一行是否相同
		for i := 1; i < m; i++ {
			if len(strs[i]) <= j || strs[i][j] != strs[0][j] {
				return strs[0][:left]
			}
		}
		left++
	}
	return strs[0][:left]
}

func main() {
	fmt.Println(matrixReshape([][]int{[]int{1, 2}, []int{3, 4}}, 4, 1))
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60}}, 11))
	fmt.Println(searchMatrix([][]int{
		{1}}, 1))
	fmt.Println(searchMatrix([][]int{
		{1}, {3}}, 3))
}
