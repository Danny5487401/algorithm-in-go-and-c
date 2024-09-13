package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 爬楼梯 https://leetcode.cn/problems/climbing-stairs/?envType=study-plan-v2&envId=top-100-liked

func TestMinCostClimbingStairs(t *testing.T) {
	convey.Convey("使用最小花费爬楼梯 :向上爬一个或者两个台阶,可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{

				[]int{10, 15, 20}, 15,
			},
			{

				[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6,
			},
		}

		for _, tst := range testCase {
			rsp := minCostClimbingStairs(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	// 表示从 0 或 1 爬到 i 的最小花费
	cache := make([]int, length+1)
	// 0 <= cost[i] <= 999
	for i := range cache {
		cache[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i <= 1 {
			// 递归边界：dfs(0)=0, dfs(1)=0
			return 0
		}
		if cache[i] != -1 {
			return cache[i]
		}
		defer func() {
			cache[i] = res
		}()

		res = min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])
		return

	}

	return dfs(length)

}
