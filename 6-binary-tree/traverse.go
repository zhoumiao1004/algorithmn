package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 257. 二叉树的所有路径
// https://leetcode.cn/problems/binary-tree-paths/description/
// 输入：root = [1,2,3,null,5]
// 输出：["1->2->5","1->3"]
// 先序遍历
func binaryTreePaths(root *TreeNode) []string {
	var results []string
	var path []string
	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 中
		path = append(path, fmt.Sprintf("%d", root.Val))
		if root.Left == nil && root.Right == nil {
			results = append(results, strings.Join(path, "->")) // 注意不能return，因为还要回溯
		}
		traverse(root.Left)  // 左
		traverse(root.Right) // 右
		path = path[:len(path)-1]
	}

	traverse(root)
	return results
}

// 129. 求根节点到叶节点数字之和
// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
// 每条从根节点到叶节点的路径都代表一个数字：
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
// 输入：root = [1,2,3]
// 输出：25
// 解释：
// 从根到叶子节点路径 1->2 代表数字 12
// 从根到叶子节点路径 1->3 代表数字 13
// 因此，数字总和 = 12 + 13 = 25
func sumNumbers(root *TreeNode) int {
	result := 0
	var path []int
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			s := 0
			for i := 0; i < len(path); i++ {
				s = 10*s + path[i]
			}
			result += s
		}
		traverse(root.Left)
		traverse(root.Right)
		path = path[:len(path)-1]
	}

	traverse(root)
	return result
}

// 199. 二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	var results []int
	if root == nil {
		return results
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		node := q[0]
		q = q[1:]
		results = append(results, q[len(q)-1].Val)
		for i := 0; i < sz; i++ {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return results
}

// 988. 从叶结点开始的最小字符串
// 给定一颗根结点为 root 的二叉树，树中的每一个结点都有一个 [0, 25] 范围内的值，分别代表字母 'a' 到 'z'。
// 返回 按字典序最小 的字符串，该字符串从这棵树的一个叶结点开始，到根结点结束。
// 输入：root = [0,1,2,3,4,3,4]
// 输出："dba"
func smallestFromLeaf(root *TreeNode) string {
	var path []byte
	result := ""
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, byte(node.Val+'a'))
		if node.Left == nil && node.Right == nil {
			tmp := append([]byte{}, path...)
			reverse(tmp)
			if result == "" || string(tmp) < result {
				result = string(tmp)
			}
		}
		traverse(node.Left)
		traverse(node.Right)
		path = path[:len(path)-1]
	}

	traverse(root)
	return result
}

func reverse(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 1022. 从根到叶的二进制数之和
// https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/description/
// 给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。
// 例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。
// 对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。
// 返回这些数字之和。题目数据保证答案是一个 32 位 整数。
// 输入：root = [1,0,1,0,1,0,1]
// 输出：22
// 解释：(100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
func sumRootToLeaf(root *TreeNode) int {
	var path []int
	result := 0
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			s := 0
			for _, val := range path {
				s = 2*s + val
			}
			result += s
		}
		traverse(node.Left)
		traverse(node.Right)
		path = path[:len(path)-1]
	}

	traverse(root)
	return result
}

// 1457. 二叉树中的伪回文路径
// https://leetcode.cn/problems/pseudo-palindromic-paths-in-a-binary-tree/
// 给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。
// 请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。
// 输入：root = [2,3,1,3,1,null,1]
// 输出：2
func pseudoPalindromicPaths(root *TreeNode) int {
	var path []int
	result := 0
	var hash [10]int
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		hash[node.Val]++
		if node.Left == nil && node.Right == nil {
			cnt := 0
			for i := 1; i <= 9; i++ {
				if hash[i]%2 == 1 {
					cnt++
				}
			}
			if cnt <= 1 {
				result++
			}
		}
		traverse(node.Left)
		traverse(node.Right)
		path = path[:len(path)-1]
		hash[node.Val]--
	}

	traverse(root)
	return result
}

// 404.左叶子之和
// https://leetcode.cn/problems/sum-of-left-leaves/
// 输入: root = [3,9,20,null,null,15,7]
// 输出: 24
// 后序遍历
func sumOfLeftLeaves(root *TreeNode) int {
	s := 0
	var traverse func(*TreeNode, bool)

	traverse = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && isLeft {
			s += node.Val
		}
		traverse(node.Left, true)
		traverse(node.Right, false)
	}

	if root == nil {
		return 0
	}
	traverse(root, false)
	return s
}

