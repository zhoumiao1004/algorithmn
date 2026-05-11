package main

import (
	"fmt"
	"sort"
)

// 0-1背包
// 明确状态：背包容量和可选择的物品；选择：装进背包 or 不装进背包
func knapsack(wt, val []int, W int) int {
	// 定义dp[i][j]: 对于前 i 个物品，当前背包的容量为 j，这种情况下可以装的最大价值是 dp[i][j]。
	N := len(wt)
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}

	// 由于数组索引从 0 开始，而我们定义中的 i 是从 1 开始计数的，所以 val[i-1] 和 wt[i-1] 表示第 i 个物品的价值和重量。
	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if j < wt[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wt[i-1]]+val[i-1]) // 和0-背包的区别是dp[i-1][j-weight[i-1]]
			}
		}
	}
	return dp[N][W]
}

func knapsack2(weight, value []int, n int) int {
	// 定义dp[i][j]数组：从下表0-i的物品中取，放进j大小的背包的价值总和
	m := len(weight)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	for j := weight[0]; j <= n; j++ {
		dp[0][j] = value[0]
	}

	for i := 1; i < m; i++ {
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

// 用滚动数组进行空间压缩
func knapsackN(weight, value []int, n int) int {
	m := len(weight)
	dp := make([]int, n+1)
	for i := 0; i < m; i++ {
		for j := n; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	// fmt.Println(dp)
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

	dp := make([]int, target+1)      // dp[j]含义：大小为j的背包，能装的最大价值
	for i := 0; i < len(nums); i++ { // 物品
		for j := target; j >= nums[i]; j-- { // 背包逆序
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}

func canPartition2(nums []int) bool {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	if s%2 == 1 {
		return false
	}
	target := s / 2

	dp := make([]bool, target+1) // dp[j]含义：能否装满大小为j的背包
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}

func canPartition3(nums []int) bool {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	if s%2 == 1 {
		return false
	}
	target := s / 2

	n := len(nums)
	dp := make([][]bool, n+1) // dp[i][j]含义：对于前 i 个物品，当前背包大小为 j，这种情况下能否装满背包 dp[i][j]。
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][target]
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

// 二维dp数组，明确状态：背包容量和可选择的物品；选择：不放第 i 个物品 or 放第 i 个物品
// 定义dp[i][j]: 对于前 i 个物品，当前背包的容量为 j，这种情况下可以装的最大价值是 dp[i][j]。
func findTargetSumWays2(nums []int, target int) int {
	n := len(nums)
	s := 0
	for i := 0; i < n; i++ {
		s += nums[i]
	}
	if (s+target)%2 != 0 {
		return 0
	}
	if target > s || target < -s {
		return 0
	}
	target = (s + target) / 2

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][target]
}

// 回溯，时间复杂度O(2^n)
func findTargetSumWays3(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := 0
	s := 0
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			if s == target {
				res++
			}
			return
		}
		s += nums[i]
		backtrack(i + 1)
		s -= nums[i]

		s -= nums[i]
		backtrack(i + 1)
		s += nums[i]
	}
	backtrack(0)
	return res
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

// 3180. 执行操作可获得的最大总奖励 I
// https://leetcode.cn/problems/maximum-total-reward-using-operations-i/
// 给你一个整数数组 rewardValues，长度为 n，代表奖励的值。
// 最初，你的总奖励 x 为 0，所有下标都是 未标记 的。你可以执行以下操作 任意次 ：
// 从区间 [0, n - 1] 中选择一个 未标记 的下标 i。
// 如果 rewardValues[i] 大于 你当前的总奖励 x，则将 rewardValues[i] 加到 x 上（即 x = x + rewardValues[i]），并 标记 下标 i。
// 以整数形式返回执行最优操作能够获得的 最大 总奖励。
// 输入：rewardValues = [1,1,3,3]
// 输出：4
// 解释：
// 依次标记下标 0 和 2，总奖励为 4，这是可获得的最大值。
// 输入：rewardValues = [1,6,4,3,2]
// 输出：11
// 分析：max(rewardValues)一定在结果中，其他元素之和必然小于它，x <= 2*max(rewardValues)-1
func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	n := len(rewardValues)
	if n == 0 {
		return 0
	}
	maxVal := rewardValues[n-1]
	// 定义：dp[i][x] 表示仅使用 rewardValues[...i] 物品，是否能凑出总价值为 x
	// 递推公式：if rewardValues[i] > dp[i-1][x-rewardValues[i]] then dp[i][x] = true
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, 2*maxVal)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		curVal := rewardValues[i-1]
		for j := 0; j < 2*maxVal; j++ {
			if j >= curVal && curVal > j-curVal {
				dp[i][j] = dp[i-1][j-curVal] || dp[i-1][j] // 用 or 不用
			} else {
				dp[i][j] = dp[i-1][j] // curVal不能装进背包
			}
		}
	}
	// 返回最大价值
	for j := maxVal*2 - 1; j >= 0; j-- {
		if dp[n][j] {
			return j
		}
	}
	return 0
}

func main() {
	fmt.Println(knapsack2([]int{1, 3, 4}, []int{15, 20, 30}, 4)) // 35 = 15+20
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
