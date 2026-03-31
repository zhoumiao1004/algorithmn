package main

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
