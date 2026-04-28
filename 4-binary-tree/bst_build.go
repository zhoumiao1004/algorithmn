package main

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

// 109. 有序链表转换二叉搜索树
// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/description/
// 给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为 平衡 二叉搜索树。
// 输入: head = [-10,-3,0,5,9]
// 输出: [0,-3,9,-10,null,5]
// 解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
// 思路1：分解问题，列表转换成数组，再转换成bst
// 思路2：分解问题，链表双指针
func sortedListToBST(head *ListNode) *TreeNode {
	var getMid func(begin, end *ListNode) *ListNode
	var build func(begin, end *ListNode) *TreeNode

	getMid = func(begin, end *ListNode) *ListNode {
		slow, fast := begin, begin
		for fast != end && fast.Next != end {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}

	build = func(begin, end *ListNode) *TreeNode {
		if begin == end {
			return nil
		}
		mid := getMid(begin, end)
		return &TreeNode{
			Val:   mid.Val,
			Left:  build(begin, mid),
			Right: build(mid.Next, end),
		}
	}

	return build(head, nil)
}

// 思路3：分解问题，中序遍历（最优）
func sortedListToBST3(head *ListNode) *TreeNode {
	root := head
	var build func(left, right int) *TreeNode

	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		leftTree := build(left, mid-1)
		rootVal := root.Val
		root = root.Next
		rightTree := build(mid+1, right)
		return &TreeNode{
			Val:   rootVal,
			Left:  leftTree,
			Right: rightTree,
		}
	}

	cnt := 0
	for cur := head; cur != nil; cur = cur.Next {
		cnt++
	}
	return build(0, cnt-1)
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
	var build func(low, high int) []*TreeNode // 明确函数定义：使用[low..high]构造二叉搜索树

	build = func(low, high int) []*TreeNode {
		var results []*TreeNode
		if low > high {
			results = append(results, nil)
			return results
		}
		// 1.穷举 root 节点的所有可能
		for i := low; i <= high; i++ {
			// 2.穷举左右子树所有bst可能
			left := build(low, i-1)
			right := build(i+1, high)
			// 3.给 root 穷举所有左右子树的组合
			for _, left := range left {
				for _, right := range right {
					results = append(results, &TreeNode{Val: i, Left: left, Right: right})
				}
			}
		}
		return results
	}

	return build(1, n) // 构造闭区间 [1, n] 组成的 BST
}

// 894. 所有可能的真二叉树
// https://leetcode.cn/problems/all-possible-full-binary-trees/
// 给你一个整数 n ，请你找出所有可能含 n 个节点的 真二叉树 ，并以列表形式返回。答案中每棵树的每个节点都必须符合 Node.val == 0 。
// 答案的每个元素都是一棵真二叉树的根节点。你可以按 任意顺序 返回最终的真二叉树列表。
// 真二叉树 是一类二叉树，树中每个节点恰好有 0 或 2 个子节点。
func allPossibleFBT(n int) []*TreeNode {
	
	memo := make(map[int][]*TreeNode)
	var build func(n int) []*TreeNode

	build = func(n int) []*TreeNode {
		var res []*TreeNode
		if n == 1 {
			res = append(res, &TreeNode{Val: 0})
			return res
		}
		if res, ok := memo[n]; ok {
			return res
		}
		for i := 1; i<n; i+=2 {
			j := n-i-1
			leftSubTree := build(i)
			rightSubTree := build(j)
			for _, left := range leftSubTree {
				for _, right := range rightSubTree {
					root := &TreeNode{Val: 0, Left: left, Right: right}
					res = append(res, root)
				}
			}
		}
		return res
	}

	if n%2 == 0 {
		return []*TreeNode{} // 题目描述的满二叉树不可能是偶数个节点
	}
	return build(n)
}
