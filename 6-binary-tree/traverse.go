package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 257. 二叉树的所有路径
// https://leetcode.cn/problems/binary-tree-paths/description/
// 输入：root = [1,2,3,null,5]
// 输出：["1->2->5","1->3"]
// 先序遍历
func binaryTreePaths(root *TreeNode) []string {
	var results []string
	var path []string
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 中
		path = append(path, fmt.Sprintf("%d", root.Val))
		if root.Left == nil && root.Right == nil {
			results = append(results, strings.Join(path, "->")) // 注意不能return，因为还要回溯
		}
		dfs(root.Left)  // 左
		dfs(root.Right) // 右
		path = path[:len(path)-1]
	}
	dfs(root)
	return results
}

// 129. 求根节点到叶节点数字之和
// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
// 每条从根节点到叶节点的路径都代表一个数字：
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
// 输入：root = [1,2,3]
// 输出：25
// 解释：
// 从根到叶子节点路径 1->2 代表数字 12
// 从根到叶子节点路径 1->3 代表数字 13
// 因此，数字总和 = 12 + 13 = 25
func sumNumbers(root *TreeNode) int {
	result := 0
	var path []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			s := 0
			for i := 0; i < len(path); i++ {
				s = 10*s + path[i]
			}
			result += s
		}
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return result
}

// 199. 二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	var results []int
	if root == nil {
		return results
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		results = append(results, q[len(q)-1].Val)
		var next []*TreeNode
		for _, node := range q {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		q = next
	}
	return results
}

// 988. 从叶结点开始的最小字符串
// 给定一颗根结点为 root 的二叉树，树中的每一个结点都有一个 [0, 25] 范围内的值，分别代表字母 'a' 到 'z'。
// 返回 按字典序最小 的字符串，该字符串从这棵树的一个叶结点开始，到根结点结束。
// 输入：root = [0,1,2,3,4,3,4]
// 输出："dba"
func smallestFromLeaf(root *TreeNode) string {
	var path []byte
	result := ""
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, byte(node.Val+'a'))
		if node.Left == nil && node.Right == nil {
			tmp := append([]byte{}, path...)
			reverse(tmp)
			if result == "" || string(tmp) < result {
				result = string(tmp)
			}
		}
		dfs(node.Left)
		dfs(node.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return result
}

func reverse(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 1022. 从根到叶的二进制数之和
// https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/description/
// 给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。
// 例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。
// 对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。
// 返回这些数字之和。题目数据保证答案是一个 32 位 整数。
// 输入：root = [1,0,1,0,1,0,1]
// 输出：22
// 解释：(100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
func sumRootToLeaf(root *TreeNode) int {
	var path []int
	result := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			s := 0
			for _, val := range path {
				s = 2*s + val
			}
			result += s
		}
		dfs(node.Left)
		dfs(node.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return result
}

// 1457. 二叉树中的伪回文路径
// https://leetcode.cn/problems/pseudo-palindromic-paths-in-a-binary-tree/
// 给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。
// 请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。
// 输入：root = [2,3,1,3,1,null,1]
// 输出：2
func pseudoPalindromicPaths(root *TreeNode) int {
	var path []int
	result := 0
	var hash [10]int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		hash[node.Val]++
		if node.Left == nil && node.Right == nil {
			cnt := 0
			for i := 1; i <= 9; i++ {
				if hash[i]%2 == 1 {
					cnt++
				}
			}
			if cnt <= 1 {
				result++
			}
		}
		dfs(node.Left)
		dfs(node.Right)
		path = path[:len(path)-1]
		hash[node.Val]--
	}
	dfs(root)
	return result
}

// 404.左叶子之和
// https://leetcode.cn/problems/sum-of-left-leaves/
// 输入: root = [3,9,20,null,null,15,7]
// 输出: 24
// 后序遍历
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	s := 0
	var dfs func(*TreeNode, bool)
	dfs = func(root *TreeNode, isLeft bool) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && isLeft {
			s += root.Val
		}
		dfs(root.Left, true)
		dfs(root.Right, false)
	}
	dfs(root, false)
	return s
}

// 递归，有左孩子时，判断一下是否是叶子节点
func sumOfLeftLeaves2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	s := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			return
		}
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			s += root.Left.Val
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return s
}

// 623. 在二叉树中增加一行
// https://leetcode.cn/problems/add-one-row-to-tree/
// 给定一个二叉树的根 root 和两个整数 val 和 depth ，在给定的深度 depth 处添加一个值为 val 的节点行。
// 注意，根节点 root 位于深度 1 。
// 加法规则如下:
// 给定整数 depth，对于深度为 depth - 1 的每个非空树节点 cur ，创建两个值为 val 的树节点作为 cur 的左子树根和右子树根。
// cur 原来的左子树应该是新的左子树根的左子树。
// cur 原来的右子树应该是新的右子树根的右子树。
// 如果 depth == 1 意味着 depth - 1 根本没有深度，那么创建一个树节点，值 val 作为整个原始树的新根，而原始树就是新根的左子树。
// 输入: root = [4,2,6,3,1,5], val = 1, depth = 2
// 输出: [4,1,1,2,null,null,6,3,1,5]
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == depth-1 {
			newLeft := &TreeNode{Val: val, Left: node.Left}
			newRight := &TreeNode{Val: val, Right: node.Right}
			node.Left = newLeft
			node.Right = newRight
			return
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dummy := &TreeNode{Left: root}
	dfs(dummy, 0)
	return dummy.Left
}

// 971. 翻转二叉树以匹配先序遍历
// 给你一棵二叉树的根节点 root ，树中有 n 个节点，每个节点都有一个不同于其他节点且处于 1 到 n 之间的值。
// 另给你一个由 n 个值组成的行程序列 voyage ，表示 预期 的二叉树 先序遍历 结果。
// 通过交换节点的左右子树，可以 翻转 该二叉树中的任意节点。例，翻转节点 1 的效果如下：
// 输入：root = [1,2], voyage = [2,1]
// 输出：[-1]
// 解释：翻转节点无法令先序遍历匹配预期行程。
func flipMatchVoyage(root *TreeNode, voyage []int) []int {

}
