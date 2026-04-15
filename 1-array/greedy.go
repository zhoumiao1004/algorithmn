package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 455.分发饼干
// https://leetcode.cn/problems/assign-cookies/description/
// 胃口值 g[i] 饼干尺寸 s[j]， 目标是满足尽可能多的孩子，并输出这个最大数值。
// 输入: g = [1,2,3], s = [1,1] 输出: 1
// 输入: g = [1,2], s = [1,2,3] 输出: 2
// 贪心，双指针。局部最优：用大饼干满足最大胃口
func findContentChildren(g []int, s []int) int {
	result := 0
	sort.Ints(g)
	sort.Ints(s)
	j := len(s) - 1
	for i := len(g) - 1; i >= 0; i-- {
		// 最大小孩的胃口是g[i]
		if j >= 0 && s[j] >= g[i] {
			result++ // 能满足当前最大胃口
			j--
		}
	}
	return result
}

// 376. 摆动序列
// https://leetcode.cn/problems/wiggle-subsequence/description/
// 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。
// 例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。
// 相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
// 返回 nums 中作为 摆动序列 的 最长子序列的长度 。
// 输入: [1,17,5,10,13,15,10,5,16,8] 输出: 7
// 解释: 这个序列包含几个长度为 7 摆动序列，其中一个可为[1,17,10,13,10,16,8]。
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	result := 1
	prediff := 0 // 向左延伸一个
	curdiff := 0
	for i := 0; i < len(nums)-1; i++ {
		curdiff = nums[i+1] - nums[i]
		if (prediff <= 0 && curdiff > 0) || (prediff >= 0 && curdiff < 0) {
			result++
			prediff = curdiff // 遇到摆动才换
		}
	}
	return result
}

// 122. 买卖股票的最佳时机 II
// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。然而，你可以在 同一天 多次买卖该股票，但要确保你持有的股票不超过一股。
// 返回 你能获得的 最大 利润 。
// 贪心思路：可以多次买卖，所以累加所有增量
func maxProfit(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		result += max(0, prices[i]-prices[i-1])
	}
	return result
}

// 55. 跳跃游戏
// https://leetcode.cn/problems/jump-game/description/
// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
// 输入：nums = [2,3,1,1,4] true
// 输入：nums = [3,2,1,0,4] true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
// 思路：覆盖范围是否包含最后一个位置
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	cover := 0
	for i := 0; i <= cover; i++ {
		// 覆盖范围下标是到i+nums[i]
		cover = max(cover, i+nums[i])
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

// 45.跳跃游戏 II
// https://leetcode.cn/problems/jump-game-ii/
// 返回到达 n - 1 的最小跳跃次数。测试用例保证可以到达 n - 1。
// 输入: nums = [2,3,1,1,4] 输出: 2
// 输入: nums = [2,3,0,1,4] 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	cnt := 0
	cur := 0 // 不断更新为本轮最远地方
	cover := 0
	for i := 0; i < len(nums); i++ {
		cover = max(cover, i+nums[i]) // 收集本层最远能跳到哪
		if i == cur {
			// 已经遍历到本层的最后一个格子
			if cur == len(nums)-1 {
				return cnt
			}
			cnt++ // 需要多走一步
			cur = cover
			if cur >= len(nums)-1 {
				return cnt
			}
		}
	}
	return cnt
}

func jump2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	cnt := 0
	cur := 0 // 本轮最远
	cover := 0
	for i := 0; i < len(nums); i++ {
		if i == cur+1 {
			cnt++ // 需要多走一步
			cur = cover
		}
		cover = max(cover, i+nums[i]) // 收集本层最远能跳到哪
	}
	return cnt
}

// 1005.K次取反后最大化的数组和
// https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/
// 给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
// 选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
// 重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
// 以这种方式修改数组后，返回数组 可能的最大和 。
// 输入：nums = [4,2,3], k = 1 输出：5
// 解释：选择下标 1 ，nums 变为 [4,-2,3] 。
// 贪心思路：1.优先选绝对值大的负数取反 2.全部是正数的时候，局部最优：最小的正数反复取反，直到把k消耗完
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	minIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
		if nums[i] < nums[minIndex] {
			minIndex = i
		}
	}
	if k%2 == 1 {
		nums[minIndex] *= -1
	}
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}

