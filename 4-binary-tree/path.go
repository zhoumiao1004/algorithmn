package main

import (
	"fmt"
	"math"
)

// 543. 二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/description/
// 给你一棵二叉树的根节点，返回该树的 直径 。
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
// 两节点之间路径的 长度 由它们之间边数表示。
// 输入：root = [1,2,3,4,5]
// 输出：3
// 解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
// 最优思路: 分解问题+后序
func diameterOfBinaryTree2(root *TreeNode) int {
	result := 0
	var maxDepth func(node *TreeNode) int

	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := maxDepth(node.Left)
		right := maxDepth(node.Right)

		// 后序位置，顺便计算最大值
		result = max(result, left+right)

		return 1 + max(left, right)
	}

	if root == nil {
		return 0
	}
	maxDepth(root)
	return result
}

// 687. 最长同值路径
// https://leetcode.cn/problems/longest-univalue-path/description/
// 给定一个二叉树的 root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。
// 两个节点之间的路径长度 由它们之间的边数表示。
// 输入：root = [5,4,5,1,1,5]
// 输出：2
// 思路1:分解问题+后序，利用函数定义，计算左右子树值为 root.val 的最长树枝长度，后序位置
func longestUnivaluePath(root *TreeNode) int {
	res := 0
	var maxLen func(node *TreeNode, parentVal int) int // 明确函数定义：返回以 root 为根的这棵二叉树中，从 root 开始值为 parentVal 的最长树枝长度

	maxLen = func(node *TreeNode, parentVal int) int {
		if node == nil {
			return 0
		}
		leftLen := maxLen(node.Left, node.Val)
		rightLen := maxLen(node.Right, node.Val)

		// 后序位置
		res = max(res, leftLen+rightLen) // 顺便计算以 node 为根的路径长度
		if node.Val != parentVal {
			return 0 // 不同值
		}
		return 1 + max(leftLen, rightLen)
	}

	if root == nil {
		return res
	}
	maxLen(root, root.Val)
	return res
}

// 112. 路径总和
// https://leetcode.cn/problems/path-sum/description/
// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
// 叶子节点 是指没有子节点的节点。
// 思路1: 分解问题
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 前序位置，不涉及中的操作，所以前中后序遍历都可以
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val) // 左右
}

// 思路2: 遍历
func hasPathSum2(root *TreeNode, targetSum int) bool {
	s := 0
	flag := false
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if flag {
			return
		}
		if root == nil {
			return
		}
		s += root.Val
		if s == targetSum && root.Left == nil && root.Right == nil {
			flag = true
			return
		}
		traverse(root.Left)
		traverse(root.Right)
		s -= root.Val
	}

	traverse(root)
	return flag
}

// 113. 路径总和 II
// https://leetcode.cn/problems/path-sum-ii/description/
// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
// 叶子节点 是指没有子节点的节点。
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：[[5,4,11,2],[5,8,4,5]]
// 思路1: 遍历
func pathSumII2(root *TreeNode, targetSum int) [][]int {
	var results [][]int
	var path []int
	s := 0
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		s += root.Val
		if s == targetSum && root.Left == nil && root.Right == nil {
			results = append(results, append([]int{}, path...))
		}
		traverse(root.Left)
		traverse(root.Right)
		path = path[:len(path)-1]
		s -= root.Val
	}

	traverse(root)
	return results
}

// 思路2: 分解
func pathSumII(root *TreeNode, targetSum int) [][]int {
	var res [][]int

	if root == nil {
		return res
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		res = append(res, []int{root.Val})
		return res
	}
	left := pathSumII(root.Left, targetSum-root.Val)
	right := pathSumII(root.Right, targetSum-root.Val)

	for _, path := range left {
		res = append(res, append([]int{root.Val}, path...))
	}
	for _, path := range right {
		res = append(res, append([]int{root.Val}, path...))
	}
	return res
}

// 技巧2:一般来说，遍历的思维模式可以帮你寻找从根节点开始的符合条件的「树枝」，但在不限制起点必须是根节点的条件下，让你寻找符合条件的「树枝」，就需要用到分解问题的思维模式了。
// 124. 二叉树中的最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/
// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
// 路径和 是路径中各节点值的总和。
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
// 思路：分解问题，明确函数定义：以 node 为根节点的二叉树，返回双边最大和
func maxPathSum(root *TreeNode) int {
	result := math.MinInt
	var oneSideMax func(node *TreeNode) int

	oneSideMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, oneSideMax(node.Left))
		right := max(0, oneSideMax(node.Right))

		// 后序位置，顺便计算双边最大路径和
		result = max(result, left+right+node.Val)

		return max(left, right) + node.Val // 返回单边最大路径和
	}

	oneSideMax(root)
	return result
}

// 437. 路径总和 III
// https://leetcode.cn/problems/path-sum-iii/description/
// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
// 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// 输出：3
// 解释：和等于 8 的路径有 3 条，如图所示。
func pathSum(root *TreeNode, targetSum int) int {
	result := 0
	if root == nil {
		return 0
	}
	preSumCount := make(map[int]int)
	preSumCount[0] = 1
	pathSum := 0
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 前序遍历位置
		pathSum += root.Val // 从根开始的前缀和
		result += preSumCount[pathSum-targetSum]
		preSumCount[pathSum]++

		traverse(root.Left)
		traverse(root.Right)

		// 后序遍历位置
		preSumCount[pathSum]--
		pathSum -= root.Val
	}

	traverse(root)
	return result
}

func main() {
	fmt.Println("hello world")
}
