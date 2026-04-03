package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 347. 前 K 个高频元素
// https://leetcode.cn/problems/top-k-frequent-elements/description/
// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
// 输入：nums = [1,1,1,2,2,3], k = 2
// 输出：[1,2]
// 思路1: hashmap统计元素频率，再按频率排序，时间复杂度O(nlogn)
func topKFrequent(nums []int, k int) []int {
	var results []int
	cntMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cntMap[nums[i]]++
	}
	for key := range cntMap {
		results = append(results, key)
	}

	sort.Slice(results, func(i, j int) bool {
		return cntMap[results[i]] > cntMap[results[j]]
	})
	return results[:k]
}

// 思路2: 小顶堆，时间复杂度: O(nlogk) 空间复杂度: O(n)
type IHeap [][2]int

func (h IHeap) Len() int            { return len(h) }
func (h IHeap) Less(i, j int) bool  { return h[i][1] < h[j][1] } // 按频率比较大小
func (h IHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent2(nums []int, k int) []int {
	m := make(map[int]int) // 统计每个元素的次数(频率)
	for _, val := range nums {
		m[val]++
	}
	// 所有元素依次入堆，堆内仅维持k个最大元素
	h := &IHeap{}
	heap.Init(h)
	for key, val := range m {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h) // 二叉堆中超过k个元素就pop最小元素
		}
	}
	result := make([]int, k) // 按频率高到低返回
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).([2]int)[0]
	}
	return result
}

// 378. 有序矩阵中第 K 小的元素
// https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/
// 给你一个 n x n 矩阵 matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
// 请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。
// 你必须找到一个内存复杂度优于 O(n2) 的解决方案。
// 输入：matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
// 输出：13
// 解释：矩阵中的元素为 [1,5,9,10,11,12,13,13,15]，第 8 小元素是 13
type NumsHeap [][]int

func (h NumsHeap) Len() int            { return len(h) }
func (h NumsHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] } // 比较第一个元素的大小
func (h NumsHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *NumsHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *NumsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 思路：小顶堆，全部入堆后，pop第k次就是第k小的元素
func kthSmallest(matrix [][]int, k int) int {
	h := &NumsHeap{}
	heap.Init(h)
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

// 23.合并 k 个有序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/
// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [1->4->5, 1->3->4, 2->6]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
// 思路1：优先级队列，时间复杂度O(Nlogk)
type ListHeap []*ListNode

func (pq ListHeap) Len() int            { return len(pq) }
func (pq ListHeap) Less(i, j int) bool  { return pq[i].Val < pq[j].Val }
func (pq ListHeap) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *ListHeap) Push(x interface{}) { *pq = append(*pq, x.(*ListNode)) }
func (pq *ListHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// 思路：小顶堆
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{}
	cur := dummy

	h := &ListHeap{}
	heap.Init(h)
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		cur.Next = node
		if node.Next != nil {
			heap.Push(h, node.Next) // 加入链表后面的元素
		}
		cur = cur.Next
	}
	return dummy.Next
}

// 思路2：分治,时间复杂度O(Nlogk), k 条链表分别遍历 O(logk) 次; 空间复杂度，该算法的空间复杂度只有递归树堆栈的开销，也就是 O(logk)，要优于优先级队列解法的 O(k)
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeKLists3(lists, 0, len(lists)-1)
}

// 定义：合并 lists[start..end] 为一个有序链表
func mergeKLists3(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	mid := start + (end-start)/2
	// 合并左半边 lists[start..mid] 为一个有序链表
	left := mergeKLists3(lists, start, mid)
	// 合并右半边 lists[mid+1..end] 为一个有序链表
	right := mergeKLists3(lists, mid+1, end)
	// 合并左右两个有序链表
	return mergeTwoLists(left, right)
}

// 373. 查找和最小的 K 对数字
// https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/
// 给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
// 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [[1,2],[1,4],[1,6]]
type NumsHeap2 [][]int

func (h NumsHeap2) Len() int            { return len(h) }
func (h NumsHeap2) Less(i, j int) bool  { return h[i][0]+h[i][1] < h[j][0]+h[j][1] } // 比较和的大小
func (h NumsHeap2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *NumsHeap2) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *NumsHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := &NumsHeap2{}
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

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }      // 小顶堆
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] } // 绑定swap方法，交换两个元素位置
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0} // 创建slice
	heap.Init(h)                                // 初始化heap
	fmt.Println(*h)
	fmt.Println(heap.Pop(h)) // 调用pop
	heap.Push(h, 6)          // 调用push
	fmt.Println(*h)
	for len(*h) > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	fmt.Println(kthSmallest([][]int{[]int{1, 5, 9}, []int{10, 11, 13}, []int{12, 13, 15}}, 8))
	fmt.Println(topKFrequent2([]int{1, 1, 1, 2, 2, 3}, 2))
}