// 递归，有左孩子时，判断一下是否是叶子节点
func sumOfLeftLeaves2(root *TreeNode) int {
	s := 0
	var traverse func(*TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			return
		}
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			s += node.Left.Val
		}
		traverse(node.Left)
		traverse(node.Right)
	}

	if root == nil {
		return 0
	}
	traverse(root)
	return s
}

// 623. 在二叉树中增加一行
// https://leetcode.cn/problems/add-one-row-to-tree/
// 给定一个二叉树的根 root 和两个整数 val 和 depth ，在给定的深度 depth 处添加一个值为 val 的节点行。
// 注意，根节点 root 位于深度 1 。
// 加法规则如下:
// 给定整数 depth，对于深度为 depth - 1 的每个非空树节点 cur ，创建两个值为 val 的树节点作为 cur 的左子树根和右子树根。
// cur 原来的左子树应该是新的左子树根的左子树。
// cur 原来的右子树应该是新的右子树根的右子树。
// 如果 depth == 1 意味着 depth - 1 根本没有深度，那么创建一个树节点，值 val 作为整个原始树的新根，而原始树就是新根的左子树。
// 输入: root = [4,2,6,3,1,5], val = 1, depth = 2
// 输出: [4,1,1,2,null,null,6,3,1,5]
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	var traverse func(node *TreeNode, level int)
	traverse = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == depth-1 {
			newLeft := &TreeNode{Val: val, Left: node.Left}
			newRight := &TreeNode{Val: val, Right: node.Right}
			node.Left = newLeft
			node.Right = newRight
			return
		}
		traverse(node.Left, level+1)
		traverse(node.Right, level+1)
	}
	dummy := &TreeNode{Left: root}
	traverse(dummy, 0)
	return dummy.Left
}

// 971. 翻转二叉树以匹配先序遍历
// 给你一棵二叉树的根节点 root ，树中有 n 个节点，每个节点都有一个不同于其他节点且处于 1 到 n 之间的值。
// 另给你一个由 n 个值组成的行程序列 voyage ，表示 预期 的二叉树 先序遍历 结果。
// 通过交换节点的左右子树，可以 翻转 该二叉树中的任意节点。例，翻转节点 1 的效果如下：
// 输入：root = [1,2], voyage = [2,1]
// 输出：[-1]
// 解释：翻转节点无法令先序遍历匹配预期行程。
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	var results []int
	canFlip := true
	i := 0
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil || !canFlip {
			return
		}
		if node.Val != voyage[i] {
			canFlip = false
			return
		}
		i++
		if node.Left != nil && node.Left.Val != voyage[i] {
			node.Left, node.Right = node.Right, node.Left
			results = append(results, node.Val)
		}
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)
	if canFlip {
		return results
	}
	return []int{-1}
}

// 987. 二叉树的垂序遍历
// https://leetcode.cn/problems/vertical-order-traversal-of-a-binary-tree/description/
// 给你二叉树的根结点 root ，请你设计算法计算二叉树的 垂序遍历 序列。
// 对位于 (row, col) 的每个结点而言，其左右子结点分别位于 (row + 1, col - 1) 和 (row + 1, col + 1) 。树的根结点位于 (0, 0) 。
// 二叉树的 垂序遍历 从最左边的列开始直到最右边的列结束，按列索引每一列上的所有结点，形成一个按出现位置从上到下排序的有序列表。如果同行同列上有多个结点，则按结点的值从小到大进行排序。
// 返回二叉树的 垂序遍历 序列。
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[9],[3,15],[20],[7]]
// 解释：
// 列 -1 ：只有结点 9 在此列中。
// 列  0 ：只有结点 3 和 15 在此列中，按从上到下顺序。
// 列  1 ：只有结点 20 在此列中。
// 列  2 ：只有结点 7 在此列中。
func verticalTraversal(root *TreeNode) [][]int {
	type Triple struct {
		row, col int
		node     *TreeNode
	}
	var nodes []*Triple
	var traverse func(node *TreeNode, i, j int)

	traverse = func(node *TreeNode, i, j int) {
		if node == nil {
			return
		}
		nodes = append(nodes, &Triple{i, j, node})
		traverse(node.Left, i+1, j-1)
		traverse(node.Right, i+1, j+1)
	}

	traverse(root, 0, 0)
	// 排序
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].col == nodes[j].col && nodes[i].row == nodes[j].row {
			return nodes[i].node.Val < nodes[j].node.Val
		}
		if nodes[i].col == nodes[j].col {
			return nodes[i].row < nodes[j].row
		}
		return nodes[i].col < nodes[j].col
	})
	// 分组
	results := [][]int{[]int{nodes[0].node.Val}}
	for i := 1; i < len(nodes); i++ {
		if nodes[i].col == nodes[i-1].col {
			results[len(results)-1] = append(results[len(results)-1], nodes[i].node.Val)
		} else {
			results = append(results, []int{nodes[i].node.Val})
		}
	}
	return results
}

