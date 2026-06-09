package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// 46.全排列
// https://leetcode.cn/problems/permutations/description/
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 思路1: 盒视角
func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	var backtrack func()
	used := make([]bool, len(nums))

	backtrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}

// 思路1: 盒视角，swap，start含义：是nums数组中每个索引位置，选择不同的元素放入这个索引位置。
// start之前的元素已经心有所属，被其他位置挑走了。所以stat位置只能从nums[start...]中选择元素
func permute2(nums []int) [][]int {
	var res [][]int
	var backtrack func(start int)

	backtrack = func(start int) {
		if start == len(nums) {
			res = append(res, append([]int{}, nums...))
			return
		}
		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			backtrack(start + 1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}

	backtrack(0)
	return res
}

// 思路2: 球视角，元素选索引
func permute3(nums []int) [][]int {
	var res [][]int
	var backtrack func()
	used := make([]bool, len(nums))
	count := 0

	backtrack = func() {
		if count == len(nums) {
			res = append(res, append([]int{}, nums...))
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
			backtrack()
			// 撤销选择
			count--
			used[swapIndex] = false
			nums[originalIndex], nums[swapIndex] = nums[swapIndex], nums[originalIndex]
		}
	}

	backtrack()
	return res
}

// 47. 全排列 II
// https://leetcode.cn/problems/permutations-ii/description/
// LCR 084. 全排列 II https://leetcode.cn/problems/7p8L0Z/description/
// 输入：nums = [1,1,2]
// 输出：[[1,1,2],[1,2,1],[2,1,1]]
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtrack func()

	backtrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
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
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}

// 967. 连续差相同的数字
// https://leetcode.cn/problems/numbers-with-same-consecutive-differences/description/
// 返回所有长度为 n 且满足其每两个连续位上的数字之间的差的绝对值为 k 的 非负整数 。
// 请注意，除了 数字 0 本身之外，答案中的每个数字都 不能 有前导零。例如，01 有一个前导零，所以是无效的；但 0 是有效的。
// 你可以按 任何顺序 返回答案。
// 输入：n = 3, k = 7
// 输出：[181,292,707,818,929]
// 解释：注意，070 不是一个有效的数字，因为它有前导零。
// 思路：元素无重可复选的排列,n 个盒子，然后有 0~9 种球（元素）可以放进盒子，每个盒子只能放一个球，但每种球的数量无限，可以使用无数次。
func numsSameConsecDiff(n int, k int) []int {
	var res []int
	var path []int
	var backtrack func()

	backtrack = func() {
		if len(path) == n {
			s := 0
			for i := 0; i < n; i++ {
				s = 10*s + path[i]
			}
			res = append(res, s)
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
			backtrack()
			path = path[:len(path)-1]
		}
	}

	backtrack()
	return res
}

// 1079. 活字印刷
// https://leetcode.cn/problems/letter-tile-possibilities/description/
// 你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。
// 注意：本题中，每个活字字模只能使用一次。
// 输入："AAB"
// 输出：8
// 解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。
// 思路: 元素可重不可复选的排列，普通排列，即并非每个元素都要参与到排列中。
func numTilePossibilities(tiles string) int {
	bs := []byte(tiles)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] }) // 先排序，让相同的元素靠在一起
	res := 0
	used := make([]bool, len(bs))
	var backtrack func()

	backtrack = func() {
		res++
		for i := 0; i < len(bs); i++ {
			if used[i] {
				continue
			}
			if i > 0 && bs[i] == bs[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			backtrack()
			used[i] = false
		}
	}

	backtrack()
	return res - 1 // 去掉空字符串
}

// 996. 平方数组的数目
// https://leetcode.cn/problems/number-of-squareful-arrays/description/
// 如果一个数组的任意两个相邻元素之和都是 完全平方数 ，则该数组称为 平方数组 。
// 给定一个整数数组 nums，返回所有属于 平方数组 的 nums 的排列数量。
// 如果存在某个索引 i 使得 perm1[i] != perm2[i]，则认为两个排列 perm1 和 perm2 不同。
// 输入：nums = [1,17,8]
// 输出：2
// 解释：[1,8,17] 和 [17,8,1] 是有效的排列。
// 思路: 元素可重不可复选的排列
func numSquarefulPerms(nums []int) int {
	sort.Ints(nums)
	res := 0
	used := make([]bool, len(nums))
	var path []int
	var isSqrt func(n int) bool
	isSqrt = func(n int) bool {
		c := int(math.Sqrt(float64(n)))
		return c*c == n
	}

	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		if len(nums) == len(path) {
			res++
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if len(path) > 0 && !isSqrt(path[len(path)-1]+nums[i]) {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtrack(nums)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack(nums)
	return res
}

// 784. 字母大小写全排列
// https://leetcode.cn/problems/letter-case-permutation/description/
// 给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。
// 返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。
// 输入：s = "a1b2"
// 输出：["a1b2", "a1B2", "A1b2", "A1B2"]
func letterCasePermutation(s string) []string {
	var res []string
	var path string
	var backtrack func(i int)

	backtrack = func(i int) {
		if i == len(s) {
			res = append(res, path)
			return
		}
		if s[i] >= '0' && s[i] <= '9' {
			path += string(s[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		} else {
			// 不转变大小写 or 转变大小写
			lower := strings.ToLower(string(s[i]))
			upper := strings.ToUpper(string(s[i]))
			for _, str := range []string{lower, upper} {
				path += str
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(numsSameConsecDiff(3, 7))
}
