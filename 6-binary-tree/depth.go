package main

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 思路1:遍历整棵树，外部变量记录递归深度
func maxDepth(root *TreeNode) int {
	result := 0
	depth := 0
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			result = max(result, depth)
			return
		}
		depth++
		traverse(node.Left)
		traverse(node.Right)
		depth--
	}

	traverse(root)
	return result
}

// 思路2:分解问题+后序
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return 1 + max(leftDepth, rightDepth)
}

// 111.二叉树的最小深度
// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
// 输入：root = [3,9,20,null,null,15,7]
// 输出：2
// 思路1:分解问题+后序
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Right)
	rightDepth := minDepth(root.Left)

	// 后序位置
	if root.Left == nil && root.Right != nil {
		return 1 + rightDepth
	}
	if root.Right == nil && root.Left != nil {
		return 1 + leftDepth
	}

	return 1 + min(leftDepth, rightDepth)
}

// 思路2: 层序遍历BFS。遍历到的第一个叶子节点的深度
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		depth++
		sz := len(q)
		for i := 0; i < sz; i++ {
			node := q[0]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return depth
}
