package main

import (
	"fmt"
	"sort"
)

// 406. 根据身高重建队列
// https://leetcode.cn/problems/queue-reconstruction-by-height/description/
// 输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
// 输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
// 解释：
// 编号为 0 的人身高为 5 ，没有身高更高或者相同的人排在他前面。
// 编号为 1 的人身高为 7 ，没有身高更高或者相同的人排在他前面。
// 编号为 2 的人身高为 5 ，有 2 个身高更高或者相同的人排在他前面，即编号为 0 和 1 的人。
// 编号为 3 的人身高为 6 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
// 编号为 4 的人身高为 4 ，有 4 个身高更高或者相同的人排在他前面，即编号为 0、1、2、3 的人。
// 编号为 5 的人身高为 7 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
// 因此 [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] 是重新构造后的队列。
// 贪心思路：2个维度，身高和人数。先确定一个纬度身高降序，再往前面插入，往前插入不影响相对位置
func reconstructQueue(people [][]int) [][]int {
	// 先按照身高降序，人数升序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	for i := 1; i < len(people); i++ {
		index := people[i][1]
		if i > index {
			// 要插入到第people[i][1]个位置，先把前面的数字往后挪
			p := people[i]
			for j := i - 1; j >= index; j-- {
				people[j+1] = people[j]
			}
			people[index] = p
		}
	}
	return people
}

// 452. 用最少数量的箭引爆气球
// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/description/
// 输入：points = [[10,16],[2,8],[1,6],[7,12]] 输出：2
// 解释：气球可以用2支箭来爆破:
// -在x = 6处射出箭，击破气球[2,8]和[1,6]。
// -在x = 11处发射箭，击破气球[10,16]和[7,12]。
// 一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。
// 贪心思路：1.先按左边界排序 2.尽量重叠，能用最少的箭。重叠后合并需要更新右边界，不重叠需要增加一枝箭
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	result := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			// 不重叠，一定需要增加一枝箭
			result++
		} else {
			// 重叠，更新右边界，箭往重叠的区域射能一箭双雕，所以这支箭的覆盖范围取min
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}
	return result
}

// 435. 无重叠区间
// https://leetcode.cn/problems/non-overlapping-intervals/description/
// 给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回 需要移除区间的最小数量，使剩余区间互不重叠 。
// 注意 只在一点上接触的区间是 不重叠的。例如 [1, 2] 和 [2, 3] 是不重叠的。
// 输入: intervals = [[1,2],[2,3],[3,4],[1,3]] 输出: 1
// 解释: 移除 [1,3] 后，剩下的区间没有重叠。
// 输入: intervals = [ [1,2], [1,2], [1,2] ] 输出: 2
// 解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	result := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			// 重叠，更新右边界，相当于删除右边界更大的区间
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
			result++
		}
	}
	return result
}

// 763.划分字母区间
// https://leetcode.cn/problems/partition-labels/description/
// 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
// 把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。例如，字符串 "ababcc" 能够被分为 ["abab", "cc"]，但类似 ["aba", "bcc"] 或 ["ab", "ab", "cc"] 的划分是非法的。
// 输入：S = "ababcbacadefegdehijhklij"
// 输出：[9,7,8] 解释： 划分结果为 "ababcbaca", "defegde", "hijhklij"。
// 每个字母最多出现在一个片段中。 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。s只包含小写字母 'a' 到 'z' 。
func partitionLabels(s string) []int {
	// 记录最大位置
	hash := make([]int, 26)
	for i := 0; i < len(s); i++ {
		hash[s[i]-'a'] = i
	}
	// fmt.Println(hash)
	var result []int
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		right = max(right, hash[s[i]-'a'])
		if i == right {
			result = append(result, right-left+1)
			left = right + 1
		}
	}
	return result
}

// 56. 合并区间
// https://leetcode.cn/problems/merge-intervals/description/
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
func merge(intervals [][]int) [][]int {
	var results [][]int
	if len(intervals) == 0 {
		return results
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > intervals[i-1][1] {
			// 不重叠
			results = append(results, intervals[i-1])
		} else {
			// 重叠，更新左右边界
			intervals[i][0] = intervals[i-1][0]
			intervals[i][1] = max(intervals[i][1], intervals[i-1][1])
		}
	}
	results = append(results, intervals[len(intervals)-1])
	return results
}

func merge2(intervals [][]int) [][]int {
	var results [][]int
	if len(intervals) == 0 {
		return results
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	results = append(results, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > results[len(results)-1][1] {
			// 不重叠
			results = append(results, intervals[i])
		} else {
			results[len(results)-1][1] = max(results[len(results)-1][1], intervals[i][1])
		}
	}
	return results
}

func main() {
	fmt.Println(reconstructQueue([][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2}}))
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
