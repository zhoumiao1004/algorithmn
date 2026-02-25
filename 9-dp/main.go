package main

import (
	"fmt"
	"math"
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

func bag_problem_01_2d(weight, value []int, n int) int {
	// dp[i][j]含义：大小为j的背包，放前i个物品，能装的最大价值
	// 2种情况：不放或者放标号为i的物品
	// 递推公式：dp[i][j] = max(dp[i-1][j], dp[i][j-weight[i]] + value[i])
	m := len(weight)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	for j := weight[0]; j <= n; j++ {
		dp[0][j] = value[0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j <= n; j++ {
			if j-weight[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-weight[i]]+value[i])
			}
		}
	}
	// fmt.Println(dp)
	return dp[m-1][n]
}

func bag_problem_01_1d(weight, value []int, n int) int {
	// dp[j]含义：大小为j的背包能装的最大价值
	// 初始化为0
	m := len(weight)
	dp := make([]int, n+1)
	// 递推公式：dp[j] = max(dp[j], dp[j-weight[i]] + value[i])
	// 遍历顺序: 先遍历物品，再逆序遍历背包
	for i := 0; i < m; i++ { // 物品
		for j := n; j >= weight[i]; j-- { // 背包逆序
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

// 139.单词拆分
// https://leetcode.cn/problems/word-break/
// 输入: s = "leetcode", wordDict = ["leet", "code"] 输出: true
// 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
// 注意：单词放入是有顺序的，所以是排列问题，不能求组合
func wordBreak(s string, wordDict []string) bool {
	// dp[j]含义：[0,j)范围的子串，能否由字典里的单词组成
	// 用集合中的物品，装大小为j的背包
	// if dp[i] = true && [i,j]区间内的字符串在字典中 : dp[j] = true
	// 遍历顺序：求排列，先遍历背包再遍历物品
	wordMap := make(map[string]bool)
	for _, w := range wordDict {
		wordMap[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for j := 1; j <= n; j++ { // 背包
		for i := 0; i <= j; i++ { // 物品
			if dp[i] && wordMap[s[i:j]] {
				dp[j] = true
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

// 300.最长递增子序列
// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
// 输入：nums = [10,9,2,5,3,7,101,18] 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
func lengthOfLIS(nums []int) int {
	// dp[i]含义：以下标i结尾的字符串最长递增子序列的长度
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	result := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				result = max(result, dp[i])
			}
		}
	}
	return result
}

// 674. 最长连续递增序列（子数组）
// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/description/
// 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
// 输入：nums = [1,3,5,4,7] 输出：3
// 解释：最长连续递增序列是 [1,3,5], 长度为3。
// 尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
func findLengthOfLCIS(nums []int) int {
	// dp[i]含义：以下标i结尾的字符串的最长连续递增序列的长度
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	result := 0
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			result = max(result, dp[i])
		}
	}
	return result
}

// 718. 最长重复子数组
// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/description/
// 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
// 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7] 输出：3
// 解释：长度最长的公共子数组是 [3,2,1] 。
func findLength(nums1 []int, nums2 []int) int {
	// dp[i][j]含义：nums1下标以i-1结尾，nums2以j-1结尾的数组的最长重复字数组长度
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				result = max(result, dp[i][j])
			}
		}
	}
	return result
}

// 1143.最长公共子序列
// https://leetcode.cn/problems/longest-common-subsequence/description/
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度
/* text1 = "abcde", text2 = "ace"
		a	c	e
	0	0	0	0
a	0	1	1	1
b	0	1	1	1
c	0	1	2	2
d	0	1	2	2
e	0	1	2	3
*/
func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	// dp[i][j]含义：[0,i-1]的text1和[0,j-1]的text2的最长公共子序列长度
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	// fmt.Println(dp)
	return dp[m][n]
}

