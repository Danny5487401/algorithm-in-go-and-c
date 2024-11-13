package _3_greedy_method

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 买卖股票的最佳时机 :https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
// 只能买卖一次

func TestMaxProfitI(t *testing.T) {
	convey.Convey("买卖股票的最佳时机: 只有买和卖一次交易次数", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{

				[]int{7, 1, 5, 3, 6, 4}, 5,
			},
			{

				[]int{7, 6, 4, 3, 1}, 0,
			},
		}

		for _, tst := range testCase {
			rsp := maxProfitI(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func maxProfitI(prices []int) int {
	/*
		- 未持有 --(买入)--> 持有   dfs(i,1)=dfs(i-1,0)-price[i]
		- 持有 --(卖出-)--> 未持有  dfs(i,0)=dfs(i-1,1)+price[i]
		- 未持有(do nothing) dfs(i,0) = dfs(i-1,0)
		- 持有(do nothing)  dfs(i,1) = dfs(i-1,1)
	*/
	//  要想获得最大利润，我们需要知道第 i 天之前，股票价格的最小值是什么，也就是从 prices[0] 到 prices[i−1] 的最小值，把它作为买入价格，这可以用一个变量 minPrice 维护

	minPrice := prices[0]
	var ans int
	for _, p := range prices {
		ans = max(ans, p-minPrice)  // 最大利润
		minPrice = min(minPrice, p) // 股票价格的最小值
	}
	return ans

}
