package main

// 222.完全二叉树的节点个数
// https://leetcode.cn/problems/count-complete-tree-nodes/description/
// 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
// 完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层（从第 0 层开始），则该层包含 1~ 2h 个节点。
// 输入：root = [1,2,3,4,5,6]
// 输出：6
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 前序位置：统计节点左右子树深度，判断以root为根的二叉树是不是完全二叉树
	var leftDepth, rightDepth int
	for cur := root; cur != nil; cur = cur.Left {
		leftDepth++
	}
	for cur := root; cur != nil; cur = cur.Right {
		rightDepth++
	}
	if leftDepth == rightDepth {
		return 1<<leftDepth - 1 // 以root为根的二叉树是完全二叉树，使用公式计算
	}
	leftNum := countNodes(root.Left)
	rightNum := countNodes(root.Right)
	// 后序位置
	return 1 + leftNum + rightNum
}
