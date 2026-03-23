package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144. 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
// 递归：遍历的思路
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val) // 中
		dfs(node.Left)                    // 左
		dfs(node.Right)                   // 右
	}
	dfs(root)
	return result
}

// 递归：分解的思路
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
func postorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}
	dfs(root)
	return result
}

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
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		result = append(result, node.Val)
		dfs(node.Right)
	}
	dfs(root)
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

// 429. N 叉树的层序遍历
// 给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
// 树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder2(root *Node) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	q := []*Node{root}
	for len(q) > 0 {
		var tmp []int
		var next []*Node
		for _, node := range q {
			tmp = append(tmp, node.Val)
			for _, c := range node.Children {
				next = append(next, c)
			}
		}
		result = append(result, tmp)
		q = next
	}
	return result
}

// 116. 填充每个节点的下一个右侧节点指针
// 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下
// func connect(root *Node) *Node {
// 	var dfs func(root *Node)
// 	dfs = func(root *Node) {
// 		if root == nil {
// 			return
// 		}
// 		if root.Left != nil {
// 			root.Left.Next = root.Right
// 		}
// 		if root.Right != nil {
// 			if root.Next != nil {
// 				root.Right.Next = root.Next.Left
// 			} else {
// 				root.Right.Next = nil
// 			}
// 		}
// 		dfs(root.Left)
// 		dfs(root.Right)
// 	}
// 	dfs(root)
// 	return root
// }

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/description/
// 后序遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)   // 左
	right := invertTree(root.Right) // 右
	root.Left = right
	root.Right = left
	return root
}

// 111.二叉树的最小深度
// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
// 输入：root = [3,9,20,null,null,15,7]
// 输出：2
// 后序遍历
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Right)
	rightDepth := minDepth(root.Left)
	if root.Left == nil && root.Right != nil {
		return 1 + rightDepth
	}
	if root.Right == nil && root.Left != nil {
		return 1 + leftDepth
	}

	return 1 + min(leftDepth, rightDepth)
}

// 222.完全二叉树的节点个数
// https://leetcode.cn/problems/count-complete-tree-nodes/description/
// 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
// 完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层（从第 0 层开始），则该层包含 1~ 2h 个节点。
// 输入：root = [1,2,3,4,5,6]
// 输出：6
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftDepth, rightDepth int
	for cur := root; cur.Left != nil; cur = cur.Left {
		leftDepth++
	}
	for cur := root; cur.Right != nil; cur = cur.Right {
		rightDepth++
	}
	if leftDepth == rightDepth {
		return 2<<leftDepth - 1
	}
	leftNum := countNodes(root.Left)
	rightNum := countNodes(root.Right)
	return 1 + leftNum + rightNum
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := 0, 0
	cur1, cur2 := root, root
	for cur1 != nil {
		cur1 = cur1.Left
		left++
	}
	for cur2 != nil {
		cur2 = cur2.Right
		right++
	}
	if left == right {
		return 1<<left - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
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

// 236. 二叉树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 后序遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root // 找到就向上返回
	}
	left := lowestCommonAncestor(root.Left, p, q)   // 左
	right := lowestCommonAncestor(root.Right, p, q) // 右
	// 中
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// 235. 二叉搜索树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// 输出: 6
// 解释: 节点 2 和节点 8 的最近公共祖先是 6。
func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root // 找到就向上返回
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorBST(root.Left, q, p)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorBST(root.Right, p, q)
	}
	return root
}

func main() {
	fmt.Println("Hello, World!")
}
