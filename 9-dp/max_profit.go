package main

import "fmt"

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

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
