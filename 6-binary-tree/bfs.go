package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 102. 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
func levelOrder(root *TreeNode) [][]int {
	var results [][]int
	if root == nil {
		return results
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		var tmp []int
		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		results = append(results, tmp)
	}
	return results
}

// 107.二叉树的层次遍历 II
// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottom(root *TreeNode) [][]int {
	results := levelOrder(root)
	left, right := 0, len(results)-1
	for left < right {
		results[left], results[right] = results[right], results[left]
		left++
		right--
	}
	return results
}

// 429. N 叉树的层序遍历
// 给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
// 树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

func levelOrderNTree(root *NTreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	q := []*NTreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		var tmp []int
		for i := 0; i < sz; i++ {
			node := q[i]
			q = q[1:]
			tmp = append(tmp, node.Val)
			for _, c := range node.Children {
				q = append(q, c)
			}
		}
		result = append(result, tmp)
	}
	return result
}

// 103. 二叉树的锯齿形层序遍历
// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/submissions/
// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
func zigzagLevelOrder(root *TreeNode) [][]int {
	var results [][]int
	if root == nil {
		return results
	}
	flag := true
	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		var tmp []int
		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if !flag {
			l, r := 0, len(tmp)-1
			for l < r {
				tmp[l], tmp[r] = tmp[r], tmp[l]
				l++
				r--
			}
		}
		flag = !flag
		results = append(results, tmp)
	}
	return results
}

// 117. 填充每个节点的下一个右侧节点指针 II
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/description/
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
// 初始状态下，所有 next 指针都被设置为 NULL 。
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		sz := len(q)

		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if i != len(q)-1 {
				node.Next = q[i+1]
			}
		}
	}
	return root
}

// 662. 二叉树最大宽度
// https://leetcode.cn/problems/maximum-width-of-binary-tree/description/
// 给你一棵二叉树的根节点 root ，返回树的 最大宽度 。
// 树的 最大宽度 是所有层中最大的 宽度 。
// 每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。将这个二叉树视作与满二叉树结构相同，两端点间会出现一些延伸到这一层的 null 节点，这些 null 节点也计入长度。
// 题目数据保证答案将会在  32 位 带符号整数范围内。
func widthOfBinaryTree(root *TreeNode) int {
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
				q = append(q, &Pair{node: curNode.Left, id: 2 * curId})
			}
			if curNode.Right != nil {
				q = append(q, &Pair{node: curNode.Right, id: 2*curId + 1})
			}
		}
		result = max(result, end-start+1)
	}
	return result
}

type Pair struct {
	node *TreeNode
	id   int
}

// 515. 在每个树行中找最大值
// 给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
// 输入: root = [1,3,2,5,3,null,9]
// 输出: [1,3,9]
func largestValues(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		maxVal := math.MinInt
		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			maxVal = max(maxVal, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append(result, maxVal)
	}
	return result
}

// 637. 二叉树的层平均值
// https://leetcode.cn/problems/average-of-levels-in-binary-tree/description/
// 给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[3.00000,14.50000,11.00000]
// 解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	if root == nil {
		return result
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		s := 0
		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			s += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append(result, float64(s)/float64(len(q)))
	}
	return result
}

// 958. 二叉树的完全性检验
// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/description/
// 给你一棵二叉树的根节点 root ，请你判断这棵树是否是一棵 完全二叉树 。
// 在一棵 完全二叉树 中，除了最后一层外，所有层都被完全填满，并且最后一层中的所有节点都尽可能靠左。最后一层（第 h 层）中可以包含 1 到 2h 个节点。
// 输入：root = [1,2,3,4,5,6]
// 输出：true
// 解释：最后一层前的每一层都是满的（即，节点值为 {1} 和 {2,3} 的两层），且最后一层中的所有节点（{4,5,6}）尽可能靠左。
func isCompleteTree(root *TreeNode) bool {
	end := false
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if node == nil {
				end = true
			} else {
				if end == true {
					return false
				}
				q = append(q, node.Left)
				q = append(q, node.Right)
			}
		}
	}
	return true
}

// 1161. 最大层内元素和
// https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/
// 给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。
// 返回总和 最大 的那一层的层号 x。如果有多层的总和一样大，返回其中 最小 的层号 x。
// 输入：root = [1,7,0,7,-8,null,null]
// 输出：2
// 解释：
// 第 1 层各元素之和为 1，
// 第 2 层各元素之和为 7 + 0 = 7，
// 第 3 层各元素之和为 7 + -8 = -1，
// 所以我们返回第 2 层的层号，它的层内元素之和最大。
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := math.MinInt
	maxSumLevel := 0
	level := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		level++
		s := 0
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			s += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if s > maxSum {
			maxSum = s
			maxSumLevel = level
		}
	}
	return maxSumLevel
}

// 1302. 层数最深叶子节点的和
// https://leetcode.cn/problems/deepest-leaves-sum/
// 给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。
// 输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
// 输出：15
func deepestLeavesSum(root *TreeNode) int {
	result := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		s := 0
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			s += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = s
	}
	return result
}
