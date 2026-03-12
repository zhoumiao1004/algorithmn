package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 230. 二叉搜索树中第 K 小的元素
// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（k 从 1 开始计数）。
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	result := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		// 中
		k--
		if k == 0 {
			result = node.Val
		}
		dfs(node.Right)
	}
	dfs(root)
	return result
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
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right) // 右
		// 中
		s += node.Val
		node.Val = s
		dfs(node.Left)
	}
	dfs(root)
	return root
}

// 98.验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/description/
// 二叉搜索树定义如下：
// 节点的左子树只包含 严格小于 当前节点的数。
// 节点的右子树只包含 严格大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var dfs func(*TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !dfs(root.Left) {
			return false
		}
		if prev != nil && root.Val <= prev.Val {
			return false
		}
		prev = root
		return dfs(root.Right)
	}
	return dfs(root)
}

// 700.二叉搜索树中的搜索
// https://leetcode.cn/problems/search-in-a-binary-search-tree/description/
// 迭代法
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

// 701.二叉搜索树中的插入操作
// https://leetcode.cn/problems/insert-into-a-binary-search-tree/description/
// 递归法
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

// 迭代法
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	prev := root
	cur := root
	for cur != nil {
		prev = cur
		if cur.Val < val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	node := &TreeNode{Val: val}
	if prev.Val < val {
		prev.Right = node
	} else {
		prev.Left = node
	}
	return root
}

// 450.删除二叉搜索树中的节点
// https://leetcode.cn/problems/delete-node-in-a-bst/description/
// 输入：root = [5,3,6,2,4,null,7], key = 3
// 输出：[5,4,6,2,null,null,7]
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		// 删除的是叶子节点
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Right == nil {
			return root.Left
		}
		// 右孩子继位，做孩子挂在右子树最左边
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		cur.Left = root.Left
		return root.Right
	}

	return root
}

// 669. 修剪二叉搜索树
// https://leetcode.cn/problems/trim-a-binary-search-tree/description/
// 给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。
// 输入：root = [1,0,2], low = 1, high = 2
// 输出：[1,null,2]
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		// 左边更小了，右子树中可能有，返回右子树中>low的节点
		return trimBST(root.Right, low, high)
	} else if root.Val > high {
		// 右边更大了，左子树中可能还有在区间内的节点
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, root.Val)
	root.Right = trimBST(root.Right, root.Val, high)
	return root
}

// 1373. 二叉搜索子树的最大键值和
// 给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。
// 二叉搜索树的定义如下：
// 任意节点的左子树中的键值都 小于 此节点的键值。
// 任意节点的右子树中的键值都 大于 此节点的键值。
// 任意节点的左子树和右子树都是二叉搜索树。
// 输入：root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
// 输出：20
// 解释：键值为 3 的子树是和最大的二叉搜索树。
func maxSumBST(root *TreeNode) int {
	var maxSum int
	var findMaxMinSum func(*TreeNode) []int
	// 计算以 root 为根的二叉树的最大值、最小值、节点和
	findMaxMinSum = func(root *TreeNode) []int {
		// base case
		if root == nil {
			return []int{1, math.MaxInt32, math.MinInt32, 0}
		}

		// 递归计算左右子树
		left := findMaxMinSum(root.Left)
		right := findMaxMinSum(root.Right)

		// ******* 后序位置 *******
		// 通过 left 和 right 推导返回值
		// 并且正确更新 maxSum 变量
		res := make([]int, 4)
		// 这个 if 在判断以 root 为根的二叉树是不是 BST
		if left[0] == 1 && right[0] == 1 &&
			root.Val > left[2] && root.Val < right[1] {
			// 以 root 为根的二叉树是 BST
			res[0] = 1
			// 计算以 root 为根的这棵 BST 的最小值
			res[1] = min(left[1], root.Val)
			// 计算以 root 为根的这棵 BST 的最大值
			res[2] = max(right[2], root.Val)
			// 计算以 root 为根的这棵 BST 所有节点之和
			res[3] = left[3] + right[3] + root.Val
			// 更新全局变量
			maxSum = max(maxSum, res[3])
		} else {
			// 以 root 为根的二叉树不是 BST
			res[0] = 0
			// 其他的值都没必要计算了，因为用不到
		}
		// ************************

		return res
	}

	findMaxMinSum(root)
	return maxSum
}
