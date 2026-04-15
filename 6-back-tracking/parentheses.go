package main

// 22. 括号生成
// https://leetcode.cn/problems/generate-parentheses/description/
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
func generateParenthesis(n int) []string {
	var results []string
	if n == 0 {
		return results
	}
	var path []byte
	var backtrack func(i, j int)
	backtrack = func(i, j int) {
		if i > j {
			return
		}
		if i < 0 || j < 0 {
			return
		}
		if i == 0 && j == 0 {
			results = append(results, string(path))
			return
		}

		for _, c := range []byte{'(', ')'} {
			path = append(path, c)
			if c == '(' {
				backtrack(i-1, j)
			} else {
				backtrack(i, j-1)
			}
			path = path[:len(path)-1]
		}
	}
	backtrack(n, n)
	return results
}
