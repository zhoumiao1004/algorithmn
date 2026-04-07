package main

// 230. 二叉搜索树中第 K 小的元素
// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（k 从 1 开始计数）。
func kthSmallest(root *TreeNode, k int) int {
	result := 0
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		// 中序位置
		k--
		if k == 0 {
			result = node.Val
		}
		traverse(node.Right)
	}

	if root == nil {
		return 0
	}
	traverse(root)
	return result
}

func kthSmallest2(root *TreeNode, k int) int {
	type MyNode struct {
		Val   int
		Size  int
		Left  *MyNode
		Right *MyNode
	}
	var build func(node *TreeNode) *MyNode
	build = func(node *TreeNode) *MyNode {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			return &MyNode{Val: node.Val, Size: 1}
		}
		size := 1
		left := build(node.Left)
		if left != nil {
			size += left.Size
		}
		right := build(node.Right)
		if right != nil {
			size += right.Size
		}
		return &MyNode{
			Val:   node.Val,
			Size:  left.Size + right.Size,
			Left:  left,
			Right: right,
		}
	}
	var find func(root *MyNode, k int) int
	find = func(root *MyNode, k int) int {
		if root == nil {
			return 0
		}
		if root.Left == nil {
			if k == 1 {
				return root.Val
			}
			return find(root.Right, k-1)
		}
		if k == root.Left.Size+1 {
			return root.Val
		} else if k < root.Left.Size+1 {
			return find(root.Left, k)
		} else {
			return find(root.Right, k-root.Left.Size-1)
		}
	}
	newRoot := build(root)
	return find(newRoot, k)
}
