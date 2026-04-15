package main

import "math"

// 1261. 在受污染的二叉树中查找元素
// https://leetcode.cn/problems/find-elements-in-a-contaminated-binary-tree/description/
// 给出一个满足下述规则的二叉树：
// root.val == 0
// 对于任意 treeNode：
// 如果 treeNode.val 为 x 且 treeNode.left != null，那么 treeNode.left.val == 2 * x + 1
// 如果 treeNode.val 为 x 且 treeNode.right != null，那么 treeNode.right.val == 2 * x + 2
// 现在这个二叉树受到「污染」，所有的 treeNode.val 都变成了 -1。
// 请你先还原二叉树，然后实现 FindElements 类：
// FindElements(TreeNode* root) 用受污染的二叉树初始化对象，你需要先把它还原。
// bool find(int target) 判断目标值 target 是否存在于还原后的二叉树中并返回结果。
// 输入：
// ["FindElements","find","find"]
// [[[-1,null,-1]],[1],[2]]
// 输出：
// [null,false,true]
// 解释：
// FindElements findElements = new FindElements([-1,null,-1]);
// findElements.find(1); // return False
// findElements.find(2); // return True
type FindElements struct {
	values map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	// 还原二叉树中的值
	fe := FindElements{values: make(map[int]bool)}
	fe.traverse(root, 0)
	return fe
}

func (this *FindElements) traverse(root *TreeNode, val int) {
	if root == nil {
		return
	}
	root.Val = val
	this.values[root.Val] = true
	this.traverse(root.Left, 2*val+1)
	this.traverse(root.Right, 2*val+2)
}

func (this *FindElements) Find(target int) bool {
	return this.values[target]
}

// 1379. 找出克隆二叉树中的相同节点
// https://leetcode.cn/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/description/
// 给你两棵二叉树，原始树 original 和克隆树 cloned，以及一个位于原始树 original 中的目标节点 target。
// 其中，克隆树 cloned 是原始树 original 的一个 副本 。
// 请找出在树 cloned 中，与 target 相同 的节点，并返回对该节点的引用（在 C/C++ 等有指针的语言中返回 节点指针，其他语言返回节点本身）。

// 1443. 收集树上所有苹果的最少时间
// https://leetcode.cn/problems/minimum-time-to-collect-all-apples-in-a-tree/description/
// 给你一棵有 n 个节点的无向树，节点编号为 0 到 n-1 ，它们中有一些节点有苹果。通过树上的一条边，需要花费 1 秒钟。你从 节点 0 出发，请你返回最少需要多少秒，可以收集到所有苹果，并回到节点 0 。
func minTime(n int, edges [][]int, hasApple []bool) int {
	graph := make(map[int][]int)
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	var collect func(graph map[int][]int, root int) int // 明确函数定义：遍历以 root 为根的多叉树，返回收集所有苹果最少步数
	collect = func(graph map[int][]int, root int) int {
		if visited[root] {
			return -1
		}
		visited[root] = true

		sum := 0
		for _, c := range graph[root] {
			subTime := collect(graph, c)
			if subTime != -1 {
				sum += subTime + 2
			}
		}
		// 后序位置
		if sum > 0 {
			return sum
		}
		if sum == 0 && hasApple[root] {
			return 0 // root 本身有苹果，子树中没有苹果
		}
		return -1
	}

	res := collect(graph, 0)
	if res == -1 {
		return 0
	}
	return res
}

