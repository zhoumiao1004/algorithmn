package main

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
