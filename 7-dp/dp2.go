package main

import (
	"math"
	"sort"
)

// 97. 交错字符串
// https://leetcode.cn/problems/interleaving-string/
// 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
// 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
// s = s1 + s2 + ... + sn
// t = t1 + t2 + ... + tm
// |n - m| <= 1
// 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
// 注意：a + b 意味着字符串 a 和 b 连接。
// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// 输出：true
// 思路1: 自顶向下递归解法。状态：i, j, dp(i, j)定义：s1[i...] 和 s2[j...] 能否凑出 s3[i+j...]
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	memo := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(i, j int) bool // s1[i...] 和 s2[j...] 能否凑出 s3[i+j...]
	dp = func(i, j int) bool {
		k := i + j
		if k == len(s3) {
			return true
		}
		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}

		res := false
		// 尝试s1[i]匹配s3
		if i < len(s1) && s1[i] == s3[k] {
			res = dp(i+1, j)
		}
		if j < len(s2) && s2[j] == s3[k] {
			res = res || dp(i, j+1)
		}
		if res {
			memo[i][j] = 1
		} else {
			memo[i][j] = 0
		}
		return res
	}

	return dp(0, 0)
}

// 从后往前推
func isInterleave2(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	memo := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if i < 0 || j < 0 {
			return false
		}
		k := i + j
		if k == 0 {
			return true
		}
		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}
		res := false
		if i > 0 && s1[i-1] == s3[k-1] {
			res = res || dp(i-1, j)
		}
		if j > 0 && s2[j-1] == s3[k-1] {
			res = res || dp(i, j-1)
		}
		if res {
			memo[i][j] = 1
		} else {
			memo[i][j] = 0
		}
		return res
	}

	return dp(m, n)
}

// 思路2: 自底向上迭代解法。状态：i, j, dp[i][j]定义：s1[...i] 和 s2[...j] 能否凑出 s3[...i+j]
func isInterleave3(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			k := i + j
			if s1[i-1] == s3[k-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if s2[j-1] == s3[k-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

// 152.乘积最大子数组
// https://leetcode.cn/problems/maximum-product-subarray/solutions/
// 输入: nums = [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n) // dp[i]含义：以i结尾的nums子数组的最大乘积
	dp[0][0] = nums[0]      // 最小乘积
	dp[0][1] = nums[0]      // 最大乘积
	res := nums[0]
	for i := 1; i < n; i++ {
		a, b := dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]
		dp[i][0] = min(min(a, b), nums[i])
		dp[i][1] = max(max(a, b), nums[i])
		res = max(res, dp[i][1])
	}
	return res
}

// 方法2:贪心
func maxProduct2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	result := nums[0]
	preMin, preMax := nums[0], nums[0]
	for i := 1; i < n; i++ {
		a := preMin * nums[i]
		b := preMax * nums[i]
		preMin = min(nums[i], min(a, b))
		preMax = max(nums[i], max(a, b))
		result = max(result, preMax)
	}
	return result
}

// 思路3: 前缀积prefix, 但不推荐，0处理起来很麻烦
func maxProduct3(nums []int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	prefix[0] = 1
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
		if prefix[i] == 0 {
			prefix[i] = 1
		}
	}

	res := math.MinInt
	// 按0分段计算
	start := 0
	for i := 0; i <= n; i++ {
		if i == n || nums[i] == 0 {
			// 处理 [start, i-1] 段
			for j := start; j < i; j++ {
				for k := j; k < i; k++ {
					res = max(res, prefix[k+1]/prefix[j])
				}
			}
			if i < n {
				res = max(res, 0) // 0本身是一个候选
			}
			start = i + 1
		}
	}
	return res
}

// 221. 最大正方形
// https://leetcode.cn/problems/maximal-square/description/
// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
// 输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m) // 定义：以 matrix[i][j] 为右下角元素的全为 1 正方形矩阵的最大边长为 dp[i][j]。
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	res := 0
	for i := 0; i < m; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
		res = max(res, dp[i][0])
	}
	for j := 0; j < n; j++ {
		dp[0][j] = int(matrix[0][j] - '0')
		res = max(res, dp[0][j])
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				continue // 不可能为正方形右下角
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			res = max(res, dp[i][j])
		}
	}
	return res * res
}

func maximalSquare2(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1) // 定义：以 matrix[i][j] 为右下角元素的全为 1 正方形矩阵的最大边长为 dp[i][j]。
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '0' {
				continue // 不可能为正方形右下角
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			res = max(res, dp[i][j])
		}
	}
	return res * res
}

// 329. 矩阵中的最长递增路径
// https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/description/
// 给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
// 对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
// 输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
// 输出：4
// 解释：最长递增路径为 [1, 2, 6, 9]。
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	res := 0
	var dp func(i, j int) int

	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}

	dp = func(i, j int) int {
		if memo[i][j] != 0 {
			return memo[i][j] // 不为0说明已经计算过
		}
		cnt := 1
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			if matrix[i][j] < matrix[x][y] {
				cnt = max(cnt, dp(x, y)+1)
			}
		}
		memo[i][j] = cnt
		return cnt
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dp(i, j))
		}
	}
	return res
}

// 1235. 规划兼职工作
// https://leetcode.cn/problems/maximum-profit-in-job-scheduling/description/
// 你打算利用空闲时间来做兼职工作赚些零花钱。
// 这里有 n 份兼职工作，每份工作预计从 startTime[i] 开始到 endTime[i] 结束，报酬为 profit[i]。
// 给你一份兼职工作表，包含开始时间 startTime，结束时间 endTime 和预计报酬 profit 三个数组，请你计算并返回可以获得的最大报酬。
// 注意，时间上出现重叠的 2 份工作不能同时进行。
// 如果你选择的工作在时间 X 结束，那么你可以立刻进行在时间 X 开始的下一份工作。
// 输入：startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
// 输出：120
// 解释：
// 我们选出第 1 份和第 4 份工作，
// 时间范围是 [1-3]+[3-6]，共获得报酬 120 = 50 + 70。
// 思路：0-1背包问题的变型，按终点排序，dp[i]定义：0-i时间区间内，最多能获得的利润是dp[i]
// 对于每个工作，有2个选择：选或者不选
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	var rgithBound func(keys []int, target int) int
	rgithBound = func(keys []int, target int) int {
		left, right := 0, len(keys)-1
		for left < right {
			mid := (left + right + 1) / 2
			if keys[mid] <= target {
				left = mid
			} else {
				right = mid - 1
			}
		}
		return left
	}

	n := len(profit)
	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	// 按结束时间排序，排序后，当处理到第 i 份工作时，所有可能与它冲突的工作都在它之后，只需判断之前的选择即可。
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})

	dp := make(map[int]int) // 定义：在 0 到 i 这个时间区间内，最多能够获得的利润是 dp[i]
	keys := []int{0}
	dp[0] = 0
	for _, job := range jobs {
		begin := job[0]
		end := job[1]
		value := job[2]
		i := rgithBound(keys, begin)
		maxProfit := max(
			dp[keys[i]]+value,     // 选择这个 job，获得的利润是当前的利润加上在开始时间之前能获得的最大利润
			dp[keys[len(keys)-1]], // 不选择，保持现有的最大利润
		)
		if maxProfit > dp[keys[len(keys)-1]] {
			dp[end] = maxProfit
			keys = append(keys, end)
		}
	}

	return dp[keys[len(keys)-1]]
}
