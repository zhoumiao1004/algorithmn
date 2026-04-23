package main

import (
	"fmt"
	"math"
)

/* 有些题目，你按照拍脑袋的方式去做，可能发现需要在递归代码中调用其他递归函数计算字数的信息。
一般来说，出现这种情况时你可以考虑用后序遍历的思维方式来优化算法，利用后序遍历传递子树的信息，避免过高的时间复杂度。
前序位置的代码只能从函数参数中获取父节点传递来的数据，而后序位置的代码不仅可以获取参数数据，还可以获取到子树通过函数返回值传递回来的数据。
一旦你发现题目和子树有关，那大概率要给函数设置合理的定义和返回值，在后序位置写代码了。
*/

// 652. 寻找重复的子树
// https://leetcode.cn/problems/find-duplicate-subtrees/description/
// 给你一棵二叉树的根节点 root ，返回所有 重复的子树 。
// 对于同一类的重复子树，你只需要返回其中任意 一棵 的根结点即可。
// 如果两棵树具有 相同的结构 和 相同的结点值 ，则认为二者是 重复 的。
// 输入：root = [1,2,3,4,null,2,4,null,null,4]
// 输出：[[2,4],[4]]
// 思路1: 后序
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	memo := make(map[string]int)
	var serialize func(node *TreeNode) string

	serialize = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		left := serialize(node.Left)
		right := serialize(node.Right)
		s := fmt.Sprintf("%d,%s,%s", node.Val, left, right)
		// 后序位置，顺便计算是否存在重复子树
		if memo[s] == 1 {
			res = append(res, node)
		}
		memo[s]++
		return s
	}

	serialize(root)
	return res
}

// 思路2: 遍历
func findDuplicateSubtrees3(root *TreeNode) []*TreeNode {
	var result []*TreeNode
	subMap := make(map[string]int)
	var serialize func(node *TreeNode) string
	var traverse func(node *TreeNode)

	serialize = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		left := serialize(node.Left)
		right := serialize(node.Right)
		return fmt.Sprintf("%d,%s,%s", node.Val, left, right)
	}

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		// 后序位置
		s := serialize(node)
		if subMap[s] == 1 {
			result = append(result, node)
		}
		subMap[s]++
	}

	traverse(root)
	return result
}

// 110. 平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/description/
// 对于树中的每个节点：左和右子树高度差不超过1
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
// 思路：分解问题
func isBalanced(root *TreeNode) bool {
	flag := true
	var maxDepth func(node *TreeNode) int // 明确函数定义：返回以 node 为根的子树的最大深度

	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := maxDepth(node.Left)
		right := maxDepth(node.Right)
		// 后序位置顺便判断是否平衡
		if math.Abs(float64(left-right)) > 1 {
			flag = false
		}
		return max(left, right) + 1
	}

	maxDepth(root)
	return flag
}

// 508. 出现次数最多的子树元素和
// https://leetcode.cn/problems/most-frequent-subtree-sum/
// 给你一个二叉树的根结点 root ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。
// 一个结点的 「子树元素和」 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。
// 输入: root = [5,2,-3]
// 输出: [2,-3,4]
func findFrequentTreeSum(root *TreeNode) []int {
	var results []int
	maxCnt := 0
	maxSumCnt := make(map[int]int)      // 记录和的次数
	var getSum func(node *TreeNode) int // 明确函数定义：返回以 node 为根的二叉树的元素和

	getSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getSum(node.Left)
		right := getSum(node.Right)
		// 后序位置顺便更新最大子树元素和
		s := node.Val + left + right
		maxSumCnt[s]++
		if maxSumCnt[s] == maxCnt {
			results = append(results, s)
		} else if maxSumCnt[s] > maxCnt {
			results = []int{s}
			maxCnt = maxSumCnt[s]
		}
		return s
	}

	getSum(root)
	return results
}

// 563. 二叉树的坡度
// https://leetcode.cn/problems/binary-tree-tilt/description/
// 给你一个二叉树的根节点 root ，计算并返回 整个树 的坡度 。
// 一个树的 节点的坡度 定义即为，该节点左子树的节点之和和右子树节点之和的 差的绝对值 。如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是 0 。
// 整个树 的坡度就是其所有节点的坡度之和。
func findTilt(root *TreeNode) int {
	result := 0
	var getSum func(node *TreeNode) int

	getSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getSum(node.Left)
		right := getSum(node.Right)
		// 后序位置顺便累加坡度和
		result += int(math.Abs(float64(left) - float64(right)))
		return left + right + node.Val
	}

	getSum(root)
	return result
}

