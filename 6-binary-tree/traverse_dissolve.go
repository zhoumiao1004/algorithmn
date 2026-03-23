package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144. 二叉树的前序遍历
// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
// 递归：遍历的思路
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

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 1.遍历的思路（回溯）
func maxDepth(root *TreeNode) int {
	result := 0
	depth := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			result = max(result, depth)
			return
		}
		depth++
		dfs(node.Left)
		dfs(node.Right)
		depth--
	}
	dfs(root)
	return result
}

// 2.分解子问题的思路(dp)后序遍历
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return 1 + max(leftDepth, rightDepth)
}

// 112. 路径总和
// https://leetcode.cn/problems/path-sum/description/
// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
// 叶子节点 是指没有子节点的节点。
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 不涉及中的操作，所以前中后序遍历都可以
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 113. 路径总和 II
// https://leetcode.cn/problems/path-sum-ii/description/
// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
// 叶子节点 是指没有子节点的节点。
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：[[5,4,11,2],[5,8,4,5]]
func pathSumII(root *TreeNode, targetSum int) [][]int {
	var results [][]int
	var path []int
	s := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		s += root.Val
		if s == targetSum && root.Left == nil && root.Right == nil {
			results = append(results, append([]int{}, path...))
		}
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]
		s -= root.Val
	}
	dfs(root)
	return results
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

// 897. 递增顺序搜索树
// https://leetcode.cn/problems/increasing-order-search-tree/
// 给你一棵二叉搜索树的 root ，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
// 输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
// 输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
// 方法1:遍历
func increasingBST2(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	cur := dummy
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left) // 左
		cur.Right = &TreeNode{Val: root.Val}
		cur = cur.Right
		dfs(root.Right) // 右
	}
	dfs(root)
	return dummy.Right
}

// 方法2:分解
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 左右子树拉平
	left := increasingBST(root.Left)
	root.Left = nil // 注意左子树置为空！
	right := increasingBST(root.Right)
	root.Right = right
	if left == nil {
		return root
	}
	// 中: root节点挂到左子树最右边的节点上
	cur := left
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = root
	return left
}

// 114. 二叉树展开为链表
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left == nil {
		return
	}
	cur := root.Left
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = root.Right
	root.Right = root.Left
	root.Left = nil // 注意需要清空左节点
}

// 938. 二叉搜索树的范围和
// https://leetcode.cn/problems/range-sum-of-bst/
// 给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
// 输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
// 输出：32
// 对比 669. 修剪二叉搜索树 https://leetcode.cn/problems/trim-a-binary-search-tree/description/
func rangeSumBST2(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	result := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val >= low && root.Val <= high {
			result += root.Val
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return result
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	left := rangeSumBST(root.Right, low, high)
	right := rangeSumBST(root.Left, low, high)
	s := left + right
	if root.Val >= low && root.Val <= high {
		s += root.Val
	}
	return s
}

// 1379. 找出克隆二叉树中的相同节点
// https://leetcode.cn/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/description/
// 给你两棵二叉树，原始树 original 和克隆树 cloned，以及一个位于原始树 original 中的目标节点 target。
// 其中，克隆树 cloned 是原始树 original 的一个 副本 。
// 请找出在树 cloned 中，与 target 相同 的节点，并返回对该节点的引用（在 C/C++ 等有指针的语言中返回 节点指针，其他语言返回节点本身）。

func main() {
	nums := []int{1, 2, 5, 3, 4, -1, 6}
	root := buildTreeByArray(nums, 0)
	fmt.Println(preorderTraversal(root))
	flatten(root)
	fmt.Println(preorderTraversal(root))
}

func buildTreeByArray(nums []int, i int) *TreeNode {
	if i >= len(nums) {
		return nil
	}
	if nums[i] == -1 {
		return nil
	}
	node := &TreeNode{Val: nums[i]}
	node.Left = buildTreeByArray(nums, 2*i+1)
	node.Right = buildTreeByArray(nums, 2*i+2)
	return node
}
