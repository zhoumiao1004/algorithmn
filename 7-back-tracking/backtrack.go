package main

import (
	"fmt"
	"math"
)

// 967. 连续差相同的数字
// https://leetcode.cn/problems/numbers-with-same-consecutive-differences/description/
// 返回所有长度为 n 且满足其每两个连续位上的数字之间的差的绝对值为 k 的 非负整数 。
// 请注意，除了 数字 0 本身之外，答案中的每个数字都 不能 有前导零。例如，01 有一个前导零，所以是无效的；但 0 是有效的。
// 你可以按 任何顺序 返回答案。
// 输入：n = 3, k = 7
// 输出：[181,292,707,818,929]
// 解释：注意，070 不是一个有效的数字，因为它有前导零。
func numsSameConsecDiff(n int, k int) []int {
	var results []int
	var path []int
	var backtrack func(n, k int)
	backtrack = func(n, k int) {
		if len(path) == n {
			s := 0
			for i := 0; i < n; i++ {
				s = 10*s + path[i]
			}
			results = append(results, s)
			return
		}
		for i := 0; i <= 9; i++ {
			if len(path) == 0 && i == 0 {
				continue
			}
			if len(path) > 0 && int(math.Abs(float64(path[len(path)-1])-float64(i))) != k {
				continue
			}
			path = append(path, i)
			backtrack(n, k)
			path = path[:len(path)-1]
		}
	}
	backtrack(n, k)
	return results
}

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
	visited := make([][]bool, m) // 标记是否已走过
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	dirs := [][2]int{
		{0, -1},  // 上
		{0, 1},   // 下
		{-1, -0}, // 左
		{1, 0},   // 右
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
		if grid[i][j] == -1 || visited[i][j] {
			return
		}
		if grid[i][j] == 2 {
			if visitedCount == totalCount {
				result++
			}
		}
		visited[i][j] = true // TODO：能否临时修改grid[i][j]=-1，遍历完再改回来？
		visitedCount++
		for _, dir := range dirs {
			dfs(grid, i+dir[0], j+dir[1])
		}
		visited[i][j] = false // TODO
		visitedCount--
	}
	dfs(grid, startx, starty)
	return result
}

// 526. 优美的排列
// https://leetcode.cn/problems/beautiful-arrangement/description/
// 假设有从 1 到 n 的 n 个整数。用这些整数构造一个数组 perm（下标从 1 开始），只要满足下述条件 之一 ，该数组就是一个 优美的排列 ：
// perm[i] 能够被 i 整除
// i 能够被 perm[i] 整除
// 给你一个整数 n ，返回可以构造的 优美排列 的 数量 。
// 输入：n = 2
// 输出：2
// 解释：
// 第 1 个优美的排列是 [1,2]：
//   - perm[1] = 1 能被 i = 1 整除
//   - perm[2] = 2 能被 i = 2 整除
//
// 第 2 个优美的排列是 [2,1]:
//   - perm[1] = 2 能被 i = 1 整除
//   - i = 2 能被 perm[2] = 1 整除
func countArrangement(n int) int {
	return 0
}

// 89. 格雷编码
// https://leetcode.cn/problems/gray-code/description/
// n 位格雷码序列 是一个由 2n 个整数组成的序列，其中：
// 每个整数都在范围 [0, 2n - 1] 内（含 0 和 2n - 1）
// 第一个整数是 0
// 一个整数在序列中出现 不超过一次
// 每对 相邻 整数的二进制表示 恰好一位不同 ，且
// 第一个 和 最后一个 整数的二进制表示 恰好一位不同
// 给你一个整数 n ，返回任一有效的 n 位格雷码序列 。
func grayCode(n int) []int {
	return []int{}
}

// 79. 单词搜索

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCCED"
// 输出：true
func exist(board [][]byte, word string) bool {
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

// 473. 火柴拼正方形
// https://leetcode.cn/problems/matchsticks-to-square/
// 你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i 个火柴棒的长度。你要用 所有的火柴棍 拼成一个正方形。你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。
// 如果你能使这个正方形，则返回 true ，否则返回 false 。
func makesquare(matchsticks []int) bool {
	return canPartitionKSubsets(matchsticks, 4)
}

// 698. 划分为k个相等的子集
// https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/description/
// 给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
// 输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
// 输出： True
// 说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	visited := make([]bool, len(nums))
	s := 0
	var backtrack func(nums []int, k, start int) bool
	backtrack = func(nums []int, k, start int) bool {
		if k == 0 {
			return true
		}
		if s == target {
			return backtrack(nums, k-1, 0)
		}
		for i := start; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			if s+nums[i] > target { // 也可以放在for条件里
				continue
			}
			visited[i] = true
			s += nums[i]
			if backtrack(nums, k, i+1) {
				return true
			}
			s -= nums[i]
			visited[i] = false
		}
		return false
	}
	return backtrack(nums, k, 0)
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
	m, n := len(grid), len(grid[0])
	result := 0
	s := 0
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var dfs func(grid [][]int, i, j int)
	dfs = func(grid [][]int, i, j int) {
		m, n := len(grid), len(grid[0])
		dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		if visited[i][j] { // 不走回头路
			return
		}
		// 回溯算法框架：进入 (i, j)，做选择
		visited[i][j] = true
		s += grid[i][j]
		result = max(result, s)
		for _, dir := range dirs {
			dfs(grid, i+dir[0], j+dir[1])
		}
		s -= grid[i][j]
		visited[i][j] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(grid, i, j)
		}
	}
	return result
}

// 1849. 将字符串拆分为递减的连续值
// https://leetcode.cn/problems/splitting-a-string-into-descending-consecutive-values/description/
// 给你一个仅由数字组成的字符串 s 。
// 请你判断能否将 s 拆分成两个或者多个 非空子字符串 ，使子字符串的 数值 按 降序 排列，且每两个 相邻子字符串 的数值之 差 等于 1 。
// 例如，字符串 s = "0090089" 可以拆分成 ["0090", "089"] ，数值为 [90,89] 。这些数值满足按降序排列，且相邻值相差 1 ，这种拆分方法可行。
// 另一个例子中，字符串 s = "001" 可以拆分成 ["0", "01"]、["00", "1"] 或 ["0", "0", "1"] 。然而，所有这些拆分方法都不可行，因为对应数值分别是 [0,1]、[0,1] 和 [0,0,1] ，都不满足按降序排列的要求。
// 如果可以按要求拆分 s ，返回 true ；否则，返回 false 。
// 子字符串 是字符串中的一个连续字符序列。
// 输入：s = "1234"
// 输出：false
// 解释：不存在拆分 s 的可行方法。
// 输入：s = "050043"
// 输出：true
// 解释：s 可以拆分为 ["05", "004", "3"] ，对应数值为 [5,4,3] 。
// 满足按降序排列，且相邻值相差 1 。
func splitString(s string) bool {
	return false
}

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
}
