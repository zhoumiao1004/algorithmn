package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
// 分析：分解+后序
func buildTreeFromPreInOrder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	// 中
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for inorder[i] != preorder[0] {
		i++
	}
	root.Left = buildTreeFromPreInOrder(preorder[1:i+1], inorder[:i])
	root.Right = buildTreeFromPreInOrder(preorder[i+1:], inorder[i+1:])
	return root
}

// 106. 从中序与后序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 输出：[3,9,20,null,null,15,7]
// 分析：分解+后序，先构造左子树和右子树，再构造根节点
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	} else if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0]}
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	// 找到根节点在中序列表中的位置
	i := 0
	for inorder[i] != root.Val {
		i++
	}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}

// 889. 根据前序和后序遍历构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/
// 给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder 是同一棵树的后序遍历，重构并返回二叉树。
// 如果存在多个答案，您可以返回其中 任何 一个。
// 输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
// 输出：[1,2,3,4,5,6,7]
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	} else if n == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	// 找到左右子树分界线
	m := make(map[int]int)
	for i := 0; i < n-1; i++ {
		m[postorder[i]] = i
	}
	root := &TreeNode{Val: preorder[0]}
	mid := m[preorder[1]]
	root.Left = constructFromPrePost(preorder[1:mid+2], postorder[:mid+1])
	root.Right = constructFromPrePost(preorder[mid+2:], postorder[mid+1:n-1])
	return root
}
