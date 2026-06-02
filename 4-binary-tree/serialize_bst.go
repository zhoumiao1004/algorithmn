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

// 449. 序列化和反序列化二叉搜索树
// https://leetcode.cn/problems/serialize-and-deserialize-bst/description/
// 序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
// 设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。
// 编码的字符串应尽可能紧凑。
type Codec struct {
	SEP string
}

func Constructor() Codec {
	return Codec{SEP: ","}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var serialize func(root *TreeNode) string

	serialize = func(root *TreeNode) string {
		if root == nil {
			return ""
		}
		left := serialize(root.Left)
		right := serialize(root.Right)
		return fmt.Sprintf("%d,%s,%s", root.Val, left, right)
	}

	return serialize(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var build func(nums []int) *TreeNode

	build = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		if nums[0] == -1 {
			return nil
		}
		root := &TreeNode{Val: nums[0]}
		i := 1
		for i < len(nums) && nums[i] < nums[0] {
			i++
		}
		root.Left = build(nums[1:i])
		root.Right = build(nums[i:])
		return root
	}

	var nums []int
	for _, s := range strings.Split(data, ",") {
		val := -1
		if s != "" {
			val, _ = strconv.Atoi(s)
		}
		nums = append(nums, val)
	}
	return build(nums)
}
