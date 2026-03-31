package main

import "strings"

// 331. 验证二叉树的前序序列化
// 序列化二叉树的一种方法是使用 前序遍历 。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
// 输入: preorder = "9,3,4,#,#,1,#,#,2,#,6,#,#"
// 输出: true
func isValidSerialization(preorder string) bool {
	edge := 1
	for _, c := range strings.Split(preorder, ",") {
		if c == "#" {
			edge--
			if edge < 0 {
				return false
			}
		} else {
			edge--
			if edge < 0 {
				return false
			}
			edge += 2
		}
	}
	return edge == 0
}

// 297. 二叉树的序列化与反序列化
// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/description/
// 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
// 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
type Codec struct {
}

func Constructor() Codec {

}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

}
