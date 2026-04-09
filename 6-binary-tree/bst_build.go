package main

// 108. 将有序数组转换为二叉搜索树
// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。
// 输入：nums = [-10,-3,0,5,9]
// 输出：[0,-3,9,-10,null,5]
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

// 109. 有序链表转换二叉搜索树
// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/description/
// 给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为 平衡 二叉搜索树。
// 输入: head = [-10,-3,0,5,9]
// 输出: [0,-3,9,-10,null,5]
// 解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head2 := slow.Next.Next
	root := &TreeNode{Val: slow.Next.Val}
	slow.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(head2)
	return root
}

// 1008. 前序遍历构造二叉搜索树
// https://leetcode.cn/problems/construct-binary-search-tree-from-preorder-traversal/description/
// 给定一个整数数组，它表示BST(即 二叉搜索树 )的 先序遍历 ，构造树并返回其根。
// 保证 对于给定的测试用例，总是有可能找到具有给定需求的二叉搜索树。
// 二叉搜索树 是一棵二叉树，其中每个节点， Node.left 的任何后代的值 严格小于 Node.val , Node.right 的任何后代的值 严格大于 Node.val。
// 二叉树的 前序遍历 首先显示节点的值，然后遍历Node.left，最后遍历Node.right。
// 输入：preorder = [8,5,1,7,10,12]
// 输出：[8,5,10,1,7,null,12]
func bstFromPreorder(preorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	val := preorder[0]
	if n == 1 {
		return &TreeNode{Val: val}
	}
	root := &TreeNode{Val: val}
	mid := 1
	for mid < n && preorder[mid] < val {
		mid++
	}
	root.Left = bstFromPreorder(preorder[1:mid])
	root.Right = bstFromPreorder(preorder[mid:])
	return root
}

// 1382.将二叉搜索树变平衡
// https://leetcode.cn/problems/balance-a-binary-search-tree/description/
// 给你一棵二叉搜索树，请你返回一棵 平衡后 的二叉搜索树，新生成的树应该与原来的树有着相同的节点值。如果有多种构造方法，请你返回任意一种。
// 如果一棵二叉搜索树中，每个节点的两棵子树高度差不超过 1 ，我们就称这棵二叉搜索树是 平衡的 。
// 输入：root = [1,null,2,null,3,null,4,null,null]
// 输出：[2,1,3,null,null,null,4]
// 解释：这不是唯一的正确答案，[3,1,4,null,2,null,null] 也是一个可行的构造方案。
func balanceBST(root *TreeNode) *TreeNode {
	var nums []int
	var traverse func(root *TreeNode)
	var buildTree func(nums []int) *TreeNode

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		nums = append(nums, root.Val)
		traverse(root.Right)
	}

	buildTree = func(nums []int) *TreeNode {
		n := len(nums)
		if n == 0 {
			return nil
		} else if n == 1 {
			return &TreeNode{Val: nums[0]}
		}
		root := &TreeNode{Val: nums[n/2]}
		root.Left = buildTree(nums[:n/2])
		root.Right = buildTree(nums[n/2+1:])
		return root
	}

	traverse(root)
	return buildTree(nums)
}
