package main

import (
	"fmt"
	"math"
)

// 301. 删除无效的括号
// https://leetcode.cn/problems/remove-invalid-parentheses/
// 给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
// 返回所有可能的结果。答案可以按 任意顺序 返回。
// 输入：s = "()())()"
// 输出：["(())()","()()()"]
func removeInvalidParentheses(s string) []string {

	var res []string
	var path string
	var isValid func(s string) bool
	var backtrack func(start int)

	isValid = func(s string) bool {
		cnt := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				cnt++
			} else if s[i] == ')' {
				cnt--
			}
			if cnt < 0 {
				return false
			}
		}
		return cnt == 0
	}
	/* var isValid2 func(s string) bool
	isValid2 = func(s string) bool {
		var st []byte
		for i := 0; i < len(s); i++ {
			if s[i] != '(' && s[i] != ')' {
				continue
			}
			if s[i] == '(' {
				st = append(st, s[i])
			} else {
				if len(st) == 0 || st[len(st)-1] != '(' {
					return false
				}
				st = st[:len(st)-1]
			}
		}
		return len(st) == 0
	}*/

	backtrack = func(start int) {
		if start == len(s) {
			if isValid(path) {
				res = append(res, path)
			}
			return
		}
		c := s[start]
		if c != '(' && c != ')' {
			// 非括号，英文字符，直接加入
			path += string(c)
			backtrack(start + 1)
			path = path[:len(path)-1]
		} else {
			// 情况1: 加入path，即不删除括号
			path += string(c)
			backtrack(start + 1)
			path = path[:len(path)-1]
			// 情况2: 不加入path，即删除括号
			backtrack(start + 1)
		}
	}

	backtrack(0)
	// 筛选出最长的有效括号字符串
	set := make(map[string]bool)
	maxLen := 0
	for _, s := range res {
		maxLen = max(maxLen, len(s))
		set[s] = true
	}

	var finalRes []string
	for s := range set {
		if len(s) == maxLen {
			finalRes = append(finalRes, s)
		}
	}
	return finalRes
}

// 2850. 将石头分散到网格图的最少移动次数
// https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid/description/
// 给你一个大小为 3 * 3 ，下标从 0 开始的二维整数矩阵 grid ，分别表示每一个格子里石头的数目。网格图中总共恰好有 9 个石头，一个格子里可能会有 多个 石头。
// 每一次操作中，你可以将一个石头从它当前所在格子移动到一个至少有一条公共边的相邻格子。
// 请你返回每个格子恰好有一个石头的 最少移动次数 。
// 输入：grid = [[1,1,0],[1,1,1],[1,2,1]]
// 输出：3
// 解释：让每个格子都有一个石头的一个操作序列为：
// 1 - 将一个石头从格子 (2,1) 移动到 (2,2) 。
// 2 - 将一个石头从格子 (2,2) 移动到 (1,2) 。
// 3 - 将一个石头从格子 (1,2) 移动到 (0,2) 。
// 总共需要 3 次操作让每个格子都有一个石头。
// 让每个格子都有一个石头的最少操作次数为 3 。
func minimumMoves(grid [][]int) int {
	minMove := math.MaxInt
	move := 0
	emptyCnt := 0
	redundant := make([][2]int, 0)
	empty := make([][2]int, 0)
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 1 {
				redundant = append(redundant, [2]int{i, j})
			} else if grid[i][j] == 0 {
				empty = append(empty, [2]int{i, j})
				emptyCnt++
			}
		}
	}

	var backtrack func()
	backtrack = func() {
		if emptyCnt == 0 {
			minMove = min(minMove, move)
			return
		}
		for _, r := range redundant {
			x1, y1 := r[0], r[1]
			if grid[x1][y1] == 1 {
				continue
			}
			for _, e := range empty {
				x2, y2 := e[0], e[1]
				if grid[x2][y2] != 0 {
					continue
				}
				step := int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
				grid[x1][y1]--
				grid[x2][y2]++
				move += step
				emptyCnt--
				backtrack()
				emptyCnt++
				move -= step
				grid[x2][y2]--
				grid[x1][y1]++
			}
		}
	}

	backtrack()
	return minMove
}

// 1723. 完成所有工作的最短时间
// https://leetcode.cn/problems/find-minimum-time-to-finish-all-jobs/
// 给你一个整数数组 jobs ，其中 jobs[i] 是完成第 i 项工作要花费的时间。
// 请你将这些工作分配给 k 位工人。所有工作都应该分配给工人，且每项工作只能分配给一位工人。工人的 工作时间 是完成分配给他们的所有工作花费时间的总和。请你设计一套最佳的工作分配方案，使工人的 最大工作时间 得以 最小化 。
// 返回分配方案中尽可能 最小 的 最大工作时间 。
// 输入：jobs = [1,2,4,7,8], k = 2
// 输出：11
// 解释：按下述方式分配工作：
// 1 号工人：1、2、8（工作时间 = 1 + 2 + 8 = 11）
// 2 号工人：4、7（工作时间 = 4 + 7 = 11）
// 最大工作时间是 11 。
func minimumTimeRequired(jobs []int, k int) int {
	workers := make([]int, k)
	res := math.MaxInt
	var backtrack func(index int)

	backtrack = func(index int) {
		if index == len(jobs) {
			// 找到一个分配方案，计算该方案的最短时间
			maxVal := 0
			for _, val := range workers {
				maxVal = max(maxVal, val)
			}
			res = min(res, maxVal)
			return
		}
		chosen := make(map[int]bool)
		// jobs[i] 可以选择 [0..k-1]
		for i := 0; i < k; i++ {
			if workers[i]+jobs[index] >= res {
				continue // 剪枝优化：如果当前工人的工作时间加上当前的工作时间已经超过了当前的最优解，那么就不用继续尝试了
			}
			if chosen[workers[i]] {
				continue // 剪枝优化：如果前面曾有工人有这个 workload，则不必把当前工作分配给他
			}
			chosen[workers[i]] = true
			workers[i] += jobs[index]
			backtrack(index + 1)
			workers[i] -= jobs[index]
		}
	}

	backtrack(0)
	return res
}

// 2305. 公平分发饼干
// https://leetcode.cn/problems/fair-distribution-of-cookies/description/
// 给你一个整数数组 cookies ，其中 cookies[i] 表示在第 i 个零食包中的饼干数量。另给你一个整数 k 表示等待分发零食包的孩子数量，所有 零食包都需要分发。在同一个零食包中的所有饼干都必须分发给同一个孩子，不能分开。
// 分发的 不公平程度 定义为单个孩子在分发过程中能够获得饼干的最大总数。
// 返回所有分发的最小不公平程度。
// 输入：cookies = [8,15,10,20,8], k = 2
// 输出：31
// 解释：一种最优方案是 [8,15,8] 和 [10,20] 。
// - 第 1 个孩子分到 [8,15,8] ，总计 8 + 15 + 8 = 31 块饼干。
// - 第 2 个孩子分到 [10,20] ，总计 10 + 20 = 30 块饼干。
// 分发的不公平程度为 max(31,30) = 31 。
// 可以证明不存在不公平程度小于 31 的分发方案。
func distributeCookies(cookies []int, k int) int {
	return minimumTimeRequired(cookies, k)
}

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
}