// 134. 加油站
// 在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
// 如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
// 说明:
// 如果题目有解，该答案即为唯一答案。
// 输入数组均为非空数组，且长度相同。
// 输入数组中的元素均为非负数。
// 示例 1: 输入: gas = [1,2,3,4,5] cost = [3,4,5,1,2] 输出：3
func canCompleteCircuit(gas []int, cost []int) int {
	result := 0
	n := len(gas)
	s := 0
	for i := 0; i < n; i++ {
		s += gas[i] - cost[i]
	}
	if s < 0 {
		return -1
	}
	s = 0
	for i := 0; i < len(gas); i++ {
		s += gas[i] - cost[i]
		if s < 0 {
			result = i + 1
			s = 0
		}
	}
	return result
}

// 135. 分发糖果
// https://leetcode.cn/problems/candy/description/
// 老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
// 你需要按照以下要求，帮助老师给这些孩子分发糖果：
// 每个孩子至少分配到 1 个糖果。
// 相邻的孩子中，评分高的孩子必须获得更多的糖果。
// 那么这样下来，老师至少需要准备多少颗糖果呢？
// 示例 1: 输入: [1,0,2] 输出: 5
// 解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
// 示例 2: 输入: [1,2,2] 输出: 4
// 解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
func candy(ratings []int) int {
	n := len(ratings)
	arr1 := make([]int, n)
	arr2 := make([]int, n)
	for i := 0; i < n; i++ {
		arr1[i] = 1
		arr2[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			arr1[i] = arr1[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			arr2[i] = arr2[i+1] + 1
		}
	}
	s := 0
	for i := 0; i < n; i++ {
		s += max(arr1[i], arr2[i])
	}
	return s
}

// 860. 柠檬水找零
// https://leetcode.cn/problems/lemonade-change/description/
// 在柠檬水摊上，每一杯柠檬水的售价为 5 美元。顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。
// 每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。
// 注意，一开始你手头没有任何零钱。
// 给你一个整数数组 bills ，其中 bills[i] 是第 i 位顾客付的账。如果你能给每位顾客正确找零，返回 true ，否则返回 false 。
// 输入：bills = [5,5,5,10,20] 输出：true
// 前 3 位顾客那里，我们按顺序收取 3 张 5 美元的钞票。
// 第 4 位顾客那里，我们收取一张 10 美元的钞票，并返还 5 美元。
// 第 5 位顾客那里，我们找还一张 10 美元的钞票和一张 5 美元的钞票。
// 由于所有客户都得到了正确的找零，所以我们输出 true。
// 贪心思路：优先给10块，没有了才给5块
func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			ten++
			five--
		} else if bills[i] == 20 {
			if ten > 0 {
				ten--
				five--
			} else {
				five -= 3
			}
		}
		if five < 0 || ten < 0 {
			return false
		}
	}
	return true
}

// 738. 单调递增的数字
// https://leetcode.cn/problems/monotone-increasing-digits/description/
// 当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。
// 给定一个整数 n ，返回 小于或等于 n 的最大数字，且数字呈 单调递增 。
// 输入: n = 10 输出: 9
// 思路：1.转成字符串处理 2.从后向前找第一个下降的地方，左减1后变9 3.标记下降处,后面都需要填9 1000 => 999
func monotoneIncreasingDigits(n int) int {
	// 转成字符串
	s := strconv.Itoa(n)
	b := []byte(s)
	k := len(b)
	for i := len(b) - 1; i > 0; i-- {
		if b[i] < b[i-1] { // 找到下降处
			b[i-1]--
			b[i] = '9'
			k = i // 标记
		}
	}
	for k < len(b) {
		b[k] = '9'
		k++
	}
	result, _ := strconv.Atoi(string(b))
	return result
}

func monotoneIncreasingDigits2(n int) int {
	s := strconv.Itoa(n)
	b := []byte(s)
	flag := len(s)
	for i := len(b) - 1; i > 0; i-- {
		if b[i] < b[i-1] {
			b[i-1] -= 1
			flag = i
		}
	}
	for i := flag; i < len(b); i++ {
		b[i] = '9'
	}
	result, _ := strconv.Atoi(string(b))
	return result
}

