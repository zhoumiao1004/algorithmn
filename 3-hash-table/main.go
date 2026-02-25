package main

import (
	"fmt"
	"math"
	"sort"
)

// 242. 有效的字母异位词
// https://leetcode.cn/problems/valid-anagram/description/
// 输入: s = "anagram", t = "nagaram"
// 输出: true
func isAnagram(s string, t string) bool {
	var arr [26]int
	for _, c := range s {
		arr[c-'a']++
	}
	for _, c := range t {
		arr[c-'a']--
	}
	for i := 0; i < 26; i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s, t string) bool {
	var arr1, arr2 [26]int
	for _, c := range s {
		arr1[c-'a']++
	}
	for _, c := range t {
		arr2[c-'a']++
	}
	return arr1 == arr2
}

// 349. 两个数组的交集
// https://leetcode.cn/problems/intersection-of-two-arrays/description/
// 给定两个数组 nums1 和 nums2 ，返回 它们的 交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2]
func intersection(nums1, nums2 []int) []int {
	var results []int
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	for _, n := range nums1 {
		m1[n] = true
	}
	for _, n := range nums2 {
		if m1[n] && !m2[n] {
			results = append(results, n)
			m2[n] = true
		}
	}
	return results
}

// 202. 快乐数
// https://leetcode.cn/problems/happy-number/description/
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
// 输入：n = 19
// 输出：true
// 解释：
// 12 + 92 = 82
// 82 + 22 = 68
// 62 + 82 = 100
// 12 + 02 + 02 = 1
func isHappy(n int) bool {
	m := make(map[int]bool)
	for n > 1 {
		if m[n] {
			return false
		}
		m[n] = true
		s := 0
		for n > 0 {
			r := n % 10
			s += r * r
			n /= 10
		}
		n = s
	}
	return true
}

// 1. 两数之和
// https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		idx, ok := m[target-val]
		if ok {
			return []int{idx, i}
		}
		m[val] = i
	}
	return []int{}
}

// 454. 四数相加 II
// https://leetcode.cn/problems/4sum-ii/
// 输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
// 输出：2
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	result := 0
	m := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	for _, v1 := range nums3 {
		for _, v2 := range nums4 {
			cnt, ok := m[-v1-v2]
			if ok {
				result += cnt
			}
		}
	}
	return result
}

// 383. 赎金信
// https://leetcode.cn/problems/ransom-note/description/
// 判断 ransomNote 能不能由 magazine 里面的字符构成。
// ransomNote 和 magazine 由小写英文字母组成
func canConstruct(ransomNote string, magazine string) bool {
	var arr [26]int
	for _, c := range magazine {
		arr[c-'a']++
	}
	for _, c := range ransomNote {
		arr[c-'a']--
		if arr[c-'a'] < 0 {
			return false
		}
	}
	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	var arr [26]int
	for _, c := range magazine {
		arr[c-'a']++
	}
	for _, c := range ransomNote {
		if arr[c-'a'] == 0 {
			return false
		}
		arr[c-'a']--
	}
	return true
}

// 15. 三数之和
// https://leetcode.cn/problems/3sum/
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 双指针、滑动窗口，关键在如何去重
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var results [][]int
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if a > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			b, c := nums[l], nums[r]
			if a+b+c == 0 {
				results = append(results, []int{a, b, c})
				for l < r && nums[l] == b {
					l++
				}
				for l < r && nums[r] == c {
					r--
				}
			} else if a+b+c < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return results
}

// 18. 四数之和
// https://leetcode.cn/problems/4sum/
// 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
// 输入：nums = [1,0,-1,0,-2,2], target = 0
// 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var results [][]int
	for i := 0; i < len(nums)-3; i++ {
		a := nums[i]
		// if a > 0 {
		// 	break
		// }
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			b := nums[j]
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				c, d := nums[l], nums[r]
				if a+b+c+d == target {
					results = append(results, []int{a, b, c, d})
					for l < r && nums[l] == c {
						l++
					}
					for l < r && nums[r] == d {
						r--
					}
				} else if a+b+c+d < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return results
}

// 205. 同构字符串
// 给定两个字符串 s 和 t，判断它们是否是同构的。
// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
// 输入：s = "egg", t = "add"
// 输出：true
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	n := len(s)
	map1 := make(map[byte]byte)
	map2 := make(map[byte]byte)
	for i := 0; i < n; i++ {
		c, ok := map1[s[i]]
		if ok && c != t[i] {
			return false
		}
		map1[s[i]] = t[i]
		c2, ok2 := map2[t[i]]
		if ok2 && c2 != s[i] {
			return false
		}
		map2[t[i]] = s[i]
	}
	return true
}

// 1002. 查找共用字符
// 给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答案。
// 输入：words = ["bella","label","roller"]
// 输出：["e","l","l"]
func commonChars(words []string) []string {
	if len(words) == 0 {
		return words
	}
	var hash [26]int
	for _, c := range words[0] {
		hash[c-'a']++
	}
	for i := 1; i < len(words); i++ {
		var hash2 [26]int
		for _, c := range words[i] {
			hash2[c-'a']++
		}
		for i := 0; i < 26; i++ {
			hash[i] = min(hash[i], hash2[i])
		}
	}
	var result []string
	for i := 0; i < 26; i++ {
		for k := hash[i]; k > 0; k-- {
			result = append(result, string([]byte{byte('a' + i)}))
		}
	}

	return result
}

func commonChars2(words []string) []string {
	if len(words) == 0 {
		return words
	}
	var hash [26]int
	for i := 0; i < 26; i++ {
		hash[i] = math.MaxInt
	}
	for i := 0; i < len(words); i++ {
		var hash2 [26]int
		for _, c := range words[i] {
			hash2[c-'a']++
		}
		for i := 0; i < 26; i++ {
			hash[i] = min(hash[i], hash2[i])
		}
	}
	var result []string
	for i := 0; i < 26; i++ {
		if hash[i] == math.MaxInt {
			continue
		}
		for k := hash[i]; k > 0; k-- {
			result = append(result, fmt.Sprint('a'+i))
		}
	}
	return result
}

// 128.最长连续序列
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
func longestConsecutive(nums []int) int {
	n := len(nums)
	m := make(map[int]bool, n)
	for _, val := range nums {
		m[val] = true
	}
	result := 1
	for val := range m {
		// 1.val+1存在：说明val不是连续序列的最大值
		// 2.val+1不存在：val是连续序列的最大值，不断尝试val-1是否存在
		if !m[val+1] {
			cnt := 1
			i := val - 1
			for m[i] {
				cnt++
				result = max(result, cnt)
				i--
			}
		}
	}
	return result
}

// 217.存在重复元素
// 给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
// 输入：nums = [1,2,3,1]
// 输出：true
// 输入：nums = [1,2,3,4]
// 输出：false
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	for _, v := range m {
		if v >= 2 {
			return true
		}
	}
	return false
}

// 219.存在重复元素 II
// 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
// 输入：nums = [1,2,3,1], k = 3
// 输出：true
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i<len(nums); i++ {
		index, ok := m[nums[i]]
		if ok && i-index <= k {
			return true
		}
		m[nums[i]] = i
	}
	return false
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}
