package _3_greedy_method

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 买卖股票的最佳时机 III: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/description/
// 不限 交易次数

func TestMaxProfitIII(t *testing.T) {
	convey.Convey("买卖股票的最佳时机III: 最多两次交易", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{

				[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6,
			},
			{

				[]int{1, 2, 3, 4, 5}, 4,
			},
			{

				[]int{7, 6, 4, 3, 1}, 0,
			},
		}

		for _, tst := range testCase {
			rsp := maxProfitIII(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func maxProfitIII(prices []int) int {
	var (
		n  = len(prices)
		dp = make([][][2]int, n+1) // 天、 最多次数、是否持股
		mn = math.MinInt32         // 避免爆int64
	)
	// 初始化
	for i := 0; i < n+1; i++ {
		dp[i] = make([][2]int, 2+2)
		for j := range dp[i] {
			dp[i][j] = [2]int{mn, mn}
		}
	}
	for j := 1; j < 2+2; j++ {
		dp[0][j][0] = 0
	}

	// 递推
	for i := 0; i < n; i++ {
		for j := 1; j < 2+2; j++ {
			dp[i+1][j][0] = max(dp[i][j][0], dp[i][j][1]+prices[i])
			dp[i+1][j][1] = max(dp[i][j][1], dp[i][j-1][0]-prices[i])
		}
	}
	return dp[n][3][0]

}

// 记忆化搜索并不是万能的，某些题目只有写成递推,才能结合数据结构等来优化时间复杂度
func maxProfitIIIBack(prices []int) int {
	/*
		- 未持有 --(买入)--> 持有   dfs(i,1)=dfs(i-1,0)-price[i]
		- 持有 --(卖出-)--> 未持有  dfs(i,0)=dfs(i-1,1)+price[i]
		- 未持有(do nothing) dfs(i,0) = dfs(i-1,0)
		- 持有(do nothing)  dfs(i,1) = dfs(i-1,1)
	*/

	length := len(prices)
	var cache = make([][][2]int, length)
	for i := range cache {
		cache[i] = make([][2]int, 2+1)
		for j := range cache[i] {
			cache[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(i int, count int, hold int) int
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
		res = cache[i][count][hold]
		if res != -1 {
			return res
		}
		defer func() { cache[i][count][hold] = res }()

		if hold == 1 {
			return max(dfs(i-1, count, hold), dfs(i-1, count-1, 0)-prices[i])
		}
		return max(dfs(i-1, count, hold), dfs(i-1, count, 1)+prices[i])
	}
	return dfs(length-1, 2, 0) // 最后一天卖出股票
}
