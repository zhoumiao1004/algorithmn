package main

import (
	"fmt"
	"math"
	"sort"
)

// 46.全排列
// https://leetcode.cn/problems/permutations/description/
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 思路1: 盒视角
func permute(nums []int) [][]int {
	var results [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		if len(path) == len(nums) {
			results = append(results, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack(nums)
	return results
}

// 使用swap更高效
func permute2(nums []int) [][]int {
	var results [][]int
	var backtrack func(nums []int, start int)

	backtrack = func(nums []int, start int) {
		if start == len(nums) {
			results = append(results, append([]int{}, nums...))
			return
		}
		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			backtrack(nums, start+1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}

	backtrack(nums, 0)
	return results
}

// 思路2: 球视角，元素选索引
func permute3(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	count := 0
	var backtrack func(nums []int)

	backtrack = func(nums []int) {
		if count == len(nums) {
			result = append(result, append([]int{}, nums...))
			return
		}
		originalIndex := -1
		swapIndex := -1
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if originalIndex == -1 {
				originalIndex = i
			}
			swapIndex = i
			// 做选择，元素 nums[originalIndex] 选择 swapIndex 位置
			nums[originalIndex], nums[swapIndex] = nums[swapIndex], nums[originalIndex]
			used[swapIndex] = true
			count++
			backtrack(nums)
			// 撤销选择
			count--
			used[swapIndex] = false
			nums[originalIndex], nums[swapIndex] = nums[swapIndex], nums[originalIndex]
		}
	}

	backtrack(nums)
	return result
}

// 47. 全排列 II
// https://leetcode.cn/problems/permutations-ii/description/
// LCR 084. 全排列 II https://leetcode.cn/problems/7p8L0Z/description/
// 输入：nums = [1,1,2]
// 输出：[[1,1,2],[1,2,1],[2,1,1]]
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var results [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtrack func(nums []int)

	backtrack = func(nums []int) {
		if len(path) == len(nums) {
			results = append(results, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue // 树层去重复
			}
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack(nums)
	return results
}

// 967. 连续差相同的数字
// https://leetcode.cn/problems/numbers-with-same-consecutive-differences/description/
// 返回所有长度为 n 且满足其每两个连续位上的数字之间的差的绝对值为 k 的 非负整数 。
// 请注意，除了 数字 0 本身之外，答案中的每个数字都 不能 有前导零。例如，01 有一个前导零，所以是无效的；但 0 是有效的。
// 你可以按 任何顺序 返回答案。
// 输入：n = 3, k = 7
// 输出：[181,292,707,818,929]
// 解释：注意，070 不是一个有效的数字，因为它有前导零。
func numsSameConsecDiff(n int, k int) []int {
	var results []int
	var path []int
	var backtrack func(n, k int)

	backtrack = func(n, k int) {
		// 满足长度n条件，收集结果
		if len(path) == n {
			s := 0
			for i := 0; i < n; i++ {
				s = 10*s + path[i]
			}
			results = append(results, s)
			return
		}
		for i := 0; i <= 9; i++ {
			if len(path) == 0 && i == 0 {
				continue // 不能前导0
			}
			if len(path) > 0 && int(math.Abs(float64(path[len(path)-1])-float64(i))) != k {
				continue // 相差不为k
			}
			path = append(path, i)
			backtrack(n, k)
			path = path[:len(path)-1]
		}
	}

	backtrack(n, k)
	return results
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
