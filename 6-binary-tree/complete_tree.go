package main

import "math"

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

// 1104. 二叉树寻路
// https://leetcode.cn/problems/path-in-zigzag-labelled-binary-tree/description/
// 在一棵无限的二叉树上，每个节点都有两个子节点，树中的节点 逐行 依次按 “之” 字形进行标记。
// 如下图所示，在奇数行（即，第一行、第三行、第五行……）中，按从左到右的顺序进行标记；
// 而偶数行（即，第二行、第四行、第六行……）中，按从右到左的顺序进行标记。
// 考察完全二叉树性质：一层的最小和最大为 2^n 2*2^n-1
func pathInZigZagTree(label int) []int {

	var log func(x int) int
	var getLevelMinMax func(n int) (int, int)
	var reverse func(nums []int)

	log = func(x int) int { return int(math.Log(float64(x)) / math.Log(float64(2))) }

	getLevelMinMax = func(n int) (int, int) {
		p := int(math.Pow(2, float64(n)))
		return p, 2*p - 1
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
		minVal, maxVal := getLevelMinMax(depth)
		label = maxVal - (label - minVal)
	}
	reverse(path)
	return path
}
