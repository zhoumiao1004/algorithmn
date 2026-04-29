package main

import (
	"fmt"
	"math"
)

/* 有些题目，你按照拍脑袋的方式去做，可能发现需要在递归代码中调用其他递归函数计算字数的信息。
一般来说，出现这种情况时你可以考虑用后序遍历的思维方式来优化算法，利用后序遍历传递子树的信息，避免过高的时间复杂度。
前序位置的代码只能从函数参数中获取父节点传递来的数据，而后序位置的代码不仅可以获取参数数据，还可以获取到子树通过函数返回值传递回来的数据。
一旦你发现题目和子树有关，那大概率要给函数设置合理的定义和返回值，在后序位置写代码了。
*/

// 652. 寻找重复的子树
// https://leetcode.cn/problems/find-duplicate-subtrees/description/
// 给你一棵二叉树的根节点 root ，返回所有 重复的子树 。
// 对于同一类的重复子树，你只需要返回其中任意 一棵 的根结点即可。
// 如果两棵树具有 相同的结构 和 相同的结点值 ，则认为二者是 重复 的。
// 输入：root = [1,2,3,4,null,2,4,null,null,4]
// 输出：[[2,4],[4]]
// 思路1: 后序
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	memo := make(map[string]int)
	var serialize func(node *TreeNode) string

	serialize = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		left := serialize(node.Left)
		right := serialize(node.Right)
		s := fmt.Sprintf("%d,%s,%s", node.Val, left, right)
		// 后序位置，顺便计算是否存在重复子树
		if memo[s] == 1 {
			res = append(res, node)
		}
		memo[s]++
		return s
	}

	serialize(root)
	return res
}

// 思路2: 遍历
func findDuplicateSubtrees3(root *TreeNode) []*TreeNode {
	var result []*TreeNode
	subMap := make(map[string]int)
	var serialize func(node *TreeNode) string
	var traverse func(node *TreeNode)

	serialize = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		left := serialize(node.Left)
		right := serialize(node.Right)
		return fmt.Sprintf("%d,%s,%s", node.Val, left, right)
	}

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		// 后序位置
		s := serialize(node)
		if subMap[s] == 1 {
			result = append(result, node)
		}
		subMap[s]++
	}

	traverse(root)
	return result
}

// 110. 平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/description/
// 对于树中的每个节点：左和右子树高度差不超过1
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
// 思路：分解问题
func isBalanced(root *TreeNode) bool {
	flag := true
	var maxDepth func(node *TreeNode) int // 明确函数定义：返回以 node 为根的子树的最大深度

	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := maxDepth(node.Left)
		right := maxDepth(node.Right)
		// 后序位置顺便判断是否平衡
		if math.Abs(float64(left-right)) > 1 {
			flag = false
		}
		return max(left, right) + 1
	}

	maxDepth(root)
	return flag
}

// 508. 出现次数最多的子树元素和
// https://leetcode.cn/problems/most-frequent-subtree-sum/
// 给你一个二叉树的根结点 root ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。
// 一个结点的 「子树元素和」 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。
// 输入: root = [5,2,-3]
// 输出: [2,-3,4]
func findFrequentTreeSum(root *TreeNode) []int {
	var res []int
	maxCnt := 0
	maxSumCnt := make(map[int]int)      // 记录和的次数
	var getSum func(node *TreeNode) int // 明确函数定义：返回以 node 为根的二叉树的元素和

	getSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getSum(node.Left)
		right := getSum(node.Right)
		// 后序位置顺便更新最大子树元素和
		s := node.Val + left + right
		maxSumCnt[s]++
		if maxSumCnt[s] == maxCnt {
			res = append(res)
		} else if maxSumCnt[s] > maxCnt {
			res = []int{s}
			maxCnt = maxSumCnt[s]
		}
		return s
	}

	getSum(root)
	return res
}

// 563. 二叉树的坡度
// https://leetcode.cn/problems/binary-tree-tilt/description/
// 给你一个二叉树的根节点 root ，计算并返回 整个树 的坡度 。
// 一个树的 节点的坡度 定义即为，该节点左子树的节点之和和右子树节点之和的 差的绝对值 。如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是 0 。
// 整个树 的坡度就是其所有节点的坡度之和。
func findTilt(root *TreeNode) int {
	result := 0
	var getSum func(node *TreeNode) int

	getSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getSum(node.Left)
		right := getSum(node.Right)
		// 后序位置顺便累加坡度和
		result += int(math.Abs(float64(left) - float64(right)))
		return left + right + node.Val
	}

	getSum(root)
	return result
}

// 814. 二叉树剪枝
// https://leetcode.cn/problems/binary-tree-pruning/description/
// 给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。
// 返回移除了所有不包含 1 的子树的原二叉树。
// 节点 node 的子树为 node 本身加上所有 node 的后代。
// 输入：root = [1,null,0,0,1]
// 输出：[1,null,0,null,1]
// 思路：分解问题，明确函数定义：返回以 root 为根的二叉树剪枝后的原二叉树
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)   // 左子树剪枝
	root.Right = pruneTree(root.Right) // 右子树剪枝
	// 后序位置
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil // return nil 相当于删除节点
	}
	return root
}

