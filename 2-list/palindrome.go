package main

import "fmt"

// 234. 回文链表
// https://leetcode.cn/problems/palindrome-linked-list/
// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
// 输入：head = [1,2,2,1]
// 输出：true
// 找到中点，反转后半部份链表，双指针比较
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	mid := getMidNode(head)
	next := mid.Next
	mid.Next = nil
	head2 := reverseList(next)
	p1, p2 := head, head2
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

// 思路2:递归+后序遍历，模仿双指针实现回文判断的功能
func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return false
	}
	left := head
	result := true
	var traverse func(right *ListNode)
	traverse = func(right *ListNode) {
		if right == nil {
			return
		}
		traverse(right.Next)
		// 后序遍历位置
		if left.Val != right.Val {
			result = false
		}
		left = left.Next
	}
	traverse(head)
	return result
}

func main() {
	fmt.Println("hello wolrd")
}
