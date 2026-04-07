package main

import "math"

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

// 662. 二叉树最大宽度
// https://leetcode.cn/problems/maximum-width-of-binary-tree/description/
// 给你一棵二叉树的根节点 root ，返回树的 最大宽度 。
// 树的 最大宽度 是所有层中最大的 宽度 。
// 每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。将这个二叉树视作与满二叉树结构相同，两端点间会出现一些延伸到这一层的 null 节点，这些 null 节点也计入长度。
// 题目数据保证答案将会在  32 位 带符号整数范围内。
func widthOfBinaryTree(root *TreeNode) int {
	type Pair struct {
		node *TreeNode
		id   int
	}

	if root == nil {
		return 0
	}
	result := 0
	q := []*Pair{{node: root, id: 1}}
	for len(q) > 0 {
		n := len(q)
		start, end := 0, 0
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]
			curNode := cur.node
			curId := cur.id
			if i == 0 {
				start = curId
			}
			if i == n-1 {
				end = curId
			}
			if curNode.Left != nil {
				q = append(q, &Pair{node: curNode.Left, id: 2 * curId}) // 完全二叉树，父节点i，左孩子=2*i，右孩子=2*i+1
			}
			if curNode.Right != nil {
				q = append(q, &Pair{node: curNode.Right, id: 2*curId + 1})
			}
		}
		result = max(result, end-start+1)
	}
	return result
}