// 1325. 删除给定值的叶子节点
// https://leetcode.cn/problems/delete-leaves-with-a-given-value/description/
// 给你一棵以 root 为根的二叉树和一个整数 target ，请你删除所有值为 target 的 叶子节点 。
// 注意，一旦删除值为 target 的叶子节点，它的父节点就可能变成叶子节点；如果新叶子节点的值恰好也是 target ，那么这个节点也应该被删除。
// 也就是说，你需要重复此过程直到不能继续删除。
// 输入：root = [1,2,3,2,null,2,4], target = 2
// 输出：[1,null,3,null,4]
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	// 后序位置
	if root.Val == target && root.Left == nil && root.Right == nil {
		return nil // return nil 相当于删除节点
	}
	return root
}

// 1026. 节点与其祖先之间的最大差值
// https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/description/
// 给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。
// （如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）
// 输入：root = [8,3,10,1,6,null,14,null,null,4,7,13]
// 输出：7
// 0 <= Node.val <= 100000
func maxAncestorDiff(root *TreeNode) int {
	res := 0
	// 定义：输入一棵二叉树，返回该二叉树中节点的最小值和最大值，
	var getMinMax func(root *TreeNode) (int, int)

	getMinMax = func(root *TreeNode) (int, int) {
		if root == nil {
			return math.MaxInt, math.MinInt
		}
		leftMin, leftMax := getMinMax(root.Left)
		rightMin, rightMax := getMinMax(root.Right)

		// 后序位置
		rootMin := min(root.Val, leftMin, rightMin)
		rootMax := max(root.Val, leftMax, rightMax)
		// 在后序位置顺便判断所有差值的最大值
		res = max(res, rootMax-root.Val, root.Val-rootMin)

		return rootMin, rootMax
	}

	getMinMax(root)
	return res
}

// 1339. 分裂二叉树的最大乘积
// https://leetcode.cn/problems/maximum-product-of-splitted-binary-tree/description/
// 给你一棵二叉树，它的根为 root 。请你删除 1 条边，使二叉树分裂成两棵子树，且它们子树和的乘积尽可能大。
// 由于答案可能会很大，请你将结果对 10^9 + 7 取模后再返回。
// 输入：root = [1,2,3,4,5,6]
// 输出：110
// 解释：删除红色的边，得到 2 棵子树，和分别为 11 和 10 。它们的乘积是 110 （11*10）
func maxProduct(root *TreeNode) int {
	res := 0
	var getTreeSum func(node *TreeNode) int
	var getSum func(node *TreeNode) int // 明确函数定义：返回以 node 为根的子树的元素和

	getTreeSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getTreeSum(node.Left)
		right := getTreeSum(node.Right)
		return node.Val + left + right
	}
	total := getTreeSum(root)

	getSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getSum(node.Left)
		right := getSum(node.Right)
		s := node.Val + left + right
		res = max(res, s*(total-s))
		return s
	}

	getSum(root)
	return res % (1e9 + 7)
}

// 1372. 二叉树中的最长交错路径
// https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree/
// 给你一棵以 root 为根的二叉树，二叉树中的交错路径定义如下：
// 选择二叉树中 任意 节点和一个方向（左或者右）。
// 如果前进方向为右，那么移动到当前节点的的右子节点，否则移动到它的左子节点。
// 改变前进方向：左变右或者右变左。
// 重复第二步和第三步，直到你在树中无法继续移动。
// 交错路径的长度定义为：访问过的节点数目 - 1（单个节点的路径长度为 0 ）。
// 请你返回给定树中最长 交错路径 的长度。
func longestZigZag(root *TreeNode) int {
	res := 0
	var getPathLen func(root *TreeNode) []int

	// 输入二叉树的根节点 root，返回两个值
	// 第一个是从 root 开始向左走的最长交错路径长度，
	// 第一个是从 root 开始向右走的最长交错路径长度
	getPathLen = func(root *TreeNode) []int {
		if root == nil {
			return []int{-1, -1}
		}
		left := getPathLen(root.Left)
		right := getPathLen(root.Right)
		// 后序位置，根据左右子树的交错路径长度推算根节点的交错路径长度
		rootPathLen1 := left[1] + 1
		rootPathLen2 := right[0] + 1
		// 更新全局最大值
		res = max(res, max(rootPathLen1, rootPathLen2))

		return []int{rootPathLen1, rootPathLen2}
	}

	getPathLen(root)
	return res
}

// 606. 根据二叉树创建字符串
// https://leetcode.cn/problems/construct-string-from-binary-tree/description/
// 给你二叉树的根节点 root ，请你采用前序遍历的方式，将二叉树转化为一个由括号和整数组成的字符串，返回构造出的字符串。
// 空节点使用一对空括号对 "()" 表示，转化后需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val)
	}
	left := tree2str(root.Left)
	right := tree2str(root.Right)
	if root.Left != nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val) + "(" + left + ")"
	} else if root.Left == nil && root.Right != nil {
		return fmt.Sprintf("%d", root.Val) + "()" + "(" + right + ")"
	}
	return fmt.Sprintf("%d", root.Val) + "(" + left + ")" + "(" + right + ")"
}

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

func main() {
	fmt.Println("hello world")
	// maxAncestorDiff()
}
