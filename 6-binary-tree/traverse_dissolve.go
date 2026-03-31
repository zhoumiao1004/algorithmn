package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 617. 合并二叉树
// https://leetcode.cn/problems/merge-two-binary-trees/description/
// 给你两棵二叉树： root1 和 root2 。
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
// 返回合并后的二叉树。
// 注意: 合并过程必须从两个树的根节点开始。
// 输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// 输出：[3,4,5,5,4,null,7]
// 思路1：分解问题
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

// 思路2：遍历
func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var traverse func(p, q *TreeNode)

	traverse = func(p, q *TreeNode) {
		if p == nil || q == nil {
			return
		}
		if p.Left == nil && q.Left != nil {
			p.Left = q.Left
			q.Left = nil
		}
		if p.Right == nil && q.Right != nil {
			p.Right = q.Right
			q.Right = nil
		}
		p.Val += q.Val
		traverse(p.Left, q.Left)
		traverse(p.Right, q.Right)
	}

	if root1 == nil {
		return root2
	}
	traverse(root1, root2)
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
// 思路1：分解问题+后序
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

// 1379. 找出克隆二叉树中的相同节点
// https://leetcode.cn/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/description/
// 给你两棵二叉树，原始树 original 和克隆树 cloned，以及一个位于原始树 original 中的目标节点 target。
// 其中，克隆树 cloned 是原始树 original 的一个 副本 。
// 请找出在树 cloned 中，与 target 相同 的节点，并返回对该节点的引用（在 C/C++ 等有指针的语言中返回 节点指针，其他语言返回节点本身）。

func main() {
	nums := []int{1, 2, 5, 3, 4, -1, 6}
	root := buildTreeByArray(nums, 0)
	// fmt.Println(preorderTraversal(root))
	flatten(root)
	// fmt.Println(preorderTraversal(root))
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
