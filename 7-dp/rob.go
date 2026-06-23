package main

import (
	"math"
	"sort"
)

// 198. 打家劫舍
// https://leetcode.cn/problems/house-robber/description/
// 相邻不能偷。dp[i]两种情况：1.偷，dp[i-2]+nums[i] 2.不偷dp[i-1]
// 输入：[1,2,3,1] 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。偷窃到的最高金额 = 1 + 3 = 4 。
// 思路1: 自顶向下的递归解法
func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = -1
	}
	var dp func(start int) int

	dp = func(start int) int {
		if start >= len(nums) {
			return 0
		}
		if memo[start] != -1 {
			return memo[start]
		}
		res := max(
			dp(start+1),
			nums[start]+dp(start+2),
		)
		memo[start] = res
		return res
	}
	return dp(0)
}

// 思路2: 自底向上的迭代解法，从后往前
func rob2(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)
	for i := n - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], dp[i+2]+nums[i])
	}
	return dp[0]
}

// 思路2: 自底向上的迭代解法，从前往后
func rob3(nums []int) int {
	n := len(nums)
	dp := make([]int, n+2)
	for i := 2; i <= n+1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-2])
	}
	return dp[n+1]
}

// 思路2: 自底向上的迭代解法，从前往后
func rob4(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
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
func robII(nums []int) int {
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

func robIII(root *TreeNode) int {
	var dp func(root *TreeNode) [2]int
	dp = func(root *TreeNode) [2]int {
		if root == nil {
			return [2]int{0, 0}
		}
		left := dp(root.Left)
		right := dp(root.Right)
		// 后序位置
		notRob := max(left[0], left[1]) + max(right[0], right[1])
		rob := left[0] + right[0] + root.Val
		return [2]int{notRob, rob}
	}
	res := dp(root)
	return max(res[0], res[1])
}

// 2140. 解决智力问题
// https://leetcode.cn/problems/solving-questions-with-brainpower/description/
// 给你一个下标从 0 开始的二维整数数组 questions ，其中 questions[i] = [pointsi, brainpoweri] 。
// 这个数组表示一场考试里的一系列题目，你需要 按顺序 （也就是从问题 0 开始依次解决），针对每个问题选择 解决 或者 跳过 操作。解决问题 i 将让你 获得  pointsi 的分数，但是你将 无法 解决接下来的 brainpoweri 个问题（即只能跳过接下来的 brainpoweri 个问题）。如果你跳过问题 i ，你可以对下一个问题决定使用哪种操作。
// 比方说，给你 questions = [[3, 2], [4, 3], [4, 4], [2, 5]] ：
// 如果问题 0 被解决了， 那么你可以获得 3 分，但你不能解决问题 1 和 2 。
// 如果你跳过问题 0 ，且解决问题 1 ，你将获得 4 分但是不能解决问题 2 和 3 。
// 请你返回这场考试里你能获得的 最高 分数。
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	memo := make([]int64, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	var dp func(i int) int64
	dp = func(i int) int64 {
		if i >= n {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		score := int64(questions[i][0])
		skip := questions[i][1]
		res := max(score+dp(i+skip+1), dp(i+1))
		memo[i] = res
		return res
	}
	return dp(0)
}

// dp数组迭代写法：从后向前遍历，依次计算 dp[i]
// nextIdx = i + skip + 1 可能越界，需要裁剪到 n
// dp[i+1] 就是"跳过当前题"的分数，score + dp[nextIdx] 是"做当前题"的分数
func mostPoints2(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1) // dp[n] = 0 是base case

	for i := n - 1; i >= 0; i-- {
		score := int64(questions[i][0])
		skip := questions[i][1]
		nextIdx := min(i+skip+1, n)
		dp[i] = max(score+dp[nextIdx], dp[i+1])
	}
	return dp[0]
}

// 2320. 统计放置房子的方式数
// https://leetcode.cn/problems/count-number-of-ways-to-place-houses/description/
// 一条街道上共有 n * 2 个 地块 ，街道的两侧各有 n 个地块。每一边的地块都按从 1 到 n 编号。每个地块上都可以放置一所房子。
// 现要求街道同一侧不能存在两所房子相邻的情况，请你计算并返回放置房屋的方式数目。由于答案可能很大，需要对 109 + 7 取余后再返回。
// 注意，如果一所房子放置在这条街某一侧上的第 i 个地块，不影响在另一侧的第 i 个地块放置房子。
func countHousePlacements(n int) int {
	memo := make([]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = -1
	}
	var dp func(i int) int // 从start到第n地块的，能放房子的方式数
	dp = func(i int) int {
		if i >= n {
			return 1
		}
		if memo[i] != -1 {
			return memo[i]
		}
		res := (dp(i+1) + dp(i+2)) % (1e9 + 7)
		memo[i] = res
		return res
	}

	res := dp(0)
	return (res * res) % (1e9 + 7)
}

// 983. 最低票价
// https://leetcode.cn/problems/minimum-cost-for-tickets/description/
// 在一个火车旅行很受欢迎的国度，你提前一年计划了一些火车旅行。在接下来的一年里，你要旅行的日子将以一个名为 days 的数组给出。每一项是一个从 1 到 365 的整数。
// 火车票有 三种不同的销售方式 ：
// 一张 为期一天 的通行证售价为 costs[0] 美元；
// 一张 为期七天 的通行证售价为 costs[1] 美元；
// 一张 为期三十天 的通行证售价为 costs[2] 美元。
// 通行证允许数天无限制的旅行。 例如，如果我们在第 2 天获得一张 为期 7 天 的通行证，那么我们可以连着旅行 7 天：第 2 天、第 3 天、第 4 天、第 5 天、第 6 天、第 7 天和第 8 天。
// 返回 你想要完成在给定的列表 days 中列出的每一天的旅行所需要的最低消费 。
// 输入：days = [1,4,6,7,8,20], costs = [2,7,15]
// 输出：11
// 解释：
// 例如，这里有一种购买通行证的方法，可以让你完成你的旅行计划：
// 在第 1 天，你花了 costs[0] = $2 买了一张为期 1 天的通行证，它将在第 1 天生效。
// 在第 3 天，你花了 costs[1] = $7 买了一张为期 7 天的通行证，它将在第 3, 4, ..., 9 天生效。
// 在第 20 天，你花了 costs[0] = $2 买了一张为期 1 天的通行证，它将在第 20 天生效。
// 你总共花了 $11，并完成了你计划的每一天旅行。
func mincostTickets(days []int, costs []int) int {
	memo := make([]int, len(days))
	for i := 0; i < len(days); i++ {
		memo[i] = -1
	}
	var dp func(start int) int
	dp = func(start int) int {
		if start >= len(days) {
			return 0
		}
		if memo[start] != -1 {
			return memo[start]
		}
		res := math.MaxInt
		// 选择买1天的票
		currentDay := days[start]
		i := start
		for i < len(days) && days[i] < currentDay+1 {
			i++
		}
		c1 := dp(i) + costs[0]
		// 选择买7天的票
		for i < len(days) && days[i] < currentDay+7 {
			i++
		}
		c2 := dp(i) + costs[1]
		// 选择买730天的票
		for i < len(days) && days[i] < currentDay+30 {
			i++
		}
		c3 := dp(i) + costs[2]
		res = min(c1, c2, c3)

		memo[start] = res
		return res
	}
	return dp(0)
}

// 740. 删除并获得点数
// https://leetcode.cn/problems/delete-and-earn/description/
// 给你一个整数数组 nums ，你可以对它进行一些操作。
// 每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。
// 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
// 输入：nums = [3,4,2]
// 输出：6
// 解释：
// 你可以执行下列步骤：
// - 删除 4 获得 4 个点数，因此 3 也被删除。nums = [2]。
// - 之后，删除 2 获得 2 个点数。nums = []。
// 总共获得 6 个点数。
func deleteAndEarn(nums []int) int {
	points := make([]int, 10001)
	for _, num := range nums {
		points[num] += num
	}
	var rob func(nums []int) int
	rob = func(nums []int) int {
		n := len(nums)
		dp := make([]int, n+2)
		for i := n - 1; i >= 0; i-- {
			dp[i] = max(dp[i+1], dp[i+2]+nums[i])
		}
		return dp[0]
	}
	return rob(points)
}

// 2611. 老鼠和奶酪
// https://leetcode.cn/problems/mice-and-cheese/description/
// 有两只老鼠和 n 块不同类型的奶酪，每块奶酪都只能被其中一只老鼠吃掉。
// 下标为 i 处的奶酪被吃掉的得分为：
// 如果第一只老鼠吃掉，则得分为 reward1[i] 。
// 如果第二只老鼠吃掉，则得分为 reward2[i] 。
// 给你一个正整数数组 reward1 ，一个正整数数组 reward2 ，和一个非负整数 k 。
// 请你返回第一只老鼠恰好吃掉 k 块奶酪的情况下，最大 得分为多少。
// 输入：reward1 = [1,1,3,4], reward2 = [4,4,1,1], k = 2
// 输出：15
// 解释：这个例子中，第一只老鼠吃掉第 2 和 3 块奶酪（下标从 0 开始），第二只老鼠吃掉第 0 和 1 块奶酪。
// 总得分为 4 + 4 + 3 + 4 = 15 。
// 15 是最高得分。
func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n := len(reward1)
	diff := make([][2]int, n)
	for i := 0; i < n; i++ {
		diff[i][0] = reward1[i] - reward2[i]
		diff[i][1] = i
	}

	sort.Slice(diff, func(i, j int) bool {
		return diff[i][0] > diff[j][0]
	})

	res := 0
	for i := 0; i < k; i++ {
		res += reward1[diff[i][1]]
	}
	for i := k; i < n; i++ {
		res += reward2[diff[i][1]]
	}
	return res
}

// 2789. 合并后数组中的最大元素
// https://leetcode.cn/problems/largest-element-in-an-array-after-merge-operations/description/
// 给你一个下标从 0 开始、由正整数组成的数组 nums 。
// 你可以在数组上执行下述操作 任意 次：
// 选中一个同时满足 0 <= i < nums.length - 1 和 nums[i] <= nums[i + 1] 的下标 i 。将元素 nums[i + 1] 替换为 nums[i] + nums[i + 1] ，并从数组中删除元素 nums[i] 。
// 返回你可以从最终数组中获得的 最大 元素的值。
// 输入：nums = [2,3,7,9,3]
// 输出：21
// 解释：我们可以在数组上执行下述操作：
// - 选中 i = 0 ，得到数组 nums = [5,7,9,3] 。
// - 选中 i = 1 ，得到数组 nums = [5,16,3] 。
// - 选中 i = 0 ，得到数组 nums = [21,3] 。
// 最终数组中的最大元素是 21 。可以证明我们无法获得更大的元素。
func maxArrayValue(nums []int) int64 {
	var res int64
	i := len(nums) - 1
	for ; i >= 0; i-- {
		blockSum := int64(nums[i])
		for i > 0 && blockSum >= int64(nums[i-1]) {
			blockSum += int64(nums[i-1])
			i--
		}
		res = max(res, blockSum)
	}
	return res
}
