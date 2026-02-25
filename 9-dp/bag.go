package main

import "fmt"

// 物品0 1 15
// 物品1 3 20
// 物品2 4 30
// weight := []int{1, 3, 4}
// value := []int{15, 20, 30}
func test_2_wei_bag_problem1(weight, value []int, n int) int {
	// 定义dp[i][j]数组：从下表0-i的物品中取，放进j大小的背包的价值总和
	m := len(weight)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	for j := weight[0]; j <= n; j++ {
		dp[0][j] = value[0]
	}
	// 初始化，从后往前，防止越界
	//for j := n; j >= weight[0]; j-- {
	//	dp[0][j] = dp[0][j-weight[0]] + value[0]
	//}
	// 递推公式
	for i := 1; i < m; i++ {
		//正序,也可以倒序
		for j := 0; j <= n; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n]
}

func test_1_wei_bag_problem(weight, value []int, n int) int {
	// 定义 and 初始化
	m := len(weight)
	dp := make([]int, n+1)
	// 递推顺序
	for i := 0; i < m; i++ {
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		for j := n; j >= weight[i]; j-- {
			// 递推公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	fmt.Println(dp)
	return dp[n]
}

// 物品0 1 15
// 物品1 3 20
// 物品2 4 30
// weight := []int{1, 3, 4}
// value := []int{15, 20, 30}
func fullPackage(weight, value []int, n int) int {
	m := len(weight)
	// dp[j]含义：大小为j的背包能装的最大价值
	dp := make([]int, n+1)
	// 初始化
	// 遍历顺序：先遍历物品再遍历背包
	for i := 0; i < m; i++ {
		for j := weight[i]; j <= n; j++ {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	fmt.Println(dp)
	return dp[n]
}
