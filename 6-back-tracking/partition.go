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
	var results [][]string
	var path []string
	var backtrack func(start int)
	var isPalindrome func(s string) bool

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
			results = append(results, append([]string{}, path...))
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
	return results
}

// 93.复原IP地址
// https://leetcode.cn/problems/restore-ip-addresses/description/
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 转换为3个.放在哪几个位置，能放[1,len(s)-1]
func restoreIpAddresses(s string) []string {
	var results []string
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

	backtrack = func(startIndex int) {
		if startIndex == len(s) && len(path) == 4 {
			results = append(results, strings.Join(path, "."))
			return
		}
		for i := startIndex; i < len(s); i++ {
			ip := s[startIndex : i+1]
			if isValidIp(ip) {
				path = append(path, ip)
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return results
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(restoreIpAddresses("25525511135"))
}
