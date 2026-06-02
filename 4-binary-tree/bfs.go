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
	var res [][]int
	if root == nil {
		return res
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
		res = append(res, tmp)
	}
	return res
}

// 107.二叉树的层次遍历 II
// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
func levelOrderBottom(root *TreeNode) [][]int {
	res := levelOrder(root)
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
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
	var res []int
	if root == nil {
		return res
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
		res = append(res, maxVal)
	}
	return res
}

// 637. 二叉树的层平均值
// https://leetcode.cn/problems/average-of-levels-in-binary-tree/description/
// 给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[3.00000,14.50000,11.00000]
// 解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		var sum float64
		for i := 0; i < sz; i++ {
			node := q[0]
			sum += float64(node.Val)
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, sum/float64(sz))
	}
	return res
}

// 958. 二叉树的完全性检验
// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/description/
// 给你一棵二叉树的根节点 root ，请你判断这棵树是否是一棵 完全二叉树 。
// 在一棵 完全二叉树 中，除了最后一层外，所有层都被完全填满，并且最后一层中的所有节点都尽可能靠左。最后一层（第 h 层）中可以包含 1 到 2h 个节点。
// 输入：root = [1,2,3,4,5,6]
// 输出：true
// 解释：最后一层前的每一层都是满的（即，节点值为 {1} 和 {2,3} 的两层），且最后一层中的所有节点（{4,5,6}）尽可能靠左。
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	hasNil := false
	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			if node == nil {
				hasNil = true
				continue
			}
			if hasNil {
				return false
			}
			q = append(q, node.Left)
			q = append(q, node.Right)
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
	res, depth, maxSum := 0, 0, math.MinInt
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		depth++
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
			res = depth
		}
	}
	return res
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
	if root == nil {
		return false
	}
	flag := true
	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			if flag && node.Val%2 == 0 {
				return false
			}
			if !flag && node.Val%2 == 1 {
				return false
			}
			if i < sz-1 {
				if flag && node.Val >= q[0].Val {
					return false
				}
				if !flag && node.Val <= q[0].Val {
					return false
				}
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		flag = !flag
	}
	return true
}

func isEvenOddTree2(root *TreeNode) bool {
	if root == nil {
		return false
	}
	q := []*TreeNode{root}
	level := 0
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			if level%2 == 0 {
				if node.Val%2 == 0 {
					return false
				}
				if i != sz-1 && node.Val >= q[0].Val {
					return false
				}
			} else {
				if node.Val%2 == 1 {
					return false
				}
				if i != sz-1 && node.Val <= q[0].Val {
					return false
				}
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		level++
	}
	return true
}

// 872. 叶子相似的树
// https://leetcode.cn/problems/leaf-similar-trees/description/
// 输入：root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
// 输出：true
// 思路1: 栈迭代 iterator
type LeafIterator struct {
	st []*TreeNode
}

func NewLeafIterator(root *TreeNode) LeafIterator {
	return LeafIterator{st: []*TreeNode{root}}
}
func (it *LeafIterator) HasNext() bool {
	return len(it.st) > 0
}
func (it *LeafIterator) Next() *TreeNode {
	for len(it.st) > 0 {
		cur := it.st[len(it.st)-1]
		it.st = it.st[:len(it.st)-1]
		if cur.Left == nil && cur.Right == nil {
			return cur
		}
		if cur.Left != nil {
			it.st = append(it.st, cur.Left)
		}
		if cur.Right != nil {
			it.st = append(it.st, cur.Right)
		}
	}
	return nil
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	it1 := NewLeafIterator(root1)
	it2 := NewLeafIterator(root2)
	for it1.HasNext() && it2.HasNext() {
		if it1.Next().Val != it2.Next().Val {
			return false
		}
	}
	return !it1.HasNext() && !it2.HasNext()
}

