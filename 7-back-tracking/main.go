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
func combine(n int, k int) [][]int {
	var results [][]int
	var path []int
	var dfs func(int, int, int)
	dfs = func(n, k, startIndex int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			results = append(results, tmp)
			return
		}
		for i := startIndex; i <= n; i++ {
			path = append(path, i)
			dfs(n, k, i+1) // 只能取一次
			path = path[:len(path)-1]
		}
	}
	dfs(n, k, 1)
	return results
}

// 剪枝优化
func combine2(n int, k int) [][]int {
	var results [][]int
	var path []int
	var dfs func(int, int, int)
	dfs = func(n, k, startIndex int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			results = append(results, tmp)
			return
		}
		// 目标是选k个数，已经选了len(path)个数，还要选k-len(path)个数，找选哪个数之后后面的数就不够了？
		for i := startIndex; i <= n; i++ {
			// 判断从i开始所有数都选上，够不够k个数，i到n有n-i+1个数
			if n-i+1+len(path) < k {
				break
			}
			path = append(path, i)
			dfs(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(n, k, 1)
	return results
}

func combine3(n int, k int) [][]int {
	var results [][]int
	var path []int
	var dfs func(int, int, int)
	dfs = func(n, k, startIndex int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			results = append(results, tmp)
			return
		}
		// 目标是选k个数，已经选了len(path)个数，还要选k-len(path)个数
		for i := startIndex; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			dfs(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(n, k, 1)
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
	var dfs func(k, n, startIndex int)
	dfs = func(k, n, startIndex int) {
		if len(path) == k && s == n {
			tmp := make([]int, len(path))
			copy(tmp, path)
			results = append(results, tmp)
			return
		}
		for i := startIndex; i <= 9 && s+i <= n; i++ {
			path = append(path, i)
			s += i
			dfs(k, n, i+1) // 每个数字只能用一次，所以i+1
			s -= i
			path = path[:len(path)-1]
		}
	}
	dfs(k, n, 1)
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
	// start记录digits中第几个数字
	var dfs func(digits string, startIndex int)
	dfs = func(digits string, startIndex int) {
		if len(path) == len(digits) {
			results = append(results, string(path))
			return
		}
		n := digits[startIndex] - '0' // 组合，所以取完前面的数字后，只能再取后面的数字
		s := arr[n]
		for i := 0; i < len(s); i++ {
			path = append(path, s[i])
			dfs(digits, startIndex+1) // 只能用一次
			path = path[:len(path)-1]
		}
	}
	dfs(digits, 0)
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
// 方法1：排序
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 排序
	var results [][]int
	var path []int
	s := 0
	var dfs func([]int, int, int)
	dfs = func(candidates []int, target int, startIndex int) {
		if s == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			results = append(results, tmp)
		}
		for i := startIndex; i < len(candidates) && s+candidates[i] <= target; i++ {
			path = append(path, candidates[i])
			s += candidates[i]
			dfs(candidates, target, i) // 可以用多次
			s -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(candidates, target, 0)
	return results
}

// 方法2：不排序
func combinationSumWithoutSort(candidates []int, target int) [][]int {
	var results [][]int
	var path []int
	s := 0
	var dfs func([]int, int, int)
	dfs = func(candidates []int, target int, startIndex int) {
		// 因为题目中说所有元素都是正整数才能这么剪枝
		if s > target {
			return
		}
		if s == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			results = append(results, tmp)
		}
		for i := startIndex; i < len(candidates); i++ {
			path = append(path, candidates[i])
			s += candidates[i]
			dfs(candidates, target, i) // 可以用多次
			s -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(candidates, target, 0)
	return results
}

// 40. 组合总和 II
// https://leetcode.cn/problems/combination-sum-ii/description/
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
	var dfs func(candidates []int, target int, startIndex int)
	dfs = func(candidates []int, target, startIndex int) {
		if s == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			results = append(results, tmp)
		}
		// 因为题目说所有元素是正整数所以可以剪枝
		for i := startIndex; i < len(candidates) && s+candidates[i] <= target; i++ {
			if i > 0 && candidates[i-1] == candidates[i] && !used[i-1] {
				continue // 树层去重，上个数没用过说明重复，处于同一层相同的两个数
			}
			path = append(path, candidates[i])
			s += candidates[i]
			used[i] = true               // 标记上一个数用过了
			dfs(candidates, target, i+1) // 只能用一次
			used[i] = false
			s -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(candidates, target, 0)
	return results
}

// 131.分割回文串
// https://leetcode.cn/problems/palindrome-partitioning/description/
// 输入：s = "aab" 输出：[["a","a","b"],["aa","b"]]
func partition(s string) [][]string {
	var results [][]string
	var path []string
	var dfs func(s string, start int)
	dfs = func(s string, start int) {
		if start == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			results = append(results, tmp)
			return
		}
		for i := start; i < len(s); i++ {
			if isPalindrome(s[start : i+1]) {
				path = append(path, s[start:i+1])
				dfs(s, i+1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(s, 0)
	return results
}

func isPalindrome(s string) bool {
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

// 93.复原IP地址
// https://leetcode.cn/problems/restore-ip-addresses/description/
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 转换为3个.放在哪几个位置，能放[1,len(s)-1]
func restoreIpAddresses(s string) []string {
	var results []string
	var path []string
	var dfs func(s string, startIndex int)
	dfs = func(s string, startIndex int) {
		if startIndex == len(s) && len(path) == 4 {
			results = append(results, strings.Join(path, "."))
			return
		}
		for i := startIndex; i < len(s); i++ {
			ip := s[startIndex : i+1]
			if isValidIp(ip) {
				path = append(path, ip)
				dfs(s, i+1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(s, 0)
	return results
}

func isValidIp(s string) bool {
	// 不能前导0
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	// 0-255之间
	n, _ := strconv.Atoi(s)
	return n <= 255
}

// 78. 子集
// https://leetcode.cn/problems/subsets/description/
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 可以取0-3个
func subsets(nums []int) [][]int {
	var results [][]int
	var path []int
	var dfs func(nums []int, start int)
	dfs = func(nums []int, start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		results = append(results, tmp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, 0)
	return results
}

// 90.子集II
// https://leetcode.cn/problems/subsets-ii/
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
func subsetsWithDup(nums []int) [][]int {
	var results [][]int
	var path []int
	sort.Ints(nums)
	used := make([]bool, len(nums))
	var dfs func(nums []int, startIndex int)
	dfs = func(nums []int, startIndes int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		results = append(results, tmp)
		for i := startIndes; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue // 树层去重
			}
			path = append(path, nums[i])
			used[i] = true
			dfs(nums, i+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(nums, 0)
	return results
}

// 491. 非递减子序列
// https://leetcode.cn/problems/non-decreasing-subsequences/description/
// 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
// 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
// 输入：nums = [4,6,7,7]
// 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
// 注：不能排序
func findSubsequences(nums []int) [][]int {
	var results [][]int
	var path []int
	var dfs func(nums []int, startIndex int)
	dfs = func(nums []int, startIndex int) {
		if len(path) > 1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			results = append(results, tmp) // 注意这里没有return，因为找到一个子序列，还能继续往树的下一层找更长的子序列
		}
		uset := make(map[int]bool) // 同层去重
		for i := startIndex; i < len(nums); i++ {
			if uset[nums[i]] {
				continue
			}
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}
			path = append(path, nums[i])
			uset[nums[i]] = true
			dfs(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, 0)
	return results
}

// 46.全排列
// https://leetcode.cn/problems/permutations/description/
func permute(nums []int) [][]int {
	var results [][]int
	var path []int
	used := make([]bool, len(nums))
	var dfs func(nums []int)
	dfs = func(nums []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			results = append(results, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(nums)
	return results
}

// 47. 全排列 II
// https://leetcode.cn/problems/permutations-ii/description/
// 输入：nums = [1,1,2]
// 输出：[[1,1,2],[1,2,1],[2,1,1]]
func permuteUnique(nums []int) [][]int {
	var results [][]int
	var path []int
	sort.Ints(nums)
	used := make([]bool, len(nums))
	var dfs func(nums []int)
	dfs = func(nums []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			results = append(results, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue // 树层去重复
			}
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(nums)
	return results
}

// 332. 重新安排行程
// https://leetcode.cn/problems/reconstruct-itinerary/
func findItinerary(tickets [][]string) []string {
	var results []string
	return results
}

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
	var dfs func(chessboard [][]byte, n int, row int)
	dfs = func(chessboard [][]byte, n int, row int) {
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
				dfs(chessboard, n, row+1)
				chessboard[row][i] = '.'
			}
		}
	}
	dfs(chessboard, n, 0)
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
	fmt.Println(combine(4, 2))
	fmt.Println(letterCombinations("23"))

	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(isPalindrome("a"))
	fmt.Println(partition("aab"))
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	// ans := solveNQueens(4)
	// for i := 0; i < len(ans); i++ {
	// 	for j := 0; j < len(ans[0]); j++ {
	// 		fmt.Printf("ans[%d][%d]=%s\n", i, j, ans[i][j])
	// 	}
	// }
}
