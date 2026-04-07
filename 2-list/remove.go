package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 27. 移除元素
// https://leetcode.cn/problems/remove-element/description/
// nums = [3,2,2,3], val = 3
// 输出: 2, nums = [2,2,_,_]
// 注意：和26题有序数组去重的解法有一个细节差异，我们这里是先给 nums[slow] 赋值然后再给 slow++，这样可以保证 nums[0..slow-1] 是不包含值为 val 的元素的，最后的结果数组长度就是 slow。
func removeElement(nums []int, val int) int {
	left := 0 // 维护 nums[0..slow] 左开右闭，为不包含val元素的结果子数组
	for right := 0; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

// 283. 移动零
// https://leetcode.cn/problems/move-zeroes/description/
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func moveZeroes(nums []int) {
	left := 0 // 维护 nums[0..slow] 左开右闭，为不包含0的结果子数组
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[left] = nums[i]
			left++
		}
	}
	for ; left < len(nums); left++ {
		nums[left] = 0
	}
}

// 203.移除链表元素
// https://leetcode.cn/problems/remove-linked-list-elements/description/
// 输入：head = [1,2,6,3,4,5,6], val = 6
// 输出：[1,2,3,4,5]
func removeElements(head *ListNode, val int) *ListNode {
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

// 更简明的方法
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

// 最简明
func removeElements3(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val != val {
			cur = cur.Next
		} else {
			cur.Next = cur.Next.Next // 跳过下个元素，即删除下个元素
		}
	}
	return dummy.Next
}

// 26. 删除有序数组中的重复项（只留一个重复元素）
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
// 考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
// nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2,_]
// 解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
// 快慢指针，注意：和27移除元素的区别，本题要求移除多余重复的元素
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0 // 维护 nums[0..slow] 左闭右闭，为不包含重复元素的结果子数组
	for fast < len(nums) {
		if nums[fast] != nums[slow] { // 发现fast是一个新的无重复元素
			// slow把fast的值复制过来
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// 83. 删除排序链表中的重复元素(只留一个重复元素)
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/
// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
// 输入：head = [1,1,2]
// 输出：[1,2]
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast // 对应数组 slow++
			slow = slow.Next // 对应数组 nums[slow] = nums[fast]
		}
		fast = fast.Next // 对应数组 fast++
	}
	slow.Next = nil // 断开与后面重复元素的连接
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next // 删除重复元素
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 80.删除有序数组中的重复项 II
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// 输入：nums = [1,1,1,2,2,3]
// 输出：5, nums = [1,1,2,2,3]
// 解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。
// 思路1:快慢指针
func removeDuplicatesII2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 快慢指针，维护 nums[0..slow] 左闭右闭，为结果子数组
	slow, fast := 0, 0
	// 记录一个元素重复的次数
	count := 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			// 此时，对于 nums[0..slow] 来说，nums[fast] 是一个新的元素，加进来
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && count < 2 {
			// 此时，对于 nums[0..slow] 来说，nums[fast] 重复次数小于 2，也加进来
			slow++
			nums[slow] = nums[fast]
		}
		fast++
		count++
		if fast < len(nums) && nums[fast] != nums[fast-1] {
			// fast 遇到新的不同的元素时，重置 count
			count = 0
		}
	}

	return slow + 1 // 数组长度为索引 + 1
}

// 思路2
func removeDuplicatesII(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// 82. 删除排序链表中的重复元素 II
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/
// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
// 输入：head = [1,2,3,3,4,4,5]
// 输出：[1,2,5]
// 和上题的区别：上题要求把多于的重复元素去掉，这道题要求把所有重复的元素全都去掉。
// 思路1:链表分解,将原链表分解为两条链表,一条不含重复元素，另一条含重复元素
func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{Val: 101}
	p1, p2 := dummy1, dummy2
	cur := head
	for cur != nil {
		next := cur.Next
		if (next != nil && cur.Val == next.Val) || cur.Val == p2.Val {
			// 重复
			p2.Next = cur
			p2 = p2.Next
		} else {
			p1.Next = cur
			p1 = p1.Next
		}
		cur.Next = nil
		cur = next
	}
	return p1
}

// 思路2: 快慢指针
func deleteDuplicatesII2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	slow := dummy // slow始终指向和下个节点值不同的节点
	fast := head  // fast在前面探路，和后一个节点值比较
	for fast != nil && fast.Next != nil {
		if fast.Val != fast.Next.Val {
			slow = slow.Next
			fast = fast.Next
			continue
		}
		// 此时fast和它的下个节点值相同，跳过这些相同的节点
		for fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}
		slow.Next = fast.Next // fast.Next是第一个值不重复的节点
	}
	return dummy.Next
}

// 思路3: 2次遍历，hashmap记录重复的元素，再遍历一次删除重复元素
func deleteDuplicatesII3(head *ListNode) *ListNode {
	// 记录次数
	m := make(map[int]int)
	for cur := head; cur != nil; cur = cur.Next {
		m[cur.Val]++
	}
	// 再遍历一次删除次数>1的节点
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur != nil {
		for cur.Next != nil && m[cur.Next.Val] > 1 {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

// 思路4: 递归解法
func deleteDuplicates3(head *ListNode) *ListNode {
	// 定义：输入一条单链表头结点，返回去重之后的单链表头结点
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		// 如果头结点和身后节点的值不同，则对之后的链表去重即可
		head.Next = deleteDuplicates3(head.Next)
		return head
	}
	// 如果如果头结点和身后节点的值相同，则说明从 head 开始存在若干重复节点
	// 越过重复节点，找到 head 之后那个不重复的节点
	for head.Next != nil && head.Val == head.Next.Val {
		head = head.Next
	}
	// 直接返回那个不重复节点开头的链表的去重结果，就把重复节点删掉了
	return deleteDuplicates3(head.Next)
}

// 75. 颜色分类
// https://leetcode.cn/problems/sort-colors/
// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
// 必须在不使用库内置的 sort 函数的情况下解决这个问题。
func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	p := 0
	for p <= p2 {
		if nums[p] == 0 {
			nums[p0], nums[p] = nums[p], nums[p0]
			p0++
		} else if nums[p] == 2 {
			nums[p], nums[p2] = nums[p2], nums[p]
			p2--
		} else if nums[p] == 1 {
			p++
		}
		// 由于p找到0就会和p0位置的数字换，所以p0一直增加，由于p0之前都是0，所以p需要>=p0
		if p < p0 {
			p = p0
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
