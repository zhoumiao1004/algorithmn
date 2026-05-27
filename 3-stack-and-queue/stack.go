package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// 1047. 删除字符串中的所有相邻重复项
// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
// 输入："abbaca"
// 输出："ca"
func removeDuplicates(s string) string {
	var st []byte
	for i := 0; i < len(s); i++ {
		if len(st) > 0 && s[i] == st[len(st)-1] {
			st = st[:len(st)-1] // 出栈
		} else {
			st = append(st, s[i])
		}
	}
	return string(st)
}

// 71. 简化路径
// https://leetcode.cn/problems/simplify-path/description/
// 输入：path = "/home/"
// 输出："/home"
// 输入：path = "/.../a/../b/c/../d/./"
// 输出："/.../b/d"
func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	var st []string
	for _, s := range strs {
		if s == "" || s == "." {
			continue
		}
		if s == ".." {
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
			continue
		}
		st = append(st, s)
	}
	res := ""
	for _, s := range st {
		res += "/" + s
	}
	if res == "" {
		return "/"
	}
	return res
}

// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses/description/
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 1.左括号必须用相同类型的右括号闭合。
// 2.左括号必须以正确的顺序闭合。
// 3.每个右括号都有一个对应的相同类型的左括号。
// 输入：s = "()" 输出：true
func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	var st []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			st = append(st, c)
			continue
		}
		if len(st) == 0 {
			return false
		}
		if m[c] != st[len(st)-1] {
			return false
		}
		st = st[:len(st)-1]
	}
	return len(st) == 0
}

// 150. 逆波兰表达式求值
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
// 有效的算符为 '+'、'-'、'*' 和 '/' 。
// 输入：tokens = ["2","1","+","3","*"]
// 输出：9
// 解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
func evalRPN(tokens []string) int {
	var st []string
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			b, _ := strconv.Atoi(st[len(st)-1])
			st = st[:len(st)-1]
			a, _ := strconv.Atoi(st[len(st)-1])
			st = st[:len(st)-1]
			res := 0
			switch tokens[i] {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			st = append(st, fmt.Sprintf("%d", res))
		} else {
			st = append(st, tokens[i])
		}
	}
	n, _ := strconv.Atoi(st[0])
	return n
}

// 388. 文件的最长绝对路径
// https://leetcode.cn/problems/longest-absolute-file-path/
// 输入：input = "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
// 输出：20
// 解释：只有一个文件，绝对路径为 "dir/subdir2/file.ext" ，路径长度 20
func lengthLongestPath(input string) int {
	st := list.New()
	maxLen := 0
	parts := strings.Split(input, "\n")
	for _, part := range parts {
		level := strings.LastIndexByte(part, '\t') + 1
		// fmt.Println(part, level)
		for st.Len() > level {
			st.Remove(st.Back()) // 让栈中只保留当前目录的父路径
		}
		st.PushBack(part[level:]) // 跳过\t，把文件夹名放进去
		if strings.Contains(part, ".") {
			sum := 0
			for e := st.Front(); e != nil; e = e.Next() {
				// sum += utf8.RuneCountInString(e.Value.(string))
				sum += len(e.Value.(string))
			}
			sum += st.Len() - 1
			maxLen = max(maxLen, sum)
		}
	}
	return maxLen
}

// 394. 字符串解码
// https://leetcode.cn/problems/decode-string/description/
// 给定一个经过编码的字符串，返回它解码后的字符串。
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
// 测试用例保证输出的长度不会超过 105。
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
func decodeString(s string) string {
	cur := ""
	k := 0
	var strStack []string
	var cntStack []int
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			k = 10*k + int(c-'0')
		} else if c == '[' {
			strStack = append(strStack, cur)
			cntStack = append(cntStack, k)
			cur = ""
			k = 0
		} else if c == ']' {
			times := cntStack[len(cntStack)-1]
			cntStack = cntStack[:len(cntStack)-1]
			prev := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			cur = prev + strings.Repeat(cur, times)
		} else {
			cur += string(c)
		}
	}
	return cur
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"))
}
