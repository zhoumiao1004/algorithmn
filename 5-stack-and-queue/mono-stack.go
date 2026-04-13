package main

import (
	"fmt"
)

/* 单调栈：通过一个栈保存已经遍历过的元素 */
// 739. 每日温度
// https://leetcode.cn/problems/daily-temperatures/description/
// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
// 输入: temperatures = [73,74,75,71,69,72,76,73]
// 输出: [1,1,4,2,1,1,0,0]
func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var st []int
	for i := 0; i < len(temperatures); i++ {
		// 比较temperature[i]和栈顶元素大小，如果大于栈顶元素，说明是右边第一个大于栈顶元素的地方
		for len(st) > 0 && temperatures[i] > temperatures[st[len(st)-1]] {
			idx := st[len(st)-1]
			result[idx] = i - idx
			st = st[:len(st)-1] // 弹出栈顶元素
		}
		st = append(st, i)
	}

	return result
}

func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, len(temperatures))
	var st []int
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && temperatures[i] >= temperatures[st[len(st)-1]] {
			st = st[:len(st)-1] // 弹出栈顶元素
		}
		if len(st) > 0 {
			result[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}
	return result
}

// 496. 下一个更大元素 I
// https://leetcode.com/problems/next-greater-element-i/description/
// 给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
// 请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
// nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
// 输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
// 输出：[-1,3,-1]
// 解释：nums1 中每个值的下一个更大元素如下所述：
// 4: 不存在下一个更大元素，所以答案是 -1 。
// 1: 下一个更大元素是 3 。
// 2: 不存在下一个更大元素，所以答案是 -1 。
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	idxMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		result[i] = -1
		idxMap[nums1[i]] = i
	}

	var st []int
	for i := 0; i < len(nums2); i++ {
		for len(st) > 0 && nums2[i] > st[len(st)-1] {
			idx, ok := idxMap[st[len(st)-1]]
			if ok {
				result[idx] = nums2[i]
			}
			st = st[:len(st)-1]
		}
		st = append(st, nums2[i])
	}
	return result
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	idxMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		result[i] = -1
		idxMap[nums1[i]] = i
	}

	var st []int
	n := len(nums2)
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] <= nums2[i] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			idx, ok := idxMap[nums2[i]]
			if ok {
				result[idx] = st[len(st)-1]
			}
		}
		st = append(st, nums2[i])
	}
	return result
}

// 503. 下一个更大元素 II
// https://leetcode.cn/problems/next-greater-element-ii/
// 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素 。
// 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1 。
// 输入: nums = [1,2,1] 输出: [2,-1,2]
// 解释: 第一个 1 的下一个更大的数是 2；
// 数字 2 找不到下一个更大的数；
// 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}
	var st []int
	for i := 0; i < 2*n; i++ {
		for len(st) > 0 && nums[i%n] > nums[st[len(st)-1]] {
			result[st[len(st)-1]] = nums[i%n]
			st = st[:len(st)-1]
		}
		st = append(st, i%n)
	}
	return result
}

func nextGreaterElements2(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	st := make([]int, 0)

	// 数组长度加倍模拟环形数组
	for i := 2*n - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] <= nums[i%n] {
			st = st[:len(st)-1] // pop element from stack
		}

		if len(st) > 0 {
			res[i%n] = st[len(st)-1]
		}

		st = append(st, nums[i%n]) // push element to stack
	}

	return res
}

// 1019. 链表中的下一个更大节点
// https://leetcode.cn/problems/next-greater-node-in-linked-list/description/
// 给定一个长度为 n 的链表 head
// 对于列表中的每个节点，查找下一个 更大节点 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 严格大于 它的值。
// 返回一个整数数组 answer ，其中 answer[i] 是第 i 个节点( 从1开始 )的下一个更大的节点的值。如果第 i 个节点没有下一个更大的节点，设置 answer[i] = 0 。
func nextLargerNodes(head *ListNode) []int {
	var nums []int
	for cur := head; cur != nil; cur = cur.Next {
		nums = append(nums, cur.Val)
	}
	n := len(nums)
	var st []int
	result := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[st[len(st)-1]] <= nums[i] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			result[i] = nums[st[len(st)-1]]
		}
		st = append(st, i)
	}
	return result
}

// 1944. 队列中可以看到的人数
// https://leetcode.cn/problems/number-of-visible-people-in-a-queue/description/
// 有 n 个人排成一个队列，从左到右 编号为 0 到 n - 1 。给你以一个整数数组 heights ，每个整数 互不相同，heights[i] 表示第 i 个人的高度。
// 一个人能 看到 他右边另一个人的条件是这两人之间的所有人都比他们两人 矮 。更正式的，第 i 个人能看到第 j 个人的条件是 i < j 且 min(heights[i], heights[j]) > max(heights[i+1], heights[i+2], ..., heights[j-1]) 。
// 请你返回一个长度为 n 的数组 answer ，其中 answer[i] 是第 i 个人在他右侧队列中能 看到 的 人数 。
// 输入：heights = [10,6,8,5,11,9]
// 输出：[3,1,2,1,1,0]
func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	var st []int
	result := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] < heights[i] {
			st = st[:len(st)-1]
			result[i]++
		}
		if len(st) > 0 {
			result[i]++
		}
		st = append(st, heights[i])
	}
	return result
}

// 1475. 商品折扣后的最终价格
// https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/description/
// 给你一个数组 prices ，其中 prices[i] 是商店里第 i 件商品的价格。
// 商店里正在进行促销活动，如果你要买第 i 件商品，那么你可以得到与 prices[j] 相等的折扣，其中 j 是满足 j > i 且 prices[j] <= prices[i] 的 最小下标 ，如果没有满足条件的 j ，你将没有任何折扣。
// 请你返回一个数组，数组中第 i 个元素是折扣后你购买商品 i 最终需要支付的价格。
// 输入：prices = [8,4,6,2,3]
// 输出：[4,2,4,2,3]
func finalPrices(prices []int) []int {
	n := len(prices)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = prices[i]
	}
	var st []int
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] > prices[i] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			result[i] = prices[i] - st[len(st)-1]
		}
		st = append(st, prices[i])
	}
	return result
}

// 901. 股票价格跨度
// https://leetcode.cn/problems/online-stock-span/
// 设计一个算法收集某些股票的每日报价，并返回该股票当日价格的 跨度 。
// 当日股票价格的 跨度 被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。
// 例如，如果未来 7 天股票的价格是 [100,80,60,70,60,75,85]，那么股票跨度将是 [1,1,1,2,1,4,6] 。
// 实现 StockSpanner 类：
// StockSpanner() 初始化类对象。
// int next(int price) 给出今天的股价 price ，返回该股票当日价格的 跨度 。
type StockSpanner struct {
	st [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		st: [][2]int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	count := 1
	for len(this.st) > 0 && this.st[len(this.st)-1][0] <= price {
		count += this.st[len(this.st)-1][1]
		this.st = this.st[:len(this.st)-1]
	}
	this.st = append(this.st, [2]int{price, count})
	return count
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4})) // 3, -1]
}
