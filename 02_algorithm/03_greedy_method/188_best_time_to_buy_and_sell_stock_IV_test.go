package _3_greedy_method

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 买卖股票的最佳时机IV : https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/
// 交易次数最多为K

func TestMaxProfitIV(t *testing.T) {
	convey.Convey("买卖股票的最佳时机 IV:交易次数最多为K ", t, func() {
		testCase := []struct {
			input  []int
			k      int
			target int
		}{
			{

				[]int{2, 4, 1}, 2, 2,
			},
			{

				[]int{3, 2, 6, 5, 0, 3}, 2, 7,
			},
		}

		for _, tst := range testCase {
			rsp := maxProfitIV(tst.k, tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func maxProfitIV(k int, prices []int) int {
	n := len(prices)

	var cache = make([][][2]int, n)
	for i := range cache {
		cache[i] = make([][2]int, k+1)
		for j := range cache[i] {
			cache[i][j] = [2]int{-1, -1} // -1 表示还没有计算过
		}

	}
	var dfs func(i int, count int, hold int) int // 交易次数 count
	dfs = func(i int, count int, hold int) (res int) {
		if count < 0 {
			return math.MinInt / 2 // 防止溢出
		}
		if i < 0 {
			if hold == 1 {
				return math.MinInt / 2 // 防止溢出
			} else {
				return 0
			}
		}
		p := &cache[i][count][hold]
		if *p != -1 {
			return *p
		}

		defer func() { *p = res }() // 记忆化
		if hold == 1 {
			// dfs[i,k,1]=  max(dfs(i-1, k,1), dfs(i-1,k-1, 0)-prices[i])
			return max(dfs(i-1, count, 1), dfs(i-1, count-1, 0)-prices[i]) // 不交易 或则 买入股票
		}
		// dfs[i,k,0]=  max(dfs(i-1, k,0), dfs(i-1,k-1, 1)+prices[i]) // 不交易 或则 卖出股票
		return max(dfs(i-1, count, 0), dfs(i-1, count, 1)+prices[i])
	}
	return dfs(n-1, k, 0) // 最后一天卖出股票
}
