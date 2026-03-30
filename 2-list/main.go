package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 24. 两两交换链表中的节点
// https://leetcode.cn/problems/swap-nodes-in-pairs/description/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(head.Next.Next)
	next.Next = head
	return next
}

// 143.重排链表
// https://leetcode.cn/problems/reorder-list/submissions/
// 输入：head = [1,2,3,4]
// 输出：[1,4,2,3]
// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
// L0 → L1 → … → Ln - 1 → Ln
// 请将其重新排列后变为：
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
func reorderList(head *ListNode) {
	mid := getMidNode(head)
	head2 := reverseList(mid.Next)
	mid.Next = nil
	p1, p2 := head, head2
	for p2 != nil {
		next1 := p1.Next
		next2 := p2.Next
		p1.Next = p2
		p2.Next = next1
		p1 = next1
		p2 = next2
	}
}

func getMidNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	return slow
}

// 2. 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	dummy := &ListNode{}
	cur := dummy
	flag := 0
	for p1 != nil || p2 != nil {
		val := flag
		if p1 != nil {
			val += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			val += p2.Val
			p2 = p2.Next
		}
		flag = val / 10
		cur.Next = &ListNode{Val: val % 10}
		cur = cur.Next
	}
	if flag == 1 {
		cur.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}

// 445. 两数相加 II
// https://leetcode.cn/problems/add-two-numbers-ii/description/
// 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
// 输入：l1 = [7,2,4,3], l2 = [5,6,4]
// 输出：[7,8,0,7]
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var st1, st2 []int
	for cur := l1; cur != nil; cur = cur.Next {
		st1 = append(st1, cur.Val)
	}
	for cur := l2; cur != nil; cur = cur.Next {
		st2 = append(st2, cur.Val)
	}
	dummy := &ListNode{}
	carry := 0
	for len(st1) > 0 || len(st2) > 0 {
		val := carry
		if len(st1) > 0 {
			val += st1[len(st1)-1]
			st1 = st1[:len(st1)-1]
		}
		if len(st2) > 0 {
			val += st2[len(st2)-1]
			st2 = st2[:len(st2)-1]
		}
		carry = val / 10
		// 头插法
		next := dummy.Next
		dummy.Next = &ListNode{Val: val % 10, Next: next}
	}
	if carry == 1 {
		next := dummy.Next
		dummy.Next = &ListNode{Val: 1, Next: next}
	}
	return dummy.Next
}

func main() {
	myList := Constructor()
	myList.AddAtTail(3)
	myList.AddAtHead(2)
	myList.AddAtHead(1)
	myList.AddAtTail(4)
	myList.AddAtTail(5)
	fmt.Println(myList.String())
	fmt.Println(myList.Get(1))
	myList.AddAtIndex(2, 9)
	fmt.Println(myList.String())
	myList.DeleteAtIndex(4)
	fmt.Println(myList.String())
}
