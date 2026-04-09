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
// 思路：后序。明确函数定义：返回以 root 为根的二叉树中，p和q的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root // 找到就向上返回
	}
	left := lowestCommonAncestor(root.Left, p, q)   // 左
	right := lowestCommonAncestor(root.Right, p, q) // 右
	// 后序位置，根据左右子树的结果，计算以 root 为根的二叉树中，p和q的最近公共祖先
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
// 思路：后序。明确函数定义：返回以 root 为根的bst中，p和q的最近公共祖先
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

// 865. 具有所有最深节点的最小子树
// https://leetcode.cn/problems/smallest-subtree-with-all-the-deepest-nodes/description/
// 1123. 最深叶节点的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/description/
// 给定一个根为 root 的二叉树，每个节点的深度是 该节点到根的最短距离 。
// 返回包含原始树中所有 最深节点 的 最小子树 。
// 如果一个节点在 整个树 的任意节点之间具有最大的深度，则该节点是 最深的 。
// 一个节点的 子树 是该节点加上它的所有后代的集合。
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4]
// 输出：[2,7,4]
// 思路：分解问题
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	// 明确函数定义：以 node 节点为根的二叉树，返回包含所有最深节点的子树和深度
	var maxDepthNode func(root *TreeNode) (*TreeNode, int)

	maxDepthNode = func(root *TreeNode) (*TreeNode, int) {
		if root == nil {
			return nil, 0
		}
		left, depth1 := maxDepthNode(root.Left)
		right, depth2 := maxDepthNode(root.Right)
		// 后序位置
		if depth1 == depth2 {
			return root, depth1 + 1 // 把自己往上报
		} else if depth1 < depth2 {
			return right, depth2 + 1 // 把右节点往上报
		}
		return left, depth1 + 1 // 把左节点往上报
	}

	node, _ := maxDepthNode(root)
	return node
}

func main() {
	fmt.Println("Hello, World!")
}
