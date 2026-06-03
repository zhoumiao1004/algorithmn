package main

import (
	"fmt"
)

// 完全背包: 计算容量为 N 的背包能装的最大价值
// 2维dp数组。状态：背包容量和可选的物品; 选择：装进 or 不装进背包
func knapsackII(weight, value []int, W int) int {
	N := len(weight)
	dp := make([][]int, N+1) // dp[i][j]含义：使用前i个物品，背包容量为j能装的最大价值
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if j < weight[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-weight[i-1]]+value[i-1]) // 0-1和完全背包的区别，0-1背包：dp[i-1][j-weight[i-1]]；完全背包dp[i][j-weight[i-1]]
			}
		}
	}
	return dp[N][W]
}

// 完全背包: 1维dp数组状态压缩
func knapsackII2(weight, value []int, W int) int {
	m := len(weight)
	dp := make([]int, W+1) // dp[j]含义：大小为j的背包能装的最大价值
	for i := 0; i < m; i++ {
		for j := weight[i]; j <= W; j++ {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[W]
}

// 518.零钱兑换II
// https://leetcode.cn/problems/coin-change-ii/
// 给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
// 输入: amount = 5, coins = [1, 2, 5] 输出: 4
// 有四种方式可以凑成总金额:
// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1
// 注意：求的是组合数
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
	dp := make([]int, amount+1) // dp[i]含义：总金额为i的总方法数
	dp[0] = 1
	// 遍历顺序：不强调顺序，求的是组合数，所以先遍历物品再遍历背包
	for i := 0; i < n; i++ { // 物品
		for j := coins[i]; j <= amount; j++ { // 背包
			// 递推公式：假设不用第i个硬币的组合数是dp[j-coins[i]],所以用上第i个硬币的组合数也是dp[j-coins[i]]
			dp[j] += dp[j-coins[i]]
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
	fmt.Println(knapsackII([]int{1, 3, 4}, []int{15, 20, 30}, 4))  // 15*4 = 60
	fmt.Println(knapsackII2([]int{1, 3, 4}, []int{15, 20, 30}, 4)) // 15*4 = 60
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(numSquares(13))
}
