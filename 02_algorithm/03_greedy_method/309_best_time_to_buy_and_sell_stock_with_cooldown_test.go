package _3_greedy_method

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 最佳买卖股票时机含冷冻期 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
func TestMaxProfitCollDown(t *testing.T) {
	convey.Convey("最佳买卖股票时机含冷冻期: 卖出股票后，你无法在第二天买入股", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{

				[]int{1, 2, 3, 0, 2}, 3,
			},
			{

				[]int{1}, 0,
			},
		}

		for _, tst := range testCase {
			rsp := maxProfitWithCoolDown(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func maxProfitWithCoolDown(prices []int) int {

	// 卖出后无法第二天买入,只有注意下第 i 天是第 i-2天买的  dfs[i,1]=  max(dfs(i-1, 1), dfs(i-2, 0)-prices[i])
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
