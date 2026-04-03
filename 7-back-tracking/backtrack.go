package main

import (
	"fmt"
)

// 526. 优美的排列
// https://leetcode.cn/problems/beautiful-arrangement/description/
// 假设有从 1 到 n 的 n 个整数。用这些整数构造一个数组 perm（下标从 1 开始），只要满足下述条件 之一 ，该数组就是一个 优美的排列 ：
// perm[i] 能够被 i 整除
// i 能够被 perm[i] 整除
// 给你一个整数 n ，返回可以构造的 优美排列 的 数量 。
// 输入：n = 2
// 输出：2
// 解释：
// 第 1 个优美的排列是 [1,2]：
//   - perm[1] = 1 能被 i = 1 整除
//   - perm[2] = 2 能被 i = 2 整除
//
// 第 2 个优美的排列是 [2,1]:
//   - perm[1] = 2 能被 i = 1 整除
//   - i = 2 能被 perm[2] = 1 整除
func countArrangement(n int) int {
	return 0
}

// 89. 格雷编码
// https://leetcode.cn/problems/gray-code/description/
// n 位格雷码序列 是一个由 2n 个整数组成的序列，其中：
// 每个整数都在范围 [0, 2n - 1] 内（含 0 和 2n - 1）
// 第一个整数是 0
// 一个整数在序列中出现 不超过一次
// 每对 相邻 整数的二进制表示 恰好一位不同 ，且
// 第一个 和 最后一个 整数的二进制表示 恰好一位不同
// 给你一个整数 n ，返回任一有效的 n 位格雷码序列 。
func grayCode(n int) []int {
	return []int{}
}

// 1849. 将字符串拆分为递减的连续值
// https://leetcode.cn/problems/splitting-a-string-into-descending-consecutive-values/description/
// 给你一个仅由数字组成的字符串 s 。
// 请你判断能否将 s 拆分成两个或者多个 非空子字符串 ，使子字符串的 数值 按 降序 排列，且每两个 相邻子字符串 的数值之 差 等于 1 。
// 例如，字符串 s = "0090089" 可以拆分成 ["0090", "089"] ，数值为 [90,89] 。这些数值满足按降序排列，且相邻值相差 1 ，这种拆分方法可行。
// 另一个例子中，字符串 s = "001" 可以拆分成 ["0", "01"]、["00", "1"] 或 ["0", "0", "1"] 。然而，所有这些拆分方法都不可行，因为对应数值分别是 [0,1]、[0,1] 和 [0,0,1] ，都不满足按降序排列的要求。
// 如果可以按要求拆分 s ，返回 true ；否则，返回 false 。
// 子字符串 是字符串中的一个连续字符序列。
// 输入：s = "1234"
// 输出：false
// 解释：不存在拆分 s 的可行方法。
// 输入：s = "050043"
// 输出：true
// 解释：s 可以拆分为 ["05", "004", "3"] ，对应数值为 [5,4,3] 。
// 满足按降序排列，且相邻值相差 1 。
func splitString(s string) bool {
	return false
}

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
}
