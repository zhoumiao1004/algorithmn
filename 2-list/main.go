package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List interface {
	Get(index int) int
	AddAtHead(val int)
	AddAtTail(val int)
	AddAtIndex(index int, val int)
	DeleteAtIndex(index int)
}

func newMyLinkedList() List {
	return &MyLinkedList{
		dummy: &ListNode{},
		Size:  0,
	}
}

// 707.设计链表
// https://leetcode.cn/problems/design-linked-list/description/
type MyLinkedList struct {
	dummy *ListNode
	Size  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummy: &ListNode{},
		Size:  0,
	}
}

func (l *MyLinkedList) String() string {
	var vals []string
	cur := l.dummy.Next
	for cur != nil {
		vals = append(vals, fmt.Sprintf("%d", cur.Val))
		cur = cur.Next
	}
	return strings.Join(vals, "->")
}
func (l *MyLinkedList) Get(index int) int {
	if l == nil || index < 0 || index >= l.Size {
		return -1
	}
	cur := l.dummy.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{Val: val}
	node.Next = l.dummy.Next
	l.dummy.Next = node
	l.Size++
}

func (l *MyLinkedList) AddAtTail(val int) {
	cur := l.dummy
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
	l.Size++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > l.Size { // 注意这里是大于
		return
	}
	cur := l.dummy
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	node := &ListNode{Val: val, Next: cur.Next}
	cur.Next = node
	l.Size++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if l == nil || index < 0 || index >= l.Size {
		return
	}
	cur := l.dummy
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	l.Size--
}

// 203.移除链表元素
// https://leetcode.cn/problems/remove-linked-list-elements/description/
// 输入：head = [1,2,6,3,4,5,6], val = 6
// 输出：[1,2,3,4,5]
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := head
	for fast != nil {
		if fast.Val != val {
			fast = fast.Next
			slow = slow.Next
		} else {
			for fast != nil && fast.Val == val {
				fast = fast.Next
			}
			slow.Next = fast
		}
	}
	return dummy.Next
}

func removeElements2(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val != val {
			cur = cur.Next
		} else {
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		}
	}
	return dummy.Next
}

// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list/
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
	cur, next := head, head.Next
	for n > 0 {
		cur.Next = prev
		prev = cur
		cur = next
		if next != nil {
			next = next.Next
		}
		n--
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

// 19.删除链表的倒数第N个节点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 面试题 02.07. 链表相交
// https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/description/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != nil || p2 != nil {
		if p1 == nil {
			p1 = headB
		}
		if p2 == nil {
			p2 = headA
		}
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return nil
}

// 234. 回文链表
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

// 后序遍历，模仿双指针实现回文判断的功能
func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return false
	}
	left := head
	result := true
	var dfs func(right *ListNode)
	dfs = func(right *ListNode) {
		if right == nil {
			return
		}
		dfs(right.Next)
		// 后序遍历位置
		if left.Val != right.Val {
			result = false
		}
		left = left.Next
	}
	dfs(head)
	return result
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
