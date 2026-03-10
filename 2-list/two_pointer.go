package main

import (
	"container/heap"
	"sort"
)

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

// 977.有序数组的平方
// https://leetcode.cn/problems/squares-of-a-sorted-array/description/
// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
func sortedSquares(nums []int) []int {
	n := len(nums)
	results := make([]int, n)
	left, right := 0, n-1
	k := n - 1
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			results[k] = nums[right] * nums[right]
			right--
		} else {
			results[k] = nums[left] * nums[left]
			left++
		}
		k--
	}
	return results
}

// 1329. 将矩阵按对角线排序
// https://leetcode.cn/problems/sort-the-matrix-diagonally/
// 输入：mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
// 输出：[[1,1,1,1],[1,2,2,2],[1,2,3,3]]
func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	diaMap := make(map[int][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := i - j
			diaMap[k] = append(diaMap[k], mat[i][j])
		}
	}
	for _, v := range diaMap {
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
	}
	// 结果回填到矩阵
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr := diaMap[i-j]
			mat[i][j] = arr[len(arr)-1]
			diaMap[i-j] = arr[:len(arr)-1]
		}
	}
	return mat
}

// 1260. 二维网格迁移
// https://leetcode.cn/problems/shift-2d-grid/description/
// 给你一个 m 行 n 列的二维网格 grid 和一个整数 k。你需要将 grid 迁移 k 次。
// 每次「迁移」操作将会引发下述活动：
// 位于 grid[i][j]（j < n - 1）的元素将会移动到 grid[i][j + 1]。
// 位于 grid[i][n - 1] 的元素将会移动到 grid[i + 1][0]。
// 位于 grid[m - 1][n - 1] 的元素将会移动到 grid[0][0]。
// 请你返回 k 次迁移操作后最终得到的 二维网格。
// 输入：grid = [[1,2,3],[4,5,6],[7,8,9]], k = 1
// 输出：[[9,1,2],[3,4,5],[6,7,8]]
// 1.除最后一列向右移1位 2.最后一列一到第一列 3.右下角移到左上角
func shiftGrid(grid [][]int, k int) [][]int {

}

// 867. 转置矩阵
// https://leetcode.cn/problems/transpose-matrix/
// 给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。
// 矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[[1,4,7],[2,5,8],[3,6,9]]
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	results := make([][]int, n)
	for i := 0; i < n; i++ {
		results[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			results[j][i] = matrix[i][j]
		}
	}
	return results
}

// 14. 最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	left := 0 // 相同的列
	m, n := len(strs), len(strs[0])
	for j := 0; j < n; j++ {
		// 第j列，对比每一行是否相同
		for i := 1; i < m; i++ {
			if len(strs[i]) <= j || strs[i][j] != strs[0][j] {
				return strs[0][:left]
			}
		}
		left++
	}
	return strs[0][:left]
}
