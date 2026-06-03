package main

import (
	"fmt"
	"strings"
)

// 动态规划和回溯算法的思维转换

// 139.单词拆分
// https://leetcode.cn/problems/word-break/
// 输入: s = "leetcode", wordDict = ["leet", "code"] 输出: true
// 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
// 注意：单词放入是有顺序的，所以是排列问题，不能求组合
// 思路1: 遍历(回溯)，通过memo记录重复子问题
func wordBreak1(s string, wordDict []string) bool {
	flag := false
	memo := make(map[string]bool)
	var path []string
	var backtrack func(start int)

	backtrack = func(start int) {
		if flag {
			return
		}
		if start == len(s) {
			flag = true
			return
		}
		suffix := s[start:]
		if memo[suffix] {
			return // 不能被切分
		}
		for _, word := range wordDict {
			if start+len(word) <= len(s) && s[start:start+len(word)] == word {
				path = append(path, word)
				backtrack(start + len(word))
				path = path[:len(path)-1]
			}
		}
		// 后序位置
		if !flag {
			memo[suffix] = true
		}
	}

	backtrack(0)
	return flag
}

// 思路2: 分解(dp) 带memo的自顶向下递归
func wordBreak2(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	n := len(s)
	memo := make([]int, n) // s[i]能否被单词拼出, -1 代表未计算，0 代表无法凑出，1 代表可以凑出
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	var dp func(start int) bool // 定义：返回s[start...] 子串是否能被单词拼出
	dp = func(start int) bool {
		if start == n {
			return true
		}
		if memo[start] != -1 {
			return memo[start] == 1
		}
		for i := start; i <= len(s); i++ {
			if wordSet[s[start:i]] && dp(i) {
				memo[start] = 1
				return true
			}
		}
		memo[start] = 0
		return false
	}
	return dp(0)
}

// 思路3: dp自底向上，也可以转换成完全背包思路：把 wordDict 中的单词放进s这个背包，有顺序要求
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1) // dp[j]定义：s的前i个字符能否被wordDict中的单词组成。我们要求的就是dp[n]
	dp[0] = true
	for j := 1; j <= n; j++ {
		// 穷举所有单词的开头，s[i...j-1]在不在 wordDict 中
		for i := 0; i <= j; i++ {
			if dp[i] && wordSet[s[i:j]] {
				dp[j] = true
			}
		}
	}
	return dp[n]
}

// 140. 单词拆分 II
// https://leetcode.cn/problems/word-break-ii/
// 给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。
// 注意：词典中的同一个单词可能在分段中被重复使用多次。
// 输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
// 输出:["cats and dog","cat sand dog"]
// 1.遍历的思路（回溯算法）
func wordBreakII(s string, wordDict []string) []string {
	var res []string
	var path []string
	var backtrack func(start int)

	backtrack = func(start int) {
		if start == len(s) {
			res = append(res, strings.Join(path, " "))
			return
		}
		for _, word := range wordDict {
			if start+len(word) <= len(s) && s[start:start+len(word)] == word {
				path = append(path, word)
				backtrack(start + len(word))
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return res
}

func wordBreakII2(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}
	n := len(s)
	var res []string
	var track []string
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == n {
			res = append(res, strings.Join(track, " "))
			return
		}
		for i := start + 1; i <= n; i++ {
			if wordSet[s[start:i]] {
				track = append(track, s[start:i])
				backtrack(i)
				track = track[:len(track)-1]
			}
		}
	}

	backtrack(0)
	return res
}

// 2.分解的思路（动态规划）
func wordBreakII3(s string, wordDict []string) []string {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	memo := make([][]string, len(s))

	// dp[start...]能由单词拼成的句子
	var dp func(s string, start int) []string
	dp = func(s string, start int) []string {
		var result []string
		if start == len(s) {
			return []string{""}
		}
		if len(memo[start]) > 0 {
			return memo[start]
		}
		// 遍历 s[start...] 的所有前缀
		for i := 1; start+i <= len(s); i++ {
			prefix := s[start : start+i]
			if wordSet[prefix] {
				for _, sentence := range dp(s, start+i) {
					if sentence == "" {
						result = append(result, prefix)
					} else {
						result = append(result, prefix+" "+sentence)
					}
				}
			}
		}
		return result
	}
	return dp(s, 0)
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
}