// 993. 二叉树的堂兄弟节点
// https://leetcode.cn/problems/cousins-in-binary-tree/description/
// 在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
// 如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
// 我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
// 只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。
// 输入：root = [1,2,3,4], x = 4, y = 3
// 输出：false
func isCousins(root *TreeNode, x int, y int) bool {
	type Node struct {
		depth  int
		father *TreeNode
	}
	var p, q *Node
	var traverse func(node, father *TreeNode, x, y, depth int)

	traverse = func(node, father *TreeNode, x, y, depth int) {
		if node == nil {
			return
		}
		depth++
		if node.Val == x {
			p = &Node{father: father, depth: depth}
		} else if node.Val == y {
			q = &Node{father: father, depth: depth}
		}
		traverse(node.Left, node, x, y, depth)
		traverse(node.Right, node, x, y, depth)
	}

	if root == nil {
		return true
	}
	traverse(root, nil, x, y, 0)
	if p == nil || q == nil {
		return false
	}
	return p.father != q.father && p.depth == q.depth
}

// 1315. 祖父节点值为偶数的节点和
// https://leetcode.cn/problems/sum-of-nodes-with-even-valued-grandparent/
// 输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
// 输出：18
func sumEvenGrandparent(root *TreeNode) int {
	result := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val%2 == 0 {
			if root.Left != nil {
				if root.Left.Left != nil {
					result += root.Left.Left.Val
				}
				if root.Left.Right != nil {
					result += root.Left.Right.Val
				}
			}
			if root.Right != nil {
				if root.Right.Left != nil {
					result += root.Right.Left.Val
				}
				if root.Right.Right != nil {
					result += root.Right.Right.Val
				}
			}
		}
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return result
}

// 1448. 统计二叉树中好节点的数目
// https://leetcode.cn/problems/count-good-nodes-in-binary-tree/description/
// 给你一棵根为 root 的二叉树，请你返回二叉树中好节点的数目。
// 「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值。
// 输入：root = [3,1,4,3,null,1,5]
// 输出：4
func goodNodes(root *TreeNode) int {
	result := 0
	preMax := math.MinInt
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		tmp := preMax
		if root.Val >= preMax {
			preMax = root.Val
			result++
		}
		dfs(root.Left)
		dfs(root.Right)
		preMax = tmp
	}
	dfs(root)
	return result
}

// 513.找树左下角的值
// https://leetcode.cn/problems/find-bottom-left-tree-value/description/
// 给定一个二叉树，在树的最后一行找到最左边的值。
// 输入: [1,2,3,4,null,5,6,null,null,7]
// 输出: 7
// 思路1:层序遍历
func findBottomLeftValue(root *TreeNode) int {
	result := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		sz := len(q)
		node := q[0]
		q = q[1:]
		result = q[0].Val
		for i := 0; i < sz; i++ {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return result
}

// 思路2:遍历整棵二叉树，用变量记录深度
func findBottomLeftValue3(root *TreeNode) int {
	maxDepth := 0
	depth := 0
	result := 0
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		depth++
		if depth > maxDepth {
			maxDepth = depth
			result = node.Val
		}
		traverse(node.Left)
		traverse(node.Right)
		depth--
	}

	if root == nil {
		return result
	}
	traverse(root)
	return result
}

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

// 386. 字典序排数
// https://leetcode.cn/problems/lexicographical-numbers/
// 给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。
// 你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。
// 输入：n = 13
// 输出：[1,10,11,12,13,2,3,4,5,6,7,8,9]
func lexicalOrder(n int) []int {
	var results []int
	var traverse func(root, n int)

	traverse = func(root, n int) {
		if root > n {
			return
		}
		results = append(results, root)
		for child := root * 10; child < root*10+10; child++ {
			traverse(child, n)
		}
	}

	for i := 1; i <= 9; i++ {
		traverse(i, n)
	}
	return results
}

