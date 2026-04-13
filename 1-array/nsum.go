package main

import (
	"fmt"
	"sort"
)

func twoSumTarget(nums []int, start, target int) [][]int {
	sort.Ints(nums) // 先排序
	var lo, hi int = start, len(nums) - 1
	var res [][]int
	for lo < hi {
		var sum = nums[lo] + nums[hi]
		var left = nums[lo]
		var right = nums[hi]
		if sum < target {
			for lo < hi && nums[lo] == left {
				lo++
			}
		} else if sum > target {
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else {
			res = append(res, []int{left, right})
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		}
	}
	return res
}

func threeSumTarget(nums []int, start int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	// 穷举第一个数
	for i := start; i < n; i++ {
		tuples := twoSumTarget(nums, i+1, target-nums[i])
		for _, tuple := range tuples {
			tuple = append(tuple, nums[i])
			res = append(res, tuple)
		}
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func fourSumTarget(nums []int, target int) [][]int {
	sort.Ints(nums) // 先排序
	n := len(nums)
	var res [][]int
	// 穷举第一个数
	for i := 0; i < n; i++ {
		triples := threeSumTarget(nums, i+1, target-nums[i]) // 对 target - nums[i] 计算 threeSum
		for _, triple := range triples {
			triple = append(triple, nums[i]) // 存在满足条件的三元组，再加上 nums[i] 就是结果四元组
			res = append(res, triple)
		}
		// fourSum 的第一个数不能重复
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

// 167. 两数之和 II - 输入有序数组
// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
// 给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
// 以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
// 你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
// 你所设计的解决方案必须只使用常量级的额外空间。
// 输入：numbers = [2,7,11,15], target = 9
// 输出：[1,2]
// 解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
func twoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{-1, -1}
}

// 1. 两数之和
// https://leetcode.cn/problems/two-sum/description/
// 思路1: hashmap
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

// 思路2: 由于排序会改变index，所以用二元组把值和原始索引关联起来，这样无论值的位置怎么变，都可以找到最初的原始索引
func twoSum2(nums []int, target int) []int {
	type Pair struct {
		Val   int
		Index int
	}
	var pairs []*Pair
	for i, val := range nums {
		pairs = append(pairs, &Pair{Val: val, Index: i})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Val < pairs[j].Val
	})
	left, right := 0, len(nums)-1
	for left < right {
		s := pairs[left].Val + pairs[right].Val
		if s == target {
			return []int{pairs[left].Index, pairs[right].Index}
		} else if s < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

// 15. 三数之和
// https://leetcode.cn/problems/3sum/
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 思路: 双指针，关键在如何去重
func threeSum2(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		a := nums[i]
		if a > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
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

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSumTarget(nums, 0, 0)) // [[-1 2 -1] [0 1 -1]]
	nums2 := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSumTarget(nums2, 0)) // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
}