// 979. 在二叉树中分配硬币
// https://leetcode.cn/problems/distribute-coins-in-binary-tree/description/
// 给你一个有 n 个结点的二叉树的根结点 root ，其中树中每个结点 node 都对应有 node.val 枚硬币。整棵树上一共有 n 枚硬币。
// 在一次移动中，我们可以选择两个相邻的结点，然后将一枚硬币从其中一个结点移动到另一个结点。移动可以是从父结点到子结点，或者从子结点移动到父结点。
// 返回使每个结点上 只有 一枚硬币所需的 最少 移动次数。
// 输入：root = [3,0,0]
// 输出：2
// 解释：一枚硬币从根结点移动到左子结点，一枚硬币从根结点移动到右子结点。
func distributeCoins(root *TreeNode) int {
	result := 0
	var getRest func(root *TreeNode) int

	getRest = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := getRest(root.Left)
		right := getRest(root.Right)
		result += int(math.Abs(float64(left)) + math.Abs(float64(right)))
		return left + right + (root.Val - 1)
	}
	getRest(root)
	return result
}

// 1080. 根到叶路径上的不足节点
// https://leetcode.cn/problems/insufficient-nodes-in-root-to-leaf-paths/description/
// 给你二叉树的根节点 root 和一个整数 limit ，请你同时删除树中所有 不足节点 ，并返回最终二叉树的根节点。
// 假如通过节点 node 的每种可能的 “根-叶” 路径上值的总和全都小于给定的 limit，则该节点被称之为 不足节点 ，需要被删除。
// 叶子节点，就是没有子节点的节点。
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	// 前序位置
	if root.Left == nil && root.Right == nil {
		if root.Val < limit {
			return nil
		}
		return root
	}
	root.Left = sufficientSubset(root.Left, limit-root.Val)
	root.Right = sufficientSubset(root.Right, limit-root.Val)
	// 后序位置
	if root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

// 2049. 统计最高分的节点数目
// https://leetcode.cn/problems/count-nodes-with-the-highest-score/description/
// 给你一棵根节点为 0 的 二叉树 ，它总共有 n 个节点，节点编号为 0 到 n - 1 。同时给你一个下标从 0 开始的整数数组 parents 表示这棵树，其中 parents[i] 是节点 i 的父节点。由于节点 0 是根，所以 parents[0] == -1 。
// 一个子树的 大小 为这个子树内节点的数目。每个节点都有一个与之关联的 分数 。求出某个节点分数的方法是，将这个节点和与它相连的边全部 删除 ，剩余部分是若干个 非空 子树，这个节点的 分数 为所有这些子树 大小的乘积 。
// 请你返回有 最高得分 节点的 数目 。
func countHighestScoreNodes(parents []int) int {
	scoreToCount := make(map[int64]int)
	var countNode func(tree [][]int, root int) int
	var buildTree func(parents []int) [][]int

	// 计算二叉树中的节点个数
	countNode = func(tree [][]int, root int) int {
		if root == -1 {
			return 0
		}
		// 二叉树中节点总数
		n := len(tree)
		leftCount := countNode(tree, tree[root][0])
		rightCount := countNode(tree, tree[root][1])

		// 后序位置，计算每个节点的「分数」
		otherCount := n - leftCount - rightCount - 1
		// 注意，这里要把 int 转化成 long，否则会产生溢出！！！
		score := int64(math.Max(float64(leftCount), 1)) *
			int64(math.Max(float64(rightCount), 1)) * int64(math.Max(float64(otherCount), 1))
		// 给分数 score 计数
		scoreToCount[score] = scoreToCount[score] + 1

		return leftCount + rightCount + 1
	}

	buildTree = func(parents []int) [][]int {
		n := len(parents)
		tree := make([][]int, n)
		for i := range tree {
			tree[i] = []int{-1, -1}
		}
		for i := 1; i < n; i++ {
			parent_i := parents[i]
			if tree[parent_i][0] == -1 {
				tree[parent_i][0] = i
			} else {
				tree[parent_i][1] = i
			}
		}
		return tree
	}

	tree := buildTree(parents)
	countNode(tree, 0)
	// 计算最大分数出现的次数
	var maxScore int64
	for score := range scoreToCount {
		maxScore = int64(math.Max(float64(maxScore), float64(score)))
	}
	return scoreToCount[maxScore]
}
