package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 897. 递增顺序搜索树
// https://leetcode.cn/problems/increasing-order-search-tree/
// 给你一棵二叉搜索树的 root ，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
// 输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
// 输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
// 思路1: 遍历
func increasingBST2(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	cur := dummy
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left) // 左

		// 中序位置
		cur.Right = &TreeNode{Val: node.Val}
		cur = cur.Right

		traverse(root.Right) // 右
	}

	traverse(root)
	return dummy.Right
}

// 思路2: 分解问题+后序
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := increasingBST(root.Left)   // 左
	root.Left = nil                    // 断掉左子树
	right := increasingBST(root.Right) // 右
	root.Right = right

	// 后序位置
	if left == nil {
		return root
	}
	// 把 root 接到左子树最右边的节点上
	cur := left
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = root
	return left
}
