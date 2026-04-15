package main

// 700.二叉搜索树中的搜索
// https://leetcode.cn/problems/search-in-a-binary-search-tree/description/
// 明确函数定义：以 root 为根的二叉树中，返回值为val的节点
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
// 注意：一旦涉及「改」，就类似二叉树的构造问题，函数要返回 TreeNode 类型，并且要对递归调用的返回值进行接收。
// 明确函数定义：返回以 root 节点为根的二叉树中，返回插入val后的根节点
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val) // 返回右子树中，插入val后的根节点
	} else {
		root.Left = insertIntoBST(root.Left, val) // 返回左子树中，插入val后的根节点
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
// 思路：分解问题，明确函数定义：返回以 root 为根节点的bst，修剪后节点值在[low..high]范围内的子树的根节点
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high) // 左边更小了，右子树中可能有，返回右子树中>low的节点
	} else if root.Val > high {
		return trimBST(root.Left, low, high) // 右边更大了，左子树中可能还有在区间内的节点
	}
	root.Left = trimBST(root.Left, low, root.Val)    // 这里用low, high也行
	root.Right = trimBST(root.Right, root.Val, high) // 这里范围用low, high也行
	return root
}
