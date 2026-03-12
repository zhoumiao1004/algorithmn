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
		path = append(path, byte(node.Val + 'a'))
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