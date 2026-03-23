package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 230. 二叉搜索树中第 K 小的元素
// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（k 从 1 开始计数）。
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	result := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		// 中
		k--
		if k == 0 {
			result = node.Val
		}
		dfs(node.Right)
	}
	dfs(root)
	return result
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
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right) // 右
		// 中
		s += node.Val
		node.Val = s
		dfs(node.Left)
	}
	dfs(root)
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

// 1373. 二叉搜索子树的最大键值和
// 给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。
// 二叉搜索树的定义如下：
// 任意节点的左子树中的键值都 小于 此节点的键值。
// 任意节点的右子树中的键值都 大于 此节点的键值。
// 任意节点的左子树和右子树都是二叉搜索树。
// 输入：root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
// 输出：20
// 解释：键值为 3 的子树是和最大的二叉搜索树。
func maxSumBST(root *TreeNode) int {
	var maxSum int
	var findMaxMinSum func(*TreeNode) []int
	// 计算以 root 为根的二叉树的最大值、最小值、节点和
	findMaxMinSum = func(root *TreeNode) []int {
		// base case
		if root == nil {
			return []int{1, math.MaxInt32, math.MinInt32, 0}
		}

		// 递归计算左右子树
		left := findMaxMinSum(root.Left)
		right := findMaxMinSum(root.Right)

		// ******* 后序位置 *******
		// 通过 left 和 right 推导返回值
		// 并且正确更新 maxSum 变量
		res := make([]int, 4)
		// 这个 if 在判断以 root 为根的二叉树是不是 BST
		if left[0] == 1 && right[0] == 1 &&
			root.Val > left[2] && root.Val < right[1] {
			// 以 root 为根的二叉树是 BST
			res[0] = 1
			// 计算以 root 为根的这棵 BST 的最小值
			res[1] = min(left[1], root.Val)
			// 计算以 root 为根的这棵 BST 的最大值
			res[2] = max(right[2], root.Val)
			// 计算以 root 为根的这棵 BST 所有节点之和
			res[3] = left[3] + right[3] + root.Val
			// 更新全局变量
			maxSum = max(maxSum, res[3])
		} else {
			// 以 root 为根的二叉树不是 BST
			res[0] = 0
			// 其他的值都没必要计算了，因为用不到
		}
		// ************************

		return res
	}

	findMaxMinSum(root)
	return maxSum
}

// 99. 恢复二叉搜索树
// https://leetcode.cn/problems/recover-binary-search-tree/description/
// 给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
func recoverTree(root *TreeNode) {
	var prev *TreeNode
	var first, second *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if prev != nil && prev.Val > root.Val {
			if first == nil {
				first = prev
			}
			second = root
		}
		prev = root
		dfs(root.Right)
	}
	dfs(root)
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
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

// 671. 二叉树中第二小的节点
// 给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。
// 正式地说，即 root.val = min(root.left.val, root.right.val) 总成立。
// 给出这样的一个二叉树，你需要输出所有节点中的 第二小的值 。
// 如果第二小的值不存在的话，输出 -1 。
// 输入：root = [2,2,5,null,null,5,7]
// 输出：5
// 解释：最小的值是 2 ，第二小的值是 5 。
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	if root.Left == nil || root.Right == nil {
		return -1
	}
	left, right := root.Left.Val, root.Right.Val
	if root.Val == left {
		left = findSecondMinimumValue(root.Left)
	}
	if root.Val == right {
		right = findSecondMinimumValue(root.Right)
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

// 653. 两数之和 IV - 输入二叉搜索树
// https://leetcode.cn/problems/two-sum-iv-input-is-a-bst/description/
// 给定一个二叉搜索树 root 和一个目标结果 k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 true。
// 输入: root = [5,3,6,2,4,null,7], k = 9
// 输出: true
func findTarget(root *TreeNode, k int) bool {
	var nums []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		nums = append(nums, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == k {
			return true
		} else if nums[left]+nums[right] < k {
			left++
		} else {
			right--
		}
	}
	return false
}

// 1008. 前序遍历构造二叉搜索树
// https://leetcode.cn/problems/construct-binary-search-tree-from-preorder-traversal/description/
// 给定一个整数数组，它表示BST(即 二叉搜索树 )的 先序遍历 ，构造树并返回其根。
// 保证 对于给定的测试用例，总是有可能找到具有给定需求的二叉搜索树。
// 二叉搜索树 是一棵二叉树，其中每个节点， Node.left 的任何后代的值 严格小于 Node.val , Node.right 的任何后代的值 严格大于 Node.val。
// 二叉树的 前序遍历 首先显示节点的值，然后遍历Node.left，最后遍历Node.right。
// 输入：preorder = [8,5,1,7,10,12]
// 输出：[8,5,10,1,7,null,12]
func bstFromPreorder(preorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	val := preorder[0]
	if n == 1 {
		return &TreeNode{Val: val}
	}
	root := &TreeNode{Val: val}
	mid := 1
	for mid < n && preorder[mid] < val {
		mid++
	}
	root.Left = bstFromPreorder(preorder[1:mid])
	root.Right = bstFromPreorder(preorder[mid:])
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

type ListNode struct {
	Val  int
	Next *ListNode
}

// 109. 有序链表转换二叉搜索树
// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/description/
// 给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为 平衡 二叉搜索树。
// 输入: head = [-10,-3,0,5,9]
// 输出: [0,-3,9,-10,null,5]
// 解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head2 := slow.Next.Next
	root := &TreeNode{Val: slow.Next.Val}
	slow.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(head2)
	return root
}

// 173. 二叉搜索树迭代器
// https://leetcode.cn/problems/binary-search-tree-iterator/description/
// 实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
// BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
// boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
// int next()将指针向右移动，然后返回指针处的数字。
// 注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。
// 你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。
// 输入
// ["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
// [[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
// 输出
// [null, 3, 7, true, 9, true, 15, true, 20, false]
type BSTIterator struct {
    st []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{st: []*TreeNode{}}
	iterator.pushLeftBranch(root)
	return iterator
}

func (this *BSTIterator) pushLeftBranch(p *TreeNode) {
	for p != nil {
		this.st = append(this.st, p)
		p = p.Left
	}
}

func (this *BSTIterator) Next() int {
	node := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]
	this.pushLeftBranch(node.Right)
	return node.Val
}

func (this *BSTIterator) Peek() int {
	node := this.st[len(this.st)-1]
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.st) > 0
}

// 1305. 两棵二叉搜索树中的所有元素
// https://leetcode.cn/problems/all-elements-in-two-binary-search-trees/
// 给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.
// 输入：root1 = [2,1,4], root2 = [1,0,3]
// 输出：[0,1,1,2,3,4]
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	t1 := Constructor(root1)
	t2 := Constructor(root2)
	var results []int
	for t1.HasNext() && t2.HasNext() {
		v1 := t1.Peek()
		v2 := t2.Peek()
		if v1 < v2 {
			results = append(results, v1)
			t1.Next()
		} else {
			results = append(results, v2)
			t2.Next()
		}
	}
	for t1.HasNext() {
		results = append(results, t1.Next())
	}
	for t2.HasNext() {
		results = append(results, t2.Next())
	}
	return results
}