// 1035.不相交的线
// 在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
// 现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足：
// nums1[i] == nums2[j]
// 且绘制的直线不与任何其他连线（非水平线）相交。
// 请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
// 以这种方法绘制线条，并返回可以绘制的最大连线数。
// 输入：nums1 = [1,4,2], nums2 = [1,2,4] 输出：2
// 解释：可以画出两条不交叉的线
// 但无法画出第三条不相交的直线，因为从 nums1[1]=4 到 nums2[2]=4 的直线将与从 nums1[2]=2 到 nums2[1]=2 的直线相交。
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	// dp[i][j]含义：以下标i-1结尾的nums1和下标j-1结尾的nums2最长公共子序列长度
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// 53. 最大子数组和
// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
func maxSubArray(nums []int) int {
	// dp[i]含义：以nums[i]结尾的最大(连续)子数组的和
	// 递推公式：dp[i] = max(dp[i-1] + nums[i], nums[i])
	n := len(nums)
	result := nums[0]
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(result, dp[i])
	}
	// fmt.Println(dp)
	return result
}

// 贪心：和为负数就放弃
func maxSubArrayGreedy(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	result := math.MinInt
	s := 0
	for i := 0; i < n; i++ {
		s = max(s+nums[i], nums[i])
		result = max(result, s)
	}
	return result
}

// 152.乘积最大子数组
// 输入: nums = [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
func maxProduct(nums []int) int {
	// dp[i]含义：以i结尾的nums子数组的最大乘积
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][0] = nums[0] // 最小乘积
	dp[0][1] = nums[0] // 最大乘积
	result := nums[0]
	for i := 1; i < n; i++ {
		a, b := dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]
		dp[i][0] = min(min(a, b), nums[i])
		dp[i][1] = max(max(a, b), nums[i])
		result = max(result, dp[i][1])
	}
	return result
}

// 方法2:贪心
func maxProduct2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	result := nums[0]
	preMin, preMax := nums[0], nums[0]
	for i := 1; i < n; i++ {
		a := preMin * nums[i]
		b := preMax * nums[i]
		preMin = min(nums[i], min(a, b))
		preMax = max(nums[i], max(a, b))
		result = max(result, preMax)
	}
	return result
}

// 392. 判断子序列
// https://leetcode.cn/problems/is-subsequence/description/
// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
// 输入：s = "abc", t = "ahbgdc" 输出：true
// [0, 0, 0, 0, 0, 0, 0]
// [0, 1, 1, 1, 1, 1, 1]
// [0, 0, 0, 2, 2, 2, 2]
// [0, 0, 0, 0, 0, 0, 3]
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	// dp[i][j]含义：[0,i-1]的s和[0,j-1]的t，相同子序列长度
	// 递推公式： if s[i-1] == s[j-1] : dp[i][j] = dp[i-1][j-1] + 1 else：dp[i][j] = dp[i][j-1]
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1] // 求的是s是不是t的子序列，t模拟删除最后一个字符
			}
		}
	}
	// fmt.Println(dp)
	return dp[m][n] == len(s)
}

// 115. 不同的子序列
// https://leetcode.cn/problems/distinct-subsequences/description/
// 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数
// 输入：s = "babgbag", t = "bag" 输出：5
// [1, 0, 0, 0]
// [1, 1, 0, 0]
// [1, 1, 1, 0]
// [1, 2, 1, 0]
// [1, 2, 1, 1]
// [1, 3, 1, 1]
// [1, 3, 4, 1]
// [1, 3, 4, 5]
func numDistinct(s string, t string) int {
	// dp[i][j]含义：[0,i-1]的s和[0,j-1]的t的个数
	// 1.s[i] == t[j]: dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
	// 2.s[i] != t[j]: dp[i][j] = dp[i-1][j]
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j] // 两边都删除的个数 + 删除s最后一个
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}

