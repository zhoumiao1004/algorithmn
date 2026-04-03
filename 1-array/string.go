package main

import (
	"fmt"
	"strings"
)

// 125. 验证回文串
// https://leetcode.cn/problems/valid-palindrome/description/
// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
// 字母和数字都属于字母数字字符。
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
// 输入: s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。
func isPalindrome(s string) bool {
	bs := []byte(s)
	// 保留小写字母
	slow := 0
	for i := 0; i < len(bs); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			bs[slow] = c
			slow++
		} else if c >= 'A' && c <= 'Z' {
			bs[slow] = c - 'A' + 'a'
			slow++
		}
	}
	// fmt.Println(string(bs[:slow]))
	left, right := 0, slow-1
	for left < right {
		if bs[left] != bs[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 541. 反转字符串 II
// https://leetcode.cn/problems/reverse-string-ii/description/
// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
// 输入：s = "abcdefg", k = 2
// 输出："bacdfeg"
func reverseStr(s string, k int) string {
	n := len(s)
	bs := []byte(s)
	for i := 0; i < n; i += 2 * k {
		// 反转前k个
		reverseString(bs[i:min(i+k, n)])
	}
	return string(bs)
}

func reverseStr2(s string, k int) string {
	ss := []byte(s)
	n := len(s)
	for i := 0; i < n; i += 2 * k {
		if n-i < k {
			reverseString(ss[i:])
		} else {
			reverseString(ss[i : i+k])
		}
	}
	return string(ss)
}

// 替换数字
// 给定一个字符串 s，它包含小写字母和数字字符，请编写一个函数，将字符串中的字母字符保持不变，而将每个数字字符替换为number。
// 例如，对于输入字符串 "a1b2c3"，函数应该将其转换为 "anumberbnumbercnumber"。
// 对于输入字符串 "a5b"，函数应该将其转换为 "anumberb"
func replaceNumber(s string) string {
	return s
}

// 151. 反转字符串中的单词
// https://leetcode.cn/problems/reverse-words-in-a-string/
// 输入：s = "the sky is blue"
// 输出："blue is sky the"
// 1.删除多余的空格 2.整体反转 3.反转每个单词
func reverseWords(s string) string {
	// removeExtraSpaces
	bs := removeExtraSpaces([]byte(s))
	// fmt.Println(string(bs))
	// reverse whole string
	reverseString(bs)
	// reverse each word
	slow := 0
	for fast := 0; fast <= len(bs); fast++ {
		if fast == len(bs) || bs[fast] == ' ' {
			reverseString(bs[slow:fast])
			slow = fast + 1
		}
	}
	return string(bs)
}

func removeExtraSpaces(s []byte) []byte {
	slow := 0
	for fast := 0; fast < len(s); fast++ {
		if s[fast] != ' ' {
			if slow > 0 { // 单词之间补空格：除了第一个单词
				s[slow] = ' '
				slow++
			}
			for fast < len(s) && s[fast] != ' ' {
				s[slow] = s[fast]
				slow++
				fast++
			}
		}
	}
	return s[:slow]
}

// 右旋字符串
// 例如，对于输入字符串 "abcdefg" 和整数 2，函数应该将其转换为 "fgabcde"

// 28. 实现 strStr()
// KMP算法
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
// 示例 1: 输入: haystack = "hello", needle = "ll" 输出: 2
// 示例 2: 输入: haystack = "aaaaa", needle = "bba" 输出: -1
// 说明: 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	next := make([]int, n)
	getNext(next, needle)
	fmt.Println(next)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}
func getNext(next []int, s string) {
	j := 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

// 459.重复的子字符串
// https://leetcode.cn/problems/repeated-substring-pattern/description/
// 给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。
// 输入: s = "abab" 输出: true
// 解释: 可由子串 "ab" 重复两次构成。
// "aba" => "abaaba" => baab 不包含 aba => false
// "abab" => "abababab" => bababa 包含 abab => true
func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}
	t := s + s
	// fmt.Println(t[1 : len(t)-1])
	return strings.Contains(t[1:len(t)-1], s)
}

// 3. 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
// 输入: s = "abcabcbb" 输出: 3
// 注意：子串不是子序列
func lengthOfLongestSubstring(s string) string {
	m := make(map[byte]int)
	n := len(s)
	//ans := 0
	l := 0
	r := 0
	left := 0
	for right := 0; right < n; right++ {
		ch := s[right]
		m[ch]++
		for m[ch] > 1 {
			ch2 := s[left]
			m[ch2]--
			left++
		}
		if right-left+1 > r-l {
			l, r = left, right
		}
	}
	return s[l : r+1]
}

// 925. 长按键入
// https://leetcode.cn/problems/long-pressed-name/description/
// 你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。
// 你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。
// 输入：name = "alex", typed = "aaleex" 输出：true
// 解释：'alex' 中的 'a' 和 'e' 被长按。
func isLongPressedName(name string, typed string) bool {
	idx := 0
	var last byte
	for i := 0; i < len(typed); i++ {
		if idx < len(name) && name[idx] == typed[i] {
			last = name[idx]
			idx++
		} else if last == typed[i] {
			continue
		} else {
			return false
		}
	}
	return idx == len(name)
}

// 844. 比较含退格的字符串
// https://leetcode.cn/problems/backspace-string-compare/description/
// 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。
// 注意：如果对空文本输入退格字符，文本继续为空。
func backspaceCompare(s string, t string) bool {
	var st1 []byte
	var st2 []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(st1) > 0 {
				st1 = st1[:len(st1)-1]
			}
		} else {
			st1 = append(st1, s[i])
		}
	}
	for i := 0; i < len(t); i++ {
		if t[i] == '#' {
			if len(st2) > 0 {
				st2 = st2[:len(st2)-1]
			}
		} else {
			st2 = append(st2, t[i])
		}
	}
	return string(st1) == string(st2)
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))

	s := "the sky is blue"
	fmt.Println(reverseWords(s))

	// “abcdbca"字符串，找出不含重复字符的最长子串的长度
	s = "abcdaefca"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(reverseWords("  the sky is blue  "))
	fmt.Println(string(removeExtraSpaces([]byte("  the sky is blue  "))))
	fmt.Println(strStr("hello", "ll"))         // 2
	fmt.Println(strStr("aabaabaaf", "aabaaf")) // 3

	// fmt.Println(string(removeExtraSpaces2([]byte("  the sky is blue  "))))
}
