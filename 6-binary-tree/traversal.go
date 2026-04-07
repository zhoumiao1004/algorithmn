package main

// 144. 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
// 思路1：遍历
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val) // 中
		traverse(node.Left)               // 左
		traverse(node.Right)              // 右
	}
	traverse(root)
	return result
}

// 思路2：分解问题
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{root.Val}
	result = append(result, preorderTraversal1(root.Left)...)
	result = append(result, preorderTraversal1(root.Right)...)
	return result
}

// 迭代法
func preorderTraversal2(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	st := []*TreeNode{root}
	for len(st) > 0 {
		node := st[len(st)-1]
		result = append(result, node.Val)
		st = st[:len(st)-1]
		if node.Right != nil {
			st = append(st, node.Right)
		}
		if node.Left != nil {
			st = append(st, node.Left)
		}
	}
	return result
}

// 145. 二叉树的后序遍历
// https://leetcode.cn/problems/binary-tree-postorder-traversal/description/
// 思路1: 遍历
func postorderTraversal(root *TreeNode) []int {
	var result []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		result = append(result, node.Val)
	}
	traverse(root)
	return result
}

// 迭代法
func postorderTraversal2(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	st := []*TreeNode{root}
	for len(st) > 0 {
		node := st[len(st)-1] // 中
		result = append(result, node.Val)
		st = st[:len(st)-1]
		if node.Left != nil {
			st = append(st, node.Left) // 左
		}
		if node.Right != nil {
			st = append(st, node.Right) // 右
		}
	}
	i, j := 0, len(result)-1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}

// 94. 二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
// 思路1: 遍历
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		result = append(result, node.Val)
		traverse(node.Right)
	}
	traverse(root)
	return result
}

// 迭代法
func inorderTraversal2(root *TreeNode) []int {
	var result []int
	var st []*TreeNode
	cur := root

	for len(st) > 0 || cur != nil {
		if cur != nil {
			st = append(st, cur)
			cur = cur.Left
		} else {
			cur = st[len(st)-1]
			result = append(result, cur.Val)
			st = st[:len(st)-1]
			cur = cur.Right
		}
	}
	return result
}
