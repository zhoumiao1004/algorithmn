package main

import (
	"fmt"
	"math"
	"sort"
)

// 509. 斐波那契数
// https://leetcode.cn/problems/fibonacci-number/description/
func fib(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 70. 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/description/
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶?
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 746. 使用最小花费爬楼梯
// https://leetcode.cn/problems/min-cost-climbing-stairs/
// cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
// 请你计算并返回达到楼梯顶部的最低花费。
// [10,15,20] => 15
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 1 {
		return 0
	} else if n == 2 {
		return min(cost[0], cost[1])
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

// 62.不同路径
// https://leetcode.cn/problems/unique-paths/description/
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？
// 输入：m = 3, n = 7 输出：28
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	// 初始化：第一行和第一列初始为1
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 递推公式
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 63. 不同路径 II
// https://leetcode.cn/problems/unique-paths-ii/description/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始化第一列
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	// 初始化第一行
	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 递推公式
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 343. 整数拆分
// https://leetcode.cn/problems/integer-break/description/
// 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
// 输入: 10 输出: 36
// 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
func integerBreak(n int) int {
	// dp[i]含义：整数i拆分成2个数后的乘积最大值
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1)
	// dp[0] = 0
	dp[1] = 0
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			// dp[i] = max(dp[i], j*max(i-j, dp[i-j]))
			dp[i] = max(dp[i], max(j*(i-j), dp[j]*(i-j)))
		}
	}
	// fmt.Println(dp)
	return dp[n]
}

// 96.不同的二叉搜索树
// https://leetcode.cn/problems/unique-binary-search-trees/description/
// 求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
// 输入：n = 3 输出：5
func numTrees(n int) int {
	// 总共n个节点，左右子树加起来n-1个节点
	// dp[i]含义：i个节点的二叉搜索树个数
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	// fmt.Println(dp)
	return dp[n]
}

// 64.最小路径和
// https://leetcode.cn/problems/minimum-path-sum/
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12
func minPathSum(grid [][]int) int {
	// dp[i][j]含义：以下标i结尾和j结尾的grid的最小路径和
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	s := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		s += grid[i][0]
		dp[i][0] = s
	}
	s = 0
	for j := 0; j < n; j++ {
		s += grid[0][j]
		dp[0][j] = s
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}

// 1262. 可被三整除的最大和
// https://leetcode.cn/problems/greatest-sum-divisible-by-three/
// 输入：nums = [3,6,5,1,8]
// 输出：18
// 解释：选出数字 3, 6, 1 和 8，它们的和是 18（可被 3 整除的最大和）。
func maxSumDivThree(nums []int) int {
	n := len(nums)
	// dp[i][j]含义：nums[0...i]中被3整除余数为j的最大和
	dp := make([][3]int, n+1)
	dp[0][0] = 0
	dp[0][1] = math.MinInt
	dp[0][2] = math.MinInt
	for i := 1; i <= n; i++ {
		if nums[i-1]%3 == 0 { // nums[i]能被3整除，必然选
			dp[i][0] = dp[i-1][0] + nums[i-1]
			dp[i][1] = dp[i-1][1] + nums[i-1]
			dp[i][2] = dp[i-1][2] + nums[i-1]
		} else if nums[i-1]%3 == 1 { // nums[i]被3整除余数为1，nums[i]有选和不选2种情况
			dp[i][0] = max(dp[i-1][0], dp[i-1][2]+nums[i-1])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]+nums[i-1])
			dp[i][2] = max(dp[i-1][2], dp[i-1][1]+nums[i-1])
		} else if nums[i-1]%3 == 2 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+nums[i-1])
			dp[i][1] = max(dp[i-1][1], dp[i-1][2]+nums[i-1])
			dp[i][2] = max(dp[i-1][2], dp[i-1][0]+nums[i-1])
		}
	}
	return dp[n][0]
}

// 120. 三角形最小路径和
// https://leetcode.cn/problems/triangle/description/
// 给定一个三角形 triangle ，找出自顶向下的最小路径和。
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
// 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11
// 解释：如下面简图所示：
//
//	  2
//	 3 4
//	6 5 7
//
// 4 1 8 3
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
func minimumTotal(triangle [][]int) int {
	// dp[i][j]为第i行j列的最小路径和
	// 递推公式：dp[i] = triangle[i][j] + min(dp[i-1][i-1], dp[i-1][j])
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				dp[i][j] = triangle[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = triangle[i][j] + min(dp[i-1][j-1], dp[i-1][j])
			}
		}
	}
	result := math.MaxInt
	for j := 0; j < n; j++ {
		result = min(result, dp[n-1][j])
	}
	return result
}

// 368. 最大整除子集
// https://leetcode.cn/problems/largest-divisible-subset/
// 给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i], answer[j]) 都应当满足：
// answer[i] % answer[j] == 0 ，或
// answer[j] % answer[i] == 0
// 如果存在多个有效解子集，返回其中任何一个均可。
// 输入：nums = [1,2,3]
// 输出：[1,2]
// 解释：[1,3] 也会被视为正确答案。
func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	sort.Ints(nums)
	// dp[i]含义：以nums[i]这个数结尾的最长符合要求的子序列长度
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = []int{}
	}
	// base case
	dp[0] = []int{nums[0]}
	// 递推公式：nums[0...n-1]中最长的子序列再加上nums[i]
	for i := 1; i < n; i++ {
		maxLen, index := 0, -1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(dp[j]) > maxLen {
				maxLen = len(dp[j])
				index = j
			}
		}
		if index != -1 {
			dp[i] = append([]int{}, dp[index]...)
		}
		dp[i] = append(dp[i], nums[i])
	}
	result := dp[0]
	for i := 1; i < n; i++ {
		if len(result) < len(dp[i]) {
			result = dp[i]
		}
	}
	return result
}

func main() {
	fmt.Println(climbStairsN(3, 2)) // 3
	fmt.Println(climbStairs(3))     // 3
	fmt.Println(integerBreak(10))   // 36 = 3 * 3 * 4
	fmt.Println(numTrees(3))
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
