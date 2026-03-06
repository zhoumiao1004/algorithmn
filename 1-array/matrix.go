package main

// 48. 旋转图像
// https://leetcode.cn/problems/rotate-image/description/
// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
func rotate(matrix [][]int) {

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
