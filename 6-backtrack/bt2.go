package main

import (
	"fmt"
	"math"
)

// 638. 大礼包
// https://leetcode.cn/problems/shopping-offers/description/
// 在 LeetCode 商店中， 有 n 件在售的物品。每件物品都有对应的价格。然而，也有一些大礼包，每个大礼包以优惠的价格捆绑销售一组物品。
// 给你一个整数数组 price 表示物品价格，其中 price[i] 是第 i 件物品的价格。另有一个整数数组 needs 表示购物清单，其中 needs[i] 是需要购买第 i 件物品的数量。
// 还有一个数组 special 表示大礼包，special[i] 的长度为 n + 1 ，其中 special[i][j] 表示第 i 个大礼包中内含第 j 件物品的数量，且 special[i][n] （也就是数组中的最后一个整数）为第 i 个大礼包的价格。
// 返回 确切 满足购物清单所需花费的最低价格，你可以充分利用大礼包的优惠活动。你不能购买超出购物清单指定数量的物品，即使那样会降低整体价格。任意大礼包可无限次购买。
// 输入：price = [2,5], special = [[3,0,5],[1,2,10]], needs = [3,2]
// 输出：14
// 解释：有 A 和 B 两种物品，价格分别为 ¥2 和 ¥5 。
// 大礼包 1 ，你可以以 ¥5 的价格购买 3A 和 0B 。
// 大礼包 2 ，你可以以 ¥10 的价格购买 1A 和 2B 。
// 需要购买 3 个 A 和 2 个 B ， 所以付 ¥10 购买 1A 和 2B（大礼包 2），以及 ¥4 购买 2A 。
func shoppingOffers(price []int, special [][]int, needs []int) int {
	minCost := math.MaxInt
	cost := 0
	var filterSpecials func(prices []int, specials [][]int) [][]int
	var canUseSpecial func(sp, needs []int) bool
	var backtrack func(start int)

	filterSpecials = func(price []int, specials [][]int) [][]int {
		var newSpecials [][]int
		for _, sp := range specials {
			cost := 0
			for j := 0; j < len(sp)-1; j++ {
				cost += sp[j] * price[j]
			}
			if cost > sp[len(sp)-1] {
				newSpecials = append(newSpecials, sp)
			}
		}
		return newSpecials
	}

	canUseSpecial = func(sp, needs []int) bool {
		for i := 0; i < len(needs); i++ {
			if sp[i] > needs[i] {
				return false
			}
		}
		return true
	}

	specials := filterSpecials(price, special)

	backtrack = func(start int) {
		if cost > minCost {
			return
		}
		useSpecial := false
		for i := start; i < len(specials); i++ {
			sp := specials[i]
			if !canUseSpecial(sp, needs) {
				continue
			}
			useSpecial = true
			// 买 specials[i] 这个大礼包
			for j := 0; j < len(needs); j++ {
				needs[j] -= sp[j]
			}
			cost += sp[len(sp)-1]
			backtrack(i) // 一个大礼包能买多次
			// 撤销买 specials[i] 这个大礼包
			for j := 0; j < len(needs); j++ {
				needs[j] += sp[j]
			}
			cost -= sp[len(sp)-1]
		}
		if !useSpecial {
			// 无法购买剩余的大礼包specials[start...]，剩下的只能单买
			s := 0
			for i := 0; i < len(needs); i++ {
				s += price[i] * needs[i]
			}
			minCost = min(minCost, cost+s)
		}

	}

	backtrack(0)
	return minCost
}

func main() {
	fmt.Println(numTilePossibilities("AAB"))
}
