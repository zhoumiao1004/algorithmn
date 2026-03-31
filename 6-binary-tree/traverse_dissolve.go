package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 思路1:遍历整棵树，外部变量记录递归深度
func maxDepth(root *TreeNode) int {
	result := 0
	depth := 0
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			result = max(result, depth)
			return
		}
		depth++
		traverse(node.Left)
		traverse(node.Right)
		depth--
	}

	traverse(root)
	return result
}

// 思路2:分解问题+后序
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return 1 + max(leftDepth, rightDepth)
}

// 617. 合并二叉树
// https://leetcode.cn/problems/merge-two-binary-trees/description/
// 给你两棵二叉树： root1 和 root2 。
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
// 返回合并后的二叉树。
// 注意: 合并过程必须从两个树的根节点开始。
// 输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// 输出：[3,4,5,5,4,null,7]
// 思路：分解问题+前序
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
// 思路1:遍历整棵树，创建一颗新树
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

// 思路2:分解问题+后序，修改原树
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 左右子树拉平
	left := increasingBST(root.Left)   // 左
	root.Left = nil                    // 注意左子树置为空！
	right := increasingBST(root.Right) // 右
	root.Right = right

	// 后序位置
	if left == nil {
		return root
	}
	cur := left
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = root // 节点挂到左子树最右边的节点上
	return left
}

// 114. 二叉树展开为链表
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
// 思路：分解问题+后序
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)  // 左
	flatten(root.Right) // 右

	// 后序位置
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

// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/description/
// 思路：分解问题+后序
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)   // 左
	right := invertTree(root.Right) // 右
	// 中序位置
	root.Left = right
	root.Right = left
	return root
}

// 111.二叉树的最小深度
// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
// 输入：root = [3,9,20,null,null,15,7]
// 输出：2
// 思路1:分解问题+后序
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

// 思路2: 层序遍历BFS。遍历到的第一个叶子节点的深度
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		depth++
		sz := len(q)
		for i := 0; i < sz; i++ {
			node := q[0]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return depth
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
