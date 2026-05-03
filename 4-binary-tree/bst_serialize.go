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

type Codec struct {
	SEP string
}

func Constructor() Codec {
	return Codec{SEP: ","}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []string
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, fmt.Sprintf("%d", root.Val))
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return strings.Join(res, this.SEP)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var nums []int
	for _, s := range strings.Split(data, this.SEP) {
		if s == "" {
			continue
		}
		val, _ := strconv.Atoi(s)
		nums = append(nums, val)
	}
	var build func(nums []int) *TreeNode
	build = func(nums []int) *TreeNode {
		n := len(nums)
		if n == 0 {
			return nil
		}
		root := &TreeNode{Val: nums[0]}
		mid := 1
		for mid < n && nums[mid] < root.Val {
			mid++
		}
		root.Left = build(nums[1:mid])
		root.Right = build(nums[mid:])
		return root
	}
	return build(nums)
}

func main() {
	fmt.Println(strings.Join([]string{}, ","))
	fmt.Println(strings.Split("", ","))
}
