package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

// 150. 逆波兰表达式求值
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
// 有效的算符为 '+'、'-'、'*' 和 '/' 。
// 输入：tokens = ["2","1","+","3","*"]
// 输出：9
// 解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
func evalRPN(tokens []string) int {
	r := 0
	var st []string
	for _, s := range tokens {
		if s == "+" || s == "-" || s == "*" || s == "/" {
			a := st[len(st)-1]
			st = st[:len(st)-1] // 出栈
			b := st[len(st)-1]
			st = st[:len(st)-1] // 出栈
			v2, _ := strconv.Atoi(a)
			v1, _ := strconv.Atoi(b)
			if s == "+" {
				r = v1 + v2
			} else if s == "-" {
				r = v1 - v2
			} else if s == "*" {
				r = v1 * v2
			} else if s == "/" {
				r = v1 / v2
			}
			st = append(st, fmt.Sprintf("%d", r))
		} else {
			st = append(st, s)
		}
	}
	if len(st) > 0 {
		r, _ = strconv.Atoi(st[0])
	}
	return r
}

func evalRPN2(tokens []string) int {
	result := 0
	var st []int
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			x, _ := strconv.Atoi(tokens[i])
			st = append(st, x)
			continue
		}
		a := st[len(st)-1]
		st = st[:len(st)-1]
		b := st[len(st)-1]
		st = st[:len(st)-1]
		switch {
		case tokens[i] == "+":
			result = a + b
		case tokens[i] == "-":
			result = b - a
		case tokens[i] == "*":
			result = a * b
		case tokens[i] == "/":
			result = b / a
		}
		st = append(st, result)
	}
	return st[0]
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

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}
