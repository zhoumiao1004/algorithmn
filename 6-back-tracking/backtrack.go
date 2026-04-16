package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 22. 括号生成
// https://leetcode.cn/problems/generate-parentheses/description/
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
func generateParenthesis(n int) []string {
	var results []string
	if n == 0 {
		return results
	}
	var path []byte
	var backtrack func(i, j int)
	backtrack = func(i, j int) {
		if i > j {
			return
		}
		if i < 0 || j < 0 {
			return
		}
		if i == 0 && j == 0 {
			results = append(results, string(path))
			return
		}

		for _, c := range []byte{'(', ')'} {
			path = append(path, c)
			if c == '(' {
				backtrack(i-1, j)
			} else {
				backtrack(i, j-1)
			}
			path = path[:len(path)-1]
		}
	}
	backtrack(n, n)
	return results
}

// 698. 划分为k个相等的子集
// https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/description/
// 给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
// 输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
// 输出： True
// 说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
// 思路：形式2: 元素有重,不可复选
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

// 473. 火柴拼正方形
// https://leetcode.cn/problems/matchsticks-to-square/
// 你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i 个火柴棒的长度。你要用 所有的火柴棍 拼成一个正方形。你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。
// 如果你能使这个正方形，则返回 true ，否则返回 false 。
func makesquare(matchsticks []int) bool {
	return canPartitionKSubsets(matchsticks, 4)
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
//
// 思路1: 索引视角, 站在索引视角选元素
func countArrangement(n int) int {
	result := 0
	used := make([]bool, n+1)
	var path []int
	var backtrack func(n, index int)
	backtrack = func(n, index int) {
		if index > n {
			result++
			return
		}

		for elem := 1; elem <= n; elem++ {
			if used[elem] {
				continue
			}
			if !(index%elem == 0 || elem%index == 0) {
				continue
			}
			// 做选择，index选elem
			used[elem] = true
			path = append(path, elem)
			backtrack(n, index+1)
			path = path[:len(path)-1]
			used[elem] = false
		}
	}

	backtrack(n, 1)
	return result
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
	used := make(map[int]bool)
	var path []int
	var result []int

	var flipBit func(x, i int) int // 把第 i 位取反（0 变 1，1 变 0）
	flipBit = func(x, i int) int {
		return x ^ (1 << i)
	}

	var traverse func(root, n int)
	traverse = func(root, n int) {
		if result != nil {
			return
		}
		if len(path) == (1 << n) {
			result = append([]int{}, path...)
			return
		}

		if _, ok := used[root]; ok {
			return
		}

		// 多叉树遍历的前序位置
		used[root] = true
		path = append(path, root)

		// 对当前数字的每个二进制位进行翻转，得到子节点
		for i := 0; i < n; i++ {
			next := flipBit(root, i)
			traverse(next, n)
		}

		// 多叉树遍历的后序位置
		delete(used, root)
		path = path[:len(path)-1]
	}

	traverse(0, n)
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
// 思路1: 站在字符的视角进行穷举
func splitString(s string) bool {
	found := false
	var path []string
	var parseInt func(s string) int64
	parseInt = func(s string) int64 {
		num, _ := strconv.ParseInt(s, 10, 64)
		return num
	}
	var backtrack func(s string, start, index int)

	backtrack = func(s string, start, index int) {
		if found {
			return
		}
		if index == len(s) {
			if len(path) >= 2 && strings.Join(path, "") == s {
				found = true
			}
			return
		}
		// 选择一，s[index] 决定切割
		subStr := s[start : index+1]
		leadingZeroCount := 0
		for j := 0; j < len(subStr); j++ {
			if subStr[j] == '0' {
				leadingZeroCount++
			} else {
				break
			}
		}
		if len(subStr)-leadingZeroCount > (len(s)+1)/2 {
			return // 剪枝逻辑，如果当前截取的子串长度大于 s 的一半，那么没必要继续截取了，肯定不可能只差一，同时可以避免溢出 long 的最大值的问题
		}

		if len(path) == 0 || parseInt(path[len(path)-1])-parseInt(subStr) == 1 {
			// 符合题目的要求，当前数字比上一个数字小 1。做选择，切割出一个子串
			path = append(path, subStr)
			backtrack(s, index+1, index+1)
			path = path[:len(path)-1]
		}

		// 选择二，s[index] 决定不切割
		backtrack(s, start, index+1)
	}

	backtrack(s, 0, 0)
	return found
}

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
}
