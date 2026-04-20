package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// 1849. 将字符串拆分为递减的连续值
// https://leetcode.cn/problems/splitting-a-string-into-descending-consecutive-values/description/
// 给你一个仅由数字组成的字符串 s 。
// 请你判断能否将 s 拆分成两个或者多个 非空子字符串 ，使子字符串的 数值 按 降序 排列，且每两个 相邻子字符串 的数值之 差 等于 1 。
// 例如，字符串 s = "0090089" 可以拆分成 ["0090", "089"] ，数值为 [90,89] 。这些数值满足按降序排列，且相邻值相差 1 ，这种拆分方法可行。
// 另一个例子中，字符串 s = "001" 可以拆分成 ["0", "01"]、["00", "1"] 或 ["0", "0", "1"] 。然而，所有这些拆分方法都不可行，因为对应数值分别是 [0,1]、[0,1] 和 [0,0,1] ，都不满足按降序排列的要求。
// 如果可以按要求拆分 s ，返回 true ；否则，返回 false 。
// 子字符串 是字符串中的一个连续字符序列。
// 输入：s = "1234"
// 输出：false
// 解释：不存在拆分 s 的可行方法。
// 输入：s = "050043"
// 输出：true
// 解释：s 可以拆分为 ["05", "004", "3"] ，对应数值为 [5,4,3] 。
// 满足按降序排列，且相邻值相差 1 。
// 思路1: 站在字符的视角进行穷举
func splitString(s string) bool {
	found := false
	var path []string
	var parseInt func(s string) int64
	parseInt = func(s string) int64 {
		num, _ := strconv.ParseInt(s, 10, 64)
		return num
	}
	var backtrack func(s string, start, index int)

	backtrack = func(s string, start, index int) {
		if found {
			return
		}
		if index == len(s) {
			if len(path) >= 2 && strings.Join(path, "") == s {
				found = true
			}
			return
		}
		// 选择一，s[index] 决定切割
		subStr := s[start : index+1]
		leadingZeroCount := 0
		for j := 0; j < len(subStr); j++ {
			if subStr[j] == '0' {
				leadingZeroCount++
			} else {
				break
			}
		}
		if len(subStr)-leadingZeroCount > (len(s)+1)/2 {
			return // 剪枝逻辑，如果当前截取的子串长度大于 s 的一半，那么没必要继续截取了，肯定不可能只差一，同时可以避免溢出 long 的最大值的问题
		}

		if len(path) == 0 || parseInt(path[len(path)-1])-parseInt(subStr) == 1 {
			// 符合题目的要求，当前数字比上一个数字小 1。做选择，切割出一个子串
			path = append(path, subStr)
			backtrack(s, index+1, index+1)
			path = path[:len(path)-1]
		}

		// 选择二，s[index] 决定不切割
		backtrack(s, start, index+1)
	}

	backtrack(s, 0, 0)
	return found
}

// 1593. 拆分字符串使唯一子字符串的数目最大
// https://leetcode.cn/problems/split-a-string-into-the-max-number-of-unique-substrings/description/
// 给你一个字符串 s ，请你拆分该字符串，并返回拆分后唯一子字符串的最大数目。
// 字符串 s 拆分后可以得到若干 非空子字符串 ，这些子字符串连接后应当能够还原为原字符串。但是拆分出来的每个子字符串都必须是 唯一的 。
// 注意：子字符串 是字符串中的一个连续字符序列。
// 输入：s = "ababccc"
// 输出：5
// 解释：一种最大拆分方法为 ['a', 'b', 'ab', 'c', 'cc'] 。像 ['a', 'b', 'a', 'b', 'c', 'cc'] 这样拆分不满足题目要求，因为其中的 'a' 和 'b' 都出现了不止一次。
// 视角1: 子串（盒）视角，切出来的子串长度可以是1,2,3..len(s)
func maxUniqueSplit(s string) int {
	set := make(map[string]bool)
	result := 0
	cnt := 0
	var backtrack func(s string, start int)
	backtrack = func(s string, start int) {
		if start == len(s) {
			result = max(result, cnt)
			return
		}
		for i := start; i < len(s); i++ {
			sub := s[start : i+1]
			if set[sub] {
				continue
			}
			cnt++
			set[sub] = true
			backtrack(s, i+1)
			delete(set, sub)
			cnt--
		}
	}

	backtrack(s, 0)
	return result
}

// 视角2: 站在索引空隙之间选择切or不切，脑海中出现一颗二叉树
func maxUniqueSplit2(s string) int {
	result := 0
	set := make(map[string]bool)
	var backtrack func(s string, index int)
	backtrack = func(s string, index int) {
		if index == len(s) {
			result = max(result, len(set))
			return
		}
		// 不切
		backtrack(s, index+1)
		// 切,把 s[0..index] 切分出来作为一个子串
		sub := s[:index+1]
		if !set[sub] {
			set[sub] = true           // 做选择
			backtrack(s[index+1:], 0) // 剩下的字符继续穷举
			delete(set, sub)          // 撤销选择
		}
	}

	backtrack(s, 0)
	return result
}

// 1079. 活字印刷
// https://leetcode.cn/problems/letter-tile-possibilities/description/
// 你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。
// 注意：本题中，每个活字字模只能使用一次。
// 输入："AAB"
// 输出：8
// 解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。
// 元素可重不可复选
func numTilePossibilities(tiles string) int {
	bs := []byte(tiles)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	used := make([]bool, len(bs))
	result := 0
	var backtrack func(bs []byte)
	backtrack = func(bs []byte) {
		result++
		for i := 0; i < len(bs); i++ {
			if used[i] {
				continue
			}
			if i > 0 && bs[i] == bs[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			backtrack(bs)
			used[i] = false
		}
	}
	backtrack(bs)
	return result - 1
}

// 996. 平方数组的数目
// https://leetcode.cn/problems/number-of-squareful-arrays/description/
// 如果一个数组的任意两个相邻元素之和都是 完全平方数 ，则该数组称为 平方数组 。
// 给定一个整数数组 nums，返回所有属于 平方数组 的 nums 的排列数量。
// 如果存在某个索引 i 使得 perm1[i] != perm2[i]，则认为两个排列 perm1 和 perm2 不同。
// 输入：nums = [1,17,8]
// 输出：2
// 解释：[1,8,17] 和 [17,8,1] 是有效的排列。
// 元素可重不可复选的排列
func numSquarefulPerms(nums []int) int {
	sort.Ints(nums)
	result := 0
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
			result++
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
	return result
}

// 784. 字母大小写全排列
// https://leetcode.cn/problems/letter-case-permutation/description/
// 给定一个字符串 s ，通过将字符串 s 中的每个字母转变大小写，我们可以获得一个新的字符串。
// 返回 所有可能得到的字符串集合 。以 任意顺序 返回输出。
// 输入：s = "a1b2"
// 输出：["a1b2", "a1B2", "A1b2", "A1B2"]
func letterCasePermutation(s string) []string {
	var result []string
	var path string
	var backtrack func(s string, i int)
	backtrack = func(s string, i int) {
		if i == len(s) {
			result = append(result, path)
			return
		}
		if s[i] >= '0' && s[i] <= '9' {
			path += string(s[i])
			backtrack(s, i+1)
			path = path[:len(path)-1]
		} else {
			// 不转变大小写 or 转变大小写
			lower := strings.ToLower(string(s[i]))
			upper := strings.ToUpper(string(s[i]))
			for _, str := range []string{lower, upper} {
				path += str
				backtrack(s, i+1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(s, 0)
	return result
}

func main() {
	fmt.Println(numTilePossibilities("AAB"))
}
