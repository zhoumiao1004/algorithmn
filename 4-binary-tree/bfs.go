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

// 层序遍历3种写法
// 写法3：假设如果每条树枝的权重可以是任意值，现在让你层序遍历整棵树，打印每个节点的路径权重和，你会怎么做？
// 这样的话，同一层节点的路径权重和就不一定相同了，写法二这样只维护一个 depth 变量就无法满足需求了。
func levelOrder3(root *TreeNode) [][]int {
	type State struct {
		node  *TreeNode
		depth int
	}
	var result [][]int
	if root == nil {
		return result
	}
	q := []State{{root, 1}}

	for len(q) > 0 {
		var tmp []int
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			tmp = append(tmp, cur.node.Val)
			// 访问 cur 节点，同时知道它的路径权重和
			// fmt.Printf("depth = %d, val = %d\n", cur.depth, cur.node.Val)

			// 把 cur 的左右子节点加入队列
			if cur.node.Left != nil {
				q = append(q, State{cur.node.Left, cur.depth + 1})
			}
			if cur.node.Right != nil {
				q = append(q, State{cur.node.Right, cur.depth + 1})
			}
		}
		result = append(result, tmp)
	}
	return result
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
// https://leetcode.cn/problems/n-ary-tree-level-order-traversal/description/
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
			node := q[0]
			q = q[1:]
			tmp = append(tmp, node.Val)
			for _, c := range node.Children {
				q = append(q, c) // 注：无需判断c != nil, 因为Children里不可能有nil
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
		var tmp []int // 保存一层的元素列表
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

// 872. 叶子相似的树
// https://leetcode.cn/problems/leaf-similar-trees/description/
// 输入：root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
// 输出：true
// 思路：遍历两颗二叉树，对比叶子节点集合
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	var getLeafVal func(node *TreeNode) []int

	getLeafVal = func(node *TreeNode) []int {
		var result []int
		if node == nil {
			return result
		}
		if node.Left == nil && node.Right == nil {
			result = append(result, node.Val)
		}
		left := getLeafVal(node.Left)
		right := getLeafVal(node.Right)
		result = append(result, left...)
		result = append(result, right...)
		return result
	}

	nums1 := getLeafVal(root1)
	nums2 := getLeafVal(root2)
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

// 863. 二叉树中所有距离为 K 的结点
// https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/description/
// 给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k ，返回到目标结点 target 距离为 k 的所有结点的值的数组。
// 答案可以以 任何顺序 返回。
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parent := make(map[int]*TreeNode) // 记录值到父节点的映射

	var traverse func(root *TreeNode, parentNode *TreeNode)
	traverse = func(root *TreeNode, parentNode *TreeNode) {
		if root == nil {
			return
		}
		parent[root.Val] = parentNode
		traverse(root.Left, root)
		traverse(root.Right, root)
	}

	traverse(root, nil)
	q := []*TreeNode{target}
	visited := make(map[int]bool)
	visited[target.Val] = true
	dist := 0
	var res []int
	for len(q) > 0 {
		if dist == k {
			for _, node := range q {
				res = append(res, node.Val)
			}
			return res
		}
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			// 向父节点、左右子节点扩散
			if parentNode, ok := parent[cur.Val]; ok && parentNode != nil && !visited[parentNode.Val] {
				visited[parentNode.Val] = true
				q = append(q, parentNode)
			}
			if cur.Left != nil && !visited[cur.Left.Val] {
				visited[cur.Left.Val] = true
				q = append(q, cur.Left)
			}
			if cur.Right != nil && !visited[cur.Right.Val] {
				visited[cur.Right.Val] = true
				q = append(q, cur.Right)
			}
		}
		dist++ // 向外扩展一圈
	}
	return res
}

// 662. 二叉树最大宽度
// https://leetcode.cn/problems/maximum-width-of-binary-tree/description/
// 给你一棵二叉树的根节点 root ，返回树的 最大宽度 。
// 树的 最大宽度 是所有层中最大的 宽度 。
// 每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。将这个二叉树视作与满二叉树结构相同，两端点间会出现一些延伸到这一层的 null 节点，这些 null 节点也计入长度。
// 题目数据保证答案将会在  32 位 带符号整数范围内。
func widthOfBinaryTree(root *TreeNode) int {
	type Pair struct {
		Node *TreeNode
		Id   int
	}
	res := 0
	q := []*Pair{{Node: root, Id: 1}}
	for len(q) > 0 {
		sz := len(q)
		start, end := 0, 0
		for i := 0; i < sz; i++ {
			obj := q[0]
			q = q[1:]
			if start == 0 {
				start = obj.Id
			}
			end = obj.Id
			if obj.Node.Left != nil {
				q = append(q, &Pair{Node: obj.Node.Left, Id: 2 * obj.Id})
			}
			if obj.Node.Right != nil {
				q = append(q, &Pair{Node: obj.Node.Right, Id: 2*obj.Id + 1})
			}
		}
		res = max(res, end-start+1)
	}

	return res
}

func main() {
	q := []*Node{nil}
	fmt.Println(len(q))
}
