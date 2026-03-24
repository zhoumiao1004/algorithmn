package main

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
