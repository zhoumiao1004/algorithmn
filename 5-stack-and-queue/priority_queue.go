package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
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

func main() {
	fmt.Println(kthSmallest([][]int{[]int{1, 5, 9}, []int{10, 11, 13}, []int{12, 13, 15}}, 8))
}
