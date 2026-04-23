package main

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/description/
// 思路：分解问题+后序
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)   // 左
	right := invertTree(root.Right) // 右
	// 后序位置
	root.Left, root.Right = right, left
	return root
}

// 思路：遍历
func invertTree2(root *TreeNode) *TreeNode {
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		node.Left, node.Right = node.Right, node.Left // 前/后序都可以
	}

	if root == nil {
		return nil
	}
	traverse(root)
	return root
}

// 1110. 删点成林
// https://leetcode.cn/problems/delete-nodes-and-return-forest/description/
// 给出二叉树的根节点 root，树上每个节点都有一个不同的值。
// 如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。
// 返回森林中的每棵树。你可以按任意顺序组织答案。
// 输入：root = [1,2,3,4,5,6,7], to_delete = [3,5]
// 输出：[[1,2,null,4],[6],[7]]
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	var results []*TreeNode
	delSet := make(map[int]bool)
	for _, val := range to_delete {
		delSet[val] = true
	}
	var doDelete func(node *TreeNode, hasParent bool) *TreeNode // 明确函数定义：对以node为根的二叉树，删除集合中的节点，hasParent用来传递父节点有没有被删除

	doDelete = func(node *TreeNode, hasParent bool) *TreeNode {
		if node == nil {
			return nil
		}
		deleted := delSet[node.Val] // 标记node的值在不在删除集合中
		if !deleted && !hasParent {
			results = append(results, node) // node不删除，父节点被删除了node就成了一颗新树
		}
		node.Left = doDelete(node.Left, !deleted)
		node.Right = doDelete(node.Right, !deleted)
		if deleted {
			return nil // 被删除，返回nil给父节点
		}
		return node
	}

	if root == nil {
		return results
	}
	doDelete(root, false) // root初始化为父节点被删除是为了把root加入结果列表
	return results
}

// 技巧1:类似于判断镜像二叉树、翻转二叉树的问题，一般也可以用分解问题的思路，无非就是把整棵树的问题（原问题）分解成子树之间的问题（子问题）。

// 100. 相同的树
// https://leetcode.cn/problems/same-tree/
// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
// 输入：p = [1,2,3], q = [1,2,3]
// 输出：true
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 101. 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
func isSymmetric(root *TreeNode) bool {
	// 定义check：返回两个子树是否对称
	var check func(p, q *TreeNode) bool

	check = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}
		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}

	if root == nil {
		return false
	}
	return check(root.Left, root.Right)
}

// 951. 翻转等价二叉树
// https://leetcode.cn/problems/flip-equivalent-binary-trees/
// 我们可以为二叉树 T 定义一个 翻转操作 ，如下所示：选择任意节点，然后交换它的左子树和右子树。
// 只要经过一定次数的翻转操作后，能使 X 等于 Y，我们就称二叉树 X 翻转 等价 于二叉树 Y。
// 这些树由根节点 root1 和 root2 给出。如果两个二叉树是否是翻转 等价 的树，则返回 true ，否则返回 false 。
// 输入：root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
// 输出：true
// 解释：我们翻转值为 1，3 以及 5 的三个节点。
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}
	if root1.Val != root2.Val {
		return false
	}
	unflip := flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)
	flip := flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
	return unflip || flip
}

// 1609. 奇偶树
// https://leetcode.cn/problems/even-odd-tree/
// 如果一棵二叉树满足下述几个条件，则可以称为 奇偶树 ：
// 二叉树根节点所在层下标为 0 ，根的子节点所在层下标为 1 ，根的孙节点所在层下标为 2 ，依此类推。
// 偶数下标 层上的所有节点的值都是 奇 整数，从左到右按顺序 严格递增
// 奇数下标 层上的所有节点的值都是 偶 整数，从左到右按顺序 严格递减
// 给你二叉树的根节点，如果二叉树为 奇偶树 ，则返回 true ，否则返回 false 。
// 输入：root = [1,10,4,3,null,7,9,12,8,6,null,null,2]
// 输出：true
// 解释：每一层的节点值分别是：
// 0 层：[1]
// 1 层：[10,4]
// 2 层：[3,7,9]
// 3 层：[12,8,6,2]
// 由于 0 层和 2 层上的节点值都是奇数且严格递增，而 1 层和 3 层上的节点值都是偶数且严格递减，因此这是一棵奇偶树。
func isEvenOddTree(root *TreeNode) bool {
	return true
}
