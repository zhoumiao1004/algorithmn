package main

// 51. N皇后
// https://leetcode.cn/problems/n-queens/description/
// 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
func solveNQueens(n int) [][]string {
	var results [][]string
	chessboard := make([][]byte, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = '.'
		}
	}
	var backtrack func(chessboard [][]byte, n int, row int)

	backtrack = func(chessboard [][]byte, n int, row int) {
		if row == n {
			tmp := make([]string, n)
			for i := 0; i < n; i++ {
				tmp[i] = string(chessboard[i])
			}
			results = append(results, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if isValid(n, row, i, chessboard) {
				chessboard[row][i] = 'Q'
				backtrack(chessboard, n, row+1)
				chessboard[row][i] = '.'
			}
		}
	}

	backtrack(chessboard, n, 0)
	return results
}

func isValid(n, row, col int, chessboard [][]byte) bool {
	// 上
	i, j := 0, 0
	for i = 0; i < row; i++ {
		if chessboard[i][col] == 'Q' {
			return false
		}
	}
	// 左上
	i, j = row-1, col-1
	for i >= 0 && j >= 0 {
		if chessboard[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	// 右上
	i, j = row-1, col+1
	for i >= 0 && j < n {
		if chessboard[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}

// 37. 解数独
// 编写一个程序，通过填充空格来解决数独问题。
// 数独的解法需 遵循如下规则：
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
// 数独部分空格内已填入了数字，空白格用 '.' 表示。
func solveSudoku(board [][]byte) {

}

func main() {
	// ans := solveNQueens(4)
	// for i := 0; i < len(ans); i++ {
	// 	for j := 0; j < len(ans[0]); j++ {
	// 		fmt.Printf("ans[%d][%d]=%s\n", i, j, ans[i][j])
	// 	}
	// }
}
