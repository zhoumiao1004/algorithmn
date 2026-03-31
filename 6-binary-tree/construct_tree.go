package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 654. 最大二叉树
// https://leetcode.cn/problems/maximum-binary-tree/description/
// 给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:
// 创建一个根节点，其值为 nums 中的最大值。
// 递归地在最大值 左边 的 子数组前缀上 构建左子树。
// 递归地在最大值 右边 的 子数组后缀上 构建右子树。
// 返回 nums 构建的 最大二叉树 。
// 输入：nums = [3,2,1,6,0,5]
// 输出：[6,3,5,null,2,0,null,null,1]
func constructMaximumBinaryTree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	} else if n == 1 {
		return &TreeNode{Val: nums[0]}
	}
	maxIndex := 0
	for i := 1; i < n; i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	root := &TreeNode{Val: nums[maxIndex]}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
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

// 114. 二叉树展开为链表
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
// 思路：分解问题+后序。由于题目要求返回原TreeNode，所以不能构建新树
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)  // 左
	flatten(root.Right) // 右

	// 后序位置
	if root.Left == nil {
		return
	}
	cur := root.Left
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = root.Right
	root.Right = root.Left
	root.Left = nil // 注意需要清空左节点
}

// 897. 递增顺序搜索树
// https://leetcode.cn/problems/increasing-order-search-tree/
// 给你一棵二叉搜索树的 root ，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
// 输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
// 输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
// 思路1:中序遍历bst
func increasingBST2(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	cur := dummy
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left) // 左

		// 中序位置
		cur.Right = &TreeNode{Val: node.Val}
		cur = cur.Right

		traverse(root.Right) // 右
	}

	traverse(root)
	return dummy.Right
}

// 思路2:分解问题+后序
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 左右子树拉平
	left := increasingBST(root.Left)   // 左
	root.Left = nil                    // 断掉左子树
	right := increasingBST(root.Right) // 右
	root.Right = right

	// 后序位置，把左
	if left == nil {
		return root
	}
	cur := left
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = root // 根节点挂到左子树最右边的节点上
	return left
}

// 617. 合并二叉树
// https://leetcode.cn/problems/merge-two-binary-trees/description/
// 给你两棵二叉树： root1 和 root2 。
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
// 返回合并后的二叉树。
// 注意: 合并过程必须从两个树的根节点开始。
// 输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// 输出：[3,4,5,5,4,null,7]
// 思路1：分解问题
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	}
	root1.Val += root2.Val                             // 中
	root1.Left = mergeTrees(root1.Left, root2.Left)    // 左
	root1.Right = mergeTrees(root1.Right, root2.Right) // 右
	return root1
}

// 思路2：遍历
func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var traverse func(p, q *TreeNode)

	traverse = func(p, q *TreeNode) {
		if p == nil || q == nil {
			return
		}
		if p.Left == nil && q.Left != nil {
			p.Left = q.Left
			q.Left = nil
		}
		if p.Right == nil && q.Right != nil {
			p.Right = q.Right
			q.Right = nil
		}
		p.Val += q.Val
		traverse(p.Left, q.Left)
		traverse(p.Right, q.Right)
	}

	if root1 == nil {
		return root2
	}
	traverse(root1, root2)
	return root1
}
