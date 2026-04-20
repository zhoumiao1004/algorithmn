package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// 131.分割回文串
// https://leetcode.cn/problems/palindrome-partitioning/description/
// 输入：s = "aab" 输出：[["a","a","b"],["aa","b"]]
func partition(s string) [][]string {
	var results [][]string
	var path []string
	var backtrack func(s string, start int)
	var isPalindrome func(s string) bool

	isPalindrome = func(s string) bool {
		left, right := 0, len(s)-1
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	backtrack = func(s string, start int) {
		if start == len(s) {
			results = append(results, append([]string{}, path...))
			return
		}
		for i := start; i < len(s); i++ {
			if !isPalindrome(s[start : i+1]) {
				continue // 剪枝：分割出的子串不是回文串
			}
			path = append(path, s[start:i+1])
			backtrack(s, i+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(s, 0)
	return results
}

// 93.复原IP地址
// https://leetcode.cn/problems/restore-ip-addresses/description/
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 转换为3个.放在哪几个位置，能放[1,len(s)-1]
func restoreIpAddresses(s string) []string {
	var results []string
	var path []string
	var isValidIp func(s string) bool
	var backtrack func(s string, startIndex int)

	isValidIp = func(s string) bool {
		// 不能前导0
		if len(s) > 1 && s[0] == '0' {
			return false
		}
		// 0-255之间
		n, _ := strconv.Atoi(s)
		return n <= 255
	}

	backtrack = func(s string, startIndex int) {
		if startIndex == len(s) && len(path) == 4 {
			results = append(results, strings.Join(path, "."))
			return
		}
		for i := startIndex; i < len(s); i++ {
			ip := s[startIndex : i+1]
			if isValidIp(ip) {
				path = append(path, ip)
				backtrack(s, i+1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(s, 0)
	return results
}

// 332. 重新安排行程
// https://leetcode.cn/problems/reconstruct-itinerary/
func findItinerary(tickets [][]string) []string {
	var results []string
	return results
}

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
// 思路：形式2: 元素有重,不可复选。桶视角选球，n个球，每个球又取或不取2种选择，k个桶，所以复杂度=k*2^n
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k

	used := make([]bool, len(nums))
	memo := make(map[string]bool)
	var backtrack func(nums []int, k, s, start int) bool
	backtrack = func(nums []int, k, s, start int) bool {
		if k == 0 {
			return true
		}
		data, _ := json.Marshal(used)
		state := string(data)
		if s == target {
			res := backtrack(nums, k-1, 0, 0)
			memo[state] = res
			return res
		}
		if _, ok := memo[state]; ok {
			return memo[state]
		}
		for i := start; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if s+nums[i] > target {
				continue
			}
			used[i] = true
			s += nums[i]
			if backtrack(nums, k, s, i+1) {
				return true
			}
			s -= nums[i]
			used[i] = false
		}
		return false
	}

	return backtrack(nums, k, 0, 0)
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

// 思路2: 元素视角, 站在元素视角选索引
func countArrangement2(n int) int {
	result := 0
	var backtrack func(n, start int)
	backtrack = func(n, start int) {

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

func main() {
	fmt.Println(countArrangement(2))
	fmt.Println(partition("aab"))
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(canPartitionKSubsets([]int{2, 2, 2, 2, 3, 4, 5}, 4))
}
