package main

import (
	"math"
)

// 509. 斐波那契数
// https://leetcode.cn/problems/fibonacci-number/description/
// 思路1: 自顶向下带备忘录的递归解法
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

// 思路2: 自底向上的迭代写法
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
// 思路1: 自顶向下递归解法，明确状态：目标金额 amount，选择：coins 数组中列出的所有硬币面额。时间复杂度 = 递归次数O(k^N) * 单次递归O(k)
func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = -2 // 初始化一个不会取到的值，代表还未计算。-1代表凑不出
	}
	var dp func(amount int) int // dp定义：凑出总金额 amount 至少需要的硬币数

	dp = func(amount int) int {
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
			subProblem := dp(amount - coin) // 计算子问题的结果
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

	return dp(amount)
}

// 思路2：自底向上迭代解法，初始化为amount+1, 处理起来最简单，没有MaxInt溢出的问题，还能直接min运算
func coinChange2(coins []int, amount int) int {
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

// 377. 组合总和 Ⅳ
// https://leetcode.cn/problems/combination-sum-iv/description/
// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
// 输入：nums = [1,2,3], target = 4 输出：7
// 所有可能的组合为：
// (1, 1, 1, 1)
// (1, 1, 2)
// (1, 2, 1)
// (1, 3)
// (2, 1, 1)
// (2, 2)
// (3, 1)
// 思路1：自顶向下的递归解法，明确状态：目标和, 选择：nums中的数
func combinationSum4(nums []int, target int) int {
	memo := make([]int, target+1)
	for i := range memo {
		memo[i] = -2 // 初始化一个不会取到的值，代表还未计算。-1代表凑不出
	}
	var dp func(amount int) int // dp定义：凑出总金额 amount 的排列数

	dp = func(amount int) int {
		if amount == 0 {
			return 1
		}
		if amount < 0 {
			return -1 // 凑不出，无解
		}
		if memo[amount] != -2 {
			return memo[amount] // 已经计算过
		}
		res := 0
		for _, val := range nums {
			subProblem := dp(amount - val) // 计算子问题的结果
			if subProblem == -1 {
				continue // 子问题无解则跳过
			}
			res += subProblem
		}
		memo[amount] = res
		return memo[amount]
	}

	return dp(target)
}

// 思路2：自底向上的迭代解法，明确状态：目标和, 选择：nums中的数
func combinationSum42(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, target+1)
	for j := 1; j <= target; j++ {
		for i := 1; i <= n; i++ {
			if j >= nums[i-1] {
				dp[j] += dp[j-nums[i-1]]
			}
		}
	}
	return dp[target]
}

// 279.完全平方数
// https://leetcode.cn/problems/perfect-squares/description/
// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
// 输入：n = 13 输出：2
// 解释：13 = 4 + 9
// 思路1: 自顶向下的递归解法，明确状态：目标和，选择：完全平方数
func numSquares(n int) int {
	var dp func(amount int) int // 定义dp函数：凑成target需要的最少完全平方数的个数
	memo := make([]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = -1
	}

	dp = func(amount int) int {
		if amount < 2 {
			return amount
		}
		if memo[amount] != -1 {
			return memo[amount]
		}
		res := amount + 1
		for i := 1; i*i <= amount; i++ {
			res = min(res, dp(amount-i*i)+1)
		}
		memo[amount] = res
		return res
	}

	return dp(n)
}

// 思路2: 自底向上的迭代解法，明确状态：目标和，选择：完全平方数
func numSquares2(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1) // dp[j]含义：组成和为j的，需要dp[j]个完全平方数
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = n + 1 // 最少的完全平方数个数范围在1-n之间，初始化为不可能的值：n+1
	}

	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n] // 因为有1这个完全平方数，所以一定有结果
}
