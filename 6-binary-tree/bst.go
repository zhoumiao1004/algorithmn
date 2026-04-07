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

// 96.不同的二叉搜索树
// https://leetcode.cn/problems/unique-binary-search-trees/description/
// 求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
// 输入：n = 3 输出：5
// 方法1: dp
func numTrees(n int) int {
	// 总共n个节点，左右子树加起来n-1个节点
	// dp[i]含义：i个节点的二叉搜索树个数
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}

// 方法2: 递归，分解问题的思路
func numTrees2(n int) int {
	var count func(low, high int, memo [][]int) int
	count = func(low, high int, memo [][]int) int {
		if low > high {
			return 1
		}
		if memo[low][high] != 0 {
			return memo[low][high]
		}

		result := 0
		for i := low; i <= high; i++ {
			// i的值作为root
			left := count(low, i-1, memo)
			right := count(i+1, high, memo)
			result += left * right
		}
		memo[low][high] = result
		return result
	}
	// 计算闭区间 [1, n] 组成的 BST 个数
	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n+1)
	}
	return count(1, n, memo)
}

// 95. 不同的二叉搜索树 II
// https://leetcode.cn/problems/unique-binary-search-trees-ii/description/
// 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
// 输入：n = 3
// 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
func generateTrees(n int) []*TreeNode {
	var build func(low, high int) []*TreeNode
	build = func(low, high int) []*TreeNode {
		var results []*TreeNode
		if low > high {
			// 这里需要装一个 null 元素，这样才能让下面的两个内层 for 循环都能进入，正确地创建出叶子节点
			// 举例来说吧，什么时候会进到这个 if 语句？当你创建叶子节点的时候，对吧。
			// 那么如果你这里不加 null，直接返回空列表，那么下面的内层两个 for 循环都无法进入
			// 你的那个叶子节点就没有创建出来，看到了吗？所以这里要加一个 null，确保下面能把叶子节点做出来
			results = append(results, nil)
			return results
		}
		// 穷举 root 节点的所有可能
		for i := low; i <= high; i++ {
			// 递推构造出左右子树的所有BST
			left := build(low, i-1)
			right := build(i+1, high)
			for _, node1 := range left {
				for _, node2 := range right {
					results = append(results, &TreeNode{Val: i, Left: node1, Right: node2})
				}
			}
		}
		return results
	}
	if n == 0 {
		return []*TreeNode{}
	}
	// 构造闭区间 [1, n] 组成的 BST
	return build(1, n)
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
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		// 中序位置
		if prev != nil && prev.Val > root.Val {
			if first == nil {
				first = prev
			}
			second = root
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
func getMinimumDifference(root *TreeNode) int {
	result := math.MaxInt
	var prev *TreeNode
	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		if prev != nil {
			result = root.Val - prev.Val
		}
		prev = root
		traverse(root.Right)
	}

	traverse(root)
	return result
}

// 653. 两数之和 IV - 输入二叉搜索树
// https://leetcode.cn/problems/two-sum-iv-input-is-a-bst/description/
// 给定一个二叉搜索树 root 和一个目标结果 k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 true。
// 输入: root = [5,3,6,2,4,null,7], k = 9
// 输出: true
// 思路1: 利用bst中序有序的特点，输出到数组+双指针
func findTarget(root *TreeNode, k int) bool {
	var nums []int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		nums = append(nums, root.Val)
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)

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

// 思路2: 分解问题（一般二叉树解法），明确函数定义：以 root 为根节点的二叉树，返回是否存在2个节点和为k
func findTarget2(root *TreeNode, k int) bool {
	m := make(map[int]bool)
	// 明确函数定义：返回以 node 节点为根的二叉树是否包含2个节点和为k
	var check func(node *TreeNode) bool

	check = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		// 前序位置
		if m[k-node.Val] {
			return true
		}
		m[node.Val] = true
		return check(node.Left) || check(node.Right)
	}

	if root == nil {
		return false
	}
	return check(root)
}

// 思路3: 遍历 + hashmap（一般二叉树解法）
func findTarget3(root *TreeNode, k int) bool {
	result := false
	m := make(map[int]bool)
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		// 中序位置
		if m[k-node.Val] {
			result = true
		}
		m[node.Val] = true
		traverse(node.Right)
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