// 814. 二叉树剪枝
// https://leetcode.cn/problems/binary-tree-pruning/description/
// 给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。
// 返回移除了所有不包含 1 的子树的原二叉树。
// 节点 node 的子树为 node 本身加上所有 node 的后代。
// 输入：root = [1,null,0,0,1]
// 输出：[1,null,0,null,1]
// 思路：分解问题，明确函数定义：返回以 root 为根的二叉树剪枝后的原二叉树
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := pruneTree(root.Left)   // 左子树剪枝
	right := pruneTree(root.Right) // 右子树剪枝
	// 后序位置
	if root.Val == 0 && left == nil && right == nil {
		return nil // return nil 相当于删除节点
	}
	root.Left = left   // 接住左子树
	root.Right = right // 接住右子树
	return root
}

// 1325. 删除给定值的叶子节点
// https://leetcode.cn/problems/delete-leaves-with-a-given-value/description/
// 给你一棵以 root 为根的二叉树和一个整数 target ，请你删除所有值为 target 的 叶子节点 。
// 注意，一旦删除值为 target 的叶子节点，它的父节点就可能变成叶子节点；如果新叶子节点的值恰好也是 target ，那么这个节点也应该被删除。
// 也就是说，你需要重复此过程直到不能继续删除。
// 输入：root = [1,2,3,2,null,2,4], target = 2
// 输出：[1,null,3,null,4]
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	left := removeLeafNodes(root.Left, target)
	right := removeLeafNodes(root.Right, target)
	// 后序位置
	if root.Val == target && left == nil && right == nil {
		return nil // return nil 相当于删除节点
	}
	root.Left = left   // 接住左子树
	root.Right = right // 接住右子树
	return root
}

// 1026. 节点与其祖先之间的最大差值
// https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/description/
// 给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。
// （如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）
// 输入：root = [8,3,10,1,6,null,14,null,null,4,7,13]
// 输出：7
// 0 <= Node.val <= 100000
func maxAncestorDiff(root *TreeNode) int {
	res := 0
	// 定义：输入一棵二叉树，返回该二叉树中节点的最小值和最大值，
	var getMinMax func(root *TreeNode) (int, int)

	getMinMax = func(root *TreeNode) (int, int) {
		if root == nil {
			return math.MaxInt, math.MinInt
		}
		leftMin, leftMax := getMinMax(root.Left)
		rightMin, rightMax := getMinMax(root.Right)

		// 后序位置
		rootMin := min(root.Val, leftMin, rightMin)
		rootMax := max(root.Val, leftMax, rightMax)
		// 在后序位置顺便判断所有差值的最大值
		res = max(res, rootMax-root.Val, root.Val-rootMin)

		return rootMin, rootMax
	}

	getMinMax(root)
	return res
}

// 1339. 分裂二叉树的最大乘积
// https://leetcode.cn/problems/maximum-product-of-splitted-binary-tree/description/
// 给你一棵二叉树，它的根为 root 。请你删除 1 条边，使二叉树分裂成两棵子树，且它们子树和的乘积尽可能大。
// 由于答案可能会很大，请你将结果对 10^9 + 7 取模后再返回。
// 输入：root = [1,2,3,4,5,6]
// 输出：110
// 解释：删除红色的边，得到 2 棵子树，和分别为 11 和 10 。它们的乘积是 110 （11*10）
func maxProduct(root *TreeNode) int {
	var result int64
	treeSum := getTreeSum(root)
	var getSum func(node *TreeNode, treeSum int) int // 明确函数定义：返回以 node 为根的子树的元素和

	getSum = func(node *TreeNode, treeSum int) int {
		if node == nil {
			return 0
		}
		left := getSum(node.Left, treeSum)
		right := getSum(node.Right, treeSum)
		rootSum := node.Val + left + right
		result = max(result, int64(rootSum)*(int64(treeSum)-int64(rootSum)))
		return rootSum
	}

	getSum(root, treeSum)
	return int(result % (1e9 + 7))
}

func getTreeSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := getTreeSum(node.Left)
	right := getTreeSum(node.Right)
	return node.Val + left + right
}

func main() {
	fmt.Println("hello world")
	// maxAncestorDiff()
}
