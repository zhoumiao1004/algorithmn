package main

import (
	"fmt"
	"math"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144. 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
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
		var next []*TreeNode
		var tmp []int
		for _, node := range q {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		results = append(results, tmp)
		q = next
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
		s := 0
		var next []*TreeNode
		for _, node := range q {
			s += node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		result = append(result, float64(s)/float64(len(q)))
		q = next
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
		maxVal := math.MinInt
		var next []*TreeNode
		for _, node := range q {
			maxVal = max(maxVal, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		result = append(result, maxVal)
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

// 117. 填充每个节点的下一个右侧节点指针 II
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
// 初始状态下，所有 next 指针都被设置为 NULL 。
// func connect(root *Node) *Node {
//     if root == nil {
//         return nil
//     }
// 	q := []*Node{root}
//     for len(q) > 0 {
//         var next []*Node
//         for i, node := range q {
//             if node.Left != nil {
//                 next = append(next, node.Left)
//             }
//             if node.Right != nil {
//                 next = append(next, node.Right)
//             }
//             if i != len(q)-1 {
//                 node.Next = q[i+1]
//             }
//         }
//         q = next
//     }
//     return root
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

// 101. 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
// 后序遍历
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

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 后序遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return 1 + max(leftDepth, rightDepth)
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

// 110. 平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/description/
// 对于树中的每个节点：左和右子树高度差不超过1
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
// 后序遍历
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getDepth(node.Left)
		if left == -1 {
			return -1
		}
		right := getDepth(node.Right)
		if right == -1 {
			return -1
		}
		if left-right > 1 || left-right < -1 {
			return -1
		}
		return max(left, right) + 1
	}
	return getDepth(root) != -1
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

// 513.找树左下角的值
// https://leetcode.cn/problems/find-bottom-left-tree-value/description/
// 给定一个二叉树，在树的最后一行找到最左边的值。
// 输入: [1,2,3,4,null,5,6,null,null,7]
// 输出: 7
// 方法1:层序遍历
func findBottomLeftValue(root *TreeNode) int {
	result := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		var next []*TreeNode
		result = q[0].Val
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
	return result
}

// 方法2:递归+回溯
func findBottomLeftValue2(root *TreeNode) int {
	maxDepth := 0 // 最大深度
	result := 0   // 最大深度最左节点的值
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		// 终止条件：遍历到叶子节点
		if node.Left == nil && node.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				result = node.Val
			}
			return
		}
		if node.Left != nil {
			depth++
			dfs(node.Left, depth)
			depth--
		}
		if node.Right != nil {
			depth++
			dfs(node.Right, depth)
			depth--
		}
	}
	dfs(root, 1)
	return result
}

func findBottomLeftValue3(root *TreeNode) int {
	maxDepth := 0
	depth := 0
	result := 0
	if root == nil {
		return result
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		depth++
		if depth > maxDepth {
			maxDepth = depth
			result = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
		depth--
	}
	dfs(root)
	return result
}

// 112. 路径总和
// https://leetcode.cn/problems/path-sum/description/
// pathSum 返回从根节点到叶子节点路径总和等于给定目标和的1个路径
// 不涉及中的操作，所以前中后序遍历都可以
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
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

// 617. 合并二叉树
// https://leetcode.cn/problems/merge-two-binary-trees/description/
// 给你两棵二叉树： root1 和 root2 。
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
// 返回合并后的二叉树。
// 注意: 合并过程必须从两个树的根节点开始。
// 输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// 输出：[3,4,5,5,4,null,7]
// 前序遍历
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	}
	root1.Val += root2.Val                             // 中
	root1.Left = mergeTrees(root1.Left, root2.Left)    // 左
	root1.Right = mergeTrees(root1.Right, root2.Right) // 右
	return root1
}

// 700.二叉搜索树中的搜索
// https://leetcode.cn/problems/search-in-a-binary-search-tree/description/
// 迭代法
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

// 98.验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/description/
// 二叉搜索树定义如下：
// 节点的左子树只包含 严格小于 当前节点的数。
// 节点的右子树只包含 严格大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var dfs func(*TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !dfs(root.Left) {
			return false
		}
		if prev != nil && root.Val <= prev.Val {
			return false
		}
		prev = root
		return dfs(root.Right)
	}
	return dfs(root)
}

// 530. 二叉搜索树的最小绝对差
// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/
func getMinimumDifference(root *TreeNode) int {
	result := math.MaxInt
	var prev *TreeNode
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if prev != nil {
			result = root.Val - prev.Val
		}
		prev = root
		dfs(root.Right)
	}
	dfs(root)
	return result
}

// 501.二叉搜索树中的众数
// https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/
func findMode(root *TreeNode) []int {
	var results []int
	maxCnt := 0
	cnt := 0
	var prev *TreeNode
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if prev != nil && prev.Val == root.Val {
			cnt++
		} else {
			cnt = 1
		}
		if cnt == maxCnt {
			results = append(results, root.Val)
		} else if cnt > maxCnt {
			results = []int{root.Val}
			maxCnt = cnt
		}
		prev = root
		dfs(root.Right)
	}
	dfs(root)
	return results
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

// 701.二叉搜索树中的插入操作
// https://leetcode.cn/problems/insert-into-a-binary-search-tree/description/
// 递归法
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

// 迭代法
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	prev := root
	cur := root
	for cur != nil {
		prev = cur
		if cur.Val < val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	node := &TreeNode{Val: val}
	if prev.Val < val {
		prev.Right = node
	} else {
		prev.Left = node
	}
	return root
}

