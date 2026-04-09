package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 116. 填充每个节点的下一个右侧节点指针
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/description/
// 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下
// 输入：root = [1,2,3,4,5,6,7]
// 输出：[1,#,2,3,#,4,5,6,7,#]
// 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。
// 思路：遍历
func connect(root *Node) *Node {
	var traverse func(p, q *Node)

	traverse = func(p, q *Node) {
		if p == nil || q == nil {
			return
		}
		p.Next = q
		traverse(p.Left, p.Right)
		traverse(p.Right, q.Left)
		traverse(q.Left, q.Right)
	}

	if root == nil {
		return nil
	}
	traverse(root.Left, root.Right)
	return root
}

func connect2(root *Node) *Node {
	var traverse func(root *Node)

	traverse = func(root *Node) {
		if root == nil {
			return
		}
		// 前序位置
		if root.Left != nil {
			root.Left.Next = root.Right
		}
		if root.Right != nil {
			if root.Next != nil {
				root.Right.Next = root.Next.Left
			} else {
				root.Right.Next = nil
			}
		}
		traverse(root.Left)  // 左
		traverse(root.Right) // 右
	}

	traverse(root)
	return root
}

// 117. 填充每个节点的下一个右侧节点指针 II
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/description/
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
// 初始状态下，所有 next 指针都被设置为 NULL 。
// 思路：层序遍历
func connectII(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		sz := len(q)

		for i := 0; i < sz; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if i != len(q)-1 {
				node.Next = q[i+1]
			}
		}
	}
	return root
}
