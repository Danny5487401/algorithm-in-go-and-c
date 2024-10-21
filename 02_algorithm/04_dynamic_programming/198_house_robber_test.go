package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 打家劫舍 https://leetcode.cn/problems/house-robber/description/
func TestRob(t *testing.T) {
	convey.Convey("打家劫舍", t, func() {
		testCase := []struct {
			input  []int
			target int
		}{
			{

				[]int{1, 2, 3, 1}, 4,
			},
			{

				[]int{2, 7, 9, 3, 1}, 12,
			},
		}

		for _, tst := range testCase {
			rsp := rob(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

// 回溯的方式
// 1 递归搜索 + 保存计算结果 = 记忆化搜索
func rob1(nums []int) int {
	// dfs[i] = max(dfs(i-1), dfs(i-2)+nums[i])
	length := len(nums)
	memo := make([]int, length) // 缓存
	for i := range memo {
		memo[i] = -1 // 因为都是正数， -1 表示没有计算过

	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i < 0 { // 递归边界（没有房子）
			return 0
		}

		if memo[i] != -1 { // 之前计算过
			return memo[i]
		}
		res := max(dfs(i-1), dfs(i-2)+nums[i])
		memo[i] = res
		return res
	}
	return dfs(length - 1) // 从最后一个房子开始思考
}

// 2 递推
func rob(nums []int) int {
	f0, f1 := 0, 0
	for _, x := range nums {
		f0, f1 = f1, max(f1, f0+x)
	}
	// 注意返回值是 f1 不是 f0
	return f1
}

func rob2(nums []int) int {
	// dfs[i] = max(dfs(i-1), dfs(i-2)+nums[i])
	// 数组转换--> f[i]=max(f[i-1],f[i-2]+num[i])
	// 因为防止负数 --> f[i+2]=max(f[i+1],f[i]+num[i])
	n := len(nums)
	f := make([]int, n+2)

	for i, x := range nums {
		f[i+2] = max(f[i+1], f[i]+x)

	}
	return f[n+1] // n-1 + 2
}
