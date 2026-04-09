package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 297. 二叉树的序列化与反序列化
// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/description/
// 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
// 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
type Codec struct {
	SEP  string
	NULL string
}

func Constructor() Codec {
	return Codec{
		SEP:  ",",
		NULL: "#",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var path []string
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			path = append(path, this.NULL)
			return
		}
		path = append(path, fmt.Sprintf("%d", node.Val))
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)
	return strings.Join(path, this.SEP)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	nodes := strings.Split(data, this.SEP)
	var build func() *TreeNode

	build = func() *TreeNode {
		if len(nodes) == 0 {
			return nil
		}
		first := nodes[0]
		nodes = nodes[1:]
		if first == this.NULL {
			return nil
		}
		val, _ := strconv.Atoi(first)
		root := &TreeNode{Val: val}
		root.Left = build()
		root.Right = build()
		return root
	}

	root := build()
	return root
}

// 思路4：层序遍历
func (this *Codec) serialize4(root *TreeNode) string {
	var path []string
	if root == nil {
		return this.NULL
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			path = append(path, this.NULL)
		} else {
			path = append(path, fmt.Sprintf("%d", node.Val))
			q = append(q, node.Left, node.Right)
		}
	}
	return strings.Join(path, this.SEP)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize4(data string) *TreeNode {

	strs := strings.Split(data, this.SEP)
	var nodes []*TreeNode
	for _, s := range strs {
		if s == this.NULL {
			nodes = append(nodes, nil)
			continue
		}
		val, _ := strconv.Atoi(s)
		nodes = append(nodes, &TreeNode{Val: val})
	}

	root := nodes[0]
	q := []*TreeNode{root}
	i := 1
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			continue
		}
		node.Left = nodes[i]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		node.Right = nodes[i+1]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		i += 2
	}
	return root
}

func main() {
	rootArr := []interface{}{1, 2, 3, nil, nil, 4, 5}
	root := BuildTree(rootArr)
	c := Constructor()
	fmt.Println(c.serialize4(root))
	// c.deserialize4()
}

// BuildTree 根据层序遍历数组构建二叉树
func BuildTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	// 创建根节点
	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		// 取出队列首个节点
		node := queue[0]
		queue = queue[1:]

		// 处理左子节点
		if i < len(arr) && arr[i] != nil {
			node.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// 处理右子节点
		if i < len(arr) && arr[i] != nil {
			node.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
