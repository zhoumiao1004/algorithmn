package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
关于完全二叉树和满二叉树的定义，中文语境和英文语境似乎有点区别。
我们说的完全二叉树对应英文 Complete Binary Tree，这个没问题，说的是同一种树。
我们说的满二叉树，按理说应该翻译成 Full Binary Tree 对吧，但其实不是，满二叉树的定义对应英文的 Perfect Binary Tree。
而英文中的 Full Binary Tree 是指一棵二叉树的所有节点要么没有孩子节点，要么有两个孩子节点。*/

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
	// 前序位置：统计节点左右子树深度，判断以root为根的二叉树是不是完全二叉树
	var leftDepth, rightDepth int
	for cur := root; cur != nil; cur = cur.Left {
		leftDepth++
	}
	for cur := root; cur != nil; cur = cur.Right {
		rightDepth++
	}
	if leftDepth == rightDepth {
		return 1<<leftDepth - 1 // 以root为根的二叉树是完全二叉树，使用公式计算
	}
	leftNum := countNodes(root.Left)
	rightNum := countNodes(root.Right)
	// 后序位置
	return 1 + leftNum + rightNum
}

// 普通二叉树的节点个数
func countNodesNormal(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftNum := countNodes(root.Left)
	rightNum := countNodes(root.Right)
	return 1 + leftNum + rightNum
}

// 1104. 二叉树寻路
// https://leetcode.cn/problems/path-in-zigzag-labelled-binary-tree/description/
// 在一棵无限的二叉树上，每个节点都有两个子节点，树中的节点 逐行 依次按 “之” 字形进行标记。
// 如下图所示，在奇数行（即，第一行、第三行、第五行……）中，按从左到右的顺序进行标记；
// 而偶数行（即，第二行、第四行、第六行……）中，按从右到左的顺序进行标记。
// 考察完全二叉树性质：一层的最小和最大为 2^n 2*2^n-1
func pathInZigZagTree(label int) []int {
	var log func(x int) int
	var reverse func(nums []int)

	log = func(x int) int {
		return int(math.Log(float64(x)) / math.Log(2))
	}
	reverse = func(nums []int) {
		left, right := 0, len(nums)-1
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	var path []int
	for label >= 1 {
		path = append(path, label)
		label /= 2
		depth := log(label)
		fmt.Println(depth)
		minVal := int(math.Pow(2, float64(depth)))
		maxVal := 2*minVal - 1
		label = maxVal - (label - minVal)
	}
	reverse(path)
	return path
}

func main() {
	fmt.Println(math.Log(4) / math.Log(2))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(pathInZigZagTree(14))
}
