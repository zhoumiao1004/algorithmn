package main

import "math"

/* 有些题目，你按照拍脑袋的方式去做，可能发现需要在递归代码中调用其他递归函数计算字数的信息。一般来说，出现这种情况时你可以考虑用后序遍历的思维方式来优化算法，利用后序遍历传递子树的信息，避免过高的时间复杂度。 */

// 110. 平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/description/
// 对于树中的每个节点：左和右子树高度差不超过1
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := maxDepth(node.Left)
		if left == -1 {
			return -1
		}
		right := maxDepth(node.Right)
		if right == -1 {
			return -1
		}
		if left-right > 1 || left-right < -1 {
			return -1
		}
		return max(left, right) + 1
	}
	return maxDepth(root) != -1
}

func isBalanced2(root *TreeNode) bool {
	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		return max(left, right) + 1
	}
	if root == nil {
		return true
	}
	// 中
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if math.Abs(float64(left-right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right) // 左右
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
	maxSumCnt := make(map[int]int) // 记录和的次数

	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := traverse(root.Left)
		right := traverse(root.Right)
		// 后序位置
		s := root.Val + left + right
		maxSumCnt[s]++
		if maxSumCnt[s] == maxCnt {
			results = append(results, s)
		} else if maxSumCnt[s] > maxCnt {
			results = []int{s}
			maxCnt = maxSumCnt[s]
		}
		return s
	}

	traverse(root)
	return results
}

// 563. 二叉树的坡度
// https://leetcode.cn/problems/binary-tree-tilt/description/
// 给你一个二叉树的根节点 root ，计算并返回 整个树 的坡度 。
// 一个树的 节点的坡度 定义即为，该节点左子树的节点之和和右子树节点之和的 差的绝对值 。如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是 0 。
// 整个树 的坡度就是其所有节点的坡度之和。
func findTilt(root *TreeNode) int {
	result := 0
	var traverse func(root *TreeNode) int

	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := traverse(root.Left)
		right := traverse(root.Right)
		// 后序位置
		result += int(math.Abs(float64(left) - float64(right)))
		return left + right + root.Val
	}

	traverse(root)
	return result
}

// 814. 二叉树剪枝
// https://leetcode.cn/problems/binary-tree-pruning/description/
// 给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。
// 返回移除了所有不包含 1 的子树的原二叉树。
// 节点 node 的子树为 node 本身加上所有 node 的后代。
// 输入：root = [1,null,0,0,1]
// 输出：[1,null,0,null,1]
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := pruneTree(root.Left)
	right := pruneTree(root.Right)
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

// 687. 最长同值路径
// https://leetcode.cn/problems/longest-univalue-path/description/
// 给定一个二叉树的 root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。
// 两个节点之间的路径长度 由它们之间的边数表示。
// 输入：root = [5,4,5,1,1,5]
// 输出：2
// 思路:分解问题+后序,这题不适合用遍历的思路，因为需要从子树获取最长路径长度
func longestUnivaluePath(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	var maxLen func(root *TreeNode, parentVal int) int
	// 定义：计算以 root 为根的这棵二叉树中，从 root 开始值为 parentVal 的最长树枝长度
	maxLen = func(root *TreeNode, parentVal int) int {
		if root == nil {
			return 0
		}
		// 利用函数定义，计算左右子树值为 root.val 的最长树枝长度
		leftLen := maxLen(root.Left, root.Val)
		rightLen := maxLen(root.Right, root.Val)

		// 后序遍历位置顺便更新全局变量
		// 同值路径就是左右同值树枝长度之和
		res = max(res, leftLen+rightLen)
		// 如果 root 本身和上级值不同，那么整棵子树都不可能有同值树枝
		if root.Val != parentVal {
			return 0
		}
		// 实现函数的定义：
		// 以 root 为根的二叉树从 root 开始值为 parentVal 的最长树枝长度
		// 等于左右子树的最长树枝长度的最大值加上 root 节点本身
		return 1 + max(leftLen, rightLen)
	}
	// 在后序遍历的位置更新 res
	maxLen(root, root.Val)
	return res
}

// 865. 具有所有最深节点的最小子树
// https://leetcode.cn/problems/smallest-subtree-with-all-the-deepest-nodes/description/
// 1123. 最深叶节点的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/description/
// 给定一个根为 root 的二叉树，每个节点的深度是 该节点到根的最短距离 。
// 返回包含原始树中所有 最深节点 的 最小子树 。
// 如果一个节点在 整个树 的任意节点之间具有最大的深度，则该节点是 最深的 。
// 一个节点的 子树 是该节点加上它的所有后代的集合。
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4]
// 输出：[2,7,4]
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	node, _ := maxDepthNode(root)
	return node
}

func maxDepthNode(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}
	left, depth1 := maxDepthNode(root.Left)
	right, depth2 := maxDepthNode(root.Right)
	if depth1 == depth2 {
		return root, depth1 + 1
	} else if depth1 < depth2 {
		return right, depth2 + 1
	}
	return left, depth1 + 1
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
	getMinMax(root, &res)
	return res
}

// 定义：输入一棵二叉树，返回该二叉树中节点的最小值和最大值，
// 第一个元素是最小值，第二个值是最大值
func getMinMax(root *TreeNode, res *int) (int, int) {
	if root == nil {
		return int(^uint(0) >> 1), -int(^uint(0)>>1) - 1 // Integer.MAX_VALUE, Integer.MIN_VALUE in Go
	}
	leftMin, leftMax := getMinMax(root.Left, res)
	rightMin, rightMax := getMinMax(root.Right, res)
	// 以 root 为根的这棵树的最大值和最小值可以通过左右子树的最大最小值推导出来
	rootMin := min3(root.Val, leftMin, rightMin)
	rootMax := max3(root.Val, leftMax, rightMax)
	// 在后序位置顺便判断所有差值的最大值
	*res = max3(*res, rootMax-root.Val, root.Val-rootMin)

	return rootMin, rootMax
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func max3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}
