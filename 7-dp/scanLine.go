package main

type Interval struct {
	Start, End int
}

// 920 · 会议室
// https://www.lintcode.com/problem/920/description
// 给定一系列的会议时间间隔，包括起始和结束时间[(s1,e1)，(s2,e2)，…(si < ei)，确定一个人是否可以参加所有会议。
// 输入: intervals = [(0,30),(5,10),(15,20)]
// 输出: false
// 解释:
// (0,30), (5,10) 和 (0,30),(15,20) 这两对会议会冲突
func CanAttendMeetings(intervals []*Interval) bool {
	// Write your code here
}

// 919 · 会议室 II
// 给定一系列的会议时间间隔intervals，包括起始和结束时间[[s1,e1],[s2,e2],...] (si < ei)，找到所需的最小的会议室数量。
// 输入: intervals = [(0,30),(5,10),(15,20)]
// 输出: 2
// 解释:
// 需要两个会议室
// 会议室1:(0,30)
// 会议室2:(5,10),(15,20)
// 分析：如果会议之间的时间有重叠，那就得额外申请会议室来开会，想求至少需要多少间会议室，就是让你计算同一时刻最多有多少会议在同时进行。
func MinMeetingRooms(intervals []*Interval) int {
	// Write your code here
}
