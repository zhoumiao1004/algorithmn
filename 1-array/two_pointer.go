package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 26. 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
// 考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
// nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2,_]
// 解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// 83. 删除排序链表中的重复元素
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
			// slow++
			slow.Next = fast
			// nums[slow] = nums[fast]
			slow = slow.Next
		}
		// fast++
		fast = fast.Next
	}
	// 断开与后面重复元素的连接
	slow.Next = nil
	return head
}

// 82. 删除排序链表中的重复元素 II
// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
// 输入：head = [1,2,3,3,4,4,5]
// 输出：[1,2,5]
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := head
	for fast != nil && fast.Next != nil {
		if fast.Val != fast.Next.Val {
			slow = slow.Next
			fast = fast.Next
			continue
		}
		for fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}
		slow.Next = fast.Next
		fast = fast.Next
	}
	return dummy.Next
}

// 27. 移除元素
// https://leetcode.cn/problems/remove-element/description/
// nums = [3,2,2,3], val = 3
// 输出: 2, nums = [2,2,_,_]
func removeElement(nums []int, val int) int {
	left := 0
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
	left := 0
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

// 167. 两数之和 II - 输入有序数组
// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
// 给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
// 以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
// 你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
// 你所设计的解决方案必须只使用常量级的额外空间。
// 输入：numbers = [2,7,11,15], target = 9
// 输出：[1,2]
// 解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{-1, -1}
}

// 344. 反转字符串
// https://leetcode.cn/problems/reverse-string/description/
// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 5. 最长回文子串
// 给你一个字符串 s，找到 s 中最长的 回文 子串。
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心的最长回文子串
		s1 := palindrome(s, i, i)
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		s2 := palindrome(s, i, i+1)
		// res = longest(res, s1, s2)
		if len(res) > len(s1) {
			res = res
		} else {
			res = s1
		}
		if len(res) > len(s2) {
			res = res
		} else {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l, r int) string {
	// 防止索引越界
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 向两边展开
		l--
		r++
	}
	// 此时 s[l+1..r-1] 就是最长回文串
	return s[l+1 : r]
}

// 80.删除有序数组中的重复项 II
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// 输入：nums = [1,1,1,2,2,3]
// 输出：5, nums = [1,1,2,2,3]
// 解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。
func removeDuplicates2(nums []int) int {
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

// 快慢指针
func removeDuplicates3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 快慢指针，维护 nums[0..slow] 为结果子数组
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
	// 数组长度为索引 + 1
	return slow + 1
}

// 125. 验证回文串
// https://leetcode.cn/problems/valid-palindrome/description/
// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
// 字母和数字都属于字母数字字符。
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
// 输入: s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。
func isPalindrome(s string) bool {
	bs := []byte(s)
	// 保留小写字母
	slow := 0
	for i := 0; i < len(bs); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			bs[slow] = c
			slow++
		} else if c >= 'A' && c <= 'Z' {
			bs[slow] = c - 'A' + 'a'
			slow++
		}
	}
	fmt.Println(string(bs[:slow]))
	left, right := 0, slow-1
	for left < right {
		if bs[left] != bs[right] {
			return false
		}
		left++
		right--
	}
	return true
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

// 88.合并两个有序数组
// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
// 注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
// 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出：[1,2,2,3,5,6]
// 解释：需要合并 [1,2,3] 和 [2,5,6] 。
// 合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	k := len(nums1) - 1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
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

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