// 450.删除二叉搜索树中的节点
// https://leetcode.cn/problems/delete-node-in-a-bst/description/
// 输入：root = [5,3,6,2,4,null,7], key = 3
// 输出：[5,4,6,2,null,null,7]
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		// 删除的是叶子节点
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Right == nil {
			return root.Left
		}
		// 右孩子继位，做孩子挂在右子树最左边
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		cur.Left = root.Left
		return root.Right
	}

	return root
}

// 669. 修剪二叉搜索树
// https://leetcode.cn/problems/trim-a-binary-search-tree/description/
// 给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。
// 输入：root = [1,0,2], low = 1, high = 2
// 输出：[1,null,2]
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		// 左边更小了，右子树中可能有，返回右子树中>low的节点
		return trimBST(root.Right, low, high)
	} else if root.Val > high {
		// 右边更大了，左子树中可能还有在区间内的节点
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, root.Val)
	root.Right = trimBST(root.Right, root.Val, high)
	return root
}

// 108. 将有序数组转换为二叉搜索树
// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。
// 输入：nums = [-10,-3,0,5,9]
// 输出：[0,-3,9,-10,null,5]
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

// 538. 把二叉搜索树转换为累加树
func convertBST(root *TreeNode) *TreeNode {
	prev := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)  // 右
		node.Val += prev // 中
		prev = node.Val
		dfs(node.Left) // 左
	}
	dfs(root)
	return root
}

func convertBST2(root *TreeNode) *TreeNode {
	var prev *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		if prev != nil {
			root.Val += prev.Val
		}
		prev = root
		dfs(root.Left)
	}
	dfs(root)
	return root
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 3}}
	//fmt.Println(maxPath(root))
	fmt.Println(maxPath2(root))
	fmt.Println("Hello, World!")
}

func pathSum(root *TreeNode, targetSum int) []*TreeNode {
	var ans []*TreeNode
	// dfs 深度优先搜索
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, targetSum int) bool {
		if node == nil {
			return false
		}

		// 将当前节点添加到路径中
		ans = append(ans, node)

		// 检查是否到达叶子节点且路径总和等于目标和
		if node.Left == nil && node.Right == nil && node.Val == targetSum {
			return true
		}

		// 继续搜索左子树和右子树
		//if dfs(node.Left, targetSum-node.Val) || dfs(node.Right, targetSum-node.Val) {
		if dfs(node.Right, targetSum-node.Val) || dfs(node.Left, targetSum-node.Val) {
			return true
		}
		ans = ans[:len(ans)-1]
		return false
	}
	dfs(root, targetSum)
	return ans
}

// pathSum2 返回从根节点到叶子节点路径总和等于给定目标和的所有路径
func pathSum2(root *TreeNode, targetSum int) [][]*TreeNode {
	var ans [][]*TreeNode
	var path []*TreeNode
	// dfs 深度优先搜索
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}

		// 将当前节点添加到路径中
		path = append(path, root)

		// 检查是否到达叶子节点且路径总和等于目标和
		if root.Left == nil && root.Right == nil && root.Val == targetSum {
			//ans = append(ans, pre)
			var tmp []*TreeNode
			for _, v := range path {
				tmp = append(tmp, v)
			}
			ans = append(ans, tmp)
		}

		// 继续搜索左子树和右子树
		dfs(root.Left, targetSum-root.Val)
		dfs(root.Right, targetSum-root.Val)
		path = path[0 : len(path)-1]
	}
	dfs(root, targetSum)
	return ans
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

// 1382.将二叉搜索树变平衡
// https://leetcode.cn/problems/balance-a-binary-search-tree/description/
// 给你一棵二叉搜索树，请你返回一棵 平衡后 的二叉搜索树，新生成的树应该与原来的树有着相同的节点值。如果有多种构造方法，请你返回任意一种。
// 如果一棵二叉搜索树中，每个节点的两棵子树高度差不超过 1 ，我们就称这棵二叉搜索树是 平衡的 。
// 输入：root = [1,null,2,null,3,null,4,null,null]
// 输出：[2,1,3,null,null,null,4]
// 解释：这不是唯一的正确答案，[3,1,4,null,2,null,null] 也是一个可行的构造方案。
func balanceBST(root *TreeNode) *TreeNode {
	var nums []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nums = append(nums, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	var buildTree func(nums []int) *TreeNode
	buildTree = func(nums []int) *TreeNode {
		n := len(nums)
		if n == 0 {
			return nil
		} else if n == 1 {
			return &TreeNode{Val: nums[0]}
		}
		root := &TreeNode{Val: nums[n/2]}
		root.Left = buildTree(nums[:n/2])
		root.Right = buildTree(nums[n/2+1:])
		return root
	}
	return buildTree(nums)
}

func maxPath(root *TreeNode) []int {
	var ans []int
	var path []int
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			if len(path) > len(ans) {
				var tmp []int
				tmp = append(tmp, path...)
				ans = tmp
			}
		}
		dfs(root.Left)
		dfs(root.Right)
		path = path[0 : len(path)-1]
	}
	dfs(root)
	return ans
}

func maxPath2(root *TreeNode) []int {
	var ans []int
	var dfs func(*TreeNode, []int)
	dfs = func(root *TreeNode, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			if len(path) > len(ans) {
				var tmp []int
				for _, v := range path {
					tmp = append(tmp, v)
				}
				ans = tmp
			}
		}
		dfs(root.Left, path)
		dfs(root.Right, path)
	}
	dfs(root, []int{})
	return ans
}
