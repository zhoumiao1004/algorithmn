package main

// 653. 两数之和 IV - 输入二叉搜索树
// https://leetcode.cn/problems/two-sum-iv-input-is-a-bst/description/
// 给定一个二叉搜索树 root 和一个目标结果 k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 true。
// 输入: root = [5,3,6,2,4,null,7], k = 9
// 输出: true
// 思路1: 利用bst中序有序的特点，输出到数组+双指针
func findTarget(root *TreeNode, k int) bool {
	var nums []int
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		nums = append(nums, root.Val)
		traverse(root.Left)
		traverse(root.Right)
	}

	traverse(root) // 中序遍历bst，输出到有序数组中
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == k {
			return true
		} else if nums[left]+nums[right] < k {
			left++
		} else {
			right--
		}
	}
	return false
}

// 思路2: 分解问题（一般二叉树解法），明确函数定义：以 root 为根节点的二叉树，返回是否存在2个节点和为k。明确函数定义：返回以 node 节点为根的二叉树是否包含2个节点和为k
func findTarget2(root *TreeNode, k int) bool {
	m := make(map[int]bool)

	var check func(node *TreeNode) bool

	check = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		// 前序位置
		if m[k-node.Val] {
			return true
		}
		m[node.Val] = true
		return check(node.Left) || check(node.Right)
	}

	if root == nil {
		return false
	}
	return check(root)
}

// 思路3: 遍历 + hashmap（一般二叉树解法）
func findTarget3(root *TreeNode, k int) bool {
	result := false
	m := make(map[int]bool)
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		// 中序位置
		if m[k-node.Val] {
			result = true
		}
		m[node.Val] = true
		traverse(node.Right)
	}

	traverse(root)
	return result
}
