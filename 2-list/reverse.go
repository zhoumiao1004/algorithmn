package main

import "fmt"

// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list/
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// 思路2:递归
func reverseListRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 92. 反转链表 II
// https://leetcode.cn/problems/reverse-linked-list-ii/
// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	// 找到第m-1个节点
	prev := head
	for i := 1; i < left-1; i++ {
		prev = prev.Next
	}
	prev.Next = reverseN(prev.Next, right-left+1)
	return head
}

// 反转链表前 N 个节点
func reverseN(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for i := 0; i < n; i++ {
		if cur == nil {
			break
		}
		next := cur.Next
		cur.Next = prev // 修改cur指向，从指后修改为指前
		prev = cur      // prev向后移
		cur = next      // cur向后移
	}
	head.Next = cur
	return prev
}

// 25. K 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	newHead := reverseN(head, k)
	head.Next = reverseKGroup(cur, k)
	return newHead
}

func main() {
	fmt.Println("hello world")
}
