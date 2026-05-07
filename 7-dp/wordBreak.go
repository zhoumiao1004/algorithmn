package main

import (
	"fmt"
	"strings"
)

// 139.单词拆分
// https://leetcode.cn/problems/word-break/
// 输入: s = "leetcode", wordDict = ["leet", "code"] 输出: true
// 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
// 注意：单词放入是有顺序的，所以是排列问题，不能求组合
// 1.遍历的思路，就是用回溯算法解决，回溯最经典的应用就是排列组合问题。时间复杂度2的N次方
func wordBreakBT(s string, wordDict []string) bool {
	// memo := make(map[string]bool)
	found := false
	var path []string
	var backtrack func(wordDict []string, start int)
	backtrack = func(wordDict []string, start int) {
		if found {
			return
		}
		if start == len(wordDict) {
			found = true
			return
		}

		for i := 0; i < len(wordDict); i++ {
			word := wordDict[i]
			if start+len(word) <= len(s) && s[start:start+len(word)] == word {
				// 做选择
				path = append(path, word)
				// 进入下一层回溯树
				backtrack(wordDict, i+len(word))
				// 撤销选择
				path = path[:len(path)-1]
			}

		}
	}
	backtrack(wordDict, 0)
	return found
}

// 2.分解的思路
func wordBreakMemo(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	memo := make([]int, len(s)) // s[i]能否被单词拼出, -1 代表未计算，0 代表无法凑出，1 代表可以凑出
	// 定义：返回s[start...] 子串是否能被单词拼出
	var dp func(s string, start int) bool
	dp = func(s string, start int) bool {
		if start == len(s) {
			return true
		}
		if memo[start] != 0 {
			return memo[start] == 1
		}
		// 遍历 s[start...] 的所有前缀，看看哪些前缀存在 wordDict 中
		for i := 1; start+i <= len(s); i++ {
			prefix := s[start : start+i]
			if wordSet[prefix] && dp(s, start+i) {
				memo[start] = 1
				return true
			}
		}
		// s[1...] 无法被拼出
		memo[start] = 0
		return false
	}
	return dp(s, 0)
}

func wordBreak(s string, wordDict []string) bool {
	// dp[j]含义：[0,j)范围的子串，能否由字典里的单词组成
	// 用集合中的物品，装大小为j的背包
	// if dp[i] = true && [i,j]区间内的字符串在字典中 : dp[j] = true
	// 遍历顺序：求排列，先遍历背包再遍历物品
	wordMap := make(map[string]bool)
	for _, w := range wordDict {
		wordMap[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for j := 1; j <= n; j++ { // 背包
		for i := 0; i <= j; i++ { // 物品
			if dp[i] && wordMap[s[i:j]] {
				dp[j] = true
			}
		}
	}
	// fmt.Println(dp)
	return dp[n]
}

// 140. 单词拆分 II
// https://leetcode.cn/problems/word-break-ii/
// 给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。
// 注意：词典中的同一个单词可能在分段中被重复使用多次。
// 输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
// 输出:["cats and dog","cat sand dog"]
// 1.遍历的思路（回溯算法）
func wordBreakIIBT(s string, wordDict []string) []string {
	var result []string
	var path []string

	var backtrack func(s string, start int)
	backtrack = func(s string, start int) {
		if start == len(s) {
			result = append(result, strings.Join(path, " "))
			return
		}
		for _, word := range wordDict {
			if start+len(word) <= len(s) && s[start:start+len(word)] == word {
				path = append(path, word)
				backtrack(s, start+len(word))
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(s, 0)
	return result
}

// 2.分解的思路（动态规划）
func wordBreakII(s string, wordDict []string) []string {
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
