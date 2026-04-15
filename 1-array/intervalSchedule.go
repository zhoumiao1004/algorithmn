package main

import (
	"fmt"
	"sort"
)

// 贪心-区间调度问题，算出这些区间中最多有几个互不相交的区间
func intervalSchedule(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 1
	xEnd := intervals[0][1]
	for _, interval := range intervals {
		start := interval[0]
		if start >= xEnd {
			count++
			xEnd = interval[1]
		}
	}
	return count
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

// 思路2: 区间调度
func eraseOverlapIntervals2(intervals [][]int) int {
	return len(intervals) - intervalSchedule(intervals)
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

func findMinArrowShots2(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1
	xEnd := points[0][1]
	for _, interval := range points {
		start := interval[0]
		if start > xEnd {
			count++
			xEnd = interval[1]
		}
	}
	return count
}

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

// 1288. 删除被覆盖区间
// https://leetcode.cn/problems/remove-covered-intervals/description/
// 给你一个区间列表，请你删除列表中被其他区间所覆盖的区间。
// 只有当 c <= a 且 b <= d 时，我们才认为区间 [a,b) 被区间 [c,d) 覆盖。
// 在完成所有删除操作后，请你返回列表中剩余区间的数目。
// 输入：intervals = [[1,4],[3,6],[2,8]]
// 输出：2
// 解释：区间 [3,6] 被区间 [2,8] 覆盖，所以它被删除了。
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1] // 按照起点升序排列，起点相同时降序排列
		}
		return intervals[i][0] < intervals[j][0]
	})
	// 记录合并区间的起点和终点
	left, right := intervals[0][0], intervals[0][1]
	cnt := 0
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		// 情况一，找到覆盖区间
		if left <= interval[0] && right >= interval[1] {
			cnt++
		}
		// 情况二，找到相交区间，合并
		if interval[0] <= right && right <= interval[1] {
			right = interval[1]
		}
		// 情况三，完全不相交，更新起点和终点
		if right < interval[0] {
			left, right = interval[0], interval[1]
		}
	}
	return len(intervals) - cnt
}

// 56. 合并区间
// https://leetcode.cn/problems/merge-intervals/description/
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
func merge(intervals [][]int) [][]int {
	var results [][]int
	sort.Slice(intervals, func(i, j int) bool {
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
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	results = append(results, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > results[len(results)-1][1] {
			results = append(results, intervals[i]) // 不重叠
		} else {
			results[len(results)-1][1] = max(results[len(results)-1][1], intervals[i][1]) // 重叠
		}
	}
	return results
}

// 986. 区间列表的交集
// https://leetcode.cn/problems/interval-list-intersections/description/
// 给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而 secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。
// 返回这 两个区间列表的交集 。
// 形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。
// 两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。
// 返回需要申请的会议室数量,比如给你输入 meetings = [[0,30],[5,10],[15,20]]，算法应该返回 2，因为后两个会议和第一个会议时间是冲突的，至少申请两个会议室才能让所有会议顺利进行。
// 输入：firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
// 输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
// 思路: 双指针
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0
	var res [][]int
	for i < len(firstList) && j < len(secondList) {
		a1, a2 := firstList[i][0], firstList[i][1]
		b1, b2 := secondList[j][0], secondList[j][1]
		if b2 >= a1 && a2 >= b1 {
			res = append(res, []int{max(a1, b1), min(a2, b2)})
		}
		if b2 < a2 {
			j++
		} else {
			i++
		}
	}
	return res
}

func minMeetingRooms(meetings [][]int) int {
	n := len(meetings)
	begin := make([]int, n)
	end := make([]int, n)
	for i := 0; i < n; i++ {
		begin[i] = meetings[i][0]
		end[i] = meetings[i][1]
	}
	sort.Slice(begin, func(i, j int) bool {
		return begin[i] < begin[j]
	})
	sort.Slice(end, func(i, j int) bool {
		return end[i] < end[j]
	})
	count := 0
	result := 0
	i, j := 0, 0
	for i < n && j < n {
		if begin[i] < end[j] {
			count++
			i++
		} else {
			count--
			j++
		}
		result = max(result, count)
	}
	return result
}

func main() {
	fmt.Println(intervalSchedule([][]int{{1, 3}, {2, 4}, {3, 6}}))
	fmt.Println(reconstructQueue([][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2}}))
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
}
