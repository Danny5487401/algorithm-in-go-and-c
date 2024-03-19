package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 目标和 https://leetcode.cn/problems/target-sum/description/

func TestFindTargetSumWays(t *testing.T) {
	convey.Convey("目标和 ", t, func() {
		testCase := []struct {
			input struct {
				nums   []int
				target int
			}
			target int
		}{
			{

				struct {
					nums   []int
					target int
				}{nums: []int{1, 1, 1, 1, 1}, target: 3}, 5,
			},
		}

		for _, tst := range testCase {
			rsp := findTargetSumWays(tst.input.nums, tst.input.target)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func findTargetSumWays(nums []int, target int) int {
	// p 添加所有正数的和
	// s 所有元素的和
	// s-p 所有负数的和
	// 目标和 t = p - (s-p)= 2p-s -->p=(s+t)/2 含义 从nums 中选择一些数字==一个数（方案数量），并且和是偶数,并且非负数

	// 计算 s+t
	for _, num := range nums {
		target += num
	}
	if target < 0 || target%2 == 1 {
		return 0
	}

	target = target / 2
	length := len(nums)

	// 缓存

	var dfs func(i int, c int) int
	dfs = func(i int, c int) int {
		// 从后往前
		if i < 0 {
			if c == 0 {
				// 代表找到了一个方案
				return 1
			}
			return 0
		}
		if c < nums[i] {
			return dfs(i-1, c)
		}
		return dfs(i-1, c) + dfs(i-1, c-nums[i])

	}
	return dfs(length-1, target)

}

// 记忆化搜索转成递推
