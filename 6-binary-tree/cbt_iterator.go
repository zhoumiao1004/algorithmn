package main

// 919. 完全二叉树插入器
// https://leetcode.cn/problems/complete-binary-tree-inserter/description/
// 完全二叉树 是每一层（除最后一层外）都是完全填充（即，节点数达到最大）的，并且所有的节点都尽可能地集中在左侧。
// 设计一种算法，将一个新节点插入到一棵完全二叉树中，并在插入后保持其完整。
// 实现 CBTInserter 类:
// CBTInserter(TreeNode root) 使用头节点为 root 的给定树初始化该数据结构；
// CBTInserter.insert(int v)  向树中插入一个值为 Node.val == val的新节点 TreeNode。使树保持完全二叉树的状态，并返回插入节点 TreeNode 的父节点的值；
// CBTInserter.get_root() 将返回树的头节点。
type CBTInserter struct {
	q    []*TreeNode
	root *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	q := []*TreeNode{}
	tmp := []*TreeNode{root}
	for len(tmp) > 0 {
		cur := tmp[0]
		tmp = tmp[1:]
		if cur.Left != nil {
			tmp = append(tmp, cur.Left)
		}
		if cur.Right != nil {
			tmp = append(tmp, cur.Right)
		}
		if cur.Right == nil || cur.Left == nil {
			q = append(q, cur) // 找到完全二叉树底部可以进行插入的节点
		}
	}
	return CBTInserter{q: q, root: root}
}

func (this *CBTInserter) Insert(val int) int {
	node := &TreeNode{Val: val}
	cur := this.q[0]
	if cur.Left == nil {
		cur.Left = node
	} else if cur.Right == nil {
		cur.Right = node
		this.q = this.q[1:]
	}
	// 新节点的左右节点也是可以插入的
	this.q = append(this.q, node)
	return cur.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