// 1104. 二叉树寻路
// https://leetcode.cn/problems/path-in-zigzag-labelled-binary-tree/description/
// 在一棵无限的二叉树上，每个节点都有两个子节点，树中的节点 逐行 依次按 “之” 字形进行标记。
// 如下图所示，在奇数行（即，第一行、第三行、第五行……）中，按从左到右的顺序进行标记；
// 而偶数行（即，第二行、第四行、第六行……）中，按从右到左的顺序进行标记。
func pathInZigZagTree(label int) []int {
	var path []int
	for label >= 1 {
		path = append(path, label)
		label /= 2
		depth := log(label)
		rangeVals := getLevelRange(depth)
		label = rangeVals[1] - (label - rangeVals[0])
	}
	reverseInts(path)
	return path
}

func log(x int) int { return int(math.Log(float64(x)) / math.Log(float64(2))) }

func getLevelRange(n int) []int {
	p := int(math.Pow(2, float64(n)))
	return []int{p, 2*p - 1}
}
func reverseInts(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// 1145. 二叉树着色游戏
// 有两位极客玩家参与了一场「二叉树着色」的游戏。游戏中，给出二叉树的根节点 root，树上总共有 n 个节点，且 n 为奇数，其中每个节点上的值从 1 到 n 各不相同。
// 最开始时：
// 「一号」玩家从 [1, n] 中取一个值 x（1 <= x <= n）；
// 「二号」玩家也从 [1, n] 中取一个值 y（1 <= y <= n）且 y != x。
// 「一号」玩家给值为 x 的节点染上红色，而「二号」玩家给值为 y 的节点染上蓝色。
// 之后两位玩家轮流进行操作，「一号」玩家先手。每一回合，玩家选择一个被他染过色的节点，将所选节点一个 未着色 的邻节点（即左右子节点、或父节点）进行染色（「一号」玩家染红色，「二号」玩家染蓝色）。
// 如果（且仅在此种情况下）当前玩家无法找到这样的节点来染色时，其回合就会被跳过。
// 若两个玩家都没有可以染色的节点时，游戏结束。着色节点最多的那位玩家获得胜利 ✌️。
// 现在，假设你是「二号」玩家，根据所给出的输入，假如存在一个 y 值可以确保你赢得这场游戏，则返回 true ；若无法获胜，就请返回 false 。
// 输入：root = [1,2,3,4,5,6,7,8,9,10,11], n = 11, x = 3
// 输出：true
// 解释：第二个玩家可以选择值为 2 的节点。
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	node := find(root, x)
	leftCount := countNode(node.Left)
	rightCount := countNode(node.Right)
	otherCount := n - 1 - leftCount - rightCount

	return max(leftCount, max(rightCount, otherCount)) > n/2
}

// 定义：在以 root 为根的二叉树中搜索值为 x 的节点并返回
func find(root *TreeNode, x int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == x {
		return root
	}
	// 去左子树找
	left := find(root.Left, x)
	if left != nil {
		return left
	}
	// 左子树找不到的话去右子树找
	return find(root.Right, x)
}

// 定义：计算以 root 为根的二叉树的节点总数
func countNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNode(root.Left) + countNode(root.Right)
}

// 572. 另一棵树的子树
// https://leetcode.cn/problems/subtree-of-another-tree/
// 输入：root = [3,4,5,1,2], subRoot = [4,1,2]
// 输出：true
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var isSameTree func(p, q *TreeNode) bool
	isSameTree = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	if root == nil {
		return false
	}
	// 中
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot) // 左右
}

// 1367. 二叉树中的链表
// https://leetcode.cn/problems/linked-list-in-binary-tree/description/
// 给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。
// 如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。
// 一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。
func isSubPath(head *ListNode, root *TreeNode) bool {
	// 思路：遍历二叉树的所有节点，每个节点用 check 函数判断是否能够将链表嵌进去。
	var check func(head *ListNode, root *TreeNode) bool
	check = func(head *ListNode, root *TreeNode) bool {
		if head == nil {
			return true
		}
		if root == nil {
			return false
		}
		if head.Val == root.Val {
			return check(head.Next, root.Left) || check(head.Next, root.Right)
		}
		return false
	}

	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	// 中
	if check(head, root) {
		return true
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right) // 左右
}

func main() {
	fmt.Println("hello world")
}