func monotoneIncreasingDigits3(n int) int {
	s := strconv.Itoa(n)
	bs := []byte(s)
	k := len(bs)
	for i := len(bs) - 2; i >= 0; i-- {
		if bs[i] > bs[i+1] {
			bs[i] -= 1
			bs[i+1] = '9'
			k = i + 1
		}
	}
	for ; k < len(bs); k++ {
		bs[k] = '9'
	}
	result, _ := strconv.Atoi(string(bs))
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 968.监控二叉树
// https://labuladong.online/zh/algo/problem-set/binary-tree-post-order-iii/#slug_binary-tree-cameras
// 给定一个二叉树，我们在树的节点上安装摄像头。
// 节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
// 计算监控树的所有节点所需的最小摄像头数量。
// 思路：从叶子节点向上，叶子的父节点放一个，向上空2个节点放一个 3种状态转移：0.无覆盖 1.有摄像头 2.有覆盖
func minCameraCover(root *TreeNode) int {
	result := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		// 空节点有覆盖，上面的叶子节点就无覆盖
		if root == nil {
			return 2
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		// 1.左右都有覆盖 => 父节点无覆盖
		if left == 2 && right == 2 {
			return 0
		}
		// 2.左或右无覆盖 => 父节点要放一个摄像头
		if left == 0 || right == 0 {
			result++
			return 1
		}
		// 3.左或右有摄像头 => 父节点有覆盖
		if left == 1 || right == 1 {
			return 2
		}
		return -1
	}
	status := dfs(root)
	if status == 0 {
		result++
	}
	return result
}

// 2149. 按符号重排数组
// https://leetcode.cn/problems/rearrange-array-elements-by-sign/description/
func rearrangeArray(arr []int) []int {
	// 统计每个数字出现的次数
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}

	// 将元素按照出现次数排序，先放出现最多的数字
	type freq struct {
		num   int
		count int
	}
	freqs := []freq{}
	for num, count := range counts {
		freqs = append(freqs, freq{num, count})
	}
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})

	// 创建结果数组并按顺序填充
	n := len(arr)
	result := make([]int, n)

	// 使用贪心算法，将数字按最大频率的顺序放置，并确保相邻不相同
	idx := 0
	for _, f := range freqs {
		for i := 0; i < f.count; i++ {
			result[idx] = f.num
			idx += 2
			if idx >= n {
				idx = 1
			}
		}
	}

	return result
}

// 1221. 分割平衡字符串
// https://leetcode.cn/problems/split-a-string-in-balanced-strings/description/
// 平衡字符串 中，'L' 和 'R' 字符的数量是相同的。
// 给你一个平衡字符串 s，请你将它分割成尽可能多的子字符串，并满足：
// 每个子字符串都是平衡字符串。
// 返回可以通过分割得到的平衡字符串的 最大数量 。
// 输入：s = "RLRRLLRLRL"
// 输出：4
// 解释：s 可以分割为 "RL"、"RRLL"、"RL"、"RL" ，每个子字符串中都包含相同数量的 'L' 和 'R' 。
func balancedStringSplit(s string) int {
	result := 0
	diff := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			diff++
		} else {
			diff--
		}
		if diff == 0 {
			result++
		}
	}
	return result
}

func main() {
	// 给出一个数组，返回相邻不相同的一个数组，例如 [1,1,2,2,2,3] 返回 [2,1,2,1,2,3]
	arr := []int{1, 1, 2, 2, 2, 3}
	result := rearrangeArray(arr)
	fmt.Println(result)                                           // 输出例如 [2, 1, 2, 1, 2, 3]
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1})) // 输出例如 [2, 1, 2, 1, 2, 3]
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))  // 2
	fmt.Println(jump2([]int{2, 3, 0, 1, 4})) // 2
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
	fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{4, 5, 2, 6, 5, 3}, []int{3, 2, 7, 3, 2, 9}))
	fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}))
	fmt.Println(monotoneIncreasingDigits(10))
}
