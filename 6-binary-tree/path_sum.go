package main

import (
	"fmt"
	"math"
)

// 112. 路径总和
// https://leetcode.cn/problems/path-sum/description/
// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
// 叶子节点 是指没有子节点的节点。
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 不涉及中的操作，所以前中后序遍历都可以
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 113. 路径总和 II
// https://leetcode.cn/problems/path-sum-ii/description/
// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
// 叶子节点 是指没有子节点的节点。
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：[[5,4,11,2],[5,8,4,5]]
func pathSumII(root *TreeNode, targetSum int) [][]int {
	var results [][]int
	var path []int
	s := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		s += root.Val
		if s == targetSum && root.Left == nil && root.Right == nil {
			results = append(results, append([]int{}, path...))
		}
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]
		s -= root.Val
	}
	dfs(root)
	return results
}

// 技巧2:一般来说，遍历的思维模式可以帮你寻找从根节点开始的符合条件的「树枝」，但在不限制起点必须是根节点的条件下，让你寻找符合条件的「树枝」，就需要用到分解问题的思维模式了。
// 124. 二叉树中的最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/
// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
// 路径和 是路径中各节点值的总和。
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
func maxPathSum(root *TreeNode) int {
	result := math.MinInt
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		// 后序位置顺便计算双边最大路径和
		maxSum := left + right + root.Val
		result = max(result, maxSum)
		// 返回单边最大路径和
		return max(left, right) + root.Val
	}
	dfs(root)
	return result
}

func main() {
	fmt.Println("hello world")
}