// 583. 两个字符串的删除操作
// https://leetcode.cn/problems/delete-operation-for-two-strings/description/
// 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
// 每步 可以删除任意一个字符串中的一个字符。
// 输入: word1 = "sea", word2 = "eat" 输出: 2
// 解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
// [0, 1, 2, 3]
// [1, 2, 3, 4]
// [2, 1, 2, 3]
// [3, 2, 2, 2]
func minDistance0(word1 string, word2 string) int {
	// dp[i][j]含义：下标为i-1的word1和下标为j-1的word2需要删除的最小次数
	// 递推公式
	// 1.相同 dp[i][j] = dp[i-1][j-1]
	// 2.不同 dp[i][j] = min(dp[i-1][j] + 1, dp[i][j-1] + 1)
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

// 方法2:
// 1.求最长公共子序列长度
// 2.len(word1) + len(word2) - 2*dp[m][n]
func minDistance1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp)
	return len(word1) + len(word2) - 2*dp[m][n]
}

/*
72. 编辑距离
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：插入、删除、替换
输入：word1 = "horse", word2 = "ros" 输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	//dp[i][j]含义：以i-1结尾的word1和j-1结尾的word2使用的最少操作数
	// 递推公式：
	// if word1[i-1] == word2[j-1]: dp[i][j] = dp[i-1][j-1]
	// else：dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}

// 647. 回文子串
// https://leetcode.cn/problems/palindromic-substrings/description/
// 输入：s = "abc" 输出：3
// 解释：三个回文子串: "a", "b", "c"
// [true false false]
// [false true false]
// [false false true]
// 输入：s = "aaa" 输出：6
// 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
func countSubstrings(s string) int {
	// dp[i][j]含义：[i,j]范围内的回文子串个数
	// 递推公式
	// 1.相同 s[i] == s[j]
	//  a.j-i <=1 dp[i][j] = true 回文个数+1
	//  b.        if dp[i+1][j-1] == true => dp[i][j] = true 回文个数+1
	// 2.不同 false
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	result := 0
	for i := n - 1; i >= 0; i-- { // 从下往上
		for j := i; j < n; j++ { // 从左往右
			if s[i] == s[j] {
				if j-i <= 1 || dp[i+1][j-1] {
					dp[i][j] = true
					result++
				}
			}
		}
	}
	// fmt.Println(dp)
	return result
}

func countSubstrings2(s string) int {
	// 1.dp[i][j]含义：左闭右闭区间s[i:j]是不是回文串
	// 2.递推公式：
	// if s[i] == s[j]:
	//   if j-i<=1 || dp[i+1][j-1]: dp[i][j] = true
	// 3.初始化：dp[i][i] = true
	// 4.遍历顺序：i从下到上，j从左到右
	n := len(s)
	dp := make([][]bool, n)
	result := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		result++
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				if j == i+1 || dp[i+1][j-1] {
					dp[i][j] = true
					result++
				}
			}
		}
	}
	return result
}

// 5.最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring/description/
// 给你一个字符串 s，找到 s 中最长的 回文 子串。
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
func longestPalindrome(s string) string {
	// dp[i][j]含义：s[i][j]是不是回文串
	// 递推公式：
	// if s[i] == s[j]:
	//   if j - i <= 1: dp[i][j] = true
	//   else dp[i][j] = dp[i+1][j-1]
	//
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		// dp[i][i] = true
	}
	maxLen := 1
	left := 0
	right := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] && j-i+1 > maxLen {
					maxLen = j - i + 1
					left = i
					right = j
				}
			}
		}
	}
	return s[left : right+1]
}

func longestPalindrome2(s string) string {
	// 1.dp[i][j]含义：左闭右闭区间[i,j]范围内的字符串s是不是回文串
	// 2.递推公式：
	// if s[i] == s[j]:
	//   if j-i==1: dp[i][j] = true
	//   else dp[ij][j] = dp[i+1][j-1]
	// 3.初始化：
	// 4.遍历顺序：i从下往上，j从左往右
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	maxLen := 0
	left, right := 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 || dp[i+1][j-1] {
					dp[i][j] = true
					if j-i+1 > maxLen {
						left, right = i, j
						maxLen = j - i + 1
					}
				}
			}
		}
	}
	return s[left : right+1]
}

// 516.最长回文子序列
// https://leetcode.cn/problems/longest-palindromic-subsequence/description/
// 输入：s = "bbbab" 输出：4
// 解释：一个可能的最长回文子序列为 "bbbb" 。
// 输入：s = "cbbd" 输出：2
// 解释：一个可能的最长回文子序列为 "bb" 。
// dp[i][j]：字符串s在[i, j]范围内最长的回文子序列的长度为dp[i][j]。
// 递推公式：如果s[i]与s[j]不相同，说明s[i]和s[j]的同时加入 并不能增加[i,j]区间回文子序列的长度，那么分别加入s[i]、s[j]看看哪一个可以组成最长的回文子序列。
func longestPalindromeSubseq(s string) int {
	// dp[i][j]含义：下标i到j范围内的字符串内最长回文子序列长度
	// 递推公式：if s[i] == s[j]: dp[i][j] = dp[i-1][j-1] + 2
	//         else: dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

// 132. 分割回文串 II
// https://leetcode.cn/problems/palindrome-partitioning-ii/description/
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串。
// 返回符合要求的 最少分割次数 。
// 输入：s = "aab"
// 输出：1
// 解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
func minCut(s string) int {
	// 预处理：先统计左闭右闭子串s[i:j]是不是回文串
	isValid := make([][]bool, len(s))
	for i := 0; i < len(isValid); i++ {
		isValid[i] = make([]bool, len(s))
		isValid[i][i] = true
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 || isValid[i+1][j-1] {
					isValid[i][j] = true
				}
			}
		}
	}
	// 1.dp[i]含义：切割字符串s[0:i]为多个回文串，最少分割次数
	// 2.递推公式：dp[i] = min(dp[i], dp[j] + 1)
	// 3.初始化：求最少，所以初始化为MaxInt
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < len(s); i++ {
		if isValid[0][i] {
			dp[i] = 0 // 0到i的子串已经是回文串了，不需要切割
			continue
		}
		// 0到i的子串不是回文串，需要切割，使用j来0到i之间尝试
		for j := 0; j < i; j++ {
			if isValid[j+1][i] { // 从j+1到i-1是回文串，可以在j后面切一刀分割
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(s)-1]
}

// 673. 最长递增子序列的个数
// https://leetcode.cn/problems/number-of-longest-increasing-subsequence/description/
// 给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。
// 注意 这个数列必须是 严格 递增的。
// 输入: [1,3,5,4,7]
// 输出: 2
// 解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	// dp[i]含义：以i结尾的nums数组最长递增子序列长度
	// count[i]：以i结尾的nums数组最长递增子序列的个数
	dp := make([]int, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
	}
	maxCount := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
				dp[i] = max(dp[i], dp[j]+1)
			}
			maxCount = max(maxCount, dp[i])
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		if dp[i] == maxCount {
			result += count[i]
		}
	}
	return result
}

func findNumberOfLIS2(nums []int) int {
	// dp[i]含义：以i结尾的nums数组最长递增子序列长度
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}
	maxCount := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i][0] = max(dp[i][0], dp[j][0]+1)
			}
			if dp[i][0] == maxCount {
				dp[i][1]++
			} else if dp[i][0] > maxCount {
				dp[i][1] = 1
				maxCount = dp[i][0]
			}
		}
	}
	fmt.Println(dp)
	fmt.Println(maxCount)
	for i := 0; i < n; i++ {
		if dp[i][0] == maxCount {
			return dp[i][1]
		}
	}
	return 1
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
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))         // 3
	fmt.Println(countSubstrings("abc"))                         // 3
	fmt.Println(findNumberOfLIS2([]int{1, 3, 5, 4, 7}))         // 2
	fmt.Println(findNumberOfLIS2([]int{2, 2, 2, 2, 2}))         // 5
	fmt.Println(longestPalindrome("babad"))                     // bab
	fmt.Println(longestPalindrome2("babad"))                    // bab
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
