package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 零钱兑换 https://leetcode.cn/problems/coin-change/?envType=study-plan-v2&envId=top-100-liked

func TestCoinChange(t *testing.T) {
	convey.Convey("零钱兑换 ", t, func() {
		testCase := []struct {
			input  []int
			amount int
			target int
		}{
			{

				[]int{1, 2, 5}, 11, 3,
			},
			{

				[]int{2}, 3, -1,
			},
			{

				[]int{1}, 0, 0,
			},
		}

		for _, tst := range testCase {
			rsp := coinChange(tst.input, tst.amount)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func coinChange(coins []int, amount int) int {
	length := len(coins)

	memo := make([][]int, length)
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有访问过
		}
	}

	var dfs func(i int, c int) int
	dfs = func(i int, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 0 // 注意不是返回 1，因为 i <0

			}
			return math.MaxInt / 2 // 因为取最小值，所以这里取无穷大，除 2 是防止下面方案 + 1 溢出
		}
		p := &memo[i][c]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()

		if c < coins[i] {
			return dfs(i-1, c)
		}
		return min(dfs(i-1, c), dfs(i, c-coins[i])+1) // 方案+1
	}
	ans := dfs(length-1, amount)
	if ans < math.MaxInt/2 {
		return ans

	}
	return -1

}
