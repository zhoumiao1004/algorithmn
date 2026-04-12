package main

// 238. 除了自身以外数组的乘积
// https://leetcode.cn/problems/product-of-array-except-self/description/
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除了 nums[i] 之外其余各元素的乘积 。
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
// prefix: [1,   2,  6, 24]
// suffix: [24, 24, 12,  4]
func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}
	
	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}

	result := make([]int, n)
	result[0] = suffix[1]
	result[n-1] = prefix[n-2]
	for i := 1; i < n-1; i++ {
		result[i] = prefix[i-1] * suffix[i+1]
	}
	return result
}

// 1352. 最后 K 个数的乘积
// https://leetcode.cn/problems/product-of-the-last-k-numbers/
// 设计一个算法，该算法接受一个整数流并检索该流中最后 k 个整数的乘积。
// 实现 ProductOfNumbers 类：
// ProductOfNumbers() 用一个空的流初始化对象。
// void add(int num) 将数字 num 添加到当前数字列表的最后面。
// int getProduct(int k) 返回当前数字列表中，最后 k 个数字的乘积。你可以假设当前列表中始终 至少 包含 k 个数字。
// 题目数据保证：任何时候，任一连续数字序列的乘积都在 32 位整数范围内，不会溢出。
// 输入：
// ["ProductOfNumbers","add","add","add","add","add","getProduct","getProduct","getProduct","add","getProduct"]
// [[],[3],[0],[2],[5],[4],[2],[3],[4],[8],[2]]
// 输出：
// [null,null,null,null,null,null,20,40,0,null,32]
type ProductOfNumbers struct {
	prefix []int
}

func Constructor3() ProductOfNumbers {
	return ProductOfNumbers{
		prefix: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if len(this.prefix) == 0 {
		this.prefix = []int{1}
	} else {
		this.prefix = append(this.prefix, this.prefix[len(this.prefix)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.prefix)
	if k > n-1 {
		return 0
	}
	return this.prefix[n-1] / this.prefix[n-k-1]
}
