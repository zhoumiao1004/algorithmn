package main

import (
	"fmt"
)

// 可以参考209. 长度最小的子数组 思路：滑动窗口
// https://leetcode.cn/problems/minimum-size-subarray-sum/description/?envType=problem-list-v2&envId=REW96uSa
// 某一个大文件被拆成了 N 个小文件，每个小文件编号从 0 至 N-1，相应大小分别记为 S(i)。给定磁盘空间为 C ，
// 试实现一个函数从 N 个文件中连续选出若干个文件拷贝到磁盘中，使得磁盘剩余空间最小。
func minDiskSpaceRemainder(nums []int, target int) int {

	left := 0
	s := 0
	maxSum := 0 // 不超过target的最大和
	for right := 0; right < len(nums); right++ {
		s += nums[right]
		for s >= target {
			s -= nums[left]
			left++
		}
		// 经过调整left，s已经满足<target条件了。求最大
		if s > maxSum {
			maxSum = s
		}
	}
	return target - maxSum
}

func main() {
	sizes := []int{100, 200, 300, 400, 500}
	capacity := 800
	fmt.Println(minDiskSpaceRemainder(sizes, capacity)) // 输出应为 100，因为选择文件 2, 3 (300 + 400 = 700)
}
