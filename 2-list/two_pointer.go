package main

import (
	"container/heap"
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
		if cur.Val < x {
			p1.Next = cur
			p1 = p1.Next
		} else {
			p2.Next = cur
			p2 = p2.Next
		}
		next := cur.Next
		cur.Next = nil
		cur = next
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}

// 23.合并 k 个有序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/
// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//
//	1->4->5,
//	1->3->4,
//	2->6
//
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{}
	cur := dummy
	pq := &PriorityQueue{}
	heap.Init(pq)
	for _, head := range lists {
		if head != nil {
			heap.Push(pq, head)
		}
	}
	for pq.Len() > 0 {
		node := heap.Pop(pq).(*ListNode)
		cur.Next = node
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
		cur = cur.Next
	}
	return dummy.Next
}

// 378. 有序矩阵中第 K 小的元素
// https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/
// 给你一个 n x n 矩阵 matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
// 请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。
// 你必须找到一个内存复杂度优于 O(n2) 的解决方案。
// 输入：matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
// 输出：13
// 解释：矩阵中的元素为 [1,5,9,10,11,12,13,13,15]，第 8 小元素是 13
type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	// 初始化优先级队列
	for i := 0; i < len(matrix); i++ {
		heap.Push(h, matrix[i])
	}
	result := -1
	for h.Len() > 0 && k > 0 {
		nums := heap.Pop(h).([]int)
		result = nums[0]
		k--
		if len(nums) > 1 {
			heap.Push(h, append([]int{}, nums[1:]...))
		}
	}
	return result
}

// 面试题 02.02. 返回倒数第 k 个节点
// https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/description/
func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return 0
	}
	p1 := head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2.Val
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
