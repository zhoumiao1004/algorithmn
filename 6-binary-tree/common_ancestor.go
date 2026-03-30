package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 236. 二叉树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 后序遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root // 找到就向上返回
	}
	left := lowestCommonAncestor(root.Left, p, q)   // 左
	right := lowestCommonAncestor(root.Right, p, q) // 右
	// 中
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// 235. 二叉搜索树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// 输出: 6
// 解释: 节点 2 和节点 8 的最近公共祖先是 6。
func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root // 找到就向上返回
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorBST(root.Left, q, p)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorBST(root.Right, p, q)
	}
	return root
}

func main() {
	fmt.Println("Hello, World!")
}