// 思路2：遍历两颗二叉树，对比叶子节点集合
func leafSimilar2(root1 *TreeNode, root2 *TreeNode) bool {

	var getLeafVal func(node *TreeNode) []int

	getLeafVal = func(node *TreeNode) []int {
		var res []int
		if node == nil {
			return res
		}
		if node.Left == nil && node.Right == nil {
			return []int{node.Val}
		}
		left := getLeafVal(node.Left)
		right := getLeafVal(node.Right)
		res = append(res, left...)
		res = append(res, right...)
		return res
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

func leafSimilar3(root1 *TreeNode, root2 *TreeNode) bool {
	var nums1, nums2 []int
	var traverse1, traverse2 func(root *TreeNode)

	traverse1 = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			nums1 = append(nums1, root.Val)
		}
		traverse1(root.Left)
		traverse1(root.Right)
	}
	traverse2 = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			nums2 = append(nums2, root.Val)
		}
		traverse2(root.Left)
		traverse2(root.Right)
	}

	traverse1(root1)
	traverse2(root2)
	if len(nums1) != len(nums2) {
		return false
	}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] != nums2[j] {
			return false
		}
		i++
		j++
	}
	return true
}

// 863. 二叉树中所有距离为 K 的结点
// https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/description/
// 给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k ，返回到目标结点 target 距离为 k 的所有结点的值的数组。
// 答案可以以 任何顺序 返回。
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var res []int
	if root == nil {
		return res
	}
	nodeToParent := make(map[int]*TreeNode)
	visited := make(map[int]bool)
	var traverse func(root *TreeNode, parent *TreeNode)
	traverse = func(root *TreeNode, parent *TreeNode) {
		if root == nil {
			return
		}
		nodeToParent[root.Val] = parent
		traverse(root.Left, root)
		traverse(root.Right, root)
	}

	traverse(root, nil)
	q := []*TreeNode{target}
	for k >= 0 && len(q) > 0 {
		sz := len(q)
		res = []int{}
		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			if visited[node.Val] {
				continue
			}
			visited[node.Val] = true
			res = append(res, node.Val)
			if nodeToParent[node.Val] != nil {
				q = append(q, nodeToParent[node.Val])
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		k--
	}
	if k >= 0 {
		return []int{}
	}
	return res
}

func distanceK2(root *TreeNode, target *TreeNode, k int) []int {
	nodeToParent := make(map[int]*TreeNode) // 记录值到父节点的映射
	var traverse func(root *TreeNode, parent *TreeNode)

	traverse = func(root *TreeNode, parent *TreeNode) {
		if root == nil {
			return
		}
		nodeToParent[root.Val] = parent
		traverse(root.Left, root)
		traverse(root.Right, root)
	}

	traverse(root, nil)
	q := []*TreeNode{target}
	visited := make(map[int]bool)
	visited[target.Val] = true
	var res []int
	for k > 0 && len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			visited[cur.Val] = true
			// 向父节点、左右子节点扩散
			if parent, ok := nodeToParent[cur.Val]; ok && parent != nil && !visited[parent.Val] {
				q = append(q, parent)
			}
			if cur.Left != nil && !visited[cur.Left.Val] {
				q = append(q, cur.Left)
			}
			if cur.Right != nil && !visited[cur.Right.Val] {
				q = append(q, cur.Right)
			}
		}
		k-- // 向外扩展一圈
	}
	for _, node := range q {
		res = append(res, node.Val)
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
		res = max(res, q[len(q)-1].Id-q[0].Id+1)
		for i := 0; i < sz; i++ {
			obj := q[0]
			q = q[1:]
			if obj.Node.Left != nil {
				q = append(q, &Pair{Node: obj.Node.Left, Id: obj.Id * 2})
			}
			if obj.Node.Right != nil {
				q = append(q, &Pair{Node: obj.Node.Right, Id: obj.Id*2 + 1})
			}
		}
	}

	return res
}

func main() {
	q := []*Node{nil}
	fmt.Println(len(q))
}
