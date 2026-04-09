package main

import "math"

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
		if k > 0 {
			result = node.Val
			k--
		}
		traverse(node.Right)
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

// 538. 把二叉搜索树转换为累加树
// https://leetcode.cn/problems/convert-bst-to-greater-tree/
// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
// 提醒一下，二叉搜索树满足下列约束条件：
// 节点的左子树仅包含键 小于 节点键的节点。
// 节点的右子树仅包含键 大于 节点键的节点。
// 左右子树也必须是二叉搜索树。
// 注意：本题和 1038: https://leetcode.cn/problems/binary-search-tree-to-greater-sum-tree/ 相同
func convertBST(root *TreeNode) *TreeNode {
	s := 0
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Right) // 右
		// 中序位置
		s += node.Val
		node.Val = s
		traverse(node.Left) // 左
	}

	traverse(root)
	return root
}

// 98.验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/description/
// 二叉搜索树定义如下：
// 节点的左子树只包含 严格小于 当前节点的数。
// 节点的右子树只包含 严格大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
// 思路1: 分解问题，返回以 node 为根的树是不是bst
func isValidBST(root *TreeNode) bool {

	var isValid func(node *TreeNode, minNode, maxNode *TreeNode) bool

	isValid = func(node *TreeNode, minNode, maxNode *TreeNode) bool {
		if node == nil {
			return true
		}
		if minNode != nil && node.Val <= minNode.Val {
			return false
		}
		if maxNode != nil && node.Val >= maxNode.Val {
			return false
		}
		return isValid(node.Left, minNode, node) && isValid(node.Right, node, maxNode)
	}

	return isValid(root, nil, nil)
}

// 思路2：中序遍历有序
func isValidBST2(root *TreeNode) bool {
	isValid := true
	var prev *TreeNode

	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		// 中序位置
		if prev != nil && prev.Val >= node.Val {
			isValid = false
		}
		prev = node
		traverse(node.Right)
	}

	traverse(root)
	return isValid
}

// 501.二叉搜索树中的众数
// https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/
// 思路：遍历，利用中序有序累计节点值个数，不断更新结果（最大个数和有最大个数的值）
func findMode(root *TreeNode) []int {
	var results []int
	maxCnt := 0
	cnt := 0
	var prev *TreeNode
	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		// 中序位置
		if prev != nil && prev.Val == root.Val {
			cnt++
		} else {
			cnt = 1
		}
		if cnt == maxCnt {
			results = append(results, root.Val)
		} else if cnt > maxCnt {
			results = []int{root.Val}
			maxCnt = cnt
		}
		prev = root
		traverse(root.Right)
	}

	traverse(root)
	return results
}

// 530. 二叉搜索树的最小绝对差
// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/
// 思路：遍历。中序位置计算相邻节点的差，不断更新结果（最小值）
func getMinimumDifference(root *TreeNode) int {
	result := math.MaxInt
	var prev *TreeNode
	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		// 中序位置
		if prev != nil {
			result = min(result, root.Val-prev.Val)
		}
		prev = root
		traverse(root.Right)
	}

	traverse(root)
	return result
}

// 99. 恢复二叉搜索树
// https://leetcode.cn/problems/recover-binary-search-tree/description/
// 给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
// 思路：遍历。中序遍历找不满足第一个和最后一个不满足有序的2个节点进行交换
func recoverTree(root *TreeNode) {
	var prev *TreeNode
	var first, second *TreeNode
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		// 中序位置
		if prev != nil && prev.Val > root.Val {
			if first == nil {
				first = prev // 记录第一个不满足有序的节点
			}
			second = root // 更新最后一个不满足有序的节点
		}
		prev = root
		traverse(root.Right)
	}

	traverse(root)
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}
