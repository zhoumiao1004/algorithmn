package main

import (
	"math"
	"strings"
)

// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
func buildTree2(preorder []int, inorder []int) *TreeNode {
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
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

// 106. 从中序与后序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 输出：[3,9,20,null,null,15,7]
// 先从postordre获取最后一个节点作为根节点，在inordre中找到所在位置
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	} else if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0]}
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
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

// 331. 验证二叉树的前序序列化
// 序列化二叉树的一种方法是使用 前序遍历 。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
// 输入: preorder = "9,3,4,#,#,1,#,#,2,#,6,#,#"
// 输出: true
func isValidSerialization(preorder string) bool {
	edge := 1
	for _, c := range strings.Split(preorder, ",") {
		if c == "#" {
			edge--
			if edge < 0 {
				return false
			}
		} else {
			edge--
			if edge < 0 {
				return false
			}
			edge += 2
		}
	}
	return edge == 0
}

// 998. 最大二叉树 II
// https://leetcode.cn/problems/maximum-binary-tree-ii/description/
// 假设 b 是 a 的副本，并在末尾附加值 val。题目数据保证 b 中的值互不相同。返回 Construct(b) 。
// 输入：root = [4,1,3,null,null,2], val = 5
// 输出：[5,4,null,1,3,null,null,2]
// 解释：a = [1,4,2,3], b = [1,4,2,3,5]
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		return &TreeNode{Val: val, Left: root}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
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
	if root == nil {
		return results
	}
	delSet := make(map[int]bool)
	for _, val := range to_delete {
		delSet[val] = true
	}
	var doDelete func(root *TreeNode, hasParent bool) *TreeNode
	doDelete = func(root *TreeNode, hasParent bool) *TreeNode {
		if root == nil {
			return nil
		}
		deleted := delSet[root.Val]
		if !deleted && !hasParent {
			results = append(results, root) // 父节点被删除了就成了一颗新树
		}
		root.Left = doDelete(root.Left, !deleted)
		root.Right = doDelete(root.Right, !deleted)
		if deleted {
			return nil // 被删除，返回nil给父节点
		}
		return root
	}
	doDelete(root, false)
	return results
}

// 技巧1:类似于判断镜像二叉树、翻转二叉树的问题，一般也可以用分解问题的思路，无非就是把整棵树的问题（原问题）分解成子树之间的问题（子问题）。

// 100. 相同的树
// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
// 输入：p = [1,2,3], q = [1,2,3]
// 输出：true
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 101. 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return compare(root.Left, root.Right)
}

func compare(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	outside := compare(left.Left, right.Right)
	inside := compare(left.Right, right.Left)
	return outside && inside
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(node1, node2 *TreeNode) bool
	dfs = func(node1, node2 *TreeNode) bool {
		if node1 == nil || node2 == nil {
			return node1 == node2
		}
		return node1.Val == node2.Val && dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
	}
	return dfs(root.Left, root.Right)
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
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	unflip := flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)
	flip := flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
	return unflip || flip
}

// 技巧2:一般来说，遍历的思维模式可以帮你寻找从根节点开始的符合条件的「树枝」，但在不限制起点必须是根节点的条件下，让你寻找符合条件的「树枝」，就需要用到分解问题的思维模式了。
// 124. 二叉树中的最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/
// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
// 路径和 是路径中各节点值的总和。
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
func maxPathSum(root *TreeNode) int {
	result := math.MinInt
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		// 后序位置顺便计算双边最大路径和
		maxSum := left + right + root.Val
		result = max(result, maxSum)
		// 返回单边最大路径和
		return max(left, right) + root.Val
	}
	dfs(root)
	return result
}
