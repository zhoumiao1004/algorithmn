package main

import (
	"fmt"
	"sort"
)

/* 贪心-区间调度问题一般都要排序，为什么按右边界排序？因为排序后，当处理到第 i 个区间时，所有可能与它有交集的区间都在它之后
场景1: 假设现在只有一个会议室，还有若干会议，你如何将尽可能多的会议安排到这个会议室里？
这个问题需要将这些会议（区间）按结束时间（右端点）排序，然后进行处理。435. 无重叠区间(https://leetcode.cn/problems/non-overlapping-intervals/description/)

场景2: 给你若干较短的视频片段，和一个较长的视频片段，请你从较短的片段中尽可能少地挑出一些片段，拼接出较长的这个片段。
这个问题需要将这些视频片段（区间）按开始时间（左端点）排序，然后进行处理。1024. 视频拼接(https://leetcode.cn/problems/video-stitching/)

场景3: 给你若干区间，其中可能有些区间比较短，被其他区间完全覆盖住了，请你删除这些被覆盖的区间。
这个问题需要将这些区间按左端点排序，然后就能找到并删除那些被完全覆盖的区间了。1288. 删除被覆盖区间(https://leetcode.cn/problems/remove-covered-intervals/description/)

场景4: 给你若干区间，请你将所有有重叠部分的区间进行合并。
这个问题需要将这些区间按左端点排序，方便找出存在重叠的区间，56. 合并区间(https://leetcode.cn/problems/merge-intervals/submissions/)

场景5: 有两个部门同时预约了同一个会议室的若干时间段，请你计算会议室的冲突时段。
这个问题就是给你两组区间列表，请你找出这两组区间的交集，这需要你将这些区间按左端点排序，986. 区间列表的交集

场景6: 假设现在只有一个会议室，还有若干会议，如何安排会议才能使这个会议室的闲置时间最少？
这个问题需要动动脑筋，说白了这就是个 0-1 背包问题的变形：
会议室可以看做一个背包，每个会议可以看做一个物品，物品的价值就是会议的时长，请问你如何选择物品（会议）才能最大化背包中的价值（会议室的使用时长）？
当然，这里背包的约束不是一个最大重量，而是各个物品（会议）不能互相冲突。把各个会议按照结束时间进行排序。253. 会议室 II

场景7: 给你若干会议，让你最小化申请会议室的数量。1235. 规划兼职工作

*/
// 求这些区间中最多有几个互不相交的区间
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
// 思路1: 不相交的区间
func eraseOverlapIntervals(intervals [][]int) int {
	return len(intervals) - intervalSchedule(intervals)
}

// 思路2: 按左边界排序(not recommend)
func eraseOverlapIntervals2(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1]) // 更新右边界
			res++
		}
	}
	return res
}

// 452. 用最少数量的箭引爆气球
// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/description/
// 输入：points = [[10,16],[2,8],[1,6],[7,12]] 输出：2
// 解释：气球可以用2支箭来爆破:
// -在x = 6处射出箭，击破气球[2,8]和[1,6]。
// -在x = 11处发射箭，击破气球[10,16]和[7,12]。
// 一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。
// 思路1: 不相交的区间，按右边界排序（recommend）
func findMinArrowShots1(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	res := 1 // 至少需要一支箭
	xEnd := points[0][1]
	for i := 1; i < len(points); i++ {
		start := points[i][0] // 由于已经按右边界排序了，所以只用比较points[i]的左边界和xEnd的大小
		if start > xEnd {
			res++ // 不相交，需要增加一支箭
			xEnd = points[i][1]
		}
	}
	return res
}

// 思路2: 贪心思路：1.先按左边界排序 2.尽量重叠，能用最少的箭。重叠后合并需要更新右边界，不重叠需要增加一枝箭
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			res++ // 不重叠，一定需要增加一枝箭
		} else {
			points[i][1] = min(points[i][1], points[i-1][1]) // 重叠，更新右边界，箭往重叠的区域射能一箭双雕，所以这支箭的覆盖范围取min
		}
	}
	return res
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
	var res []int
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		right = max(right, hash[s[i]-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = right + 1
		}
	}
	return res
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

func removeCoveredIntervals2(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	cnt := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] <= intervals[i-1][1] {
			cnt++
			intervals[i][1] = max(intervals[i-1][1], intervals[i][1])
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
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for i := 1; i < n; i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
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
		if a1 <= b2 && b1 <= a2 {
			res = append(res, []int{max(a1, b1), min(a2, b2)})
		}
		if a2 < b2 {
			i++
		} else {
			j++
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

// 1024. 视频拼接
// https://leetcode.cn/problems/video-stitching/description/
// 你将会获得一系列视频片段，这些片段来自于一项持续时长为 time 秒的体育赛事。这些片段可能有所重叠，也可能长度不一。
// 使用数组 clips 描述所有的视频片段，其中 clips[i] = [starti, endi] 表示：某个视频片段开始于 starti 并于 endi 结束。
// 甚至可以对这些片段自由地再剪辑：
// 例如，片段 [0, 7] 可以剪切成 [0, 1] + [1, 3] + [3, 7] 三部分。
// 我们需要将这些片段进行再剪辑，并将剪辑后的内容拼接成覆盖整个运动过程的片段（[0, time]）。返回所需片段的最小数目，如果无法完成该任务，则返回 -1 。
// 输入：clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], time = 10
// 输出：3
// 解释：
// 选中 [0,2], [8,10], [1,9] 这三个片段。
// 然后，按下面的方案重制比赛片段：
// 将 [1,9] 再剪辑为 [1,2] + [2,8] + [8,9] 。
// 现在手上的片段为 [0,2] + [2,8] + [8,10]，而这些覆盖了整场比赛 [0, 10]。
func videoStitching(clips [][]int, T int) int {
	if T == 0 {
		return 0
	}
	// 按起点升序排列，起点相同的降序排列
	// PS：其实起点相同的不用降序排列也可以，不过我觉得这样更清晰
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		}
		return clips[i][0] < clips[j][0]
	})
	// 记录选择的短视频个数
	res := 0

	curEnd, nextEnd := 0, 0
	i, n := 0, len(clips)
	for i < n && clips[i][0] <= curEnd {
		// 在第 res 个视频的区间内贪心选择下一个视频
		for i < n && clips[i][0] <= curEnd {
			nextEnd = max(nextEnd, clips[i][1])
			i++
		}
		// 找到下一个视频，更新 curEnd
		res++
		curEnd = nextEnd
		if curEnd >= T {
			// 已经可以拼出区间 [0, T]
			return res
		}
	}
	// 无法连续拼出区间 [0, T]
	return -1
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
