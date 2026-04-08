package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 897. 递增顺序搜索树
// https://leetcode.cn/problems/increasing-order-search-tree/
// 给你一棵二叉搜索树的 root ，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
// 输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
// 输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
// 思路1: 遍历
func increasingBST2(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	cur := dummy
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left) // 左

		// 中序位置
		cur.Right = &TreeNode{Val: node.Val}
		cur = cur.Right

		traverse(root.Right) // 右
	}

	traverse(root)
	return dummy.Right
}

// 思路2: 分解问题+后序
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := increasingBST(root.Left)   // 左
	root.Left = nil                    // 断掉左子树
	right := increasingBST(root.Right) // 右
	root.Right = right

	// 后序位置
	if left == nil {
		return root
	}
	// 把 root 接到左子树最右边的节点上
	cur := left
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = root
	return left
}

// 538. 把二叉搜索树转换为累加树
// https://leetcode.cn/problems/convert-bst-to-greater-tree/
// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
// 提醒一下，二叉搜索树满足下列约束条件：
// 节点的左子树仅包含键 小于 节点键的节点。
// 节点的右子树仅包含键 大于 节点键的节点。
// 左右子树也必须是二叉搜索树。
// 注意：本题和 1038: https://leetcode.cn/problems/binary-search-tree-to-greater-sum-tree/ 相同
func convertBST(root *TreeNode) *TreeNode {
	s := 0
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Right) // 右
		// 中序位置
		s += node.Val
		node.Val = s
		traverse(node.Left) // 左
	}

	traverse(root)
	return root
}

// 98.验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/description/
// 二叉搜索树定义如下：
// 节点的左子树只包含 严格小于 当前节点的数。
// 节点的右子树只包含 严格大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	// 定义函数：返回以 node 节点为根的二叉树是不是bst
	var isBST func(node *TreeNode) bool

	isBST = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !isBST(node.Left) {
			return false
		}
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		return isBST(node.Right)
	}

	return isBST(root)
}

// 700.二叉搜索树中的搜索
// https://leetcode.cn/problems/search-in-a-binary-search-tree/description/
// 明确函数定义：以 root 为根的二叉树中，返回值为val的节点
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

// 701.二叉搜索树中的插入操作
// https://leetcode.cn/problems/insert-into-a-binary-search-tree/description/
// 明确函数定义：返回以 root 节点为根的二叉树中，返回插入val后的根节点
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val) // 返回右子树中，插入val后的根节点
	} else {
		root.Left = insertIntoBST(root.Left, val) // 返回左子树中，插入val后的根节点
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

// 1373. 二叉搜索子树的最大键值和

// 给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。
// 二叉搜索树的定义如下：
// 任意节点的左子树中的键值都 小于 此节点的键值。
// 任意节点的右子树中的键值都 大于 此节点的键值。
// 任意节点的左子树和右子树都是二叉搜索树。
// 输入：root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
// 输出：20
// 解释：键值为 3 的子树是和最大的二叉搜索树。
// 思路：分解问题，明确函数定义：返回以 root 为根的二叉树是不是bst、最大值、最小值、节点和
func maxSumBST(root *TreeNode) int {
	var maxSum int
	var findMaxMinSum func(*TreeNode) []int

	findMaxMinSum = func(root *TreeNode) []int {
		// base case
		if root == nil {
			return []int{1, math.MaxInt32, math.MinInt32, 0}
		}

		left := findMaxMinSum(root.Left)
		right := findMaxMinSum(root.Right)

		// 后序位置
		res := make([]int, 4)
		if left[0] == 1 && right[0] == 1 &&
			root.Val > left[2] && root.Val < right[1] {
			res[0] = 1                             // 以 root 为根的二叉树是不是 BST
			res[1] = min(left[1], root.Val)        // 以 root 为根的这棵 BST 的最小值
			res[2] = max(right[2], root.Val)       // 以 root 为根的这棵 BST 的最大值
			res[3] = left[3] + right[3] + root.Val // 以 root 为根的这棵 BST 所有节点之和
			maxSum = max(maxSum, res[3])           // 顺便统计节点之和的最大值
		} else {
			res[0] = 0 // 以 root 为根的二叉树不是 BST，其他的值都没必要计算了，因为用不到
		}
		return res
	}

	findMaxMinSum(root)
	return maxSum
}

// 99. 恢复二叉搜索树
// https://leetcode.cn/problems/recover-binary-search-tree/description/
// 给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
// 思路：遍历。中序遍历找不满足第一个和最后一个不满足有序的2个节点进行交换
func recoverTree(root *TreeNode) {
	var prev *TreeNode
	var first, second *TreeNode
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		// 中序位置
		if prev != nil && prev.Val > root.Val {
			if first == nil {
				first = prev // 记录第一个不满足有序的节点
			}
			second = root // 更新最后一个不满足有序的节点
		}
		prev = root
		traverse(root.Right)
	}

	traverse(root)
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

// 669. 修剪二叉搜索树
// https://leetcode.cn/problems/trim-a-binary-search-tree/description/
// 给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。
// 输入：root = [1,0,2], low = 1, high = 2
// 输出：[1,null,2]
// 思路：分解问题，明确函数定义：返回以 root 为根节点的bst，修剪后节点值在[low..high]范围内的子树的根节点
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high) // 左边更小了，右子树中可能有，返回右子树中>low的节点
	} else if root.Val > high {
		return trimBST(root.Left, low, high) // 右边更大了，左子树中可能还有在区间内的节点
	}
	root.Left = trimBST(root.Left, low, root.Val)
	root.Right = trimBST(root.Right, root.Val, high)
	return root
}

// 671. 二叉树中第二小的节点
// 给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。
// 正式地说，即 root.val = min(root.left.val, root.right.val) 总成立。
// 给出这样的一个二叉树，你需要输出所有节点中的 第二小的值 。
// 如果第二小的值不存在的话，输出 -1 。
// 输入：root = [2,2,5,null,null,5,7]
// 输出：5
// 解释：最小的值是 2 ，第二小的值是 5 。
// 思路：分解问题。根是最小的节点，第二小的节点可能在左子树也可能在右子树中。明确函数定义：返回以 root 为根的第二小的节点值，没有的话返回-1
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	if root.Left == nil || root.Right == nil {
		return -1 // 叶子节点也没有第二小的节点
	}
	left, right := root.Left.Val, root.Right.Val
	if root.Val == left {
		left = findSecondMinimumValue(root.Left)
	}
	if root.Val == right {
		right = findSecondMinimumValue(root.Right) //
	}
	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}
	return min(left, right)
}

// 501.二叉搜索树中的众数
// https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/
// 思路：遍历，利用中序有序累计节点值个数，不断更新结果（最大个数和有最大个数的值）
func findMode(root *TreeNode) []int {
	var results []int
	maxCnt := 0
	cnt := 0
	var prev *TreeNode
	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		// 中序位置
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
		traverse(root.Right)
	}

	traverse(root)
	return results
}

// 530. 二叉搜索树的最小绝对差
// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/
// 思路：遍历。中序位置计算相邻节点的差，不断更新结果（最小值）
func getMinimumDifference(root *TreeNode) int {
	result := math.MaxInt
	var prev *TreeNode
	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		// 中序位置
		if prev != nil {
			result = min(result, root.Val-prev.Val)
		}
		prev = root
		traverse(root.Right)
	}

	traverse(root)
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
	var traverse func(root *TreeNode)
	var buildTree func(nums []int) *TreeNode

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		nums = append(nums, root.Val)
		traverse(root.Right)
	}

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

	traverse(root)
	return buildTree(nums)
}
