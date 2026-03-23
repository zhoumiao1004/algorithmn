package main

import (
	"fmt"
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

// 198. 打家劫舍
// https://leetcode.cn/problems/house-robber/description/
// 相邻不能偷。dp[i]两种情况：1.偷，dp[i-2]+nums[i] 2.不偷dp[i-1]
// 输入：[1,2,3,1] 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。偷窃到的最高金额 = 1 + 3 = 4 。
func rob(nums []int) int {
	// dp[i] 含义：偷到第i个房屋的最大金额
	// 2种情况：1.不偷上间房屋，最大金额=dp[i-2]+nums[i] 2.偷上间房屋，最大金额=dp[i-1]
	// 递推公式：dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

// 213.打家劫舍II
// https://leetcode.cn/problems/house-robber-ii/description/
// 围成一圈,相邻不能偷
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	return max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}

// 337.打家劫舍 III
// https://leetcode.cn/problems/house-robber-iii/
// 房间连成树，相邻不能偷
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob3(root *TreeNode) int {
	var dfs func(root *TreeNode) [2]int
	dfs = func(root *TreeNode) [2]int {
		// dp数组含义：dp[0]代表不偷本节点的最大金额，dp[1]代表偷本节点的最大金额
		var dp [2]int
		if root == nil {
			return dp
		}
		// 后序遍历
		left := dfs(root.Left)
		right := dfs(root.Right)
		// 递推公式
		// 不偷：dp[0] = max(left[0], left[1]) + max(right[0], right[1])
		// 偷：dp[1] = left[0] + right[0] + root.Val
		dp[0] = max(left[0], left[1]) + max(right[0], right[1])
		dp[1] = left[0] + right[0] + root.Val
		return dp
	}
	dp := dfs(root)
	return max(dp[0], dp[1])
}

// 121. 买卖股票的最佳时机
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
// 只能买卖一次
// 输入：[7,1,5,3,6,4] 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
// 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票
func maxProfit(prices []int) int {
	// 只有两种状态：持有股票，不持有股票
	// dp含义：dp[i][0]代表第i天不持有股票的最大利润 dp[i][1]代表第i天持有股票的最大利润
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		// 状态转移递推公式
		// 第i天不持有股票有2种情况可以得到：1.保持第i-1天不持有 2.第i-1天持有，第i天卖出
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 第i天持有股票有2种情况可以得到：1.保持第i天持有 2.第i天第一次买入
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	// fmt.Println(dp)
	return dp[n-1][0]
}

// 贪心解法
func maxProfitGreedy(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	result := 0
	low := prices[0]
	for i := 1; i < len(prices); i++ {
		result = max(result, prices[i]-low)
		if prices[i] < low {
			low = prices[i]
		}
	}
	return result
}

// 122.买卖股票的最佳时机II
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// 可以多次买卖，同时只能有一只股票
// prices = [7,1,5,3,6,4] 输出：7
func maxProfit2(prices []int) int {
	// dp[i][0]代表不持有股票的最大金额，dp[i][1]代表持有股票的最大金额
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		// 第i天不持有有2种情况可以推出：1.保持第i-1天不持有 2.第i-1天持有，第i天卖出
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 第i天持有有2种情况可以推出：1.保持第i-1天持有 2.第i-1天不持有，第i天买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

// 贪心解法：收集所有上升段
func maxProfit2Greedy(prices []int) int {
	s := 0
	for i := 1; i < len(prices); i++ {
		s += max(0, prices[i]-prices[i-1])
	}
	return s
}

// 123.买卖股票的最佳时机III
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 输入：prices = [3,3,5,0,0,3,1,4] 输出：6
// 解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
// 随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3。
func maxProfit3(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][5]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]                            // 不持有
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) // 持有第一次
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i]) // 不持有（第一次持有后卖出）
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i]) // 持有第二次
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i]) // 不持有（第二次持有后卖出）
	}
	return dp[n-1][4]
}

// 188.买卖股票的最佳时机IV
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/
// 最多可以买卖 k 次。输入：k = 2, prices = [2,4,1] 输出：2
// 解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
func maxProfit4(k int, prices []int) int {
	// dp[i][0]代表不持有股票，dp[i][j]代表第j次持有股票，dp[i][j+1]代表第j次不持有
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2*k+1)
	}
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0] // 第j次持有
	}
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		// j控制第几次买卖
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])   // 持有:保持第i-1天持有或第i-1天不持有在第i天第j次买入
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i]) // 不持有:保持第i-1不持有或第i-1天持有在第i天第j次卖出
		}
	}
	return dp[n-1][2*k]
}

// 309. 买卖股票的最佳时机含冷冻期
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// 可以多次买卖。卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 状态转移：1.持有股票 2.保持卖出 3.卖出不持有 4.处于冷冻期
func maxProfit5(prices []int) int {
	m := len(prices)
	dp := make([][4]int, m)
	dp[0][0] = -prices[0] // 持有（保持昨天持有 或 昨天未持有且未卖出今天买入）
	dp[0][1] = 0          // 不持有且非冷冻期且未卖出（保持前一天的不持有 或 前一天冷冻期）
	dp[0][2] = 0          // 不持有且卖出股票 （昨天持有今天卖出）
	dp[0][3] = 0          // 不持有且冷冻期（昨天刚卖出）
	for i := 1; i < m; i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1], dp[i-1][3])-prices[i]) // 持有 = 1.昨天持有今天保持持有 2.昨天冷冻期或不持有，今天买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])                            // 不持有 = 1.保持不持有 2.昨天冷冻期
		dp[i][2] = dp[i-1][0] + prices[i]                                 // 当天卖出 = 昨天持有今天卖出
		dp[i][3] = dp[i-1][2]                                             // 冷冻期
	}
	return max(max(dp[m-1][1], dp[m-1][2]), dp[m-1][3])
}

// 714.买卖股票的最佳时机含手续费
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
// 输入：prices = [1, 3, 2, 8, 4, 9], fee = 2 输出：8
// 总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8
func maxProfit6(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][0] = -prices[0] // 持有
	dp[0][1] = 0          // 不持有
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])     // 持有 = 1.保持昨天持有 2.昨天不持有，今天买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee) // 不持有 = 1.保持昨天不持有 2.昨天持有，今天卖出，交手续费
	}
	return dp[n-1][1]
}

// 64.最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。
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

func main() {
	fmt.Println(climbStairsN(3, 2)) // 3
	fmt.Println(climbStairs(3))     // 3

	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	// fmt.Println(bag_problem_01_2d(weight, value, 4)) //35
	fmt.Println(bag_problem_01_1d(weight, value, 4)) //35
	//fmt.Println(package22(weight, value, 4))
	//fmt.Println(package1(weight, value, 5))
	// fmt.Println(fullPackage([]int{1, 3, 4}, []int{15, 20, 30}, 4))

	fmt.Println(change(5, []int{1, 2, 5}))

	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(longestCommonSubsequence("abced", "ace")) // 3

	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{-1}
	fmt.Println(maxSubArray(nums)) // 6 连续子数组 [4,-1,2,1] 的和最大，为 6 。

	//输入：s = "abc", t = "ahbgdc"
	//输出：true
	//[0 0 0 0 0 0 0]
	//[0 1 1 1 1 1 1]
	//[0 0 0 2 2 2 2]
	//[0 0 0 0 0 0 3]
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(integerBreak(10)) // 36 = 3 * 3 * 4
	fmt.Println("------")
	fmt.Println(numTrees(3))
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(numSquares(13))
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(numDistinct("babgbag", "bag"))
	// fmt.Println(minDistance("horse", "ros"))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
