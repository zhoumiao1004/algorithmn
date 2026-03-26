package main

import (
	"fmt"
	"math"
)

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
	// fmt.Println(dp)
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

// 416. 分割等和子集
// https://leetcode.cn/problems/partition-equal-subset-sum/description/
// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
// 输入：nums = [1,5,11,5] 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
// 01背包的应用：问能不能分成相等的两部份。转换为背包问题：任意选择物品能否正好装满大小为s/2的背包
func canPartition(nums []int) bool {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	if s%2 == 1 {
		return false
	}
	target := s / 2
	// 转换为01背包问题，大小为target的背包尽量装，能否装价值为target的物品
	// dp[j]含义：大小为j的背包，能装的最大价值
	// 递推公式：dp[j] = max(dp[j], dp[j-nums[i]] + nums[i])
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ { // 物品
		for j := target; j >= nums[i]; j-- { // 背包逆序
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}

// 1049.最后一块石头的重量II
// https://leetcode.cn/problems/last-stone-weight-ii/description/
// 输入：stones = [2,7,4,1,8,1] 输出：1
// 解释：
// 组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
// 组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
// 组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
// 组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
// 01背包的应用：分成两部分，尽量分成近似相等的两块。转换为背包问题：大小为s/2的背包，最多能装的物品重量
func lastStoneWeightII(stones []int) int {
	s := 0
	for i := 0; i < len(stones); i++ {
		s += stones[i]
	}
	target := s / 2
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ { // 物品
		for j := target; j >= stones[i]; j-- { // 背包逆序
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	// 2 * dp[target] + x = s
	return s - 2*dp[target]
}

// 494.目标和
// https://leetcode.cn/problems/target-sum/description/
// 给你一个非负整数数组 nums 和一个整数 target 。
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
// 输入：nums = [1,1,1,1,1], target = 3 => 输出：5
// 解释：一共有 5 种方法让最终目标和为 3 。
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3
// 01背包的应用：分成两部分，需要保证两部分差值为target。转换为背包问题：装满大小为(s+target)/2的背包，有几种装法
func findTargetSumWays(nums []int, target int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	if s+target < 0 || (s+target)%2 == 1 {
		return 0
	}
	m := (s + target) / 2
	// dp[j]含义：装满容量为j的背包有几种方法
	dp := make([]int, m+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := m; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[m]
}

// 474. 一和零
// https://leetcode.cn/problems/ones-and-zeroes/description/
// 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
// 请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
// 输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
// 输出：4
// 解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
// 其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3
func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j] 含义：i个0和j个1的容器，能装的字符串最大个数
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs { // 遍历物品，每个物品有2个维度
		zeroNum := 0
		oneNum := 0
		for _, c := range s {
			if c == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}
		// 逆序遍历背包
		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}

	return dp[m][n]
}

/* 完全背包
求组合数：先遍历物品，再遍历背包
求排列数：先遍历背包再遍历物品
*/
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
func change(amount int, coins []int) int {
	// dp[i]含义：总金额为i的总方法数
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1
	// 遍历顺序：不强调顺序，求的是组合数，所以先遍历物品再遍历背包
	for i := 0; i < n; i++ { // 物品
		for j := coins[i]; j <= amount; j++ { // 背包
			// 递推公式：假设不用第i个硬币的组合数是dp[j-coins[i]],所以用上第i个硬币的组合数也是dp[j-coins[i]]
			dp[j] += dp[j-coins[i]]
		}
	}
	// fmt.Println(dp)
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
// 注意：求的是排列数，所以遍历顺序是先遍历背包，再遍历物品
func combinationSum4(nums []int, target int) int {
	// dp[j]含义：组成总和为j有n种排列
	dp := make([]int, target+1)
	dp[0] = 1
	// 求排列，先遍历背包，再遍物品
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

// 322. 零钱兑换
// https://leetcode.cn/problems/coin-change/description/
// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
// 输入：coins = [1, 2, 5], amount = 11 输出：3
// 解释：11 = 5 + 5 + 1
func coinChange(coins []int, amount int) int {
	// dp[j]含义：组成总金额为j的硬币数最少为dp[j]
	// 求组合中最少硬币个数，递推公式：dp[j] = min(dp[j], dp[j-coins[i]]+1)
	dp := make([]int, amount+1)
	dp[0] = 0
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
	fmt.Println(dp)
	return dp[amount]
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
	// dp[j]含义：组成和为j的，需要dp[j]个完全平方数
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt
	}
	// 递推公式：dp[j] = min(dp[j], dp[j-i*i]+1)
	// 遍历顺序：无所谓顺序，先遍历物品再遍历背包
	for i := 1; i*i <= n; i++ { // 物品
		for j := i * i; j <= n; j++ { // 背包
			if dp[j-i*i] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}
	}
	// fmt.Println(dp)
	return dp[n]
}

// 每次可以爬 1 、 2、.....、m 个台阶。问有多少种不同的方法可以爬到楼顶呢？
// 转换为完全背包问题：装满大小为n的背包。可以装1/2/3/4...m,有几种方式
// 递推公式：dp[j] += dp[j-i]
func climbStairsN(n, m int) int {
	// dp[j]含义：爬到j个台阶的方法数
	dp := make([]int, n+1)
	dp[0] = 1
	// 求的是排列数，遍历顺序：先遍历背包，再遍历物品
	for j := 0; j <= n; j++ { // 背包
		for i := 1; i <= m; i++ { // 物品
			if j >= i {
				dp[j] += dp[j-i]
			}
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(numSquares(13))

	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
}
