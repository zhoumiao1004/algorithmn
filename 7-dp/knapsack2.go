package main

import (
	"fmt"
)

// 完全背包，求背包容量为 W 能装的最大价值。明确状态：背包容量，选择：可选择的物品
// https://labuladong.online/zh/problem/core/knapsack-unbounded/description/
func unboundedKnapsack(W int, wt []int, val []int) int {
	N := len(wt)
	dp := make([][]int, N+1) // 状态：背包容量和选择的物品；选择：装进背包 or 不装进背包
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if j < wt[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-wt[i-1]]+val[i-1]) // 和0-1背包唯一区别
			}
		}
	}
	return dp[N][W]
}

// 压缩成一维数组
func unboundedKnapsack2(W int, wt []int, val []int) int {
	dp := make([]int, W+1) // dp[j]含义：大小为j的背包能装的最大价值
	for j := 0; j <= W; j++ {
		for i := 0; i < len(wt); i++ {
			if j >= wt[i] {
				dp[j] = max(dp[j], dp[j-wt[i]]+val[i])
			}
		}
	}
	return dp[W]
}

// 完全背包-组合数
// https://labuladong.online/zh/problem/core/knapsack-unbounded-count/description/
func unboundedKnapsackCount(W int, wt []int) int64 {
	N := len(wt)
	dp := make([][]int64, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int64, W+1)
		dp[i][0] = 1
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if j < wt[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = (dp[i-1][j] + dp[i][j-wt[i-1]]) % (1e9 + 7)
			}
		}
	}
	return dp[N][W]
}

// 完全背包问题：能否装满
func unboundedKnapsackExist(W int, wt []int) bool {
	N := len(wt)
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if j < wt[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-wt[i-1]]+wt[i-1])
			}
		}
	}
	return dp[N][W] == W
}

// 518.零钱兑换II (求的是组合数)
// https://leetcode.cn/problems/coin-change-ii/
// 给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
// 输入: amount = 5, coins = [1, 2, 5] 输出: 4
// 有四种方式可以凑成总金额:
// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1
// 思路1：二维dp数组，明确状态：背包容量和可选择的物品；选择：装进背包 or 不装进背包
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1) // dp[i][j]定义：使用 coins 中的前 i 个（i 从 1 开始计数）硬币，若想凑出金额 j，有 dp[i][j] 种凑法。
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j] // 装不下，只能不把第i个物品放入背包
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]] // 不把第i个物品放入背包 or 把第i个物品放入背包
			}
		}
	}
	return dp[n][amount]
}

// 思路2: 一维dp数组状态压缩
func change2(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1) // dp[i]含义：总金额为i的组合数
	dp[0] = 1
	for i := 0; i < n; i++ { // 物品
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]] // 假设不用第i个硬币的组合数是dp[j-coins[i]],所以用上第i个硬币的组合数也是dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// 每次可以爬 1 、 2、.....、m 个台阶。问有多少种不同的方法可以爬到楼顶呢？
// 转换为完全背包问题：装满大小为n的背包。可以装1/2/3/4...m,有几种排列方式
// 递推公式：dp[j] += dp[j-i]
func climbStairsN(n, m int) int {
	dp := make([]int, n+1) // dp[j]含义：爬到j个台阶的方法数
	dp[0] = 1              // base case
	for j := 0; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if j >= i {
				dp[j] += dp[j-i]
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(unboundedKnapsack(4, []int{1, 3, 4}, []int{15, 20, 30})) // 15*4 = 60
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(numSquares(13))
}
