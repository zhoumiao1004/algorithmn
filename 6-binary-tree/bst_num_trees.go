package main

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
