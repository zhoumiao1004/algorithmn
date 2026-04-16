package main

// 980. 不同路径 III
// https://leetcode.cn/problems/unique-paths-iii/description/
// 在二维网格 grid 上，有 4 种类型的方格：
// 1 表示起始方格。且只有一个起始方格。
// 2 表示结束方格，且只有一个结束方格。
// 0 表示我们可以走过的空方格。
// -1 表示我们无法跨越的障碍。
// 返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目。
// 输入：[
// [1,0,0,0],
// [0,0,0,0],
// [0,0,2,-1]]
// 输出：2
// 解释：我们有以下两条路径：
// 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
// 2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
func uniquePathsIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	/* 优化: 临时修改grid[i][j]=-1，遍历完再改回来
	visited := make([][]bool, m) // 标记是否已走过
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	} */
	dirs := [][2]int{
		{0, -1},
		{0, 1},
		{-1, -0},
		{1, 0},
	}
	// 找到起始位置
	result := 0
	visitedCount, totalCount := 0, 0
	startx, starty := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				startx, starty = i, j
			}
			if grid[i][j] == 1 || grid[i][j] == 0 {
				totalCount++
			}
		}
	}
	var dfs func(grid [][]int, i, j int)
	dfs = func(grid [][]int, i, j int) {
		m, n := len(grid), len(grid[0])
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == -1 {
			return
		}
		if grid[i][j] == 2 && visitedCount == totalCount {
			result++
			return
		}
		// visited[i][j] = true
		tmp := grid[i][j]
		visitedCount++
		for _, dir := range dirs {
			dfs(grid, i+dir[0], j+dir[1])
		}
		// visited[i][j] = false // TODO
		grid[i][j] = tmp
		visitedCount--
	}

	dfs(grid, startx, starty)
	return result
}

// 79. 单词搜索
// https://leetcode.cn/problems/word-search/
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCCED"
// 输出：true
// 思路: dfs, 穷举所有起点
func exist(board [][]byte, word string) bool {
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	m, n := len(board), len(board[0])
	k := 0
	found := false

	var dfs func(board [][]byte, i, j int)
	dfs = func(board [][]byte, i, j int) {
		if found {
			return
		}
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if board[i][j] != word[k] {
			return
		}
		if k == len(word)-1 {
			found = true
			return
		}
		k++
		tmp := board[i][j]
		board[i][j] = '-'
		for _, dir := range dirs {
			dfs(board, i+dir[0], j+dir[1])
		}
		board[i][j] = tmp
		k--
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			found = false
			dfs(board, i, j)
			if found {
				return true
			}
		}
	}
	return false
}

func exist2(board [][]byte, word string) bool {
	found := false
	var dfs func(board [][]byte, i, j, p int)

	dfs = func(board [][]byte, i, j, p int) {
		if p == len(word) {
			found = true
			return
		}
		if found {
			return
		}
		m, n := len(board), len(board[0])
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if board[i][j] != word[p] {
			return
		}
		// 做选择
		original := board[i][j]
		board[i][j] = '-'       // 技巧：起到了visted数组的作用，标记已走过，不走回头路
		dfs(board, i-1, j, p+1) // 可以改写为for循环，p+1隐藏了回溯过程：可以改为全局变量进入节点时++，出节点时--
		dfs(board, i+1, j, p+1)
		dfs(board, i, j-1, p+1)
		dfs(board, i, j+1, p+1)
		// 撤销选择
		board[i][j] = original
	}

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			found = false
			dfs(board, i, j, 0)
			if found {
				return true
			}
		}
	}
	return false
}

// 1219. 黄金矿工
// https://leetcode.cn/problems/path-with-maximum-gold/description/
// 你要开发一座金矿，地质勘测学家已经探明了这座金矿中的资源分布，并用大小为 m * n 的网格 grid 进行了标注。每个单元格中的整数就表示这一单元格中的黄金数量；如果该单元格是空的，那么就是 0。
// 为了使收益最大化，矿工需要按以下规则来开采黄金：
// 每当矿工进入一个单元，就会收集该单元格中的所有黄金。
// 矿工每次可以从当前位置向上下左右四个方向走。
// 每个单元格只能被开采（进入）一次。
// 不得开采（进入）黄金数目为 0 的单元格。
// 矿工可以从网格中 任意一个 有黄金的单元格出发或者是停止。
// 输入：grid = [[0,6,0],[5,8,7],[0,9,0]]
// 输出：24
// 解释：
// [[0,6,0],
// [5,8,7],
// [0,9,0]]
// 一种收集最多黄金的路线是：9 -> 8 -> 7。
func getMaximumGold(grid [][]int) int {
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid), len(grid[0])
	result := 0
	s := 0
	var dfs func(grid [][]int, i, j int)

	dfs = func(grid [][]int, i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		s += grid[i][j]
		tmp := grid[i][j]
		grid[i][j] = 0
		result = max(result, s)
		for _, dir := range dirs {
			dfs(grid, i+dir[0], j+dir[1])
		}
		grid[i][j] = tmp
		s -= grid[i][j]
	}

	// 穷举从所有可能起点出发
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(grid, i, j)
		}
	}
	return result
}
