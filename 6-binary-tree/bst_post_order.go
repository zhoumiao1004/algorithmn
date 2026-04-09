package main

import "math"

// 1373. 二叉搜索子树的最大键值和
// https://leetcode.cn/problems/maximum-sum-bst-in-binary-tree/?show=1
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
			return []int{1, math.MaxInt32, math.MinInt32, 0} // 这里nil节点最小值初始化为最大整数，最大值初始化为最小整数，方便之处在于下面可以通过min()和max()处理2种情况：1.左右孩子有空节点 2.左右孩子非空
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
