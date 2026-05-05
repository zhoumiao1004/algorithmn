package main

// 1109. 航班预订统计
// 这里有 n 个航班，它们分别从 1 到 n 进行编号。
// 有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
// 请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
// 输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
// 输出：[10,55,45,25,25]
// 解释：
// 航班编号        1   2   3   4   5
// 预订记录 1 ：   10  10
// 预订记录 2 ：       20  20
// 预订记录 3 ：       25  25  25  25
// 总座位数：      10  55  45  25  25
// 因此，answer = [10,55,45,25,25]
func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	d := NewDifference(nums)
	for _, b := range bookings {
		i, j, val := b[0]-1, b[1]-1, b[2]
		d.increment(i, j, val)
	}
	return d.result()
}

type Difference struct {
	diff []int
}

func NewDifference(nums []int) *Difference {
	n := len(nums)
	diff := make([]int, n)
	diff[0] = nums[0]
	for i := 1; i < n; i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{diff: diff}
}

func (d *Difference) increment(i, j, val int) {
	d.diff[i] += val
	n := len(d.diff)
	if j+1 < n {
		d.diff[j+1] -= val
	}
}

func (d *Difference) result() []int {
	diff := d.diff
	n := len(diff)
	nums := make([]int, n)
	nums[0] = diff[0]
	for i := 1; i < n; i++ {
		nums[i] = nums[i-1] + diff[i]
	}
	return nums
}

// 1094. 拼车
// https://leetcode.cn/problems/car-pooling/description/
// 车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）
// 给定整数 capacity 和一个数组 trips ,  trips[i] = [numPassengersi, fromi, toi] 表示第 i 次旅行有 numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。
// 当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。
// 输入：trips = [[2,1,5],[3,3,7]], capacity = 4
// 输出：false
func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001)
	df := NewDifference(nums)
	for _, trip := range trips {
		val, i, j := trip[0], trip[1], trip[2]-1
		df.increment(i, j, val)
	}
	res := df.result()
	for i := 0; i<len(res); i++ {
		if res[i] > capacity {
			return false
		}
	} 
	return true
}
