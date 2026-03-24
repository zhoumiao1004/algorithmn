package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*1、合并两个有序链表
2、链表的分解
3、合并 k 个有序链表
4、寻找单链表的倒数第 k 个节点
5、寻找单链表的中点
6、判断单链表是否包含环并找出环起点
7、判断两个单链表是否相交并找出交点*/

// 21. 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy *ListNode
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

// 86. 分隔链表
// https://leetcode.cn/problems/partition-list/description/
// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你应当 保留 两个分区中每个节点的初始相对位置。
// 输入：head = [1,4,3,2,5,2], x = 3
// 输出：[1,2,2,4,3,5]
func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	p1, p2 := dummy1, dummy2
	cur := head
	for cur != nil {
		next := cur.Next // 记录下个需要遍历的节点，防止丢失
		if cur.Val < x {
			p1.Next = cur
			p1 = p1.Next
		} else {
			p2.Next = cur
			p2 = p2.Next
		}
		cur.Next = nil // 修改节点的指针，断掉和原链表的联系，节点加入p1/p2中
		cur = next
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}

// 面试题 02.02. 返回倒数第 k 个节点
// https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/description/
func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return 0
	}
	// fast先走k步
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	// slow 和 fast 同时走 n - k 步
	slow := head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

// 19.删除链表的倒数第N个节点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
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

// 思路2: 删除倒数第 n 个，要先找倒数第 n + 1 个节点（要删除节点前面一个节点）
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head} // 前面加一个dummy节点
	node := findFromLast(dummy, n+1)
	node.Next = node.Next.Next
	return dummy.Next
}

func findFromLast(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// fast先走k步
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	// slow 和 fast 同时走 n - k 步
	slow := head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// LCR 140. 训练计划 II
// 给定一个头节点为 head 的链表用于记录一系列核心肌群训练项目编号，请查找并返回倒数第 cnt 个训练项目编号。
// 输入：head = [2,4,7,8], cnt = 1
// 输出：8
func trainingPlan(head *ListNode, cnt int) *ListNode {
	if head == nil {
		return nil
	}
	p1 := head
	for i := 0; i < cnt; i++ {
		p1 = p1.Next
	}
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

// 876. 链表的中间结点
// https://leetcode.cn/problems/middle-of-the-linked-list/description/
// 输入：head = [1,2,3,4,5]
// 输出：[3,4,5]
// 解释：链表只有一个中间结点，值为 3 。
// 输入：head = [1,2,3,4,5,6]
// 输出：[4,5,6]
// 解释：该链表有两个中间结点，值分别为 3 和 4 ，返回第二个结点。
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 141. 环形链表
// https://leetcode.cn/problems/linked-list-cycle/
// 给你一个链表的头节点 head ，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
// 如果链表中存在环 ，则返回 true 。 否则，返回 false 。
// 输入：head = [3,2,0,-4], pos = 1
// 输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 142. 环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii/
// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
// 不允许修改 链表。
// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点
// 解释：链表中有一个环，其尾部连接到第二个节点。
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			cur := head
			for cur != slow {
				cur = cur.Next
				slow = slow.Next
			}
			return cur
		}
	}
	return nil
}

// 160. 两个链表是否相交
// https://leetcode.cn/problems/intersection-of-two-linked-lists/
// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	if p1 == nil || p2 == nil {
		return nil
	}
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

func main() {
	fmt.Println(kthSmallest([][]int{[]int{1, 5, 9}, []int{10, 11, 13}, []int{12, 13, 15}}, 8))
}
