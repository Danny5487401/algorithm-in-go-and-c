package _3_greedy_method

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 买卖股票的最佳时机 II : https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
// 不限 交易次数

func TestMaxProfit(t *testing.T) {
	convey.Convey("买卖股票的最佳时机II: 不限交易次数", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{

				[]int{7, 1, 5, 3, 6, 4}, 7,
			},
			{

				[]int{1, 2, 3, 4, 5}, 4,
			},
			{

				[]int{7, 6, 4, 3, 1}, 0,
			},
		}

		for _, tst := range testCase {
			rsp := maxProfit(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func maxProfit(prices []int) int {
	/*
		- 未持有 --(买入)--> 持有   dfs(i,1)=dfs(i-1,0)-price[i]
		- 持有 --(卖出-)--> 未持有  dfs(i,0)=dfs(i-1,1)+price[i]
		- 未持有(do nothing) dfs(i,0) = dfs(i-1,0)
		- 持有(do nothing)  dfs(i,1) = dfs(i-1,1)
	*/
	n := len(prices)

	var cache = make([][2]int, n)
	for i := range cache {
		cache[i] = [2]int{-1, -1} // -1 表示还没有计算过
	}
	var dfs func(i int, hold int) int
	dfs = func(i int, hold int) (res int) {
		if i < 0 {
			if hold == 1 { // dfs(-1.1)= -∞ 第0天开始不可能持有股票
				return math.MinInt64
			} else {
				return 0 // dfs(-1.0)= 0 第0天开始未持有股票
			}
		}
		p := &cache[i][hold]
		if *p != -1 {
			return *p
		}

		defer func() { *p = res }() // 记忆化
		if hold == 1 {              // 计算持有股票的状态
			// dfs[i,1]=  max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
			return max(dfs(i-1, 1), dfs(i-1, 0)-prices[i]) // 不交易 或则 买入股票
		}
		// dfs[i,0]=  max(dfs(i-1, 0), dfs(i-1, 1)+prices[i]) // 不交易 或则 卖出股票
		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
	}
	return dfs(n-1, 0) // 最后一天卖出股票 dfs(n-1, 1) < dfs(n-1, 0)
}
