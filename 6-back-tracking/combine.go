package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/* 回溯: 解决组合、切割、子集、排列、棋盘
1.递归函数和返回值
2.确定终止条件
3.单层递归逻辑（横向取数，纵向递归）
*/

// 77. 组合 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 输入：n = 4, k = 2 输出：[[2,4],[3,4],[2,3],[1,2],[1,3],[1,4]]
// 求组合，每个数只能取一次。
// 情况1: 元素无重不可复选
func combine(n int, k int) [][]int {
	var results [][]int
	var path []int
	var backtrack func(int, int, int)

	backtrack = func(n, k, startIndex int) {
		if len(path) == k {
			results = append(results, append([]int{}, path...))
			return
		}
		for i := startIndex; i <= n; i++ {
			if n-i+1+len(path) < k {
				break // 剪枝优化
			}
			path = append(path, i)
			backtrack(n, k, i+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(n, k, 1)
	return results
}

// 216.组合总和III
// https://leetcode.cn/problems/combination-sum-iii/
// 只使用数字1到9,每个数字最多使用一次
// 输入: k = 3, n = 7 输出: [[1,2,4]]
// 解释: 1 + 2 + 4 = 7
func combinationSum3(k int, n int) [][]int {
	var results [][]int
	var path []int
	s := 0
	var backtrack func(k, n, startIndex int)

	backtrack = func(k, n, startIndex int) {
		if len(path) == k && s == n {
			results = append(results, append([]int{}, path...))
			return
		}
		for i := startIndex; i <= 9 && s+i <= n; i++ {
			path = append(path, i)
			s += i
			backtrack(k, n, i+1) // 每个数字只能用一次，所以i+1
			s -= i
			path = path[:len(path)-1]
		}
	}

	backtrack(k, n, 1)
	return results
}

// 17.电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
func letterCombinations(digits string) []string {
	arr := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var results []string
	var path []byte
	var backtrack func(digits string, startIndex int)

	backtrack = func(digits string, startIndex int) {
		if len(path) == len(digits) {
			results = append(results, string(path))
			return
		}
		n := digits[startIndex] - '0' // 组合，所以取完前面的数字后，只能再取后面的数字
		s := arr[n]
		for i := 0; i < len(s); i++ {
			path = append(path, s[i])
			backtrack(digits, startIndex+1) // 只能用一次
			path = path[:len(path)-1]
		}
	}

	backtrack(digits, 0)
	return results
}

// 39. 组合总和
// https://leetcode.cn/problems/combination-sum/description/
// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 1.所有数字（包括 target）都是正整数。 2.解集不能包含重复的组合。
// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
// 注意：候选值可以选多次，目标和为target
// 思路1：排序
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 排序
	var results [][]int
	var path []int
	s := 0
	var backtrack func([]int, int, int)

	backtrack = func(candidates []int, target int, startIndex int) {
		if s == target {
			results = append(results, append([]int{}, path...))
		}
		for i := startIndex; i < len(candidates) && s+candidates[i] <= target; i++ {
			path = append(path, candidates[i])
			s += candidates[i]
			backtrack(candidates, target, i) // 可以用多次
			s -= candidates[i]
			path = path[:len(path)-1]
		}
	}

	backtrack(candidates, target, 0)
	return results
}

// 思路2：不排序,利用元素都是正数来剪枝，但剪枝效率不如排序，不推荐
func combinationSumWithoutSort(candidates []int, target int) [][]int {
	var results [][]int
	var path []int
	s := 0
	var backtrack func([]int, int, int)

	backtrack = func(candidates []int, target int, startIndex int) {
		if s > target {
			return // 因为题目中说所有元素都是正整数才能这么剪枝
		}
		if s == target {
			results = append(results, append([]int{}, path...))
		}
		for i := startIndex; i < len(candidates); i++ {
			path = append(path, candidates[i])
			s += candidates[i]
			backtrack(candidates, target, i) // 可以用多次
			s -= candidates[i]
			path = path[:len(path)-1]
		}
	}

	backtrack(candidates, target, 0)
	return results
}

// 40. 组合总和 II
// https://leetcode.cn/problems/combination-sum-ii/description/
// LCR 082. 组合总和 II https://leetcode.cn/problems/4sjJUc/description/
// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用 一次 。
// 注意：解集不能包含重复的组合。
// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 1 <= candidates[i] <= 50
// 输出:[[1,1,6],[1,2,5],[1,7],[2,6]]
// 说明：所有数字（包括目标数）都是正整数。解集不能包含重复的组合。
// 注意：1.选过的不能再选 2.组合不能重复
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	// fmt.Println(candidates)
	var results [][]int
	var path []int
	used := make([]bool, len(candidates))
	s := 0
	var backtrack func(candidates []int, target int, startIndex int)

	backtrack = func(candidates []int, target, startIndex int) {
		if s == target {
			results = append(results, append([]int{}, path...))
		}
		// 因为题目说所有元素是正整数所以可以剪枝
		for i := startIndex; i < len(candidates) && s+candidates[i] <= target; i++ {
			if i > 0 && candidates[i-1] == candidates[i] && !used[i-1] {
				continue // 树层去重，上个数没用过说明重复，处于同一层相同的两个数
			}
			path = append(path, candidates[i])
			s += candidates[i]
			used[i] = true                     // 标记上一个数用过了
			backtrack(candidates, target, i+1) // 只能用一次
			used[i] = false
			s -= candidates[i]
			path = path[:len(path)-1]
		}
	}

	backtrack(candidates, target, 0)
	return results
}

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

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(letterCombinations("23"))

	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(partition("aab"))
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}
