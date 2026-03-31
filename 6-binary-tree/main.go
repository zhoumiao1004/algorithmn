package main

import "strings"

// 331. 验证二叉树的前序序列化
// 序列化二叉树的一种方法是使用 前序遍历 。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
// 输入: preorder = "9,3,4,#,#,1,#,#,2,#,6,#,#"
// 输出: true
func isValidSerialization(preorder string) bool {
	edge := 1
	for _, c := range strings.Split(preorder, ",") {
		if c == "#" {
			edge--
			if edge < 0 {
				return false
			}
		} else {
			edge--
			if edge < 0 {
				return false
			}
			edge += 2
		}
	}
	return edge == 0
}

// 543. 二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/description/
// 给你一棵二叉树的根节点，返回该树的 直径 。
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
// 两节点之间路径的 长度 由它们之间边数表示。
// 输入：root = [1,2,3,4,5]
// 输出：3
// 解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
// 思路1: 遍历+前序，遍历整棵二叉树，对每个节点计算左右子树最大深度之和
func diameterOfBinaryTree(root *TreeNode) int {
	// 定义函数返回以node节点为根的二叉树的最大深度
	result := 0
	var maxDepth func(node *TreeNode) int
	var traverse func(node *TreeNode)

	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := maxDepth(node.Left)
		right := maxDepth(node.Right)
		return 1 + max(left, right)
	}

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = max(result, maxDepth(node.Left)+maxDepth(node.Right))
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

// 思路2: 分解问题+后序
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
// 思路1:分解问题+后序
func longestUnivaluePath(root *TreeNode) int {
	res := 0
	var maxLen func(node *TreeNode, parentVal int) int
	// 定义：计算以 root 为根的这棵二叉树中，从 root 开始值为 parentVal 的最长树枝长度
	maxLen = func(node *TreeNode, parentVal int) int {
		if node == nil {
			return 0
		}

		// 利用函数定义，计算左右子树值为 root.val 的最长树枝长度
		leftLen := maxLen(node.Left, node.Val)
		rightLen := maxLen(node.Right, node.Val)

		// 后序位置
		if node.Val != parentVal {
			return 0
		}
		res = max(res, leftLen+rightLen)

		return 1 + max(leftLen, rightLen)
	}

	if root == nil {
		return res
	}
	maxLen(root, root.Val)
	return res
}
