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
// 思路：二维dp数组，明确状态：背包容量和可选择的物品；选择：装进背包 or 不装进背包
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

// 一维dp数组状态压缩
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
// 思路：二维dp数组，明确状态：背包容量和可选择的物品；选择：装进背包 or 不装进背包
func combinationSum4(nums []int, target int) int {
	n := len(nums)
	dp := make([][]int, n+1) // dp[i][j]含义：使用前 i 个数，组成总和为 j 有n种排列
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]] // 不装第 i 个数的排列数 + 装第 i 个数的排列数
			}
		}
	}
	return dp[n][target]
}

// 状态压缩写法：1维dp数组，注意：求的是排列数，所以遍历顺序是先遍历背包，再遍历物品
func combinationSum42(nums []int, target int) int {
	dp := make([]int, target+1) // dp[j]含义：组成总和为j有n种排列
	dp[0] = 1
	for j := 0; j <= target; j++ { // 背包
		for i := 0; i < len(nums); i++ { // 物品
			if j >= nums[i] {
				// 递推公式 dp[j] += dp[j-nums[i]]
				dp[j] += dp[j-nums[i]]
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
func numSquares(n int) int {
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
