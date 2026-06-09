package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 131.分割回文串
// https://leetcode.cn/problems/palindrome-partitioning/description/
// 输入：s = "aab" 输出：[["a","a","b"],["aa","b"]]
func partition(s string) [][]string {
	var res [][]string
	var path []string
	var isPalindrome func(s string) bool
	var backtrack func(start int)

	isPalindrome = func(s string) bool {
		left, right := 0, len(s)-1
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	backtrack = func(start int) {
		if start == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := start; i < len(s); i++ {
			if !isPalindrome(s[start : i+1]) {
				continue // 剪枝：分割出的子串不是回文串
			}
			path = append(path, s[start:i+1])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}

// 93.复原IP地址
// https://leetcode.cn/problems/restore-ip-addresses/description/
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 转换为3个.放在哪几个位置，能放[1,len(s)-1]
func restoreIpAddresses(s string) []string {
	var res []string
	var path []string
	var isValidIp func(s string) bool
	var backtrack func(startIndex int)

	isValidIp = func(s string) bool {
		// 不能前导0
		if len(s) > 1 && s[0] == '0' {
			return false
		}
		// 0-255之间
		n, _ := strconv.Atoi(s)
		return n <= 255
	}

	backtrack = func(start int) {
		if start == len(s) && len(path) == 4 {
			res = append(res, strings.Join(path, "."))
			return
		}
		for i := start; i < len(s); i++ {
			ip := s[start : i+1]
			if isValidIp(ip) {
				path = append(path, ip)
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return res
}

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

// 视角2: 站在子串的视角进行穷举
func splitString2(s string) bool {
	found := false
	var path []string
	var backtrack func(start int)

	backtrack = func(start int) {
		if found {
			return
		}
		if start == len(s) && len(path) > 1 {
			found = true
			return
		}
		for i := start; i < len(s); i++ {
			val, _ := strconv.Atoi(s[start : i+1])
			if len(path) > 0 {
				last, _ := strconv.Atoi(path[len(path)-1])
				if last-val != 1 {
					continue
				}
			}
			path = append(path, s[start:i+1])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
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
	res := 0
	uset := make(map[string]bool)
	var path []string
	var backtrack func(start int)

	backtrack = func(start int) {
		if start == len(s) {
			res = max(res, len(path))
			return
		}
		for i := start; i < len(s); i++ {
			sub := s[start : i+1]
			if uset[sub] {
				continue
			}
			uset[sub] = true
			path = append(path, sub)
			backtrack(i + 1)
			path = path[:len(path)-1]
			uset[sub] = false
		}
	}

	backtrack(0)
	return res
}

// 视角2: 站在索引空隙之间选择切or不切，脑海中出现一颗二叉树
func maxUniqueSplit2(s string) int {
	res := 0
	uset := make(map[string]bool)
	var backtrack func(s string, index int)
	backtrack = func(s string, index int) {
		if index == len(s) {
			res = max(res, len(uset))
			return
		}
		// 不切
		backtrack(s, index+1)
		// 切,把 s[0..index] 切分出来作为一个子串
		sub := s[:index+1]
		if !uset[sub] {
			uset[sub] = true          // 做选择
			backtrack(s[index+1:], 0) // 剩下的字符继续穷举
			delete(uset, sub)         // 撤销选择
		}
	}

	backtrack(s, 0)
	return res
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(restoreIpAddresses("25525511135"))
	val, _ := strconv.Atoi("05")
	fmt.Println(val)
}
