package main

// 938. 二叉搜索树的范围和
// https://leetcode.cn/problems/range-sum-of-bst/
// 给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
// 输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
// 输出：32
// 对比 669. 修剪二叉搜索树 https://leetcode.cn/problems/trim-a-binary-search-tree/description/
// 思路1: 分解问题，利用bst的特性（推荐）
func rangeSumBST2(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val < low {
		return rangeSumBST2(root.Right, low, high)
	} else if root.Val > high {
		return rangeSumBST2(root.Left, low, high)
	} else {
		left := rangeSumBST2(root.Left, low, root.Val)
		right := rangeSumBST2(root.Right, root.Val, high)
		return root.Val + left + right
	}
}

// 思路2: 遍历，利用bst的特性
func rangeSumBST(root *TreeNode, low int, high int) int {
	result := 0
	var traverse func(node *TreeNode, low, high int)

	traverse = func(node *TreeNode, low, high int) {
		if node == nil {
			return
		}
		if node.Val < low {
			traverse(node.Right, low, high)
		} else if node.Val > high {
			traverse(node.Left, low, high)
		} else {
			result += node.Val
			traverse(node.Left, low, node.Val-1)
			traverse(node.Right, node.Val+1, high)
		}
	}

	traverse(root, low, high)
	return result
}

// 普通二叉树的范围和
// 思路1：遍历
func rangeSum(root *TreeNode, low int, high int) int {
	result := 0
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val >= low && root.Val <= high {
			result += root.Val
		}
		traverse(root.Left)
		traverse(root.Right)
	}

	if root == nil {
		return 0
	}
	traverse(root)
	return result
}

// 思路2: 分解问题的思路，明确函数定义：返回以 root 节点为根的二叉树在[low...high]区间内的节点
func rangeSumBST4(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	left := rangeSumBST4(root.Right, low, high)
	right := rangeSumBST4(root.Left, low, high)
	// 后序位置
	s := left + right
	if root.Val >= low && root.Val <= high {
		s += root.Val
	}
	return s
}
