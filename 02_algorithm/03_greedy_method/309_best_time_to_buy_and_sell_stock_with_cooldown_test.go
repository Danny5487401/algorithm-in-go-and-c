package _3_greedy_method

import "math"

// 最佳买卖股票时机含冷冻期 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/

func maxProfitWithCoolDown(prices []int) int {

	// 卖出后无法第二天买入
	n := len(prices)

	var cache = make([][2]int, n)
	for i := range cache {
		cache[i] = [2]int{-1, -1} // -1 表示还没有计算过
	}
	var dfs func(i int, hold int) int
	dfs = func(i int, hold int) (res int) {
		if i < 0 {
			if hold == 1 {
				return math.MinInt64
			} else {
				return 0
			}
		}
		p := &cache[i][hold]
		if *p != -1 {
			return *p
		}

		defer func() { *p = res }() // 记忆化
		if hold == 1 {
			// dfs[i,1]=  max(dfs(i-1, 1), dfs(i-2, 0)-prices[i])
			return max(dfs(i-1, 1), dfs(i-2, 0)-prices[i]) // 不交易 或则 买入股票
		}
		// dfs[i,0]=  max(dfs(i-1, 0), dfs(i-1, 1)+prices[i]) // 不交易 或则 卖出股票
		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
	}
	return dfs(n-1, 0) // 最后一天卖出股票
}
