package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	var strs []string
	for cur := l; cur != nil; cur = cur.Next {
		strs = append(strs, fmt.Sprintf("%d", cur.Val))
	}
	fmt.Println(strings.Join(strs, "->"))
}

func (l *ListNode) String() string {
	var strs []string
	for cur := l; cur != nil; cur = cur.Next {
		strs = append(strs, fmt.Sprintf("%d", cur.Val))
	}
	return strings.Join(strs, "->")
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
// 思路1：反转后半部分链表，再拉拉链合并2个列表
func reorderList(head *ListNode) {
	var getMidNode func(node *ListNode) *ListNode
	var reverseList func(node *ListNode) *ListNode

	getMidNode = func(node *ListNode) *ListNode {
		if node == nil {
			return nil
		}
		dummy := &ListNode{Next: node}
		slow, fast := dummy, dummy
		// slow, fast := head, head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}

	reverseList = func(node *ListNode) *ListNode {
		if node == nil {
			return nil
		}
		var prev *ListNode
		cur := node
		for cur != nil {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}
		return prev
	}

	mid := getMidNode(head)
	fmt.Println(mid.Val)
	head2 := reverseList(mid.Next)
	mid.Next = nil
	fmt.Println(head.Val, head2.Val)
	p1, p2 := head, head2
	for p1 != nil && p2 != nil {
		next1 := p1.Next
		next2 := p2.Next
		p2.Next = next1
		p1.Next = p2
		p1 = next1
		p2 = next2
	}
}

// 思路2：栈
func reorderList2(head *ListNode) {
	var st []*ListNode
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		st = append(st, cur)
		n++
	}

	cur := head
	for cur != nil {
		last := st[len(st)-1]
		st = st[:len(st)-1]
		next := cur.Next                       // 保存下个节点
		if last == next || last.Next == next { // 奇数：last == next, 偶数：last.Next == next
			last.Next = nil
			break
		}
		last.Next = next
		cur.Next = last
		cur = next
	}
}

func main() {
	head := buildList([]int{1, 2, 3, 4, 5})
	head.Print()
	reorderList2(head)
	head.Print()
}

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return dummy.Next
}
