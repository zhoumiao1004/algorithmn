package main

import "container/heap"

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

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Val < pq[j].Val }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

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

// 373. 查找和最小的 K 对数字
// https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/
// 给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
// 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [[1,2],[1,4],[1,6]]
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := &IntHeap2{}
	heap.Init(pq)
	for i := 0; i < len(nums1); i++ {
		// 存储三元组 (num1[i], nums2[i], i)
		// i 记录 nums2 元素的索引位置，用于生成下一个节点
		heap.Push(pq, []int{nums1[i], nums2[0], 0})
	}
	var result [][]int
	// 执行合并多个有序链表的逻辑
	for pq.Len() > 0 && k > 0 {
		cur := heap.Pop(pq).([]int)
		k--
		// 链表中的下一个节点加入优先级队列
		nextIndex := cur[2] + 1
		if nextIndex < len(nums2) {
			heap.Push(pq, []int{cur[0], nums2[nextIndex], nextIndex})
		}
		result = append(result, []int{cur[0], cur[1]})
	}
	return result
}

type IntHeap2 [][]int

func (h IntHeap2) Len() int           { return len(h) }
func (h IntHeap2) Less(i, j int) bool { return h[i][0]+h[i][1] < h[j][0]+h[j][1] }
func (h IntHeap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap2) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *IntHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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

// 287. 寻找重复数
// https://leetcode.cn/problems/find-the-duplicate-number/description/
// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
// 输入：nums = [1,3,4,2,2]
// 输出：2
func findDuplicate(nums []int) int {
    slow, fast := 0, 0
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if fast == slow {
			break
		}
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
		if slow == fast {
			return slow
		}
	}
	return -1
}