package main

import "math"

// 509. 斐波那契数
// https://leetcode.cn/problems/fibonacci-number/description/
// 思路1: 暴力递归解法
func fib(n int) int {
	memo := make([]int, n+1)
	var dp func(n int) int
	dp = func(n int) int {
		if n < 2 {
			return n
		}
		if memo[n] != 0 {
			return memo[n]
		}
		memo[n] = dp(n-1) + dp(n-2)
		return memo[n]
	}
	return dp(n)
}

func fib2(n int) int {
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

// 322. 零钱兑换
// https://leetcode.cn/problems/coin-change/description/
// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
// 输入：coins = [1, 2, 5], amount = 11 输出：3
// 解释：11 = 5 + 5 + 1
// 思路1: 暴力递归解法，时间复杂度 = 递归次数O(k^N) * 单次递归O(k)
// 状态：目标金额 amount
// 选择：coins 数组中列出的所有硬币面额
func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = -2 // 初始化一个不会取到的值，代表还未计算。-1代表凑不出
	}
	var dp func(coins []int, amount int) int // dp定义：凑出总金额 amount 至少需要的硬币数

	dp = func(coins []int, amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 1 {
			return -1 // 凑不出，无解
		}
		if memo[amount] != -2 {
			return memo[amount] // 已经计算过
		}
		res := math.MaxInt
		for _, coin := range coins {
			subProblem := dp(coins, amount-coin) // 计算子问题的结果
			if subProblem == -1 {
				continue // 子问题无解则跳过
			}
			res = min(res, subProblem+1)
		}
		if res == math.MaxInt {
			memo[amount] = -1
		} else {
			memo[amount] = res
		}
		return memo[amount]
	}

	return dp(coins, amount)
}

// 思路2：1维dp最推荐的写法，初始化为amount+1, 处理起来最简单，没有MaxInt溢出的问题，还能直接min运算
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0 // base case

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 变换状态遍历顺序也可以：外层遍历背包，内层遍历物品
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0 // base case

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func coinChange4(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[j]含义：组成总金额为j的硬币数最少为dp[j]
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}
	// 不强调顺序，求组合数，先遍历物品，再遍历背包
	for i := 0; i < len(coins); i++ { // 物品
		for j := coins[i]; j <= amount; j++ { // 背包
			if dp[j-coins[i]] != math.MaxInt { // 条件判断
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

// 931.下降路径最小和
// https://leetcode.com/problems/minimum-falling-path-sum/
// 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
// 输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
// 输出：13
// 暴力思路：定义dp函数
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	memos := make([][]int, n)
	for i := 0; i < n; i++ {
		memos[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memos[i][j] = 10001
		}
	}

	var dp func(i, j int) int // dp函数含义：从0行下落，落到matrix[i][j]的最小路径和

	dp = func(i int, j int) int {
		n := len(matrix)
		if j < 0 || j >= n {
			return math.MaxInt
		}
		if i == 0 {
			return matrix[i][j]
		}
		if memos[i][j] != 10001 {
			return memos[i][j]
		}
		memos[i][j] = matrix[i][j] + min(
			dp(i-1, j-1),
			dp(i-1, j),
			dp(i-1, j+1),
		)
		return memos[i][j]
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		res = min(res, dp(n-1, i))
	}
	return res
}

// 自底向上迭代：dp数组
func minFallingPathSum2(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n) // dp[i][j]含义：下降路径最小和
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = matrix[0][j] // 初始化第一行
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			minVal := dp[i-1][j]
			if j > 0 {
				minVal = min(minVal, dp[i-1][j-1])
			}
			if j < n-1 {
				minVal = min(minVal, dp[i-1][j+1])
			}
			dp[i][j] = matrix[i][j] + minVal
		}
	}

	res := math.MaxInt
	for j := 0; j < n; j++ {
		res = min(res, dp[n-1][j]) // 结果在最后一行
	}
	return res
}
