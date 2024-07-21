package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 爬楼梯 https://leetcode.cn/problems/climbing-stairs/?envType=study-plan-v2&envId=top-100-liked

func TestClimbStairs(t *testing.T) {
	convey.Convey("爬楼梯 ", t, func() {
		testCase := []struct {
			input  int
			target int
		}{
			{

				2, 2,
			},
			{

				3, 3,
			},
		}

		for _, tst := range testCase {
			rsp := climbStairs(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func climbStairs(n int) int {
	// dfs(n)= dfs(n-1)+dfs(n-2)

	var dfs func(i int) int

	var cache = make([]int, n+1) // 注意 开始为 n
	dfs = func(i int) (res int) {
		if i <= 1 {
			return 1
		}
		p := &cache[i]
		if *p != 0 {
			return cache[i]
		}
		res = dfs(i-1) + dfs(i-2)
		*p = res
		return

	}

	return dfs(n)
}

func climbStairs2(n int) int {
	// n >= 1
	// f[n]= f[n-1]+f[n-2]
	// f[n+2]= f[n+1]+f[n]
	var f = make([]int, n+2)
	f[0] = 1 // 注意这里是 1， 不是 0
	f[1] = 1
	for i := 0; i < n; i++ {
		f[i+2] = f[i+1] + f[i]
	}
	return f[n] // 注意这里是 f[n], 不是 f[n-1]

}
